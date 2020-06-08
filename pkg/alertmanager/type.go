package alertmanager

// {
//     "version": "4",
//     "groupKey": "test",
//     "status": "firing",
//     "receiver": "test",
//     "groupLabels": {
//         "alertname": "test"
//     },
//     "commonLabels": {
//         "alertname": "test"
//     },
//     "commonAnnotations": {
//         "alertname": "test",
//         "summary": "Summary Title name"
//     },
//     "externalURL": "test",
//     "alerts": [
//         {
//             "labels": {
//                 "alertname": "KubeStatefulSetReplicasMismatch2",
//                 "endpoint": "http",
//                 "instance": "10.112.130.213:8080",
//                 "job": "kube-state-metrics",
//                 "namespace": "elk",
//                 "pod": "prometheus-operator-kube-state-metrics-7f979567df-lx9zw",
//                 "prometheus": "prometheus/prometheus-operator-prometheus",
//                 "service": "prometheus-operator-kube-state-metrics",
//                 "severity": "critical",
//                 "statefulset": "elk-ingest"
//             },
//             "annotations": {
//                 "description": "StatefulSet elk/elk-ingest has not matched the expected number of replicas for longer than 15 minutes.",
//                 "runbook_url": "https://github.com/kubernetes-monitoring/kubernetes-mixin/tree/master/runbook.md#alert-name-kubestatefulsetreplicasmismatch"
//             },
//             "startsAt": "2020-05-31T01:25:01.044Z",
//             "endsAt": "2020-05-31T12:01:01.044Z",
//             "generatorURL": "https://prometheus.localhost:9090/graph?g0.expr=%28kube_statefulset_status_replicas_ready%7Bjob%3D%22kube-state-metrics%22%2Cnamespace%3D~%22.%2A%22%7D+%21%3D+kube_statefulset_status_replicas%7Bjob%3D%22kube-state-metrics%22%2Cnamespace%3D~%22.%2A%22%7D%29+and+%28changes%28kube_statefulset_status_replicas_updated%7Bjob%3D%22kube-state-metrics%22%2Cnamespace%3D~%22.%2A%22%7D%5B5m%5D%29+%3D%3D+0%29\u0026g0.tab=1",
//             "status": "firing",
//             "receivers": [
//                 "null"
//             ],
//             "fingerprint": "03ae847ecccdcce4"
//         }
//     ]
// }

// AlertManOut json structure from Alertmanager
type AlertManOut struct {
	Version     string `json:"version,omitempty"`
	GroupKey    string `json:"groupKey,omitempty"`
	Status      string `json:"status,omitempty"`
	Receiver    string `json:"receiver,omitempty"`
	GroupLabels struct {
		Alertname string `json:"alertname,omitempty"`
	} `json:"groupLabels,omitempty"`
	CommonLabels struct {
		Alertname string `json:"alertname,omitempty"`
	} `json:"commonLabels,omitempty"`
	CommonAnnotations struct {
		Summary string `json:"summary,omitempty"`
	} `json:"commonAnnotations,omitempty"`
	ExternalURL string           `json:"externalURL,omitempty"`
	Alerts      []*AlertManAlert `json:"alerts,omitempty"`
}

// AlertManAlert json structure from Alertmanager
type AlertManAlert struct {
	Annotations struct {
		Description string `json:"description,omitempty"`
		Summary     string `json:"summary,omitempty"`
	} `json:"annotations,omitempty"`
	Status       string            `json:"status,omitempty"`
	Labels       map[string]string `json:"labels,omitempty"`
	GeneratorURL string            `json:"generatorURL,omitempty"`
	StartsAt     string            `json:"startsAt,omitempty"`
	EndsAt       string            `json:"endsAt,omitempty"`
}
