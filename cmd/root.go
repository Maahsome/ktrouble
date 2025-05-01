package cmd

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"ktrouble/cmd/add"
	"ktrouble/cmd/edit"
	"ktrouble/cmd/get"
	"ktrouble/cmd/remove"
	"ktrouble/cmd/set"
	"ktrouble/cmd/update"
	"ktrouble/common"
	"ktrouble/config"
	"ktrouble/defaults"
	"ktrouble/gitupstream"
	"ktrouble/kubernetes"
	"ktrouble/objects"

	"ktrouble/help"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	// "gopkg.in/yaml.v3"
	// "sigs.k8s.io/yaml"
)

type (
	Project struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Path   string `json:"path"`
		SSHURL string `json:"ssh_url_to_repo"`
	}
)

var (
	cfgFile      string
	semVer       string
	gitCommit    string
	gitRef       string
	buildDate    string
	changelogURL string

	// semVerReg - gets the semVer portion only, cutting off any other release details
	semVerReg = regexp.MustCompile(`(v[0-9]+\.[0-9]+\.[0-9]+).*`)

	c = &config.Config{}
)

var deleteHelp = help.DeleteCmd{}
var launchHelp = help.LaunchCmd{}
var migrateHelp = help.MigrateCmd{}
var pullHelp = help.PullCmd{}
var pushHelp = help.PushCmd{}
var statusHelp = help.StatusCmd{}
var changelogHelp = help.ChangelogCmd{}
var fieldsHelp = help.FieldsCmd{}
var genhelpHelp = help.GenHelpCmd{}
var versionHelp = help.VersionCmd{}

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "ktrouble",
	Short: "A tool for launching PODs from a curated list of utility PODs",
	Long: `EXAMPLE:
  Simply run the 'launch' command and you will be prompted for all of the
  required details.
    - Utility Pod Selection
    - Namespace
    - Service Account
    - Node Selector
    - Resource Sizing

  > ktrouble launch
`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {

		var err error

		logFile, _ := cmd.Flags().GetString("log-file")
		logLevel, _ := cmd.Flags().GetString("log-level")
		ll := "Info"
		switch strings.ToLower(logLevel) {
		case "trace":
			ll = "Trace"
		case "debug":
			ll = "Debug"
		case "info":
			ll = "Info"
		case "warning":
			ll = "Warning"
		case "error":
			ll = "Error"
		case "fatal":
			ll = "Fatal"
		}

		common.NewLogger(ll, logFile)

		c.VersionDetail.SemVer = semVer
		c.VersionDetail.BuildDate = buildDate
		c.VersionDetail.GitCommit = gitCommit
		c.VersionDetail.GitRef = gitRef
		c.VersionJSON = fmt.Sprintf("{\"SemVer\": \"%s\", \"BuildDate\": \"%s\", \"GitCommit\": \"%s\", \"GitRef\": \"%s\"}", semVer, buildDate, gitCommit, gitRef)
		c.Semver, err = config.ParseSemver(semVer)
		if err != nil {
			common.Logger.WithError(err).Fatal("Failed to parse the semver")
		}

		// Upgrade check and upgrade config
		if c.ConfigVersion != fmt.Sprintf("v%d", c.Semver.Major) {
			common.Logger.Warnf("Config version %s is not the same as the current application version %s, migrating...", c.ConfigVersion, fmt.Sprintf("v%d", c.Semver.Major))
			migrate := c.MigrateLocal(fmt.Sprintf("v%d", c.Semver.Major))
			if migrate {
				c.ConfigVersion = fmt.Sprintf("v%d", c.Semver.Major)
				viper.Set("configVersion", c.ConfigVersion)
				verr := viper.WriteConfig()
				if verr != nil {
					logrus.WithError(verr).Info("Failed to write config")
				}
			}
		}

		if c.OutputFormat != "" {
			c.FormatOverridden = true
			c.NoHeaders = false
			c.OutputFormat = strings.ToLower(c.OutputFormat)
			switch c.OutputFormat {
			case "json", "gron", "yaml", "text", "table", "raw":
				break
			default:
				fmt.Println("Valid options for -o are [json|gron|text|table|yaml|raw]")
				os.Exit(1)
			}
		}

		if os.Args[1] == "pull" || os.Args[1] == "push" || os.Args[1] == "status" || os.Args[1] == "diff" || os.Args[1] == "migrate" {
			gitUser := viper.GetString("gitUser")
			if len(gitUser) == 0 {
				common.Logger.Fatal("gitUser is not set, use 'ktrouble set config --help'")
			}
			gitTokenVar := ""
			gitToken := viper.GetString("gitToken")
			if len(gitToken) == 0 {
				gitTokenVar = viper.GetString("gitTokenVar")
				if len(gitTokenVar) == 0 {
					common.Logger.Fatal("gitToken or gitTokenVar config option is not set, use 'ktrouble set config --help'")
				}
				gitToken = os.Getenv(gitTokenVar)
			}
			gitURL := viper.GetString("gitURL")
			if len(gitURL) == 0 {
				common.Logger.Fatal("gitURL is not set, use 'ktrouble set config --help'")
			}

			if len(gitToken) == 0 {
				common.Logger.Fatalf("no git token set, gitToken or %s ENV VAR is not set, use 'ktrouble set config --help'", gitTokenVar)
			}

			common.Logger.Tracef("Semver.Major: %d", c.Semver.Major)
			c.GitUpstream = gitupstream.New(gitUser, gitToken, gitURL, fmt.Sprintf("v%d", c.Semver.Major))
		}
		subCmd := ""
		if len(os.Args) > 2 {
			subCmd = os.Args[2]
		}
		if needKubernetes(os.Args[1], subCmd) {
			common.Logger.Trace("Creating Kubernetes client")
			c.Client = kubernetes.New()
		}
	},
}

