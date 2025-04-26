package defaults

import (
	"ktrouble/objects"
)

func UtilityDefinitions() []objects.UtilityPod {
	utilityDefinitionHints := map[string]string{}

	utilityDefinitionHints["dnsutils"] = `Image has basic DNS investigation tools like nslookup and dig`
	utilityDefinitionHints["psql-curl"] = `Image has a psql client and curl installed`
	utilityDefinitionHints["psqlutils15"] = `A debian image with a version 15 psql client`
	utilityDefinitionHints["psqlutils14"] = `A debian image with a version 14 psql client`
	utilityDefinitionHints["awscli"] = `The "latest" tagged aws-cli image from amazon`
	utilityDefinitionHints["gcloudutils"] = `The "latest" tagged cloud-sdk image from google`
	utilityDefinitionHints["azutils"] = `The "latest" tagged azure-cli image from microsoft`
	utilityDefinitionHints["mysqlutils5"] = `A debian image with mysql 5.7.40 versioned client`
	utilityDefinitionHints["mysqlutils8"] = `A debian image with mysql 8.n.n versioned client`
	utilityDefinitionHints["redis6"] = `An alpine image with redis-cli v6.2 installed`
	utilityDefinitionHints["curl"] = `Just the "latest" curl command in a small image`
	utilityDefinitionHints["basic-tools"] = `A small image with curl, jq, yq, and others`

	return []objects.UtilityPod{
		{
			Name:              "dnsutils",
			Repository:        "gcr.io/kubernetes-e2e-test-images/dnsutils:1.3",
			ExecCommand:       "/bin/sh",
			Source:            "ktrouble-utils",
			ExcludeFromShare:  false,
			Hidden:            false,
			RequireSecrets:    false,
			RequireConfigmaps: false,
			Hint:              utilityDefinitionHints["dnsutils"],
		},
		{
			Name:              "psql-curl",
			Repository:        "barrypiccinni/psql-curl:latest",
			ExecCommand:       "/bin/bash",
			Source:            "ktrouble-utils",
			ExcludeFromShare:  false,
			Hidden:            false,
			RequireSecrets:    false,
			RequireConfigmaps: false,
			Hint:              utilityDefinitionHints["psql-curl"],
		},
		{
			Name:              "psqlutils15",
			Repository:        "postgres:15-bullseye",
			ExecCommand:       "/bin/bash",
			Source:            "ktrouble-utils",
			ExcludeFromShare:  false,
			Hidden:            false,
			RequireSecrets:    false,
			RequireConfigmaps: false,
			Hint:              utilityDefinitionHints["psqlutils15"],
		},
		{
			Name:              "psqlutils14",
			Repository:        "postgres:14-bullseye",
			ExecCommand:       "/bin/bash",
			Source:            "ktrouble-utils",
			ExcludeFromShare:  false,
			Hidden:            false,
			RequireSecrets:    false,
			RequireConfigmaps: false,
			Hint:              utilityDefinitionHints["psqlutils14"],
		},
		{
			Name:              "awscli",
			Repository:        "amazon/aws-cli:latest",
			ExecCommand:       "/bin/bash",
			Source:            "ktrouble-utils",
			ExcludeFromShare:  false,
			Hidden:            false,
			RequireSecrets:    false,
			RequireConfigmaps: false,
			Hint:              utilityDefinitionHints["awscli"],
		},
		{
			Name:              "gcloudutils",
			Repository:        "google/cloud-sdk:latest",
			ExecCommand:       "/bin/bash",
			Source:            "ktrouble-utils",
			ExcludeFromShare:  false,
			Hidden:            false,
			RequireSecrets:    false,
			RequireConfigmaps: false,
			Hint:              utilityDefinitionHints["gcloudutils"],
		},
		{
			Name:              "azutils",
			Repository:        "mcr.microsoft.com/azure-cli",
			ExecCommand:       "/bin/bash",
			Source:            "ktrouble-utils",
			ExcludeFromShare:  false,
			Hidden:            false,
			RequireSecrets:    false,
			RequireConfigmaps: false,
			Hint:              utilityDefinitionHints["azutils"],
		},
		{
			Name:              "mysqlutils5",
			Repository:        "mysql:5.7.40-debian",
			ExecCommand:       "/bin/bash",
			Source:            "ktrouble-utils",
			ExcludeFromShare:  false,
			Hidden:            false,
			RequireSecrets:    false,
			RequireConfigmaps: false,
			Hint:              utilityDefinitionHints["mysqlutils5"],
		},
		{
			Name:              "mysqlutils8",
			Repository:        "mysql:8-debian",
			ExecCommand:       "/bin/bash",
			Source:            "ktrouble-utils",
			ExcludeFromShare:  false,
			Hidden:            false,
			RequireSecrets:    false,
			RequireConfigmaps: false,
			Hint:              utilityDefinitionHints["mysqlutils8"],
		},
		{
			Name:              "redis6",
			Repository:        "cmaahs/redis-cli:6.2",
			ExecCommand:       "/bin/bash",
			Source:            "ktrouble-utils",
			ExcludeFromShare:  false,
			Hidden:            false,
			RequireSecrets:    false,
			RequireConfigmaps: false,
			Hint:              utilityDefinitionHints["redis6"],
		},
		{
			Name:              "curl",
			Repository:        "curlimages/curl:latest",
			ExecCommand:       "/bin/sh",
			Source:            "ktrouble-utils",
			ExcludeFromShare:  false,
			Hidden:            false,
			RequireSecrets:    false,
			RequireConfigmaps: false,
			Hint:              utilityDefinitionHints["curl"],
		},
		{
			Name:              "basic-tools",
			Repository:        "cmaahs/basic-tools:v0.0.1",
			ExecCommand:       "/bin/bash",
			Source:            "ktrouble-utils",
			ExcludeFromShare:  false,
			Hidden:            false,
			RequireSecrets:    false,
			RequireConfigmaps: false,
			Hint:              utilityDefinitionHints["basic-tools"],
		},
	}

}
