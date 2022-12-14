/*
Copyright 2022 The Tekton Authors
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package bundle

import (
	"context"
	"fmt"

	"github.com/google/go-containerregistry/pkg/name"
	"github.com/tektoncd/pipeline/pkg/resolution/resolver/framework"
)

// ParamServiceAccount is the parameter defining what service
// account name to use for bundle requests.
const ParamServiceAccount = "serviceAccount"

// ParamBundle is the parameter defining what the bundle image url is.
const ParamBundle = "bundle"

// ParamName is the parameter defining what the layer name in the bundle
// image is.
const ParamName = "name"

// ParamKind is the parameter defining what the layer kind in the bundle
// image is.
const ParamKind = "kind"

// OptionsFromParams parses the params from a resolution request and
// converts them into options to pass as part of a bundle request.
func OptionsFromParams(ctx context.Context, params map[string]string) (RequestOptions, error) {
	opts := RequestOptions{}
	conf := framework.GetResolverConfigFromContext(ctx)

	saVal, ok := params[ParamServiceAccount]
	sa := ""
	if !ok || saVal == "" {
		if saString, ok := conf[ConfigServiceAccount]; ok {
			sa = saString
		} else {
			return opts, fmt.Errorf("default Service Account  was not set during installation of the bundle resolver")
		}
	} else {
		sa = saVal
	}

	bundleVal, ok := params[ParamBundle]
	if !ok || bundleVal == "" {
		return opts, fmt.Errorf("parameter %q required", ParamBundle)
	}
	if _, err := name.ParseReference(bundleVal); err != nil {
		return opts, fmt.Errorf("invalid bundle reference: %w", err)
	}

	nameVal, ok := params[ParamName]
	if !ok || nameVal == "" {
		return opts, fmt.Errorf("parameter %q required", ParamName)
	}

	kindVal, ok := params[ParamKind]
	kind := ""
	if !ok || kindVal == "" {
		if kindString, ok := conf[ConfigKind]; ok {
			kind = kindString
		} else {
			return opts, fmt.Errorf("default resource Kind  was not set during installation of the bundle resolver")
		}
	} else {
		kind = kindVal
	}

	opts.ServiceAccount = sa
	opts.Bundle = bundleVal
	opts.EntryName = nameVal
	opts.Kind = kind

	return opts, nil
}