func containsAlias(v string, a []string) bool {
	for _, i := range a {
		if i == v {
			return true
		}
	}
	return false
}

func needKubernetes(arg string, sub string) bool {

	// handle the TOP level kubernetes commands
	switch arg {
	case "attach", "delete", "launch":
		return true
	}
	switch {
	case containsAlias(arg, defaults.AttachAliases),
		containsAlias(arg, defaults.DeleteAliases),
		containsAlias(arg, defaults.LaunchAliases):
		return true
	}

	// handle the get 'subcommands' and subcommand aliases
	if arg == "get" {
		switch sub {
		case "attachments", "ingresses", "namespace", "node", "nodelabels", "running", "serviceaccount", "services":
			return true
		}
		switch {
		case containsAlias(sub, defaults.GetIngressesAliases),
			containsAlias(sub, defaults.GetNamespacesAliases),
			containsAlias(sub, defaults.GetNodesAliases),
			containsAlias(sub, defaults.GetNodeLabelsAliases),
			containsAlias(sub, defaults.GetRunningAliases),
			containsAlias(sub, defaults.GetServiceAccountsAliases),
			containsAlias(sub, defaults.GetServicesAliases):
			return true
		}
	}

	return false
}

func buildRootCmd() *cobra.Command {
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.splicectl/config.yml)")
	RootCmd.PersistentFlags().StringVarP(&c.OutputFormat, "output", "o", "", "output types: json, text, yaml, gron, raw")
	RootCmd.PersistentFlags().BoolVar(&c.NoHeaders, "no-headers", false, "Suppress header output in Text output")
	RootCmd.PersistentFlags().StringVarP(&c.LogLevel, "log-level", "v", "", "Set the logging level: trace,debug,info,warning,error,fatal")
	RootCmd.PersistentFlags().StringVar(&c.LogFile, "log-file", "", "Set the logging level: trace,debug,info,warning,error,fatal")
	RootCmd.PersistentFlags().StringVarP(&c.Namespace, "namespace", "n", "", "Specify the namespace to run in, ENV NAMESPACE then -n for preference")
	RootCmd.PersistentFlags().BoolVarP(&c.ShowHidden, "show-hidden", "s", false, "Show entries with the 'hidden' property set to 'true'")
	RootCmd.PersistentFlags().StringSliceVarP(&c.Fields, "fields", "f", []string{}, "Specify an array of field names: eg, --fields 'NAME,REPOSITORY'")
	RootCmd.PersistentFlags().StringVarP(&c.TemplateFile, "template", "t", "default", "Specify the template file to use to render the POD manifest")
	RootCmd.PersistentFlags().StringVar(&c.ServiceTemplateFile, "service-template", "default-service", "Specify the service template file to use to render the SERVICE manifest, for --create-ingress option")
	RootCmd.PersistentFlags().StringVar(&c.IngressTemplateFile, "ingress-template", "default-ingress", "Specify the ingress template file to use to render the INGRESS manifest, for --create-ingress option")
	return RootCmd
}

