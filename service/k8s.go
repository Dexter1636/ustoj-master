package service

import (
	"context"
	"ustoj-master/common"
	"ustoj-master/model"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
)

var logger = common.LogInstance()

var c model.Cluster

func InitCluster(masterUrl string, masterConfigPath string) {
	c.InitKube(masterUrl, masterConfigPath)

}

func ListNode() (*v1.NodeList, error) {

	list, err := c.ListNodes()

	return list, err
}

func ListJob() (*v1.PodList, error) {
	podClient, err := c.GetPodClient("default")
	if err != nil {
		return nil, err
	}

	labelSelector := metav1.LabelSelector{MatchLabels: map[string]string{"ustoj": "job"}}
	list, err := podClient.List(context.Background(), metav1.ListOptions{
		LabelSelector: labels.Set(labelSelector.MatchLabels).String(),
	})

	if err != nil {
		return nil, err
	}

	return list, nil
}

func ListRunningJob() (*v1.PodList, error) {
	podClient, err := c.GetPodClient("default")
	if err != nil {
		return nil, err
	}

	labelSelector := metav1.LabelSelector{MatchLabels: map[string]string{"ustoj": "job"}}
	fieldSelector := fields.SelectorFromSet(fields.Set{"status.phase": string(v1.PodRunning)})
	list, err := podClient.List(context.Background(), metav1.ListOptions{
		TypeMeta:      metav1.TypeMeta{},
		LabelSelector: labels.Set(labelSelector.MatchLabels).String(),
		FieldSelector: fieldSelector.String(),
	})

	if err != nil {
		return nil, err
	}

	return list, nil
}
