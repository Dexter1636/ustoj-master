package service

import (
	"context"
	"ustoj-master/common"
	"ustoj-master/model"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	corev1 "k8s.io/client-go/applyconfigurations/core/v1"
	applyconfv1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

var logger = common.LogInstance()

var c model.Cluster

func InitCluster(masterUrl string, masterConfigPath string) error {
	c.InitKube(masterUrl, masterConfigPath)

	err := CreateJob()
	if err != nil {
		logger.Errorln(err)
	}
	return nil
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
		LabelSelector: labels.Set(labelSelector.MatchLabels).String(),
		FieldSelector: fieldSelector.String(),
	})

	if err != nil {
		return nil, err
	}

	return list, nil
}

func CreateJob() error {
	podClient, err := c.GetPodClient("default")
	if err != nil {
		return err
	}

	var kind = new(string)
	*kind = "Pod"
	var apiVer = new(string)
	*apiVer = "v1"
	var podName = new(string)
	*podName = "job-test"
	var imageName = new(string)
	*imageName = "debian"
	container := corev1.ContainerApplyConfiguration{
		Name:    podName,
		Image:   imageName,
		Command: []string{"/bin/bash", "-c", "--"},
		Args:    []string{"while true; do echo 1;sleep 3600;done;"},
		// WorkingDir:             new(string),
		// Env:                    []corev1.EnvVarApplyConfiguration{},
		// VolumeMounts:           []corev1.VolumeMountApplyConfiguration{},
		// LivenessProbe:          &corev1.ProbeApplyConfiguration{},
		// StartupProbe:           &corev1.ProbeApplyConfiguration{},
		// ImagePullPolicy:          &"",
	}
	_, err = podClient.Apply(context.Background(), &corev1.PodApplyConfiguration{
		TypeMetaApplyConfiguration: applyconfv1.TypeMetaApplyConfiguration{
			Kind:       kind,
			APIVersion: apiVer,
		},
		ObjectMetaApplyConfiguration: &applyconfv1.ObjectMetaApplyConfiguration{
			Name: podName,
			// Namespace:                  new(string),
			// ResourceVersion:            new(string),
			Labels: map[string]string{"ustoj": "job"},
		},
		Spec: &corev1.PodSpecApplyConfiguration{
			// Volumes:                       []corev1.VolumeApplyConfiguration{},
			Containers: []corev1.ContainerApplyConfiguration{container},
			// RestartPolicy:       &"",
			NodeSelector: map[string]string{},
		},
	}, metav1.ApplyOptions{
		TypeMeta:     metav1.TypeMeta{},
		DryRun:       []string{},
		Force:        false,
		FieldManager: "ustoj",
	})

	if err != nil {
		return err
	}

	return nil
}