func addSubCommands() {
	RootCmd.AddCommand(
		// from 'import ktrouble/cmd/<subcommand:package>'
		// <package>.InitSubCommands(c),
		get.InitSubCommands(c),
		add.InitSubCommands(c),
		edit.InitSubCommands(c),
		remove.InitSubCommands(c),
		update.InitSubCommands(c),
		set.InitSubCommands(c),
	)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	buildRootCmd()
	cobra.OnInitialize(initConfig)
	addSubCommands()
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	home, herr := os.UserHomeDir()
	cobra.CheckErr(herr)
	confDir := fmt.Sprintf("%s/.config/ktrouble", home)
	tmplDir := fmt.Sprintf("%s/.config/ktrouble/templates", home)
	envCfgFile := os.Getenv("KTROUBLE_CONFIG")
	if envCfgFile != "" {
		logrus.Debug("Using KTROUBLE_CONFIG")
		configFile := fmt.Sprintf("%s/%s", confDir, envCfgFile)
		createRestrictedConfigFile(configFile)
		viper.SetConfigFile(configFile)
	} else {
		if cfgFile != "" {
			// Use config file from the flag.
			viper.SetConfigFile(cfgFile)
		} else {
			// Find home directory.
			if _, err := os.Stat(tmplDir); err != nil {
				if os.IsNotExist(err) {
					mkerr := os.MkdirAll(tmplDir, os.ModePerm)
					if mkerr != nil {
						logrus.WithError(mkerr).Fatal("Error creating ~/.config/ktrouble/templates directory")
					}
				}
			}
			if stat, err := os.Stat(confDir); err == nil && stat.IsDir() {
				configFile := fmt.Sprintf("%s/%s", confDir, "config.yaml")
				createRestrictedConfigFile(configFile)
				viper.SetConfigFile(configFile)
			} else {
				logrus.Info("The ~/.config/ktrouble path is a file and not a directory, please remove the 'ktrouble' file.")
				os.Exit(1)
			}
		}
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		logrus.Warn("Failed to read viper config file.")
	}

	// ConfigVersion
	if viper.IsSet("configVersion") {
		c.ConfigVersion = viper.GetString("configVersion")
	} else {
		viper.Set("configVersion", "v0")
		c.ConfigVersion = "v0"
		verr := viper.WriteConfig()
		if verr != nil {
			logrus.WithError(verr).Info("Failed to write config")
		}
	}

	// Environment Definitions
	err := viper.UnmarshalKey("environments", &c.EnvDefs)
	if err != nil {
		logrus.Fatal("Error unmarshalling environment defs...")
	}
	if len(c.EnvDefs) > 0 {
		c.EnvMap = make(map[string]objects.Environment, len(c.EnvDefs))
		for _, v := range c.EnvDefs {
			c.EnvMap[v.Name] = v
		}
	}

	// Utility Definitions
	uerr := viper.UnmarshalKey("utilityDefinitions", &c.UtilDefs)
	if uerr != nil {
		logrus.Fatal("Error unmarshalling utility defs...")
	}
	if len(c.UtilDefs) == 0 {
		logrus.Warn("Adding default utility definitions to config.yaml")
		seedDefs := defaults.UtilityDefinitions()
		viper.Set("utilityDefinitions", seedDefs)
		c.UtilDefs = defaults.UtilityDefinitions()
		c.UtilMap = make(map[string]objects.UtilityPod, len(c.UtilDefs))
		for _, v := range c.UtilDefs {
			c.UtilMap[v.Name] = v
		}
		verr := viper.WriteConfig()
		if verr != nil {
			logrus.WithError(verr).Info("Failed to write config")
		}
	} else {
		updatedSources := false
		defaultDefs := defaults.UtilityDefinitions()
		c.UtilMap = make(map[string]objects.UtilityPod, len(c.UtilDefs))
		for i, v := range c.UtilDefs {
			c.UtilMap[v.Name] = v
			if len(v.Source) == 0 {
				c.UtilDefs[i].Source = whichSource(defaultDefs, v.Name)
				updatedSources = true
			}
		}
		if updatedSources {
			viper.Set("utilityDefinitions", c.UtilDefs)
			verr := viper.WriteConfig()
			if verr != nil {
				logrus.WithError(verr).Info("Failed to write config")
			}
		}
	}

	// Size Definitions
	serr := viper.UnmarshalKey("resourceSizing", &c.SizeDefs)
	if serr != nil {
		logrus.Fatal("Error unmarshalling resource sizing...")
	}
	if len(c.SizeDefs) == 0 {
		logrus.Warn("Adding default resource sizing to config.yaml")
		seedSizes := defaults.ResourceSizingList()
		viper.Set("resourceSizing", seedSizes)
		c.SizeDefs = defaults.ResourceSizingList()
		c.SizeMap = make(map[string]objects.ResourceSize, len(c.SizeDefs))
		for _, v := range c.SizeDefs {
			c.SizeMap[v.Name] = v
		}
		verr := viper.WriteConfig()
		if verr != nil {
			logrus.WithError(verr).Info("Failed to write config")
		}
	} else {
		c.SizeMap = make(map[string]objects.ResourceSize, len(c.SizeDefs))
		for _, v := range c.SizeDefs {
			c.SizeMap[v.Name] = v
		}
	}

	// Ephemeral Sleep Definitions
	eserr := viper.UnmarshalKey("ephemeralSleep", &c.EphemeralSleepDefs)
	if eserr != nil {
		logrus.Fatal("Error unmarshalling ephemeral sleep...")
	}
	if len(c.EphemeralSleepDefs) == 0 {
		logrus.Warn("Adding default ephemeral sleep to config.yaml")
		seedEphemeralSleep := defaults.EphemeralSleepList()
		viper.Set("ephemeralSleep", seedEphemeralSleep)
		c.EphemeralSleepDefs = defaults.EphemeralSleepList()
		c.EphemeralSleepMap = make(map[string]objects.EphemeralSleep, len(c.EphemeralSleepDefs))
		for _, v := range c.EphemeralSleepDefs {
			c.EphemeralSleepMap[v.Name] = v
		}
		verr := viper.WriteConfig()
		if verr != nil {
			logrus.WithError(verr).Info("Failed to write config")
		}
	} else {
		c.EphemeralSleepMap = make(map[string]objects.EphemeralSleep, len(c.EphemeralSleepDefs))
		for _, v := range c.EphemeralSleepDefs {
			c.EphemeralSleepMap[v.Name] = v
		}
	}

	// Node Selector Labels
	nerr := viper.UnmarshalKey("nodeSelectorLabels", &c.NodeSelectorLabels)
	if nerr != nil {
		logrus.Fatal("Error unmarshalling node selector labels...")
	}
	if len(c.NodeSelectorLabels) == 0 {
		logrus.Warn("Adding default node selector labels to config.yaml")
		seedLabels := defaults.Labels()
		viper.Set("nodeSelectorLabels", seedLabels)
		c.NodeSelectorLabels = defaults.Labels()
		verr := viper.WriteConfig()
		if verr != nil {
			logrus.WithError(verr).Info("Failed to write config")
		}
	}

	// Unique ID Length
	if viper.IsSet("uniqIdLength") {
		c.UniqIdLength = viper.GetInt("uniqIdLength")
	} else {
		// Set the default
		viper.Set("uniqIdLength", 6)
		c.UniqIdLength = 6
		verr := viper.WriteConfig()
		if verr != nil {
			logrus.WithError(verr).Info("Failed to write config")
		}
	}

	// EnableBashLinks
	if viper.IsSet("enableBashLinks") {
		c.EnableBashLinks = viper.GetBool("enableBashLinks")
	} else {
		// Set the default
		viper.Set("enableBashLinks", false)
		c.EnableBashLinks = false
		verr := viper.WriteConfig()
		if verr != nil {
			logrus.WithError(verr).Info("Failed to write config")
		}
	}

	// GitURL
	if viper.IsSet("gitURL") {
		c.GitURL = viper.GetString("gitURL")
	} else {
		// Set the default
		viper.Set("gitURL", defaults.GitURL())
		c.GitURL = defaults.GitURL()
		verr := viper.WriteConfig()
		if verr != nil {
			logrus.WithError(verr).Info("Failed to write config")
		}
	}

	// GitUser
	if viper.IsSet("gitUser") {
		c.GitUser = viper.GetString("gitUser")
	} else {
		// Set the default
		viper.Set("gitUser", os.Getenv("USER"))
		c.GitUser = os.Getenv("USER")
		verr := viper.WriteConfig()
		if verr != nil {
			logrus.WithError(verr).Info("Failed to write config")
		}
	}

	// GitTokenVar
	if !viper.IsSet("GitTokenVar") {
		tokenVar := defaults.GitTokenVar()
		tv := os.Getenv(tokenVar)
		if len(tv) == 0 {
			tokenVar = "GIT_TOKEN"
		}
		// Set the default
		viper.Set("GitTokenVar", tokenVar)
		verr := viper.WriteConfig()
		if verr != nil {
			logrus.WithError(verr).Info("Failed to write config")
		}
	}

	// PromptForSecrets
	if viper.IsSet("promptForSecrets") {
		c.PromptForSecrets = viper.GetBool("promptForSecrets")
	} else {
		// Set the default
		viper.Set("promptForSecrets", false)
		c.PromptForSecrets = false
		verr := viper.WriteConfig()
		if verr != nil {
			logrus.WithError(verr).Info("Failed to write config")
		}
	}

	// PromptForConfigMaps
	if viper.IsSet("promptForConfigMaps") {
		c.PromptForConfigMaps = viper.GetBool("promptForConfigMaps")
	} else {
		// Set the default
		viper.Set("promptForConfigMaps", false)
		c.PromptForConfigMaps = false
		verr := viper.WriteConfig()
		if verr != nil {
			logrus.WithError(verr).Info("Failed to write config")
		}
	}

	// Create 'default' template in the config templates directory
	templateFile := fmt.Sprintf("%s/default", tmplDir)
	createDefaultTemplateFile(templateFile)

	// Create 'default' service template in the config templates directory
	serviceTemplateFile := fmt.Sprintf("%s/default-service", tmplDir)
	createDefaultServiceTemplateFile(serviceTemplateFile)

	// Create 'default' ingress template in the config templates directory
	ingressTemplateFile := fmt.Sprintf("%s/default-ingress", tmplDir)
	createDefaultIngressTemplateFile(ingressTemplateFile)
}

