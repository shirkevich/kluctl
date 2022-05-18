package kluctl_project

import (
	"context"
	"fmt"
	"github.com/kluctl/kluctl/v2/pkg/deployment"
	"github.com/kluctl/kluctl/v2/pkg/jinja2"
	"github.com/kluctl/kluctl/v2/pkg/k8s"
	"github.com/kluctl/kluctl/v2/pkg/status"
	"github.com/kluctl/kluctl/v2/pkg/types"
	"github.com/kluctl/kluctl/v2/pkg/utils"
	"github.com/kluctl/kluctl/v2/pkg/utils/uo"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd/api"
	"path/filepath"
)

type TargetContext struct {
	KluctlProject        *LoadedKluctlProject
	Target               *types.Target
	ClusterConfig        *types.ClusterConfig
	K                    *k8s.K8sCluster
	DeploymentProject    *deployment.DeploymentProject
	DeploymentCollection *deployment.DeploymentCollection
}

func (p *LoadedKluctlProject) NewTargetContext(ctx context.Context, targetName string, clusterName *string, dryRun bool, args map[string]string, forSeal bool, images *deployment.Images, inclusion *utils.Inclusion, renderOutputDir string) (*TargetContext, error) {
	deploymentDir, err := filepath.Abs(p.DeploymentDir)
	if err != nil {
		return nil, err
	}

	var target *types.Target
	if targetName != "" {
		t, err := p.FindDynamicTarget(targetName)
		if err != nil {
			return nil, err
		}
		target = t.Target

		for _, fi := range target.Images {
			images.AddFixedImage(fi)
		}
	}

	var contextName *string
	if clusterName == nil && target != nil {
		clusterName = target.Cluster
		contextName = target.Context
	}

	clusterConfig, clientConfig, err := p.LoadClusterConfig(clusterName, contextName)
	if err != nil {
		return nil, err
	}

	var k *k8s.K8sCluster
	if clientConfig != nil {
		s := status.Start(ctx, fmt.Sprintf("Initializing k8s client for context %s", clusterConfig.Cluster.Context))
		k, err = k8s.NewK8sCluster(ctx, clientConfig, dryRun)
		if err != nil {
			s.Failed()
			return nil, err
		}
		s.Success()
	}

	varsCtx := jinja2.NewVarsCtx(p.J2, p.grc)
	err = varsCtx.UpdateChildFromStruct("cluster", clusterConfig.Cluster)
	if err != nil {
		return nil, err
	}
	targetVars, err := uo.FromStruct(target)
	if err != nil {
		return nil, err
	}
	varsCtx.UpdateChild("target", targetVars)

	allArgs := uo.New()

	if target != nil {
		for argName, argValue := range args {
			err = p.CheckDynamicArg(target, argName, argValue)
			if err != nil {
				return nil, err
			}
		}
	}
	allArgs.Merge(deployment.ConvertArgsToVars(args))
	if target != nil {
		if target.Args != nil {
			allArgs.Merge(target.Args)
		}
		if forSeal {
			if target.SealingConfig.Args != nil {
				allArgs.Merge(target.SealingConfig.Args)
			}
		}
	}

	err = deployment.CheckRequiredDeployArgs(deploymentDir, varsCtx, allArgs)
	if err != nil {
		return nil, err
	}

	varsCtx.UpdateChild("args", allArgs)

	d, err := deployment.NewDeploymentProject(ctx, k, varsCtx, deploymentDir, p.sealedSecretsDir, nil)
	if err != nil {
		return nil, err
	}

	d.DefaultSealedSecretsOutputPattern = targetName

	c, err := deployment.NewDeploymentCollection(ctx, d, images, inclusion, renderOutputDir, forSeal)
	if err != nil {
		return nil, err
	}

	targetCtx := &TargetContext{
		KluctlProject:        p,
		Target:               target,
		ClusterConfig:        clusterConfig,
		K:                    k,
		DeploymentProject:    d,
		DeploymentCollection: c,
	}

	return targetCtx, nil
}

func (p *LoadedKluctlProject) LoadClusterConfig(clusterName *string, contextName *string) (*types.ClusterConfig, *rest.Config, error) {
	var err error
	var clusterConfig *types.ClusterConfig

	if clusterName != nil {
		clusterConfig, err = types.LoadClusterConfig(p.ClustersDir, *clusterName)
		if err != nil {
			return nil, nil, err
		}
	}

	var clientConfig *rest.Config
	if clusterConfig != nil {
		clientConfig, _, err = p.loadArgs.ClientConfigGetter(&clusterConfig.Cluster.Context)
		if err != nil {
			return nil, nil, err
		}
	} else {
		var rawConfig *api.Config
		clientConfig, rawConfig, err = p.loadArgs.ClientConfigGetter(contextName)
		if err != nil {
			return nil, nil, err
		}
		ctx, ok := rawConfig.Contexts[rawConfig.CurrentContext]
		if !ok {
			return nil, nil, fmt.Errorf("context %s not found", rawConfig.CurrentContext)
		}
		clusterConfig = &types.ClusterConfig{
			Cluster: &types.ClusterConfig2{
				Name:    ctx.Cluster,
				Context: rawConfig.CurrentContext,
				Vars:    uo.New(),
			},
		}
	}

	return clusterConfig, clientConfig, nil
}
