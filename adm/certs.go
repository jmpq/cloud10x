/*
Copyright 2017 The Kubernetes Authors.

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

package main

import (
	//"fmt"

	"github.com/spf13/cobra"

	//kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	"k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/validation"
	//cmdutil "k8s.io/kubernetes/cmd/kubeadm/app/cmd/util"
	//kubeadmconstants "k8s.io/kubernetes/cmd/kubeadm/app/constants"
	//"k8s.io/kubernetes/pkg/util/normalizer"
	kubeadmapiext "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/v1alpha1"
	certsphase "k8s.io/kubernetes/cmd/kubeadm/app/phases/certs"
	kubeadmutil "k8s.io/kubernetes/cmd/kubeadm/app/util"
	configutil "k8s.io/kubernetes/cmd/kubeadm/app/util/config"
	//"k8s.io/kubernetes/pkg/api/legacyscheme"
)

func createPKIAssets(cmd *cobra.Command, args []string, cfgPath *string, cfg *kubeadmapiext.MasterConfiguration) error {
	if err := validation.ValidateMixedArguments(cmd.Flags()); err != nil {
		kubeadmutil.CheckErr(err)
	}

	// This call returns the ready-to-use configuration based on the configuration file that might or might not exist and the default cfg populated by flags
	internalcfg, err := configutil.ConfigFileAndDefaultsToInternalConfig(*cfgPath, cfg)
	kubeadmutil.CheckErr(err)

	// Execute the cmdFunc

	err = certsphase.CreatePKIAssets(internalcfg)
	kubeadmutil.CheckErr(err)

	return err
}