// whichSource returns 'ktrouble-utils' if the utility name is in the default list
// otherwise it returns 'local' which would be something that was added locally.
// this function is to bring the config.yaml up to date with new properties added
func whichSource(defList []objects.UtilityPod, name string) string {

	source := "local"

	for _, v := range defList {
		if name == v.Name {
			source = "ktrouble-utils"
			break
		}
	}

	return source
}

func createDefaultTemplateFile(fileName string) {
	if _, err := os.Stat(fileName); err != nil {
		if os.IsNotExist(err) {
			file, ferr := os.Create(fileName)
			if ferr != nil {
				logrus.Fatalf("Unable to create the default template file: %s", fileName)
			}
			mode := int(0600)
			if cherr := file.Chmod(os.FileMode(mode)); cherr != nil {
				logrus.Error("Chmod for default template file failed, please set the mode to 0600.")
			}
			_, werr := file.WriteString(defaults.DefaultTemplate())
			if werr != nil {
				logrus.Error("failed to write the default template")
			}
		}
	} else {
		// determine if we can update this file
		currentTemplateData, rerr := os.ReadFile(fileName)
		if rerr != nil {
			logrus.Fatal("Unable to read from the existing template file, permission issue?")
		}

		if string(currentTemplateData) != defaults.DefaultTemplate() {
			backupFile := fmt.Sprintf("%s.saved-%s", fileName, time.Now().Format("20060102150405"))
			logrus.Warnf("current default template has been updated, the previous has been saved as %s.", backupFile)
			// create the backup
			if fileExists(backupFile) {
				logrus.Fatalf("The backup file, %s, already exists, we cannot proceed since an update to the default template is needed.  Please remove the file.", backupFile)
			}
			mode := int(0600)
			os.WriteFile(backupFile, currentTemplateData, os.FileMode(mode))
			// overwrite the current file
			file, ferr := os.Create(fileName)
			if ferr != nil {
				logrus.Fatalf("Unable to create the default template file: %s", fileName)
			}
			if cherr := file.Chmod(os.FileMode(mode)); cherr != nil {
				logrus.Error("Chmod for default template file failed, please set the mode to 0600.")
			}
			_, werr := file.WriteString(defaults.DefaultTemplate())
			if werr != nil {
				logrus.Error("failed to write the default template")
			}
		}
	}
}

