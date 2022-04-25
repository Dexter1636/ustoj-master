package model

import (
	"ustoj-master/common"

	"k8s.io/client-go/kubernetes"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var logger = common.LogInstance()

type Cluster struct {
	KubeMaster      string
	KubeConfFile    string
	KubeQPS         float32
	KubeBurst       int
	KubeContentType string
}

func (c *Cluster) InitKube(masterUrl string, masterConfigPath string) {

	c.KubeMaster = masterUrl
	c.KubeConfFile = masterConfigPath
	c.KubeQPS = float32(5.000000)
	c.KubeBurst = 10
	c.KubeContentType = "application/vnd.kubernetes.protobuf"

	logger.Infof("kubeMaster = %s, kubeConfigPath = %s", c.KubeMaster, c.KubeConfFile)
}

// KubeConfig from flags
func (c *Cluster) KubeConfig() (conf *rest.Config, err error) {
	kubeConfig, err := clientcmd.BuildConfigFromFlags(c.KubeMaster, c.KubeConfFile)
	if err != nil {
		return nil, err
	}
	kubeConfig.QPS = c.KubeQPS
	kubeConfig.Burst = c.KubeBurst
	kubeConfig.ContentType = c.KubeContentType
	return kubeConfig, err
}

func (c *Cluster) GetNodeClient() (corev1.NodeInterface, error) {
	kubeConfig, err := c.KubeConfig()
	if err != nil {
		logger.Error("Failed to create KubeConfig , error : %v", err)
		return nil, err
	}

	clientSet, err := kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		logger.Error("Failed to create clientset , error : %v", err)
		return nil, err
	}

	return clientSet.CoreV1().Nodes(), nil
}
