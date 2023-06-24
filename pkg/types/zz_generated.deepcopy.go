//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023.

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

// Code generated by controller-gen. DO NOT EDIT.

package types

import (
	"github.com/kluctl/kluctl/v2/pkg/types/k8s"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeleteObjectItemConfig) DeepCopyInto(out *DeleteObjectItemConfig) {
	*out = *in
	if in.Group != nil {
		in, out := &in.Group, &out.Group
		*out = new(string)
		**out = **in
	}
	if in.Kind != nil {
		in, out := &in.Kind, &out.Kind
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeleteObjectItemConfig.
func (in *DeleteObjectItemConfig) DeepCopy() *DeleteObjectItemConfig {
	if in == nil {
		return nil
	}
	out := new(DeleteObjectItemConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentArg) DeepCopyInto(out *DeploymentArg) {
	*out = *in
	if in.Default != nil {
		in, out := &in.Default, &out.Default
		*out = new(v1.JSON)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentArg.
func (in *DeploymentArg) DeepCopy() *DeploymentArg {
	if in == nil {
		return nil
	}
	out := new(DeploymentArg)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentItemConfig) DeepCopyInto(out *DeploymentItemConfig) {
	*out = *in
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
	if in.Include != nil {
		in, out := &in.Include, &out.Include
		*out = new(string)
		**out = **in
	}
	if in.Git != nil {
		in, out := &in.Git, &out.Git
		*out = new(GitProject)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
	if in.Vars != nil {
		in, out := &in.Vars, &out.Vars
		*out = make([]*VarsSource, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(VarsSource)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.DeleteObjects != nil {
		in, out := &in.DeleteObjects, &out.DeleteObjects
		*out = make([]DeleteObjectItemConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RenderedHelmChartConfig != nil {
		in, out := &in.RenderedHelmChartConfig, &out.RenderedHelmChartConfig
		*out = new(HelmChartConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.RenderedObjects != nil {
		in, out := &in.RenderedObjects, &out.RenderedObjects
		*out = make([]k8s.ObjectRef, len(*in))
		copy(*out, *in)
	}
	if in.RenderedInclude != nil {
		in, out := &in.RenderedInclude, &out.RenderedInclude
		*out = new(DeploymentProjectConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentItemConfig.
func (in *DeploymentItemConfig) DeepCopy() *DeploymentItemConfig {
	if in == nil {
		return nil
	}
	out := new(DeploymentItemConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentProjectConfig) DeepCopyInto(out *DeploymentProjectConfig) {
	*out = *in
	if in.Vars != nil {
		in, out := &in.Vars, &out.Vars
		*out = make([]*VarsSource, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(VarsSource)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.SealedSecrets != nil {
		in, out := &in.SealedSecrets, &out.SealedSecrets
		*out = new(SealedSecretsConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Deployments != nil {
		in, out := &in.Deployments, &out.Deployments
		*out = make([]*DeploymentItemConfig, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(DeploymentItemConfig)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.CommonLabels != nil {
		in, out := &in.CommonLabels, &out.CommonLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.CommonAnnotations != nil {
		in, out := &in.CommonAnnotations, &out.CommonAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.OverrideNamespace != nil {
		in, out := &in.OverrideNamespace, &out.OverrideNamespace
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.IgnoreForDiff != nil {
		in, out := &in.IgnoreForDiff, &out.IgnoreForDiff
		*out = make([]*IgnoreForDiffItemConfig, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(IgnoreForDiffItemConfig)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentProjectConfig.
func (in *DeploymentProjectConfig) DeepCopy() *DeploymentProjectConfig {
	if in == nil {
		return nil
	}
	out := new(DeploymentProjectConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalProject) DeepCopyInto(out *ExternalProject) {
	*out = *in
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(GitProject)
		(*in).DeepCopyInto(*out)
	}
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalProject.
func (in *ExternalProject) DeepCopy() *ExternalProject {
	if in == nil {
		return nil
	}
	out := new(ExternalProject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FixedImage) DeepCopyInto(out *FixedImage) {
	*out = *in
	if in.DeployedImage != nil {
		in, out := &in.DeployedImage, &out.DeployedImage
		*out = new(string)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.Object != nil {
		in, out := &in.Object, &out.Object
		*out = new(k8s.ObjectRef)
		**out = **in
	}
	if in.Deployment != nil {
		in, out := &in.Deployment, &out.Deployment
		*out = new(string)
		**out = **in
	}
	if in.Container != nil {
		in, out := &in.Container, &out.Container
		*out = new(string)
		**out = **in
	}
	if in.DeployTags != nil {
		in, out := &in.DeployTags, &out.DeployTags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DeploymentDir != nil {
		in, out := &in.DeploymentDir, &out.DeploymentDir
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FixedImage.
func (in *FixedImage) DeepCopy() *FixedImage {
	if in == nil {
		return nil
	}
	out := new(FixedImage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FixedImagesConfig) DeepCopyInto(out *FixedImagesConfig) {
	*out = *in
	if in.Images != nil {
		in, out := &in.Images, &out.Images
		*out = make([]FixedImage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FixedImagesConfig.
func (in *FixedImagesConfig) DeepCopy() *FixedImagesConfig {
	if in == nil {
		return nil
	}
	out := new(FixedImagesConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitProject) DeepCopyInto(out *GitProject) {
	*out = *in
	in.Url.DeepCopyInto(&out.Url)
	if in.Ref != nil {
		in, out := &in.Ref, &out.Ref
		*out = new(GitRef)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitProject.
func (in *GitProject) DeepCopy() *GitProject {
	if in == nil {
		return nil
	}
	out := new(GitProject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitRef) DeepCopyInto(out *GitRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitRef.
func (in *GitRef) DeepCopy() *GitRef {
	if in == nil {
		return nil
	}
	out := new(GitRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitRepoKey) DeepCopyInto(out *GitRepoKey) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitRepoKey.
func (in *GitRepoKey) DeepCopy() *GitRepoKey {
	if in == nil {
		return nil
	}
	out := new(GitRepoKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitUrl.
func (in *GitUrl) DeepCopy() *GitUrl {
	if in == nil {
		return nil
	}
	out := new(GitUrl)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalSealedSecretsConfig) DeepCopyInto(out *GlobalSealedSecretsConfig) {
	*out = *in
	if in.Bootstrap != nil {
		in, out := &in.Bootstrap, &out.Bootstrap
		*out = new(bool)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.ControllerName != nil {
		in, out := &in.ControllerName, &out.ControllerName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalSealedSecretsConfig.
func (in *GlobalSealedSecretsConfig) DeepCopy() *GlobalSealedSecretsConfig {
	if in == nil {
		return nil
	}
	out := new(GlobalSealedSecretsConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmChartConfig) DeepCopyInto(out *HelmChartConfig) {
	*out = *in
	in.HelmChartConfig2.DeepCopyInto(&out.HelmChartConfig2)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmChartConfig.
func (in *HelmChartConfig) DeepCopy() *HelmChartConfig {
	if in == nil {
		return nil
	}
	out := new(HelmChartConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmChartConfig2) DeepCopyInto(out *HelmChartConfig2) {
	*out = *in
	if in.CredentialsId != nil {
		in, out := &in.CredentialsId, &out.CredentialsId
		*out = new(string)
		**out = **in
	}
	if in.UpdateConstraints != nil {
		in, out := &in.UpdateConstraints, &out.UpdateConstraints
		*out = new(string)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.Output != nil {
		in, out := &in.Output, &out.Output
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmChartConfig2.
func (in *HelmChartConfig2) DeepCopy() *HelmChartConfig2 {
	if in == nil {
		return nil
	}
	out := new(HelmChartConfig2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IgnoreForDiffItemConfig) DeepCopyInto(out *IgnoreForDiffItemConfig) {
	*out = *in
	if in.FieldPath != nil {
		in, out := &in.FieldPath, &out.FieldPath
		*out = make(SingleStringOrList, len(*in))
		copy(*out, *in)
	}
	if in.FieldPathRegex != nil {
		in, out := &in.FieldPathRegex, &out.FieldPathRegex
		*out = make(SingleStringOrList, len(*in))
		copy(*out, *in)
	}
	if in.Group != nil {
		in, out := &in.Group, &out.Group
		*out = new(string)
		**out = **in
	}
	if in.Kind != nil {
		in, out := &in.Kind, &out.Kind
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IgnoreForDiffItemConfig.
func (in *IgnoreForDiffItemConfig) DeepCopy() *IgnoreForDiffItemConfig {
	if in == nil {
		return nil
	}
	out := new(IgnoreForDiffItemConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KluctlProject) DeepCopyInto(out *KluctlProject) {
	*out = *in
	if in.Targets != nil {
		in, out := &in.Targets, &out.Targets
		*out = make([]*Target, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Target)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]*DeploymentArg, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(DeploymentArg)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.SecretsConfig != nil {
		in, out := &in.SecretsConfig, &out.SecretsConfig
		*out = new(SecretsConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KluctlProject.
func (in *KluctlProject) DeepCopy() *KluctlProject {
	if in == nil {
		return nil
	}
	out := new(KluctlProject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SealedSecretsConfig) DeepCopyInto(out *SealedSecretsConfig) {
	*out = *in
	if in.OutputPattern != nil {
		in, out := &in.OutputPattern, &out.OutputPattern
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SealedSecretsConfig.
func (in *SealedSecretsConfig) DeepCopy() *SealedSecretsConfig {
	if in == nil {
		return nil
	}
	out := new(SealedSecretsConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SealingConfig) DeepCopyInto(out *SealingConfig) {
	*out = *in
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = (*in).DeepCopy()
	}
	if in.SecretSets != nil {
		in, out := &in.SecretSets, &out.SecretSets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CertFile != nil {
		in, out := &in.CertFile, &out.CertFile
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SealingConfig.
func (in *SealingConfig) DeepCopy() *SealingConfig {
	if in == nil {
		return nil
	}
	out := new(SealingConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretSet) DeepCopyInto(out *SecretSet) {
	*out = *in
	if in.Vars != nil {
		in, out := &in.Vars, &out.Vars
		*out = make([]*VarsSource, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(VarsSource)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretSet.
func (in *SecretSet) DeepCopy() *SecretSet {
	if in == nil {
		return nil
	}
	out := new(SecretSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretsConfig) DeepCopyInto(out *SecretsConfig) {
	*out = *in
	if in.SealedSecrets != nil {
		in, out := &in.SealedSecrets, &out.SealedSecrets
		*out = new(GlobalSealedSecretsConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.SecretSets != nil {
		in, out := &in.SecretSets, &out.SecretSets
		*out = make([]SecretSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretsConfig.
func (in *SecretsConfig) DeepCopy() *SecretsConfig {
	if in == nil {
		return nil
	}
	out := new(SecretsConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in SingleStringOrList) DeepCopyInto(out *SingleStringOrList) {
	{
		in := &in
		*out = make(SingleStringOrList, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SingleStringOrList.
func (in SingleStringOrList) DeepCopy() SingleStringOrList {
	if in == nil {
		return nil
	}
	out := new(SingleStringOrList)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Target) DeepCopyInto(out *Target) {
	*out = *in
	if in.Context != nil {
		in, out := &in.Context, &out.Context
		*out = new(string)
		**out = **in
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = (*in).DeepCopy()
	}
	if in.SealingConfig != nil {
		in, out := &in.SealingConfig, &out.SealingConfig
		*out = new(SealingConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Images != nil {
		in, out := &in.Images, &out.Images
		*out = make([]FixedImage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Target.
func (in *Target) DeepCopy() *Target {
	if in == nil {
		return nil
	}
	out := new(Target)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VarsSource) DeepCopyInto(out *VarsSource) {
	*out = *in
	if in.IgnoreMissing != nil {
		in, out := &in.IgnoreMissing, &out.IgnoreMissing
		*out = new(bool)
		**out = **in
	}
	if in.NoOverride != nil {
		in, out := &in.NoOverride, &out.NoOverride
		*out = new(bool)
		**out = **in
	}
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = (*in).DeepCopy()
	}
	if in.File != nil {
		in, out := &in.File, &out.File
		*out = new(string)
		**out = **in
	}
	if in.Git != nil {
		in, out := &in.Git, &out.Git
		*out = new(VarsSourceGit)
		(*in).DeepCopyInto(*out)
	}
	if in.ClusterConfigMap != nil {
		in, out := &in.ClusterConfigMap, &out.ClusterConfigMap
		*out = new(VarsSourceClusterConfigMapOrSecret)
		(*in).DeepCopyInto(*out)
	}
	if in.ClusterSecret != nil {
		in, out := &in.ClusterSecret, &out.ClusterSecret
		*out = new(VarsSourceClusterConfigMapOrSecret)
		(*in).DeepCopyInto(*out)
	}
	if in.SystemEnvVars != nil {
		in, out := &in.SystemEnvVars, &out.SystemEnvVars
		*out = (*in).DeepCopy()
	}
	if in.Http != nil {
		in, out := &in.Http, &out.Http
		*out = new(VarsSourceHttp)
		(*in).DeepCopyInto(*out)
	}
	if in.AwsSecretsManager != nil {
		in, out := &in.AwsSecretsManager, &out.AwsSecretsManager
		*out = new(VarsSourceAwsSecretsManager)
		(*in).DeepCopyInto(*out)
	}
	if in.Vault != nil {
		in, out := &in.Vault, &out.Vault
		*out = new(VarsSourceVault)
		**out = **in
	}
	if in.RenderedVars != nil {
		in, out := &in.RenderedVars, &out.RenderedVars
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VarsSource.
func (in *VarsSource) DeepCopy() *VarsSource {
	if in == nil {
		return nil
	}
	out := new(VarsSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VarsSourceAwsSecretsManager) DeepCopyInto(out *VarsSourceAwsSecretsManager) {
	*out = *in
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Profile != nil {
		in, out := &in.Profile, &out.Profile
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VarsSourceAwsSecretsManager.
func (in *VarsSourceAwsSecretsManager) DeepCopy() *VarsSourceAwsSecretsManager {
	if in == nil {
		return nil
	}
	out := new(VarsSourceAwsSecretsManager)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VarsSourceClusterConfigMapOrSecret) DeepCopyInto(out *VarsSourceClusterConfigMapOrSecret) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VarsSourceClusterConfigMapOrSecret.
func (in *VarsSourceClusterConfigMapOrSecret) DeepCopy() *VarsSourceClusterConfigMapOrSecret {
	if in == nil {
		return nil
	}
	out := new(VarsSourceClusterConfigMapOrSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VarsSourceGit) DeepCopyInto(out *VarsSourceGit) {
	*out = *in
	in.Url.DeepCopyInto(&out.Url)
	if in.Ref != nil {
		in, out := &in.Ref, &out.Ref
		*out = new(GitRef)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VarsSourceGit.
func (in *VarsSourceGit) DeepCopy() *VarsSourceGit {
	if in == nil {
		return nil
	}
	out := new(VarsSourceGit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VarsSourceHttp) DeepCopyInto(out *VarsSourceHttp) {
	*out = *in
	in.Url.DeepCopyInto(&out.Url)
	if in.Method != nil {
		in, out := &in.Method, &out.Method
		*out = new(string)
		**out = **in
	}
	if in.Body != nil {
		in, out := &in.Body, &out.Body
		*out = new(string)
		**out = **in
	}
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.JsonPath != nil {
		in, out := &in.JsonPath, &out.JsonPath
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VarsSourceHttp.
func (in *VarsSourceHttp) DeepCopy() *VarsSourceHttp {
	if in == nil {
		return nil
	}
	out := new(VarsSourceHttp)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VarsSourceVault) DeepCopyInto(out *VarsSourceVault) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VarsSourceVault.
func (in *VarsSourceVault) DeepCopy() *VarsSourceVault {
	if in == nil {
		return nil
	}
	out := new(VarsSourceVault)
	in.DeepCopyInto(out)
	return out
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new YamlUrl.
func (in *YamlUrl) DeepCopy() *YamlUrl {
	if in == nil {
		return nil
	}
	out := new(YamlUrl)
	in.DeepCopyInto(out)
	return out
}
