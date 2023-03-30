package monitoring

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/openshift/hypershift/support/metrics"
	prometheusoperatorv1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
)

func ReconcileKubeAPIServerServiceMonitor(serviceMonitor *prometheusoperatorv1.ServiceMonitor) error {
	serviceMonitor.Spec.NamespaceSelector = prometheusoperatorv1.NamespaceSelector{
		MatchNames: []string{"default"},
	}
	serviceMonitor.Spec.Selector = metav1.LabelSelector{
		MatchLabels: map[string]string{
			"component": "apiserver",
		},
	}
	serviceMonitor.Spec.Endpoints = []prometheusoperatorv1.Endpoint{
		{
			BearerTokenFile: "/var/run/secrets/kubernetes.io/serviceaccount/token",
			TLSConfig: &prometheusoperatorv1.TLSConfig{
				SafeTLSConfig: prometheusoperatorv1.SafeTLSConfig{
					ServerName: "kubernetes.default.svc",
				},
				CAFile: "/var/run/secrets/kubernetes.io/serviceaccount/ca.crt",
			},
			Scheme:               "https",
			Port:                 "https",
			Path:                 "/metrics",
			MetricRelabelConfigs: metrics.KASRelabelConfigs(metrics.MetricsSetAll),
		},
	}
	return nil
}
