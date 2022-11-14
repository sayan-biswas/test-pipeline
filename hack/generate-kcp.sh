KUBE_CLIENT_GENERATOR=$(go env GOBIN)/client-gen
KUBE_LISTER_GENERATOR=$(go env GOBIN)/lister-gen
KUBE_INFORMER_GENERATOR=$(go env GOBIN)/informer-gen
KCP_CODE_GENERATOR=$(go env GOBIN)/kcp-code-generator

#${KUBE_CLIENT_GENERATOR} \
#  --go-header-file ./hack/boilerplate/boilerplate.go.txt \
#  --input-base github.com/tektoncd/pipeline/pkg/apis \
#  --input pipeline/v1alpha1 \
#  --input pipeline/v1beta1 \
#  --input pipeline/v1 \
#  --input resolution/v1alpha1 \
#  --input resolution/v1beta1 \
#  --input resource/v1alpha1 \
#  --output-base . \
#  --output-package github.com/tektoncd/pipeline/pkg/client/generated/clientset \
#  --trim-path-prefix github.com/tektoncd/pipeline
#
#${KUBE_LISTER_GENERATOR} \
#  --go-header-file ./hack/boilerplate/boilerplate.go.txt \
#  --input-dirs github.com/tektoncd/pipeline/pkg/apis/pipeline/v1,github.com/tektoncd/pipeline/pkg/apis/pipeline/v1alpha1,github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1, github.com/tektoncd/pipeline/pkg/apis/resolution/v1alpha1,github.com/tektoncd/resolution/pkg/apis/pipeline/v1beta1, github.com/tektoncd/pipeline/pkg/apis/resource/v1alpha1 \
#  --output-base . \
#  --output-package github.com/tektoncd/pipeline/pkg/client/generated/listers \
#  --trim-path-prefix github.com/tektoncd/pipeline
#
#${KUBE_INFORMER_GENERATOR} \
#  --versioned-clientset-package github.com/tektoncd/pipeline/pkg/client/generated/clientset \
#  --listers-package github.com/tektoncd/pipeline/pkg/client/generated/listers \
#  --go-header-file ./hack/boilerplate/boilerplate.go.txt \
#  --input-dirs github.com/tektoncd/pipeline/pkg/apis/pipeline/v1,github.com/tektoncd/pipeline/pkg/apis/pipeline/v1alpha1,github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1, github.com/tektoncd/pipeline/pkg/apis/resolution/v1alpha1,github.com/tektoncd/resolution/pkg/apis/pipeline/v1beta1, github.com/tektoncd/pipeline/pkg/apis/resource/v1alpha1 \
#  --output-base . \
#  --output-package github.com/tektoncd/pipeline/pkg/client/generated/informers \
#  --trim-path-prefix github.com/tektoncd/pipeline


#${KCP_CODE_GENERATOR} \
#  "client:externalOnly=true,standalone=true,outputPackagePath=github.com/tektoncd/pipeline/pkg/client/cluster,apiPackagePath=github.com/tektoncd/pipeline/pkg/apis/pipeline,singleClusterClientPackagePath=github.com/tektoncd/pipeline/pkg/client/clientset/versioned,headerFile=./hack/boilerplate/boilerplate.go.txt" \
#  "lister:apiPackagePath=github.com/tektoncd/pipeline/pkg/apis/pipeline,singleClusterListerPackagePath=github.com/tektoncd/pipeline/pkg/client/listers,headerFile=./hack/boilerplate/boilerplate.go.txt" \
#  "informer:externalOnly=true,standalone=true,outputPackagePath=github.com/tektoncd/pipeline/pkg/client/cluster,apiPackagePath=github.com/tektoncd/pipeline/pkg/apis/pipeline,singleClusterListerPackagePath=github.com/tektoncd/pipeline/pkg/client/listers,singleClusterInformerPackagePath=github.com/tektoncd/pipeline/pkg/client/informers,headerFile=./hack/boilerplate/boilerplate.go.txt" \
#  "paths=./pkg/apis/pipeline/..." \
#  "output:dir=./pkg/client/cluster"

#${KCP_CODE_GENERATOR} \
#  "client:externalOnly=true,standalone=true,outputPackagePath=github.com/tektoncd/pipeline/pkg/client/resolution/cluster,apiPackagePath=github.com/tektoncd/pipeline/pkg/apis/resolution,singleClusterClientPackagePath=github.com/tektoncd/pipeline/pkg/client/resolution/clientset/versioned,headerFile=./hack/boilerplate/boilerplate.go.txt" \
#  "lister:apiPackagePath=github.com/tektoncd/pipeline/pkg/apis/resolution,singleClusterListerPackagePath=github.com/tektoncd/pipeline/pkg/client/resolution/listers,headerFile=./hack/boilerplate/boilerplate.go.txt" \
#  "informer:externalOnly=true,standalone=true,outputPackagePath=github.com/tektoncd/pipeline/pkg/client/resolution/cluster,apiPackagePath=github.com/tektoncd/pipeline/pkg/apis/resolution,singleClusterListerPackagePath=github.com/tektoncd/pipeline/pkg/client/resolution/listers,singleClusterInformerPackagePath=github.com/tektoncd/pipeline/pkg/client/resolution/informers,headerFile=./hack/boilerplate/boilerplate.go.txt" \
#  "paths=./pkg/apis/resolution/..." \
#  "output:dir=./pkg/client/resolution/cluster"
#
${KCP_CODE_GENERATOR} \
  "client:externalOnly=true,standalone=true,outputPackagePath=github.com/tektoncd/pipeline/pkg/client/resource/cluster,apiPackagePath=github.com/tektoncd/pipeline/pkg/apis/resource,singleClusterClientPackagePath=github.com/tektoncd/pipeline/pkg/client/resource/clientset/versioned,headerFile=./hack/boilerplate/boilerplate.go.txt" \
  "lister:apiPackagePath=github.com/tektoncd/pipeline/pkg/apis/resource,singleClusterListerPackagePath=github.com/tektoncd/pipeline/pkg/client/resource/listers,headerFile=./hack/boilerplate/boilerplate.go.txt" \
  "informer:externalOnly=true,standalone=true,outputPackagePath=github.com/tektoncd/pipeline/pkg/client/resource/cluster,apiPackagePath=github.com/tektoncd/pipeline/pkg/apis/resource,singleClusterListerPackagePath=github.com/tektoncd/pipeline/pkg/client/resource/listers,singleClusterInformerPackagePath=github.com/tektoncd/pipeline/pkg/client/resource/informers,headerFile=./hack/boilerplate/boilerplate.go.txt" \
  "paths=./pkg/apis/resource/..." \
  "output:dir=./pkg/client/resource/cluster"