module github.com/tektoncd/pipeline

go 1.18

require (
	github.com/ahmetb/gen-crd-api-reference-docs v0.3.1-0.20220720053627-e327d0730470 // Waiting for https://github.com/ahmetb/gen-crd-api-reference-docs/pull/43/files to merge
	github.com/cloudevents/sdk-go/v2 v2.11.0
	github.com/containerd/containerd v1.6.8
	github.com/go-git/go-git/v5 v5.4.2
	github.com/google/go-cmp v0.5.9
	github.com/google/go-containerregistry v0.8.1-0.20220216220642-00c59d91847c
	github.com/google/go-containerregistry/pkg/authn/k8schain v0.0.0-20220328141311-efc62d802606
	github.com/google/uuid v1.3.0
	github.com/hashicorp/errwrap v1.1.0
	github.com/hashicorp/go-multierror v1.1.1
	github.com/hashicorp/golang-lru v0.5.4
	github.com/jenkins-x/go-scm v1.11.19
	github.com/mitchellh/go-homedir v1.1.0
	github.com/opencontainers/image-spec v1.0.3-0.20220114050600-8b9d41f48198
	github.com/pkg/errors v0.9.1
	github.com/tektoncd/plumbing v0.0.0-20220817140952-3da8ce01aeeb
	go.opencensus.io v0.23.0
	go.uber.org/zap v1.23.0
	golang.org/x/oauth2 v0.0.0-20220411215720-9780585627b5
	gomodules.xyz/jsonpatch/v2 v2.2.0
	k8s.io/api v0.24.3
	k8s.io/apimachinery v0.24.3
	k8s.io/client-go v0.24.3
	k8s.io/code-generator v0.26.2
	k8s.io/klog v1.0.0
	k8s.io/kube-openapi v0.0.0-20221012153701-172d655c2280
	knative.dev/pkg v0.0.0-20220818004048-4a03844c0b15
	sigs.k8s.io/yaml v1.3.0
)

require github.com/benbjohnson/clock v1.1.0 // indirect

require (
	code.gitea.io/sdk/gitea v0.15.1
	github.com/goccy/kpoward v0.1.0
	github.com/kcp-dev/apimachinery v0.0.0-20221102195355-d65878bc16be
	github.com/kcp-dev/client-go v0.0.0-20221103171446-a51d1144350f
	github.com/kcp-dev/logicalcluster/v2 v2.0.0-alpha.3
	k8s.io/utils v0.0.0-20220728103510-ee6ede2d64ed
)

require (
	github.com/Microsoft/go-winio v0.5.1 // indirect
	github.com/ProtonMail/go-crypto v0.0.0-20210428141323-04723f9f07d7 // indirect
	github.com/acomagu/bufpipe v1.0.3 // indirect
	github.com/emirpasic/gods v1.12.0 // indirect
	github.com/go-git/gcfg v1.5.0 // indirect
	github.com/go-git/go-billy/v5 v5.3.1
	github.com/jbenet/go-context v0.0.0-20150711004518-d14ea06fba99 // indirect
	github.com/kevinburke/ssh_config v0.0.0-20201106050909-4977a11b4351 // indirect
	github.com/sergi/go-diff v1.1.0 // indirect
	github.com/xanzy/ssh-agent v0.3.0 // indirect
	gopkg.in/warnings.v0 v0.1.2 // indirect
)

