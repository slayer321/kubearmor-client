package log

import (
	"context"
	"fmt"

	"github.com/kubearmor/kubearmor-client/k8s"
	"github.com/rs/zerolog/log"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func SelectLabel(o Options) map[string]string {
	var nameMap map[string]string

	client, err := k8s.ConnectK8sClient()
	if err != nil {
		log.Error().Msgf("unable to create Kubernetes clients: %s", err.Error())
		return nil
	}

	fmt.Printf("Selector passed by user: %s\n", o.Selector)
	nameMap = make(map[string]string)

	if len(o.Selector) <= 1 {

		strVal := string(o.Selector[0])

		listPod, err := client.K8sClientset.CoreV1().Pods(o.Namespace).List(context.Background(), metav1.ListOptions{LabelSelector: strVal})
		if err != nil {
			log.Error().Msgf("unable to list pods: %s", err.Error())
			return nil
		}

		for _, i := range listPod.Items {
			nameMap[i.Name] = i.Name

		}

	} else if len(o.Selector) >= 2 {

		for _, val := range o.Selector {

			listPod, err := client.K8sClientset.CoreV1().Pods(o.Namespace).List(context.Background(), metav1.ListOptions{LabelSelector: val})
			if err != nil {
				log.Error().Msgf("unable to list pods: %s", err.Error())
				return nil
			}
			for _, i := range listPod.Items {
				nameMap[i.Name] = i.Name

			}

		}

	}
	return nameMap
}
