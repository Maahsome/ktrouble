package config

import (
	"fmt"
	"io"
	"ktrouble/common"
	"ktrouble/defaults"
	"ktrouble/objects"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func (c *Config) MigrateLocal(toVer string) bool {

	migrationSuccess := false
	switch c.ConfigVersion {
	case "v0":
		for v := 1; v <= c.Semver.Major; v++ {
			switch v {
			case 1:
				common.Logger.Warnf("Migrating config from v0 to v1")
				migrationSuccess = c.UpgradeToV1()
			}
		}
	}

	return migrationSuccess
}

func (c *Config) UpgradeToV1() bool {
	// Copy the config file as a backup
	err := BackupConfig(viper.ConfigFileUsed(), fmt.Sprintf("%s-v1-upgrade-%s.bak", viper.ConfigFileUsed(), time.Now().Format("20060102-150405")))
	if err != nil {
		common.Logger.WithError(err).Fatal("Failed to copy config file")
	}

	// Environment Definitions are new to v1, so no migration is needed

	// Utility Definitions
	utilDefsV0 := objects.UtilityPodListV0{}
	uerr := viper.UnmarshalKey("utilityDefinitions", &utilDefsV0)
	if uerr != nil {
		common.Logger.Fatal("Error unmarshalling utility defs...")
		return false
	}
	if len(utilDefsV0) == 0 {
		common.Logger.Warn("Adding default utility definitions to config.yaml")
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
		common.Logger.Warn("Migrating utility definitions to v1 format")
		utilDefs := objects.UtilityPodList{}
		uerr := viper.UnmarshalKey("utilityDefinitions", &utilDefs)
		if uerr != nil {
			common.Logger.Fatal("Error unmarshalling utility defs...")
		}
		for i := range utilDefs {
			c.UtilDefs[i].RemoveUpstream = false
			c.UtilDefs[i].Environments = []string{}
		}
		viper.Set("utilityDefinitions", c.UtilDefs)
		verr := viper.WriteConfig()
		if verr != nil {
			logrus.WithError(verr).Info("Failed to write config")
		}
	}

	return true
}

func BackupConfig(src, dest string) error {
	common.Logger.Warnf("Backing up config file from %s to %s", src, dest)
	srcFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open source file: %w", err)
	}
	defer srcFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %w", err)
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return fmt.Errorf("failed to copy file: %w", err)
	}

	return nil
}
