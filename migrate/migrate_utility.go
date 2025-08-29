package migrate

import (
	"ktrouble/objects"
	"strings"
)

func UpdateUtilityV2(util objects.UtilityPodV1) objects.UtilityPod {
	if strings.Contains(util.Repository, ":") {
		// Split the repository into image and tags
		imageAndTag := strings.Split(util.Repository, ":")
		return objects.UtilityPod{
			Name:              util.Name,
			Image:             imageAndTag[0],
			Tags:              []string{imageAndTag[1]},
			ExecCommand:       util.ExecCommand,
			RequireSecrets:    util.RequireSecrets,
			RequireConfigmaps: util.RequireConfigmaps,
			Hint:              util.Hint,
			Environments:      util.Environments,
			RemoveUpstream:    false,
		}
	} else {
		return objects.UtilityPod{
			Name:              util.Name,
			Image:             util.Repository,
			Tags:              []string{"latest"},
			ExecCommand:       util.ExecCommand,
			RequireSecrets:    util.RequireSecrets,
			RequireConfigmaps: util.RequireConfigmaps,
			Hint:              util.Hint,
			Environments:      util.Environments,
			RemoveUpstream:    false,
		}
	}
}
