module github.com/ibrokethecloud/kube-bench-metrics

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751 // indirect
	github.com/alecthomas/units v0.0.0-20190910110746-680d30ca3117 // indirect
	github.com/aquasecurity/kube-bench v0.5.0
	github.com/aws/aws-sdk-go v1.35.28 // indirect
	github.com/googleapis/gnostic v0.3.1 // indirect
	github.com/gregjones/httpcache v0.0.0-20190611155906-901d90724c79 // indirect
	github.com/peterbourgon/diskv v2.0.1+incompatible // indirect
	github.com/prometheus/client_golang v1.1.0
	github.com/sirupsen/logrus v1.4.2
	github.com/urfave/cli v1.22.1
	gopkg.in/inf.v0 v0.9.1 // indirect
	// Manually added
	k8s.io/api v0.0.0-20190409021203-6e4e0e4f393b
	k8s.io/apimachinery v0.0.0-20190927035529-0104e33c351d
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/klog v1.0.0 // indirect
	sigs.k8s.io/yaml v1.1.0 // indirect

)

replace (
	k8s.io/api => k8s.io/api v0.0.0-20190409021203-6e4e0e4f393b
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190404173353-6a84e37a896d
	k8s.io/client-go => k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
)

go 1.13
