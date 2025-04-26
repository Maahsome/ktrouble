package cmd

import (
	"fmt"
	"os"

	"ktrouble/ask"
	"ktrouble/common"
	"ktrouble/objects"

	v1 "k8s.io/api/core/v1"

	"github.com/muesli/termenv"
)

func standardAttach(utility string, sa string) {

	termFormatter := termenv.NewOutput(os.Stdout)
	if c.Client != nil {
		utilMap := make(map[string]objects.UtilityPod)
		for _, v := range c.UtilDefs {
			utilMap[v.Name] = objects.UtilityPod{
				Name:              v.Name,
				Repository:        v.Repository,
				ExecCommand:       v.ExecCommand,
				RequireSecrets:    v.RequireSecrets,
				RequireConfigmaps: v.RequireConfigmaps,
				Hint:              v.Hint,
			}
		}

		if utility == "" {
			utility = ask.PromptForUtility(c.UtilDefs, c.ShowHidden)
		}

		// Display the HINT
		if len(utilMap[utility].Hint) > 0 {
			fmt.Println(utilMap[utility].Hint)
		}

		utilRepository := utilMap[utility].Repository

		namespace := c.Client.DetermineNamespace(c.Namespace)

		namespacePods := c.Client.GetNamespacePods(namespace)
		if len(namespacePods.Items) == 0 {
			common.Logger.Warn("No pods found in namespace")
			return
		}

		selectedPod := ask.PromptForPod(namespacePods, "Choose a pod to attach to:")
		if selectedPod.Name == "" {
			common.Logger.Warn("No pod selected")
			return
		}

		sleepTime := ask.PromptForEphemeralSleep(c.EphemeralSleepDefs)
		if sleepTime == "" {
			common.Logger.Warn("No sleep time selected")
			return
		}

		// Get a list of mounts in the selected pod
		mounts := c.Client.GetPodMounts(namespace, selectedPod.Name)

		selectedMounts := []v1.VolumeMount{}
		if len(mounts) > 0 {
			selectedMounts = ask.PromptForEphemeralMounts(mounts)
		}
		shortUniq := randSeq(c.UniqIdLength)
		containerName := fmt.Sprintf("%s-%s", utility, shortUniq)
		aerr := c.Client.AttachContainerToPod(namespace, selectedPod.Name, containerName,
			utilRepository, sleepTime, selectedMounts)
		if aerr != nil {
			common.Logger.WithError(aerr).Fatal("Failed to attach container to pod")
		}

		if c.EnableBashLinks {
			hl := fmt.Sprintf("<bash:kubectl -n %s exec -it -c %s %s -- %s>", namespace, containerName, selectedPod.Name, utilMap[utility].ExecCommand)
			tx := fmt.Sprintf("kubectl -n %s exec -it -c %s %s -- %s", namespace, containerName, selectedPod.Name, utilMap[utility].ExecCommand)
			fmt.Println(termFormatter.Hyperlink(hl, tx))
		} else {
			fmt.Printf("kubectl -n %s exec -it -c %s %s -- %s\n", namespace, containerName, selectedPod.Name, utilMap[utility].ExecCommand)
		}
	} else {
		common.Logger.Warn("Cannot launch a pod, no valid kubernetes context")
	}
}
