package kubernetes

import (
	"fmt"
	"os"

	"ktrouble/ask"
	"ktrouble/common"

	homedir "github.com/mitchellh/go-homedir"
	v1 "k8s.io/api/core/v1"
	kofficial "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type KubernetesClient interface {
	// GetProperty(property string) string
	// SetProperty(property string, value string) string
	CreatePod(podJSON string, namespace string)
	GetCreatedPods() *v1.PodList
	GetNamespaces() *v1.NamespaceList
	GetNodes() *v1.NodeList
	GetServiceAccounts(namespace string) *v1.ServiceAccountList
	GetSecrets(namespace string) *v1.SecretList
	DeletePod(pod ask.PodDetail) error
	DetermineNamespace(nsParam string) string
}

type kubernetesClient struct {
	Client *kofficial.Clientset
}

// New generate a new kubernetes client
func New() KubernetesClient {

	cfg, err := restConfig()
	if err != nil {
		common.Logger.WithError(err).Debug("could not get config")
		return nil
	}
	if cfg == nil {
		common.Logger.Debug("failed to determine kubernetes config")
		return nil
	}

	client, err := kofficial.NewForConfig(cfg)
	if err != nil {
		common.Logger.WithError(err).Debug("could not create client from config")
		return nil
	}

	return &kubernetesClient{
		Client: client,
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
			common.Logger.Info("No KUBECONFIG ENV")
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
					common.Logger.Info("Could not locate the KUBECONFIG file, normally ~/.kube/config")
					os.Exit(1)
				}
				return nil, nil
			}
		}
		kubeCfg, err = clientcmd.BuildConfigFromFlags("", kubeFile)
		if err != nil {
			common.Logger.WithError(err).Error("Failed to build KUBECONFIG")
		}
	}
	return kubeCfg, nil
}