func createDefaultServiceTemplateFile(fileName string) {
	if _, err := os.Stat(fileName); err != nil {
		if os.IsNotExist(err) {
			file, ferr := os.Create(fileName)
			if ferr != nil {
				logrus.Fatalf("Unable to create the default service template file: %s", fileName)
			}
			mode := int(0600)
			if cherr := file.Chmod(os.FileMode(mode)); cherr != nil {
				logrus.Error("Chmod for default service template file failed, please set the mode to 0600.")
			}
			_, werr := file.WriteString(defaults.DefaultServiceTemplate())
			if werr != nil {
				logrus.Error("failed to write the default template")
			}
		}
	} else {
		// determine if we can update this file
		currentTemplateData, rerr := os.ReadFile(fileName)
		if rerr != nil {
			logrus.Fatal("Unable to read from the existing service template file, permission issue?")
		}

		if string(currentTemplateData) != defaults.DefaultServiceTemplate() {
			backupFile := fmt.Sprintf("%s.saved-%s", fileName, time.Now().Format("20060102150405"))
			logrus.Warnf("current default service template has been updated, the previous has been saved as %s.", backupFile)
			// create the backup
			if fileExists(backupFile) {
				logrus.Fatalf("The backup file, %s, already exists, we cannot proceed since an update to the default service template is needed.  Please remove the file.", backupFile)
			}
			mode := int(0600)
			os.WriteFile(backupFile, currentTemplateData, os.FileMode(mode))
			// overwrite the current file
			file, ferr := os.Create(fileName)
			if ferr != nil {
				logrus.Fatalf("Unable to create the default service template file: %s", fileName)
			}
			if cherr := file.Chmod(os.FileMode(mode)); cherr != nil {
				logrus.Error("Chmod for default service template file failed, please set the mode to 0600.")
			}
			_, werr := file.WriteString(defaults.DefaultServiceTemplate())
			if werr != nil {
				logrus.Error("failed to write the default service template")
			}
		}
	}
}

