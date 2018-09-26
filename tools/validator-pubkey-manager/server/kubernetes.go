package main

import (
	"flag"
	"os"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8s "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	outOfCluster  = flag.Bool("out-of-cluster", false, "Whether or not the service should run with an out of cluster config")
	namespace     = "default"
	configMapName = "validator-pubkey-config"
)

type kubernetes struct {
	clientset *k8s.Clientset
}

func newKubernetesStorage() *kubernetes {
	var config *rest.Config
	var err error
	if *outOfCluster {
		kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
	} else {
		config, err = rest.InClusterConfig()
	}

	if err != nil {
		panic(err.Error())
	}

	return &kubernetes{
		clientset: k8s.NewForConfigOrDie(config),
	}
}

func (k *kubernetes) PubkeyMap() (map[string][]byte, error) {
	cmap, err := k.clientset.CoreV1().ConfigMaps(namespace).Get(configMapName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return cmap.BinaryData, nil
}

func (k *kubernetes) SetPubkey(pod string, pkey []byte) error {
	cmap, err := k.clientset.CoreV1().ConfigMaps(namespace).Get(configMapName, metav1.GetOptions{})
	if err != nil {
		return err
	}

	cmap.BinaryData[pod] = pkey

	if _, err := k.clientset.CoreV1().ConfigMaps(namespace).Update(cmap); err != nil {
		return err
	}

	return nil
}

func (k *kubernetes) RemovePod(pod string) error {
	cmap, err := k.clientset.CoreV1().ConfigMaps(namespace).Get(configMapName, metav1.GetOptions{})
	if err != nil {
		return err
	}

	delete(cmap.BinaryData, pod)

	if _, err := k.clientset.CoreV1().ConfigMaps(namespace).Update(cmap); err != nil {
		return err
	}

	return nil
}
