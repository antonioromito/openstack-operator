module github.com/openstack-k8s-operators/openstack-operator

go 1.21

require (
	github.com/cert-manager/cert-manager v1.14.7
	github.com/go-logr/logr v1.4.2
	github.com/go-playground/validator/v10 v10.24.0
	github.com/google/uuid v1.6.0
	github.com/iancoleman/strcase v0.3.0
	github.com/k8snetworkplumbingwg/network-attachment-definition-client v1.7.5
	github.com/onsi/ginkgo/v2 v2.20.1
	github.com/onsi/gomega v1.34.1
	github.com/openshift/api v3.9.0+incompatible
	github.com/openstack-k8s-operators/barbican-operator/api v0.5.1-0.20250224124409-84d45eee6a53
	github.com/openstack-k8s-operators/cinder-operator/api v0.5.1-0.20250225073106-03e654436dc8
	github.com/openstack-k8s-operators/designate-operator/api v0.5.1-0.20250224124655-8e302536f39c
	github.com/openstack-k8s-operators/glance-operator/api v0.5.1-0.20250224115614-90f4be40dbe8
	github.com/openstack-k8s-operators/heat-operator/api v0.5.1-0.20250225040412-85d27d3c6729
	github.com/openstack-k8s-operators/horizon-operator/api v0.5.1-0.20250225012432-71eddeaa4e91
	github.com/openstack-k8s-operators/infra-operator/apis v0.5.1-0.20250226120707-ee102640a845
	github.com/openstack-k8s-operators/ironic-operator/api v0.5.1-0.20250224101430-49ab890b73aa
	github.com/openstack-k8s-operators/keystone-operator/api v0.5.1-0.20250224101200-743a4b3fb404
	github.com/openstack-k8s-operators/lib-common/modules/ansible v0.5.1-0.20250224130558-eb5a5d85bfdb
	github.com/openstack-k8s-operators/lib-common/modules/certmanager v0.5.1-0.20250224130558-eb5a5d85bfdb
	github.com/openstack-k8s-operators/lib-common/modules/common v0.5.1-0.20250225220922-3ead38b7284f
	github.com/openstack-k8s-operators/lib-common/modules/storage v0.5.1-0.20250224130558-eb5a5d85bfdb
	github.com/openstack-k8s-operators/lib-common/modules/test v0.5.1-0.20250224130558-eb5a5d85bfdb
	github.com/openstack-k8s-operators/manila-operator/api v0.5.1-0.20250224103102-19721ac4f257
	github.com/openstack-k8s-operators/mariadb-operator/api v0.5.1-0.20250224101716-b42af3d653ba
	github.com/openstack-k8s-operators/neutron-operator/api v0.5.1-0.20250224121612-3a76df1432c8
	github.com/openstack-k8s-operators/nova-operator/api v0.5.1-0.20250226193946-9c6c30c7718e
	github.com/openstack-k8s-operators/octavia-operator/api v0.5.1-0.20250224125936-003d70aa4d9a
	github.com/openstack-k8s-operators/openstack-baremetal-operator/api v0.5.1-0.20250225064626-1d6b34a75b72
	github.com/openstack-k8s-operators/openstack-operator/apis v0.0.0-20240531084739-3b4c0451297c
	github.com/openstack-k8s-operators/ovn-operator/api v0.5.1-0.20250224101201-53811064bfea
	github.com/openstack-k8s-operators/placement-operator/api v0.5.1-0.20250224095804-36c58e753869
	github.com/openstack-k8s-operators/swift-operator/api v0.5.1-0.20250224103631-98c415d47733
	github.com/openstack-k8s-operators/telemetry-operator/api v0.5.1-0.20250224173700-b5088bcec0d9
	github.com/openstack-k8s-operators/test-operator/api v0.5.1-0.20250225094237-2d20179fe5a1
	github.com/pkg/errors v0.9.1
	github.com/rabbitmq/cluster-operator/v2 v2.11.0
	go.uber.org/zap v1.27.0
	golang.org/x/exp v0.0.0-20240719175910-8a7402abbf56
	gopkg.in/yaml.v3 v3.0.1
	k8s.io/api v0.29.14
	k8s.io/apimachinery v0.29.14
	k8s.io/client-go v0.29.14
	k8s.io/utils v0.0.0-20240711033017-18e509b52bc8
	sigs.k8s.io/controller-runtime v0.17.6
)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/emicklei/go-restful/v3 v3.12.0 // indirect
	github.com/evanphx/json-patch/v5 v5.9.0 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/gabriel-vasile/mimetype v1.4.8 // indirect
	github.com/go-logr/zapr v1.3.0 // indirect
	github.com/go-openapi/jsonpointer v0.21.0 // indirect
	github.com/go-openapi/jsonreference v0.21.0 // indirect
	github.com/go-openapi/swag v0.23.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-task/slim-sprig/v3 v3.0.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/gnostic-models v0.6.9-0.20230804172637-c7be7c783f49 // indirect
	github.com/google/go-cmp v0.7.0 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/google/pprof v0.0.0-20240727154555-813a5fbdbec8 // indirect
	github.com/gophercloud/gophercloud v1.14.1 // indirect
	github.com/imdario/mergo v0.3.16 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/metal3-io/baremetal-operator/apis v0.6.3 // indirect
	github.com/metal3-io/baremetal-operator/pkg/hardwareutils v0.5.1 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/openstack-k8s-operators/lib-common/modules/openstack v0.5.1-0.20250211124440-b44f6585b825 // indirect
	github.com/prometheus/client_golang v1.19.0 // indirect
	github.com/prometheus/client_model v0.6.0 // indirect
	github.com/prometheus/common v0.53.0 // indirect
	github.com/prometheus/procfs v0.13.0 // indirect
	github.com/rhobs/obo-prometheus-operator/pkg/apis/monitoring v0.71.0-rhobs1 // indirect
	github.com/rhobs/observability-operator v0.3.1 // indirect
	github.com/robfig/cron/v3 v3.0.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/crypto v0.32.0 // indirect
	golang.org/x/mod v0.20.0 // indirect
	golang.org/x/net v0.34.0 // indirect
	golang.org/x/oauth2 v0.18.0 // indirect
	golang.org/x/sys v0.30.0 // indirect
	golang.org/x/term v0.29.0 // indirect
	golang.org/x/text v0.22.0 // indirect
	golang.org/x/time v0.5.0 // indirect
	golang.org/x/tools v0.24.0 // indirect
	gomodules.xyz/jsonpatch/v2 v2.4.0 // indirect
	google.golang.org/appengine v1.6.8 // indirect
	google.golang.org/protobuf v1.34.1 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	k8s.io/apiextensions-apiserver v0.29.14 // indirect
	k8s.io/component-base v0.29.14 // indirect
	k8s.io/klog/v2 v2.120.1 // indirect
	k8s.io/kube-openapi v0.0.0-20240322212309-b815d8309940 // indirect
	sigs.k8s.io/gateway-api v1.0.0 // indirect
	sigs.k8s.io/json v0.0.0-20221116044647-bc3834ca7abd // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.4.1 // indirect
	sigs.k8s.io/yaml v1.4.0 // indirect
)

replace github.com/openstack-k8s-operators/openstack-operator/apis => ./apis

// mschuppert: map to latest commit from release-4.16 tag
// must consistent within modules and service operators
replace github.com/openshift/api => github.com/openshift/api v0.0.0-20240830023148-b7d0481c9094 //allow-merging

// custom RabbitmqClusterSpecCore for OpenStackControlplane (v2.9.0_patches_tag)
replace github.com/rabbitmq/cluster-operator/v2 => github.com/openstack-k8s-operators/rabbitmq-cluster-operator/v2 v2.6.1-0.20241017142550-a3524acedd49 //allow-merging
