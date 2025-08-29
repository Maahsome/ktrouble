package config

import (
	"fmt"
	"io"
	"ktrouble/common"
	"ktrouble/defaults"
	"ktrouble/migrate"
	"ktrouble/objects"
	"os"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func (c *Config) MigrateLocal(toVer string) bool {

	migrationSuccess := false
	startMajor := 1
	switch c.ConfigVersion {
	case "v0":
		startMajor = 1
	case "v1":
		startMajor = 2
	}
	for v := startMajor; v <= c.Semver.Major; v++ {
		switch v {
		case 1:
			common.Logger.Warnf("Migrating config from v0 to v1")
			migrationSuccess = c.UpgradeToV1()
		case 2:
			common.Logger.Warnf("Migrating config from v1 to v2")
			migrationSuccess = c.UpgradeToV2()
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

func (c *Config) UpgradeToV2() bool {
	// Copy the config file as a backup
	err := BackupConfig(viper.ConfigFileUsed(), fmt.Sprintf("%s-v2-upgrade-%s.bak", viper.ConfigFileUsed(), time.Now().Format("20060102-150405")))
	if err != nil {
		common.Logger.WithError(err).Fatal("Failed to copy config file")
	}

	// Tags Definitions are new to v2, so no migration is needed
	// Renamed "Repository" to "Image" for v2, so we will need to split "Repository" into "Image" and "Tags"

	// Output Fields Definitions
	common.Logger.Warn("Modifying output fields for utility definitions to v2 format")
	newOutputFieldDefs := make([]objects.OutputFields, 0, len(c.OutputFieldsDefs))
	for _, v := range c.OutputFieldsDefs {
		if strings.Contains(v.Fields, "REPOSITORY") && v.Name == "utility" {
			newFields := strings.ReplaceAll(v.Fields, "REPOSITORY", "IMAGE,TAGS")
			newOutputFieldDefs = append(newOutputFieldDefs, objects.OutputFields{
				Name:   v.Name,
				Fields: newFields,
			})
		} else {
			newOutputFieldDefs = append(newOutputFieldDefs, v)
		}
	}
	c.OutputFieldsDefs = newOutputFieldDefs

	// Utility Definitions
	utilDefsV1 := objects.UtilityPodListV1{}
	uerr := viper.UnmarshalKey("utilityDefinitions", &utilDefsV1)
	if uerr != nil {
		common.Logger.Fatal("Error unmarshalling utility defs...")
		return false
	}
	if len(utilDefsV1) == 0 {
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
		common.Logger.Warn("Migrating utility definitions to v2 format")
		c.UtilDefs = objects.UtilityPodList{}
		for i := range utilDefsV1 {
			c.UtilDefs = append(c.UtilDefs, migrate.UpdateUtilityV2(utilDefsV1[i]))
		}
		viper.Set("utilityDefinitions", c.UtilDefs)
		viper.Set("outputFields", c.OutputFieldsDefs)
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