require (
	cloud.google.com/go/compute v1.5.0 // indirect
	contrib.go.opencensus.io/exporter/ocagent v0.7.1-0.20200907061046-05415f1de66d // indirect
	contrib.go.opencensus.io/exporter/prometheus v0.4.0 // indirect
	github.com/Azure/azure-sdk-for-go v63.3.0+incompatible // indirect
	github.com/Azure/go-autorest v14.2.0+incompatible // indirect
	github.com/Azure/go-autorest/autorest v0.11.27 // indirect
	github.com/Azure/go-autorest/autorest/adal v0.9.20 // indirect
	github.com/Azure/go-autorest/autorest/azure/auth v0.5.11 // indirect
	github.com/Azure/go-autorest/autorest/azure/cli v0.4.5 // indirect
	github.com/Azure/go-autorest/autorest/date v0.3.0 // indirect
	github.com/Azure/go-autorest/logger v0.2.1 // indirect
	github.com/Azure/go-autorest/tracing v0.6.0 // indirect
	github.com/PuerkitoBio/purell v1.1.1 // indirect
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 // indirect
	github.com/aws/aws-sdk-go-v2 v1.14.0 // indirect
	github.com/aws/aws-sdk-go-v2/config v1.14.0 // indirect
	github.com/aws/aws-sdk-go-v2/credentials v1.9.0 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.11.0 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.5 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.3.0 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.3.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/ecr v1.15.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/ecrpublic v1.12.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.8.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.10.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.15.0 // indirect
	github.com/aws/smithy-go v1.11.0 // indirect
	github.com/awslabs/amazon-ecr-credential-helper/ecr-login v0.0.0-20220228164355-396b2034c795 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/blang/semver/v4 v4.0.0 // indirect
	github.com/blendle/zapdriver v1.3.1 // indirect
	github.com/bluekeyes/go-gitdiff v0.4.0 // indirect
	github.com/census-instrumentation/opencensus-proto v0.3.0 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/chrismellard/docker-credential-acr-env v0.0.0-20220119192733-fe33c00cee21 // indirect
	github.com/containerd/stargz-snapshotter/estargz v0.11.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dimchansky/utfbom v1.1.1 // indirect
	github.com/docker/cli v20.10.12+incompatible // indirect
	github.com/docker/distribution v2.8.0+incompatible // indirect
	github.com/docker/docker v20.10.12+incompatible // indirect
	github.com/docker/docker-credential-helpers v0.6.4 // indirect
	github.com/emicklei/go-restful/v3 v3.9.0 // indirect
	github.com/evanphx/json-patch v5.6.0+incompatible // indirect
	github.com/evanphx/json-patch/v5 v5.6.0 // indirect
	github.com/go-kit/log v0.1.0 // indirect
	github.com/go-logfmt/logfmt v0.5.0 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.20.0 // indirect
	github.com/go-openapi/swag v0.19.15 // indirect
	github.com/gobuffalo/flect v0.2.4 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang-jwt/jwt/v4 v4.3.0 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/gnostic v0.5.7-v3refs // indirect
	github.com/google/go-containerregistry/pkg/authn/kubernetes v0.0.0-20220301182634-bfe2ffc6b6bd // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	github.com/hashicorp/go-version v1.3.0 // indirect
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/kelseyhightower/envconfig v1.4.0 // indirect
	github.com/klauspost/compress v1.14.4 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.2-0.20181231171920-c182affec369 // indirect
	github.com/mitchellh/copystructure v1.0.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.1 // indirect
	github.com/moby/spdystream v0.2.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/openzipkin/zipkin-go v0.3.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_golang v1.12.1 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.32.1 // indirect
	github.com/prometheus/procfs v0.7.3 // indirect
	github.com/prometheus/statsd_exporter v0.21.0 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/shurcooL/githubv4 v0.0.0-20190718010115-4ba037080260 // indirect
	github.com/shurcooL/graphql v0.0.0-20181231061246-d48a9a75455f // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	github.com/spf13/pflag v1.0.6-0.20210604193023-d5e0c0615ace // indirect
	github.com/stretchr/testify v1.8.0 // indirect
	github.com/vbatts/tar-split v0.11.2 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/automaxprocs v1.4.0 // indirect
	go.uber.org/goleak v1.1.12 // indirect
	go.uber.org/multierr v1.7.0 // indirect
	golang.org/x/crypto v0.1.0 // indirect
	golang.org/x/mod v0.6.0 // indirect
	golang.org/x/net v0.7.0 // indirect
	golang.org/x/sync v0.0.0-20220722155255-886fb9371eb4 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/term v0.5.0 // indirect
	golang.org/x/text v0.7.0 // indirect
	golang.org/x/time v0.0.0-20220224211638-0e9765cccd65 // indirect
	golang.org/x/tools v0.2.0 // indirect
	google.golang.org/api v0.70.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20220502173005-c8bf987b8c21 // indirect
	google.golang.org/grpc v1.47.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k8s.io/apiextensions-apiserver v0.24.3 // indirect
	k8s.io/gengo v0.0.0-20220902162205-c0856e24416d // indirect
	k8s.io/klog/v2 v2.80.1 // indirect
	sigs.k8s.io/json v0.0.0-20220713155537-f223a00ba0e2 // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.2.3 // indirect
)

replace github.com/ahmetb/gen-crd-api-reference-docs => github.com/tektoncd/ahmetb-gen-crd-api-reference-docs v0.3.1-0.20220729140133-6ce2d5aafcb4 // Waiting for https://github.com/ahmetb/gen-crd-api-reference-docs/pull/43/files to merge

replace knative.dev/pkg => ../../knative/pkg-temp

