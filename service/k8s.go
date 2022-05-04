package service

import (
	"context"
	"strconv"
	"ustoj-master/common"
	"ustoj-master/dto"
	"ustoj-master/model"

	"github.com/google/uuid"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	batchv1 "k8s.io/client-go/applyconfigurations/batch/v1"
	appcorev1 "k8s.io/client-go/applyconfigurations/core/v1"
	applyconfv1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

var logger = common.LogInstance()

var c model.Cluster

var PodJobConvMap = map[corev1.PodPhase]model.SubmitJobStatus{
	corev1.PodPending:   model.JobPending,
	corev1.PodRunning:   model.JobRunning,
	corev1.PodSucceeded: model.JobSuccess,
	corev1.PodFailed:    model.JobFailed,
	corev1.PodUnknown:   model.JobUnknown,
}

func InitCluster(masterUrl string, masterConfigPath string) error {
	c.InitKube(masterUrl, masterConfigPath)

	return nil
}

func ListNode() (*corev1.NodeList, error) {

	list, err := c.ListNodes()

	return list, err
}

func ListJobById(submitId []int) ([]dto.SubmissionDto, error) {
	podClient, err := c.GetPodClient("default")
	if err != nil {
		return nil, err
	}
	submitIdStr := []string{}
	for _, id := range submitId {
		submitIdStr = append(submitIdStr, strconv.Itoa(id))
	}

	labelSelector := metav1.LabelSelector{
		MatchLabels: map[string]string{
			"ustoj": "job",
		},
		MatchExpressions: []metav1.LabelSelectorRequirement{
			{
				Key:      "submit_id",
				Operator: metav1.LabelSelectorOpIn,
				Values:   submitIdStr,
			},
		},
	}
	labelMap, err := metav1.LabelSelectorAsMap(&labelSelector)
	if err != nil {
		return nil, err
	}
	list, err := podClient.List(context.Background(), metav1.ListOptions{
		LabelSelector: labels.SelectorFromSet(labelMap).String(),
	})

	if err != nil {
		return nil, err
	}

	var result []dto.SubmissionDto
	for _, pod := range list.Items {
		logger.Infoln("!!!!!")
		logger.Infoln(pod.Labels)
		logger.Infoln("!!!!!")
		sub_id, err := strconv.Atoi(pod.Labels["submit_id"])
		if err != nil {
			logger.Errorln(sub_id)
			return nil, err
		}
		dto := dto.SubmissionDto{
			SubmissionID: sub_id,
			Status:       PodJobConvMap[pod.Status.Phase],
		}
		result = append(result, dto)
	}
	return result, nil
}

func ListJob() ([]dto.SubmissionDto, error) {
	podClient, err := c.GetPodClient("default")
	if err != nil {
		return nil, err
	}

	labelSelector := metav1.LabelSelector{
		MatchLabels: map[string]string{"ustoj": "job"},
	}
	list, err := podClient.List(context.Background(), metav1.ListOptions{
		LabelSelector: labels.Set(labelSelector.MatchLabels).String(),
	})

	if err != nil {
		return nil, err
	}

	var result []dto.SubmissionDto
	for _, pod := range list.Items {
		sub_id, err := strconv.Atoi(pod.Labels["submit_id"])
		if err != nil {
			logger.Errorln(sub_id)
			return nil, err
		}
		dto := dto.SubmissionDto{
			SubmissionID: sub_id,
			Status:       PodJobConvMap[pod.Status.Phase],
		}
		result = append(result, dto)
	}
	return result, nil
}

func ListRunningJob() ([]dto.SubmissionDto, error) {
	podClient, err := c.GetPodClient("default")
	if err != nil {
		return nil, err
	}

	labelSelector := metav1.LabelSelector{MatchLabels: map[string]string{"ustoj": "job"}}
	fieldSelector := fields.SelectorFromSet(fields.Set{"status.phase": string(corev1.PodRunning)})
	list, err := podClient.List(context.Background(), metav1.ListOptions{
		LabelSelector: labels.Set(labelSelector.MatchLabels).String(),
		FieldSelector: fieldSelector.String(),
	})
	if err != nil {
		return nil, err
	}

	var result []dto.SubmissionDto
	for _, pod := range list.Items {
		sub_id, err := strconv.Atoi(pod.Labels["submit_id"])
		if err != nil {
			logger.Errorln(sub_id)
			return nil, err
		}
		dto := dto.SubmissionDto{
			SubmissionID: sub_id,
			Status:       PodJobConvMap[pod.Status.Phase],
		}
		result = append(result, dto)
	}

	return result, nil
}

func CreateJob(submitId int, caseList []string, language string) error {
	submitIdStr := strconv.Itoa(submitId)
	podClient, err := c.GetPodClient("default")
	if err != nil {
		return err
	}
	var kind = new(string)
	var apiVer = new(string)
	var podName = new(string)
	var imageName = new(string)
	var pullPolicy = new(corev1.PullPolicy)
	var restartPolicy = new(corev1.RestartPolicy)

	*kind = "Pod"
	*apiVer = "v1"
	*podName = "job-sumbission-" + submitIdStr
	*imageName = "python"
	*pullPolicy = corev1.PullIfNotPresent
	*restartPolicy = corev1.RestartPolicyOnFailure
	caseArrayStr := "("
	for _, _case := range caseList {
		caseArrayStr += _case + " "
	}
	caseArrayStr += ")"

	container := appcorev1.ContainerApplyConfiguration{
		Name:    podName,
		Image:   imageName,
		Command: []string{"/bin/bash", "-c", "--"},
		Args: []string{
			`
			array=` + caseArrayStr + `
			for element in ${array[@]}
			do
			python run.py $element
			done
			`,
			"echo 1",
		},
		// WorkingDir:             new(string),
		Env:          []appcorev1.EnvVarApplyConfiguration{},
		VolumeMounts: []appcorev1.VolumeMountApplyConfiguration{},
		// LivenessProbe:          &corev1.ProbeApplyConfiguration{},
		// StartupProbe:           &corev1.ProbeApplyConfiguration{},
		ImagePullPolicy: pullPolicy,
	}
	_, err = podClient.Apply(context.Background(), &appcorev1.PodApplyConfiguration{
		TypeMetaApplyConfiguration: applyconfv1.TypeMetaApplyConfiguration{
			Kind:       kind,
			APIVersion: apiVer,
		},
		ObjectMetaApplyConfiguration: &applyconfv1.ObjectMetaApplyConfiguration{
			Name: podName,
			// Namespace:                  new(string),
			// ResourceVersion:            new(string),
			Labels: map[string]string{
				"ustoj":     "job",
				"submit_id": submitIdStr,
			},
		},
		Spec: &appcorev1.PodSpecApplyConfiguration{
			// Volumes:                       []corev1.VolumeApplyConfiguration{},
			Containers: []appcorev1.ContainerApplyConfiguration{container},
			Volumes: []appcorev1.VolumeApplyConfiguration{
				{
					Name: podName,
					VolumeSourceApplyConfiguration: appcorev1.VolumeSourceApplyConfiguration{
						HostPath:              &appcorev1.HostPathVolumeSourceApplyConfiguration{},
						EmptyDir:              &appcorev1.EmptyDirVolumeSourceApplyConfiguration{},
						GCEPersistentDisk:     &appcorev1.GCEPersistentDiskVolumeSourceApplyConfiguration{},
						AWSElasticBlockStore:  &appcorev1.AWSElasticBlockStoreVolumeSourceApplyConfiguration{},
						GitRepo:               &appcorev1.GitRepoVolumeSourceApplyConfiguration{},
						Secret:                &appcorev1.SecretVolumeSourceApplyConfiguration{},
						NFS:                   &appcorev1.NFSVolumeSourceApplyConfiguration{},
						ISCSI:                 &appcorev1.ISCSIVolumeSourceApplyConfiguration{},
						Glusterfs:             &appcorev1.GlusterfsVolumeSourceApplyConfiguration{},
						PersistentVolumeClaim: &appcorev1.PersistentVolumeClaimVolumeSourceApplyConfiguration{},
						RBD:                   &appcorev1.RBDVolumeSourceApplyConfiguration{},
						FlexVolume:            &appcorev1.FlexVolumeSourceApplyConfiguration{},
						Cinder:                &appcorev1.CinderVolumeSourceApplyConfiguration{},
						CephFS:                &appcorev1.CephFSVolumeSourceApplyConfiguration{},
						Flocker:               &appcorev1.FlockerVolumeSourceApplyConfiguration{},
						DownwardAPI:           &appcorev1.DownwardAPIVolumeSourceApplyConfiguration{},
						FC:                    &appcorev1.FCVolumeSourceApplyConfiguration{},
						AzureFile:             &appcorev1.AzureFileVolumeSourceApplyConfiguration{},
						ConfigMap:             &appcorev1.ConfigMapVolumeSourceApplyConfiguration{},
						VsphereVolume:         &appcorev1.VsphereVirtualDiskVolumeSourceApplyConfiguration{},
						Quobyte:               &appcorev1.QuobyteVolumeSourceApplyConfiguration{},
						AzureDisk:             &appcorev1.AzureDiskVolumeSourceApplyConfiguration{},
						PhotonPersistentDisk:  &appcorev1.PhotonPersistentDiskVolumeSourceApplyConfiguration{},
						Projected:             &appcorev1.ProjectedVolumeSourceApplyConfiguration{},
						PortworxVolume:        &appcorev1.PortworxVolumeSourceApplyConfiguration{},
						ScaleIO:               &appcorev1.ScaleIOVolumeSourceApplyConfiguration{},
						StorageOS:             &appcorev1.StorageOSVolumeSourceApplyConfiguration{},
						CSI:                   &appcorev1.CSIVolumeSourceApplyConfiguration{},
						Ephemeral:             &appcorev1.EphemeralVolumeSourceApplyConfiguration{},
					},
				},
			},
			RestartPolicy: restartPolicy,
			NodeSelector: map[string]string{
				"ustoj": "worker",
			},
		},
	}, metav1.ApplyOptions{
		FieldManager: "ustoj",
	})
	if err != nil {
		return err
	}

	return nil
}

// === the following methods are still in test

func TestCreateJob(submitId int, caseList []string, language string) error {
	submitIdStr := strconv.Itoa(submitId)
	jobClent, err := c.GetJobClient("default")
	if err != nil {
		return err
	}

	var kind = new(string)
	var apiVer = new(string)
	var jobName = new(string)
	var podName = new(string)
	var imageName = new(string)
	var pullPolicy = new(corev1.PullPolicy)
	var backoffLimit = new(int32)
	var restartPolicy = new(corev1.RestartPolicy)

	*kind = "Job"
	*apiVer = "batch/v1"
	*jobName = "job-sumbission-" + submitIdStr
	uuid := uuid.New()
	key := uuid.String()
	*podName = *jobName + key
	*imageName = "debian"
	*pullPolicy = corev1.PullIfNotPresent
	*restartPolicy = corev1.RestartPolicyOnFailure
	caseArrayStr := "("
	for _, _case := range caseList {
		caseArrayStr += _case + " "
	}
	caseArrayStr += ")"
	*backoffLimit = 10

	container := appcorev1.ContainerApplyConfiguration{
		Name:    podName,
		Image:   imageName,
		Command: []string{"/bin/bash", "-c", "--"},
		Args: []string{
			// `
			// array=` + caseArrayStr + `
			// for element in ${array[@]}
			// do
			// echo $element
			// done
			// `,
			"echo 1",
		},
		// WorkingDir:             new(string),
		// VolumeMounts:           []corev1.VolumeMountApplyConfiguration{},
		// LivenessProbe:          &corev1.ProbeApplyConfiguration{},
		// StartupProbe:           &corev1.ProbeApplyConfiguration{},
		ImagePullPolicy: pullPolicy,
	}
	_, err = jobClent.Apply(context.Background(),
		&batchv1.JobApplyConfiguration{
			TypeMetaApplyConfiguration: applyconfv1.TypeMetaApplyConfiguration{
				Kind:       kind,
				APIVersion: apiVer,
			},
			ObjectMetaApplyConfiguration: &applyconfv1.ObjectMetaApplyConfiguration{
				Name: jobName,
				Labels: map[string]string{
					"ustoj":     "job",
					"submit_id": submitIdStr,
				},
			},
			Spec: &batchv1.JobSpecApplyConfiguration{
				BackoffLimit: backoffLimit,
				Template: &appcorev1.PodTemplateSpecApplyConfiguration{
					Spec: &appcorev1.PodSpecApplyConfiguration{
						Containers: []appcorev1.ContainerApplyConfiguration{container},
						NodeSelector: map[string]string{
							"ustoj": "worker",
						},
						RestartPolicy: restartPolicy,
					},
				},
			},
		},
		metav1.ApplyOptions{
			FieldManager: "kubectl",
		})

	if err != nil {
		logger.Errorln(err)
		return err
	}

	return nil
}

func TestListJob() ([]dto.SubmissionDto, error) {
	var result []dto.SubmissionDto
	jobClent, err := c.GetJobClient("default")
	if err != nil {
		return nil, err
	}

	labelSelector := metav1.LabelSelector{MatchLabels: map[string]string{"ustoj": "job"}}
	list, err := jobClent.List(context.Background(), metav1.ListOptions{
		LabelSelector: labels.Set(labelSelector.MatchLabels).String(),
	})

	if err != nil {
		return nil, err
	}

	for _, job := range list.Items {
		sub_id, err := strconv.Atoi(job.Labels["submit_id"])
		if err != nil {
			logger.Errorln(sub_id)
			return nil, err
		}
		dto := dto.SubmissionDto{
			SubmissionID: sub_id,
			// Status:       PodJobConvMap[job.],
		}
		logger.Infoln("!!!")
		// logger.Infoln(pod.Status.Phase)
		logger.Infoln("!!!")
		result = append(result, dto)
	}
	return result, nil
}
