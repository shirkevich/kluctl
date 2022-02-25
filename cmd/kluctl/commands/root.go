/*
Copyright © 2022 Alexander Block <ablock84@gmail.com>

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
package commands

import (
	"github.com/alecthomas/kong"
	"github.com/codablock/kluctl/pkg/utils"
	"github.com/codablock/kluctl/pkg/utils/uo"
	"github.com/codablock/kluctl/pkg/version"
	"github.com/codablock/kluctl/pkg/yaml"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const latestReleaseUrl = "https://api.github.com/repos/codablock/kluctl/releases/latest"

type cli struct {
	Verbosity     string `group:"global" short:"v" help:"Log level (debug, info, warn, error, fatal, panic)." default:"info"`
	NoUpdateCheck bool   `group:"global" help:"Disable update check on startup"`

	CheckImageUpdates checkImageUpdatesCmd `cmd:"" help:"Render deployment and check if any images have new tags available"`
	Delete            deleteCmd            `cmd:"" help:"Delete a target (or parts of it) from the corresponding cluster"`
	Deploy            deployCmd            `cmd:"" help:"Deploys a target to the corresponding cluster"`
	Diff              diffCmd              `cmd:"" help:"Perform a diff between the locally rendered target and the already deployed target"`
	Downscale         downscaleCmd         `cmd:"" help:"Downscale all deployments"`
	HelmPull          helmPullCmd          `cmd:"" help:"Recursively searches for 'helm-chart.yml' files and pulls the specified Helm charts"`
	HelmUpdate        helmUpdateCmd        `cmd:"" help:"Recursively searches for 'helm-chart.yml'' files and checks for new available versions"`
	ListImages        listImagesCmd        `cmd:"" help:"Renders the target and outputs all images used via 'images.get_image(...)"`
	ListTargets       listTargetsCmd       `cmd:"" help:"Outputs a yaml list with all target, including dynamic targets"`
	Prune             pruneCmd             `cmd:"" help:"Searches the target cluster for prunable objects and deletes them"`
	Render            renderCmd            `cmd:"" help:"Renders all resources and configuration files"`
	Seal              sealCmd              `cmd:"" help:"Seal secrets based on target's sealingConfig"`
	Validate          validateCmd          `cmd:"" help:"Validates the already deployed deployment"`
}

var flagGroups = []kong.Group{
	{Key: "project", Title: "Project arguments:", Description: ""},
	{Key: "images", Title: "Image arguments:", Description: ""},
	{Key: "inclusion", Title: "Inclusion/Exclusion arguments:", Description: ""},
	{Key: "misc", Title: "Misc arguments:", Description: ""},
	{Key: "global", Title: "Global arguments:", Description: ""},
}
var globalFlagGroup = &flagGroups[len(flagGroups)-1]

func (c *cli) setupLogs() error {
	lvl, err := log.ParseLevel(c.Verbosity)
	if err != nil {
		return err
	}
	log.SetLevel(lvl)
	return nil
}

type VersionCheckState struct {
	LastVersionCheck time.Time `yaml:"lastVersionCheck"`
}

func (c *cli) checkNewVersion() {
	if c.NoUpdateCheck {
		return
	}
	if version.Version == "0.0.0" {
		return
	}

	versionCheckPath := filepath.Join(utils.GetTmpBaseDir(), "version_check.yml")
	var versionCheckState VersionCheckState
	err := yaml.ReadYamlFile(versionCheckPath, &versionCheckState)
	if err == nil {
		if time.Now().Sub(versionCheckState.LastVersionCheck) < time.Hour {
			return
		}
	}

	versionCheckState.LastVersionCheck = time.Now()
	_ = yaml.WriteYamlFile(versionCheckPath, &versionCheckState)

	log.Debugf("Checking for new kluctl version")

	r, err := http.Get(latestReleaseUrl)
	if err != nil {
		return
	}
	defer r.Body.Close()

	var release uo.UnstructuredObject
	err = yaml.ReadYamlStream(r.Body, &release)
	if err != nil {
		return
	}

	latestVersionStr, ok, _ := release.GetNestedString("tag_name")
	if !ok {
		return
	}
	if strings.HasPrefix(latestVersionStr, "v") {
		latestVersionStr = latestVersionStr[1:]
	}
	latestVersion := utils.LooseVersion(latestVersionStr)
	localVersion := utils.LooseVersion(version.Version)
	if localVersion.Less(latestVersion, true) {
		log.Warningf("You are using an outdated version (%v) of kluctl. You should update soon to version %v", localVersion, latestVersion)
	}
}

func (c *cli) BeforeApply() error {
	if err := c.setupLogs(); err != nil {
		return err
	}
	c.checkNewVersion()
	return nil
}

func (c *cli) Help() string {
	return `
Deploy and manage complex deployments on Kubernetes

The missing glue to put together large Kubernetes deployments,
composed of multiple smaller parts (Helm/Kustomize/...) in a manageable and unified way.`
}

func ParseArgs(args []string, options ...kong.Option) (*kong.Context, error) {
	var cli cli

	helpOption := kong.HelpOptions{
		Compact: true,
		Summary: true,
		WrapUpperBound: 120,
	}

	var options2 []kong.Option
	options2 = append(options2, helpOption)
	options2 = append(options2, kong.ExplicitGroups(flagGroups))
	options2 = append(options2, options...)

	parser, err := kong.New(&cli, options2...)
	if err != nil {
		panic(err)
	}
	parser.Model.HelpFlag.Group = globalFlagGroup
	ctx, err := parser.Parse(args)
	return ctx, err
}

func Execute() {
	confOption := kong.Configuration(kong.JSON, "/etc/kluctl.json", "~/.kluctl/config.json")
	ctx, err := ExecuteWithArgs(os.Args[1:], confOption)
	if err != nil && ctx == nil {
		log.Fatal(err)
	}
	ctx.FatalIfErrorf(err)
}

func ExecuteWithArgs(args []string, options ...kong.Option) (*kong.Context, error) {
	ctx, err := ParseArgs(args, options...)
	if err != nil {
		return nil, err
	}
	return ctx, ctx.Run()
}
