module github.com/ibrokethecloud/kube-bench-metrics

require (
	github.com/aquasecurity/kube-bench v0.5.0
	github.com/googleapis/gnostic v0.3.1 // indirect
	github.com/prometheus/client_golang v1.11.1
	github.com/sirupsen/logrus v1.6.0
	github.com/urfave/cli v1.22.1
	// Manually added
	k8s.io/api v0.0.0-20190409021203-6e4e0e4f393b
	k8s.io/apimachinery v0.0.0-20190927035529-0104e33c351d
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/klog v1.0.0 // indirect

)

replace (
	k8s.io/api => k8s.io/api v0.0.0-20190409021203-6e4e0e4f393b
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190404173353-6a84e37a896d
	k8s.io/client-go => k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
)

go 1.13