//replace (
//	k8s.io/api => github.com/kcp-dev/kubernetes/staging/src/k8s.io/api v0.0.0-20221021135508-401e4f0fc370
//	k8s.io/apiextensions-apiserver => github.com/kcp-dev/kubernetes/staging/src/k8s.io/apiextensions-apiserver v0.0.0-20221021135508-401e4f0fc370
//	k8s.io/apimachinery => github.com/kcp-dev/kubernetes/staging/src/k8s.io/apimachinery v0.0.0-20221021135508-401e4f0fc370
//	k8s.io/apiserver => github.com/kcp-dev/kubernetes/staging/src/k8s.io/apiserver v0.0.0-20221021135508-401e4f0fc370
//	k8s.io/cli-runtime => github.com/kcp-dev/kubernetes/staging/src/k8s.io/cli-runtime v0.0.0-20221021135508-401e4f0fc370
//	k8s.io/client-go => github.com/kcp-dev/kubernetes/staging/src/k8s.io/client-go v0.0.0-20221021135508-401e4f0fc370
//	k8s.io/cloud-provider => github.com/kcp-dev/kubernetes/staging/src/k8s.io/cloud-provider v0.0.0-20221021135508-401e4f0fc370
//	k8s.io/cluster-bootstrap => github.com/kcp-dev/kubernetes/staging/src/k8s.io/cluster-bootstrap v0.0.0-20221021135508-401e4f0fc370
//	k8s.io/code-generator => github.com/kcp-dev/kubernetes/staging/src/k8s.io/code-generator v0.0.0-20221021135508-401e4f0fc370
//	k8s.io/component-base => github.com/kcp-dev/kubernetes/staging/src/k8s.io/component-base v0.0.0-20221021135508-401e4f0fc370
//	k8s.io/component-helpers => github.com/kcp-dev/kubernetes/staging/src/k8s.io/component-helpers v0.0.0-20221021135508-401e4f0fc370
//	k8s.io/controller-manager => github.com/kcp-dev/kubernetes/staging/src/k8s.io/controller-manager v0.0.0-20221021135508-401e4f0fc370
//	k8s.io/cri-api => github.com/kcp-dev/kubernetes/staging/src/k8s.io/cri-api v0.0.0-20221021135508-401e4f0fc370
//	k8s.io/csi-translation-lib => github.com/kcp-dev/kubernetes/staging/src/k8s.io/csi-translation-lib v0.0.0-20221021135508-401e4f0fc370
//	k8s.io/kube-aggregator => github.com/kcp-dev/kubernetes/staging/src/k8s.io/kube-aggregator v0.0.0-20221021135508-401e4f0fc370
//	k8s.io/kube-controller-manager => github.com/kcp-dev/kubernetes/staging/src/k8s.io/kube-controller-manager v0.0.0-20221021135508-401e4f0fc370
//	k8s.io/kube-proxy => github.com/kcp-dev/kubernetes/staging/src/k8s.io/kube-proxy v0.0.0-20221021135508-401e4f0fc370
//	k8s.io/kube-scheduler => github.com/kcp-dev/kubernetes/staging/src/k8s.io/kube-scheduler v0.0.0-20221021135508-401e4f0fc370
//	k8s.io/kubectl => github.com/kcp-dev/kubernetes/staging/src/k8s.io/kubectl v0.0.0-20221021135508-401e4f0fc370
//	k8s.io/kubelet => github.com/kcp-dev/kubernetes/staging/src/k8s.io/kubelet v0.0.0-20221021135508-401e4f0fc370
//	k8s.io/kubernetes => github.com/kcp-dev/kubernetes v0.0.0-20221021135508-401e4f0fc370
//	k8s.io/legacy-cloud-providers => github.com/kcp-dev/kubernetes/staging/src/k8s.io/legacy-cloud-providers v0.0.0-20221021135508-401e4f0fc370
//	k8s.io/metrics => github.com/kcp-dev/kubernetes/staging/src/k8s.io/metrics v0.0.0-20221021135508-401e4f0fc370
//	k8s.io/mount-utils => github.com/kcp-dev/kubernetes/staging/src/k8s.io/mount-utils v0.0.0-20221021135508-401e4f0fc370
//	k8s.io/pod-security-admission => github.com/kcp-dev/kubernetes/staging/src/k8s.io/pod-security-admission v0.0.0-20221021135508-401e4f0fc370
//	k8s.io/sample-apiserver => github.com/kcp-dev/kubernetes/staging/src/k8s.io/sample-apiserver v0.0.0-20221021135508-401e4f0fc370
//	knative.dev/pkg => /Users/sayan/Code/github.com/knative/pkg
//)