func createDefaultIngressTemplateFile(fileName string) {
	if _, err := os.Stat(fileName); err != nil {
		if os.IsNotExist(err) {
			file, ferr := os.Create(fileName)
			if ferr != nil {
				logrus.Fatalf("Unable to create the default ingress template file: %s", fileName)
			}
			mode := int(0600)
			if cherr := file.Chmod(os.FileMode(mode)); cherr != nil {
				logrus.Error("Chmod for default template file failed, please set the mode to 0600.")
			}
			_, werr := file.WriteString(defaults.DefaultIngressTemplate())
			if werr != nil {
				logrus.Error("failed to write the default ingress template")
			}
		}
	} else {
		// determine if we can update this file
		currentTemplateData, rerr := os.ReadFile(fileName)
		if rerr != nil {
			logrus.Fatal("Unable to read from the existing ingress template file, permission issue?")
		}

		if string(currentTemplateData) != defaults.DefaultIngressTemplate() {
			backupFile := fmt.Sprintf("%s.saved-%s", fileName, time.Now().Format("20060102150405"))
			logrus.Warnf("current default ingress template has been updated, the previous has been saved as %s.", backupFile)
			// create the backup
			if fileExists(backupFile) {
				logrus.Fatalf("The backup file, %s, already exists, we cannot proceed since an update to the default ingress template is needed.  Please remove the file.", backupFile)
			}
			mode := int(0600)
			os.WriteFile(backupFile, currentTemplateData, os.FileMode(mode))
			// overwrite the current file
			file, ferr := os.Create(fileName)
			if ferr != nil {
				logrus.Fatalf("Unable to create the default ingress template file: %s", fileName)
			}
			if cherr := file.Chmod(os.FileMode(mode)); cherr != nil {
				logrus.Error("Chmod for default ingress template file failed, please set the mode to 0600.")
			}
			_, werr := file.WriteString(defaults.DefaultIngressTemplate())
			if werr != nil {
				logrus.Error("failed to write the default ingress template")
			}
		}
	}
}

func createRestrictedConfigFile(fileName string) {
	if _, err := os.Stat(fileName); err != nil {
		if os.IsNotExist(err) {
			file, ferr := os.Create(fileName)
			if ferr != nil {
				logrus.Info("Unable to create the configfile.")
				os.Exit(1)
			}
			mode := int(0600)
			if cherr := file.Chmod(os.FileMode(mode)); cherr != nil {
				logrus.Info("Chmod for config file failed, please set the mode to 0600.")
			}
		}
	}
}

// ClientSemVer - returns the full semVer as the first string and the numerical
// portion as the second string, they may be identical. One example where they
// would not be is:
//
//	semVer: v0.1.1-cacert -> (v0.1.1-cacert, v0.1.1).
func ClientSemVer() (string, string) {
	submatches := semVerReg.FindStringSubmatch(semVer)
	if submatches == nil || len(submatches) < 2 {
		logrus.Fatalf("the semver in the current build is not valid: %s", semVer)
	}
	return submatches[0], submatches[1]
}

// fileExists checks if file exists
func fileExists(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil
}
