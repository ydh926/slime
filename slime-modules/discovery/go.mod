module slime.io/slime/slime-modules/discovery

go 1.13

require (
	github.com/go-logr/logr v0.4.0
	github.com/gogo/protobuf v1.3.2
	github.com/onsi/ginkgo v1.12.1
	github.com/onsi/gomega v1.10.5
	github.com/orcaman/concurrent-map v0.0.0-20210106121528-16402b402231
	github.com/samuel/go-zookeeper v0.0.0-20201211165307-7117e9ea2414
	istio.io/api v0.0.0-20210322145030-ec7ef4cd6eaf
	k8s.io/apimachinery v0.20.2
	k8s.io/client-go v0.17.2
	sigs.k8s.io/controller-runtime v0.5.0
	slime.io/slime/slime-framework v0.0.0-00010101000000-000000000000
)

replace (
	k8s.io/api => k8s.io/api v0.17.2
	k8s.io/apimachinery => k8s.io/apimachinery v0.17.2
	slime.io/slime/slime-framework => ../../slime-framework
)
