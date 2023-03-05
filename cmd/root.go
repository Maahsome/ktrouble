package cmd

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"ktrouble/common"
	"ktrouble/config"
	"ktrouble/objects"

	homedir "github.com/mitchellh/go-homedir"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

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

	c        = &config.Config{}
	utilDefs []objects.UtilityPod
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ktrouble",
	Short: "",
	Long: `EXAMPLE:

  TODO: add description

  > ktrouble

`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {

		rand.Seed(time.Now().UnixNano())

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

		// if os.Args[1] != "version" { // && os.Args[1] != "config" {
		// }
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ktrouble.yaml)")
	rootCmd.PersistentFlags().StringVarP(&c.OutputFormat, "output", "o", "", "Set an output format: json, text, yaml, gron, md")
	rootCmd.PersistentFlags().StringVarP(&c.Namespace, "namespace", "n", "", "Specify the namespace to run in, ENV NAMESPACE then -n for preference")
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
	err := viper.UnmarshalKey("utilityDefinitions", &utilDefs)
	if err != nil {
		logrus.Fatal("Error unmarshalling utility defs...")
	}
	if len(utilDefs) == 0 {
		logrus.Warn("Adding default utility definitions to config.yaml")
		seedDefs := defaultUtilityDefinitions()
		viper.Set("utilityDefinitions", seedDefs)
		verr := viper.WriteConfig()
		if verr != nil {
			logrus.WithError(verr).Info("Failed to write config")
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
		viper.SetConfigFile(fileName)
		defaultLabels := defaultLabelList()
		logrus.Warn("Writing default labels to config.yaml...")
		viper.Set("nodeSelectorLabels", defaultLabels)
		verr := viper.WriteConfig()
		if verr != nil {
			logrus.WithError(verr).Info("Failed to write config")
		}
	}
}

func Commands() []*cobra.Command {
	return rootCmd.Commands()
}

func defaultLabelList() []string {
	return []string{
		"kubernetes.io/arch",
		"kubernetes.io/os",
		"node.kubernetes.io/instance-type",
		"node_pool",
	}
}

func restConfig() (*rest.Config, error) {
	// We aren't likely to run this INSIDE the K8s cluster, this routine
	// simply picks up the config from the file system of a running POD.
	// kubeCfg, err := rest.InClusterConfig()
	var kubeCfg *rest.Config
	var err error

	if kubeconfig := os.Getenv("KUBECONFIG"); kubeconfig != "" {
		kubeCfg, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			logrus.Info("No KUBECONFIG ENV")
			return nil, err
		}
	} else {
		// ENV KUBECONFIG not set, check for ~/.kube/config
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		kubeFile := fmt.Sprintf("%s/%s", home, ".kube/config")
		if _, err := os.Stat(kubeFile); err != nil {
			if os.IsNotExist(err) {
				if os.Args[1] != "version" {
					logrus.Info("Could not locate the KUBECONFIG file, normally ~/.kube/config")
					os.Exit(1)
				}
				return nil, nil
			}
		}
		kubeCfg, err = clientcmd.BuildConfigFromFlags("", kubeFile)
	}
	return kubeCfg, nil
}
