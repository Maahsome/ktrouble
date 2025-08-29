package defaults

import (
	"ktrouble/objects"
)

func UtilityDefinitions() []objects.UtilityPod {
	utilityDefinitionHints := map[string]string{}

	utilityDefinitionHints["dnsutils"] = `Image has basic DNS investigation tools like nslookup and dig`
	utilityDefinitionHints["psql-curl"] = `Image has a psql client and curl installed`
	utilityDefinitionHints["psqlutils"] = `A debian image with psql client`
	utilityDefinitionHints["awscli"] = `The "latest" tagged aws-cli image from amazon`
	utilityDefinitionHints["gcloudutils"] = `The "latest" tagged cloud-sdk image from google`
	utilityDefinitionHints["azutils"] = `The "latest" tagged azure-cli image from microsoft`
	utilityDefinitionHints["mysqlutils"] = `A debian image with mysql client`
	utilityDefinitionHints["redis"] = `An alpine image with redis-cli installed`
	utilityDefinitionHints["curl"] = `Just the "latest" curl command in a small image`
	utilityDefinitionHints["basic-tools"] = `A small image with curl, jq, yq, and others`

	return []objects.UtilityPod{
		{
			Name:              "dnsutils",
			Image:             "gcr.io/kubernetes-e2e-test-images/dnsutils",
			Tags:              []string{"1.3"},
			ExecCommand:       "/bin/sh",
			ExcludeFromShare:  false,
			Hidden:            false,
			RequireSecrets:    false,
			RequireConfigmaps: false,
			Hint:              utilityDefinitionHints["dnsutils"],
			RemoveUpstream:    false,
			Environments:      []string{},
		},
		{
			Name:              "psql-curl",
			Image:             "barrypiccinni/psql-curl",
			Tags:              []string{"latest"},
			ExecCommand:       "/bin/bash",
			ExcludeFromShare:  false,
			Hidden:            false,
			RequireSecrets:    false,
			RequireConfigmaps: false,
			Hint:              utilityDefinitionHints["psql-curl"],
			RemoveUpstream:    false,
			Environments:      []string{},
		},
		{
			Name:              "psqlutils",
			Image:             "postgres",
			Tags:              []string{"14-bullseye", "15-bullseye"},
			ExecCommand:       "/bin/bash",
			ExcludeFromShare:  false,
			Hidden:            false,
			RequireSecrets:    false,
			RequireConfigmaps: false,
			Hint:              utilityDefinitionHints["psqlutils"],
			RemoveUpstream:    false,
			Environments:      []string{},
		},
		{
			Name:              "awscli",
			Image:             "amazon/aws-cli",
			Tags:              []string{"latest"},
			ExecCommand:       "/bin/bash",
			ExcludeFromShare:  false,
			Hidden:            false,
			RequireSecrets:    false,
			RequireConfigmaps: false,
			Hint:              utilityDefinitionHints["awscli"],
			RemoveUpstream:    false,
			Environments:      []string{},
		},
		{
			Name:              "gcloudutils",
			Image:             "google/cloud-sdk",
			Tags:              []string{"latest"},
			ExecCommand:       "/bin/bash",
			ExcludeFromShare:  false,
			Hidden:            false,
			RequireSecrets:    false,
			RequireConfigmaps: false,
			Hint:              utilityDefinitionHints["gcloudutils"],
			RemoveUpstream:    false,
			Environments:      []string{},
		},
		{
			Name:              "azutils",
			Image:             "mcr.microsoft.com/azure-cli",
			Tags:              []string{"latest"},
			ExecCommand:       "/bin/bash",
			ExcludeFromShare:  false,
			Hidden:            false,
			RequireSecrets:    false,
			RequireConfigmaps: false,
			Hint:              utilityDefinitionHints["azutils"],
			RemoveUpstream:    false,
			Environments:      []string{},
		},
		{
			Name:              "mysqlutils",
			Image:             "mysql",
			Tags:              []string{"5.7.40-debian", "8-debian"},
			ExecCommand:       "/bin/bash",
			ExcludeFromShare:  false,
			Hidden:            false,
			RequireSecrets:    false,
			RequireConfigmaps: false,
			Hint:              utilityDefinitionHints["mysqlutils"],
			RemoveUpstream:    false,
			Environments:      []string{},
		},
		{
			Name:              "redis",
			Image:             "cmaahs/redis-cli",
			Tags:              []string{"6.2"},
			ExecCommand:       "/bin/bash",
			ExcludeFromShare:  false,
			Hidden:            false,
			RequireSecrets:    false,
			RequireConfigmaps: false,
			Hint:              utilityDefinitionHints["redis"],
			RemoveUpstream:    false,
			Environments:      []string{},
		},
		{
			Name:              "curl",
			Image:             "curlimages/curl",
			Tags:              []string{"latest"},
			ExecCommand:       "/bin/sh",
			ExcludeFromShare:  false,
			Hidden:            false,
			RequireSecrets:    false,
			RequireConfigmaps: false,
			Hint:              utilityDefinitionHints["curl"],
			RemoveUpstream:    false,
			Environments:      []string{},
		},
		{
			Name:              "basic-tools",
			Image:             "cmaahs/basic-tools",
			Tags:              []string{"v0.0.1"},
			ExecCommand:       "/bin/bash",
			ExcludeFromShare:  false,
			Hidden:            false,
			RequireSecrets:    false,
			RequireConfigmaps: false,
			Hint:              utilityDefinitionHints["basic-tools"],
			RemoveUpstream:    false,
			Environments:      []string{},
		},
	}

}
