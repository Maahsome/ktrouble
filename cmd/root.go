package cmd

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"ktrouble/cmd/add"
	"ktrouble/cmd/get"
	"ktrouble/cmd/remove"
	"ktrouble/common"
	"ktrouble/config"
	"ktrouble/defaults"
	"ktrouble/kubernetes"
	"ktrouble/objects"

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
	cfgFile   string
	semVer    string
	gitCommit string
	gitRef    string
	buildDate string

	// semVerReg - gets the semVer portion only, cutting off any other release details
	semVerReg = regexp.MustCompile(`(v[0-9]+\.[0-9]+\.[0-9]+).*`)

	c = &config.Config{}
)

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

		logFile, _ := cmd.Flags().GetString("log-file")
		logLevel, _ := cmd.Flags().GetString("log-level")
		ll := "Warning"
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

		if os.Args[1] != "version" {
			c.Client = kubernetes.New()
			// if c.Client == nil {
			// 	common.Logger.Warn("failed to create a kubernetes context")
			// }
		}
	},
}

func buildRootCmd() *cobra.Command {
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.splicectl/config.yml)")
	RootCmd.PersistentFlags().StringVarP(&c.OutputFormat, "output", "o", "", "output types: json, text, yaml, gron, raw")
	RootCmd.PersistentFlags().BoolVar(&c.NoHeaders, "no-headers", false, "Suppress header output in Text output")
	// RootCmd.PersistentFlags().BoolVar(&c., "no-headers", false, "Suppress header output in Text output")
	RootCmd.PersistentFlags().StringVarP(&c.LogLevel, "log-level", "v", "", "Set the logging level: trace,debug,info,warning,error,fatal")
	RootCmd.PersistentFlags().StringVar(&c.LogFile, "log-file", "", "Set the logging level: trace,debug,info,warning,error,fatal")
	RootCmd.PersistentFlags().StringVarP(&c.Namespace, "namespace", "n", "", "Specify the namespace to run in, ENV NAMESPACE then -n for preference")
	RootCmd.PersistentFlags().BoolVarP(&c.ShowHidden, "show-hidden", "s", false, "Show entries with the 'hidden' property set to 'true'")

	return RootCmd
}

func addSubCommands() {
	RootCmd.AddCommand(
		// from 'import ktrouble/cmd/<subcommand:package>'
		// <package>.InitSubCommands(c),
		get.InitSubCommands(c),
		add.InitSubCommands(c),
		remove.InitSubCommands(c),
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
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		workDir := fmt.Sprintf("%s/.config/ktrouble", home)
		if _, err := os.Stat(workDir); err != nil {
			if os.IsNotExist(err) {
				mkerr := os.MkdirAll(workDir, os.ModePerm)
				if mkerr != nil {
					logrus.Fatal("Error creating ~/.config/ktrouble directory", mkerr)
				}
			}
		}
		if stat, err := os.Stat(workDir); err == nil && stat.IsDir() {
			configFile := fmt.Sprintf("%s/%s", workDir, "config.yaml")
			createRestrictedConfigFile(configFile)
			viper.SetConfigFile(configFile)
		} else {
			logrus.Info("The ~/.config/ktrouble path is a file and not a directory, please remove the 'ktrouble' file.")
			os.Exit(1)
		}
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		logrus.Warn("Failed to read viper config file.")
	}

	// Utility Definitions
	err := viper.UnmarshalKey("utilityDefinitions", &c.UtilDefs)
	if err != nil {
		logrus.Fatal("Error unmarshalling utility defs...")
	}
	if len(c.UtilDefs) == 0 {
		logrus.Warn("Adding default utility definitions to config.yaml")
		seedDefs := defaults.UtilityDefinitions()
		viper.Set("utilityDefinitions", seedDefs)
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

	// Node Selector Labels
	nerr := viper.UnmarshalKey("nodeSelectorLabels", &c.NodeSelectorLabels)
	if nerr != nil {
		logrus.Fatal("Error unmarshalling node selector labels...")
	}
	if len(c.NodeSelectorLabels) == 0 {
		logrus.Warn("Adding default node selector labels to config.yaml")
		seedLabels := defaults.Labels()
		viper.Set("nodeSelectorLabels", seedLabels)
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
		verr := viper.WriteConfig()
		if verr != nil {
			logrus.WithError(verr).Info("Failed to write config")
		}
	}

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
