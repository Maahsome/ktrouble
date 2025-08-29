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

	selectedUtility := objects.SelectedUtilityPod{}
	termFormatter := termenv.NewOutput(os.Stdout)
	if c.Client != nil {
		utilMap := objects.GetUtilityMap(c.UtilDefs, c.EnvMap)

		if utility == "" {
			utility, selectedUtility = ask.PromptForUtility(c.UtilDefs, c.EnvMap, c.ShowHidden)
		}

		// Display the HINT
		if len(utilMap[utility].Hint) > 0 {
			fmt.Println(utilMap[utility].Hint)
		}

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
		containerName := fmt.Sprintf("%s-%s", utilMap[utility].Name, shortUniq)
		common.Logger.Tracef("Utility selected: %#v", selectedUtility)
		image := fmt.Sprintf("%s:%s", utilMap[utility].Image, utilMap[utility].Tags[0])
		if p.BuildCommand {
			fmt.Printf("TODO: parameterize the attach command\n")
			fmt.Printf("ktrouble attach \n")
		} else {
			aerr := c.Client.AttachContainerToPod(namespace, selectedPod.Name, containerName,
				image, sleepTime, selectedMounts)
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
		}
	} else {
		common.Logger.Warn("Cannot launch a pod, no valid kubernetes context")
	}
}
