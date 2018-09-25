package main

type kubernetes struct {
	// clientset *kubernetes.ClientSet
}

func newKubernetesStorage() *kubernetes {
	return &kubernetes{}
}

func (k *kubernetes) GetConfigMap() (map[string][]byte, error) {
	m := make(map[string][]byte)
	return m, nil
}

func (k *kubernetes) SetPubkey(pod string, pkey []byte) error {
	return nil
}

func (k *kubernetes) RemovePod(pod string) error {
	return nil
}
