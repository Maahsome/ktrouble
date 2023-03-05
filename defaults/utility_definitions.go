package defaults

import (
	"ktrouble/objects"
)

func UtilityDefinitions() []objects.UtilityPod {

	return []objects.UtilityPod{
		{
			Name:        "dnsutils",
			Repository:  "gcr.io/kubernetes-e2e-test-images/dnsutils:1.3",
			ExecCommand: "/bin/sh",
		},
		{
			Name:        "psql-curl",
			Repository:  "barrypiccinni/psql-curl:latest",
			ExecCommand: "/bin/bash",
		},
		{
			Name:        "psqlutils15",
			Repository:  "postgres:15-bullseye",
			ExecCommand: "/bin/bash",
		},
		{
			Name:        "psqlutils14",
			Repository:  "postgres:14-bullseye",
			ExecCommand: "/bin/bash",
		},
		{
			Name:        "awscli",
			Repository:  "amazon/aws-cli:latest",
			ExecCommand: "/bin/bash",
		},
		{
			Name:        "gcloudutils",
			Repository:  "google/cloud-sdk:latest",
			ExecCommand: "/bin/bash",
		},
		{
			Name:        "azutils",
			Repository:  "mcr.microsoft.com/azure-cli",
			ExecCommand: "/bin/bash",
		},
		{
			Name:        "mysqlutils5",
			Repository:  "mysql:5.7.40-debian",
			ExecCommand: "/bin/bash",
		},
		{
			Name:        "mysqlutils8",
			Repository:  "mysql:8-debian",
			ExecCommand: "/bin/bash",
		},
		{
			Name:        "redis6",
			Repository:  "cmaahs/redis-cli:6.2",
			ExecCommand: "/bin/bash",
		},
		{
			Name:        "curl",
			Repository:  "curlimages/curl:latest",
			ExecCommand: "/bin/sh",
		},
		{
			Name:        "basic-tools",
			Repository:  "cmaahs/basic-tools:v0.0.1",
			ExecCommand: "/bin/bash",
		},
	}

}
