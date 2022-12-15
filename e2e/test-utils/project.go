package test_utils

import (
	"bytes"
	"context"
	"fmt"
	"github.com/fluxcd/go-git/v5"
	"github.com/huandu/xstrings"
	"github.com/jinzhu/copier"
	"github.com/kluctl/kluctl/v2/cmd/kluctl/commands"
	git2 "github.com/kluctl/kluctl/v2/pkg/git"
	"github.com/kluctl/kluctl/v2/pkg/status"
	"github.com/kluctl/kluctl/v2/pkg/utils"
	"github.com/kluctl/kluctl/v2/pkg/utils/uo"
	"github.com/kluctl/kluctl/v2/pkg/yaml"
	registry2 "helm.sh/helm/v3/pkg/registry"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
	"sync"
	"testing"
)

type TestProject struct {
	t *testing.T

	extraEnv   []string
	useProcess bool

	gitServer *git2.TestGitServer
}

type TestProjectOption func(p *TestProject)

func WithUseProcess(useProcess bool) TestProjectOption {
	return func(p *TestProject) {
		p.useProcess = useProcess
	}
}

func NewTestProject(t *testing.T, opts ...TestProjectOption) *TestProject {
	p := &TestProject{
		t: t,
	}

	for _, o := range opts {
		o(p)
	}

	p.gitServer = git2.NewTestGitServer(t)
	p.gitServer.GitInit("kluctl-project")

	p.UpdateKluctlYaml(func(o *uo.UnstructuredObject) error {
		return nil
	})
	p.UpdateDeploymentYaml(".", func(c *uo.UnstructuredObject) error {
		return nil
	})
	return p
}

func (p *TestProject) TestSlug() string {
	n := p.t.Name()
	n = xstrings.ToKebabCase(n)
	n = strings.ReplaceAll(n, "/", "-")
	return n
}

func (p *TestProject) AddExtraEnv(e string) {
	p.extraEnv = append(p.extraEnv, e)
}

func (p *TestProject) UpdateKluctlYaml(update func(o *uo.UnstructuredObject) error) {
	p.UpdateYaml(".kluctl.yml", update, "")
}

func (p *TestProject) UpdateDeploymentYaml(dir string, update func(o *uo.UnstructuredObject) error) {
	p.UpdateYaml(filepath.Join(dir, "deployment.yml"), func(o *uo.UnstructuredObject) error {
		if dir == "." {
			o.SetNestedField(p.TestSlug(), "commonLabels", "project_name")
		}
		return update(o)
	}, "")
}

func (p *TestProject) UpdateYaml(path string, update func(o *uo.UnstructuredObject) error, message string) {
	p.gitServer.UpdateYaml("kluctl-project", path, func(o map[string]any) error {
		u := uo.FromMap(o)
		err := update(u)
		if err != nil {
			return err
		}
		_ = copier.CopyWithOption(&o, &u.Object, copier.Option{DeepCopy: true})
		return nil
	}, message)
}

func (p *TestProject) UpdateFile(path string, update func(f string) (string, error), message string) {
	p.gitServer.UpdateFile("kluctl-project", path, update, message)
}

func (p *TestProject) GetYaml(path string) *uo.UnstructuredObject {
	o, err := uo.FromFile(filepath.Join(p.LocalRepoDir(), path))
	if err != nil {
		p.t.Fatal(err)
	}
	return o
}

func (p *TestProject) GetDeploymentYaml(dir string) *uo.UnstructuredObject {
	return p.GetYaml(filepath.Join(dir, "deployment.yml"))
}

func (p *TestProject) ListDeploymentItemPathes(dir string, fullPath bool) []string {
	var ret []string
	o := p.GetDeploymentYaml(dir)
	l, _, err := o.GetNestedObjectList("deployments")
	if err != nil {
		p.t.Fatal(err)
	}
	for _, x := range l {
		pth, ok, _ := x.GetNestedString("path")
		if ok {
			x := pth
			if fullPath {
				x = filepath.Join(dir, pth)
			}
			ret = append(ret, x)
		}
		pth, ok, _ = x.GetNestedString("include")
		if ok {
			ret = append(ret, p.ListDeploymentItemPathes(filepath.Join(dir, pth), fullPath)...)
		}
	}
	return ret
}

func (p *TestProject) UpdateKustomizeDeployment(dir string, update func(o *uo.UnstructuredObject, wt *git.Worktree) error) {
	wt := p.gitServer.GetWorktree("kluctl-project")

	pth := filepath.Join(dir, "kustomization.yml")
	p.UpdateYaml(pth, func(o *uo.UnstructuredObject) error {
		return update(o, wt)
	}, fmt.Sprintf("Update kustomization.yml for %s", dir))
}

func (p *TestProject) UpdateTarget(name string, cb func(target *uo.UnstructuredObject)) {
	p.UpdateNamedListItem(uo.KeyPath{"targets"}, name, cb)
}

func (p *TestProject) UpdateSecretSet(name string, cb func(secretSet *uo.UnstructuredObject)) {
	p.UpdateNamedListItem(uo.KeyPath{"secretsConfig", "secretSets"}, name, cb)
}

func (p *TestProject) UpdateNamedListItem(path uo.KeyPath, name string, cb func(item *uo.UnstructuredObject)) {
	if cb == nil {
		cb = func(target *uo.UnstructuredObject) {}
	}

	p.UpdateKluctlYaml(func(o *uo.UnstructuredObject) error {
		l, _, _ := o.GetNestedObjectList(path...)
		var newList []*uo.UnstructuredObject
		found := false
		for _, item := range l {
			n, _, _ := item.GetNestedString("name")
			if n == name {
				cb(item)
				found = true
			}
			newList = append(newList, item)
		}
		if !found {
			n := uo.FromMap(map[string]interface{}{
				"name": name,
			})
			cb(n)
			newList = append(newList, n)
		}

		_ = o.SetNestedObjectList(newList, path...)
		return nil
	})
}

func (p *TestProject) UpdateDeploymentItems(dir string, update func(items []*uo.UnstructuredObject) []*uo.UnstructuredObject) {
	p.UpdateDeploymentYaml(dir, func(o *uo.UnstructuredObject) error {
		items, _, _ := o.GetNestedObjectList("deployments")
		items = update(items)
		return o.SetNestedField(items, "deployments")
	})
}

func (p *TestProject) AddDeploymentItem(dir string, item *uo.UnstructuredObject) {
	p.UpdateDeploymentItems(dir, func(items []*uo.UnstructuredObject) []*uo.UnstructuredObject {
		for _, x := range items {
			if reflect.DeepEqual(x, item) {
				return items
			}
		}
		items = append(items, item)
		return items
	})
}

func (p *TestProject) AddDeploymentInclude(dir string, includePath string, tags []string) {
	n := uo.FromMap(map[string]interface{}{
		"include": includePath,
	})
	if len(tags) != 0 {
		n.SetNestedField(tags, "tags")
	}
	p.AddDeploymentItem(dir, n)
}

func (p *TestProject) AddDeploymentIncludes(dir string) {
	var pp []string
	for _, x := range strings.Split(dir, "/") {
		if x != "." {
			p.AddDeploymentInclude(filepath.Join(pp...), x, nil)
		}
		pp = append(pp, x)
	}
}

func (p *TestProject) AddKustomizeDeployment(dir string, resources []KustomizeResource, tags []string) {
	deploymentDir := filepath.Dir(dir)
	if deploymentDir != "" {
		p.AddDeploymentIncludes(deploymentDir)
	}

	absKustomizeDir := filepath.Join(p.LocalRepoDir(), dir)

	err := os.MkdirAll(absKustomizeDir, 0o700)
	if err != nil {
		p.t.Fatal(err)
	}

	p.UpdateKustomizeDeployment(dir, func(o *uo.UnstructuredObject, wt *git.Worktree) error {
		o.SetNestedField("kustomize.config.k8s.io/v1beta1", "apiVersion")
		o.SetNestedField("Kustomization", "kind")
		return nil
	})

	p.AddKustomizeResources(dir, resources)
	p.UpdateDeploymentYaml(deploymentDir, func(o *uo.UnstructuredObject) error {
		d, _, _ := o.GetNestedObjectList("deployments")
		n := uo.FromMap(map[string]interface{}{
			"path": filepath.Base(dir),
		})
		if len(tags) != 0 {
			n.SetNestedField(tags, "tags")
		}
		d = append(d, n)
		_ = o.SetNestedObjectList(d, "deployments")
		return nil
	})
}

func (p *TestProject) AddHelmDeployment(dir string, repoUrl string, chartName, version string, releaseName string, namespace string, values map[string]any) {
	localPath := ""
	if u, err := url.Parse(repoUrl); err != nil || u.Host == "" {
		localPath = repoUrl
		repoUrl = ""
	} else if registry2.IsOCI(repoUrl) {
		repoUrl += "/" + chartName
		chartName = ""
	}

	p.AddKustomizeDeployment(dir, []KustomizeResource{
		{Name: "helm-rendered.yaml"},
	}, nil)

	p.UpdateYaml(filepath.Join(dir, "helm-chart.yaml"), func(o *uo.UnstructuredObject) error {
		*o = *uo.FromMap(map[string]interface{}{
			"helmChart": map[string]any{
				"repo":         repoUrl,
				"path":         localPath,
				"chartVersion": version,
				"releaseName":  releaseName,
				"namespace":    namespace,
			},
		})
		if chartName != "" {
			_ = o.SetNestedField(chartName, "helmChart", "chartName")
		}
		return nil
	}, "")

	if values != nil {
		p.UpdateYaml(filepath.Join(dir, "helm-values.yaml"), func(o *uo.UnstructuredObject) error {
			*o = *uo.FromMap(values)
			return nil
		}, "")
	}
}

func (p *TestProject) convertInterfaceToList(x interface{}) []interface{} {
	var ret []interface{}
	if l, ok := x.([]interface{}); ok {
		return l
	}
	if l, ok := x.([]*uo.UnstructuredObject); ok {
		for _, y := range l {
			ret = append(ret, y)
		}
		return ret
	}
	if l, ok := x.([]map[string]interface{}); ok {
		for _, y := range l {
			ret = append(ret, y)
		}
		return ret
	}
	return []interface{}{x}
}

type KustomizeResource struct {
	Name     string
	FileName string
	Content  interface{}
}

func (p *TestProject) AddKustomizeResources(dir string, resources []KustomizeResource) {
	p.UpdateKustomizeDeployment(dir, func(o *uo.UnstructuredObject, wt *git.Worktree) error {
		l, _, _ := o.GetNestedList("resources")
		for _, r := range resources {
			l = append(l, r.Name)
			fileName := r.FileName
			if fileName == "" {
				fileName = r.Name
			}
			if r.Content != nil {
				x := p.convertInterfaceToList(r.Content)
				err := yaml.WriteYamlAllFile(filepath.Join(p.LocalRepoDir(), dir, fileName), x)
				if err != nil {
					return err
				}
				_, err = wt.Add(filepath.Join(dir, fileName))
				if err != nil {
					return err
				}
			}
		}
		o.SetNestedField(l, "resources")
		return nil
	})
}

func (p *TestProject) DeleteKustomizeDeployment(dir string) {
	deploymentDir := filepath.Dir(dir)
	p.UpdateDeploymentItems(deploymentDir, func(items []*uo.UnstructuredObject) []*uo.UnstructuredObject {
		var newItems []*uo.UnstructuredObject
		for _, item := range items {
			pth, _, _ := item.GetNestedString("path")
			if pth == filepath.Base(dir) {
				continue
			}
			newItems = append(newItems, item)
		}
		return newItems
	})
}

func (p *TestProject) GitUrl() string {
	return p.gitServer.GitUrl("kluctl-project")
}

func (p *TestProject) LocalRepoDir() string {
	return p.gitServer.LocalRepoDir("kluctl-project")
}

func (p *TestProject) GetGitRepo() *git.Repository {
	return p.gitServer.GetGitRepo("kluctl-project")
}

func (p *TestProject) KluctlProcess(argsIn ...string) (string, string, error) {
	var args []string
	args = append(args, argsIn...)
	args = append(args, "--no-update-check")

	cwd := p.LocalRepoDir()

	args = append(args, "--debug")

	env := os.Environ()
	env = append(env, p.extraEnv...)

	// this will cause the init() function from call_kluctl_hack.go to invoke the kluctl root command and then exit
	env = append(env, "CALL_KLUCTL=true")
	env = append(env, fmt.Sprintf("KLUCTL_BASE_TMP_DIR=%s", p.t.TempDir()))

	p.t.Logf("Runnning kluctl: %s", strings.Join(args, " "))

	testExe, err := os.Executable()
	if err != nil {
		panic(err)
	}

	cmd := exec.Command(testExe, args...)
	cmd.Dir = cwd
	cmd.Env = env

	stdout, stderr, err := runHelper(p.t, cmd)
	return stdout, stderr, err
}

func (p *TestProject) KluctlProcessMust(argsIn ...string) (string, string) {
	stdout, stderr, err := p.KluctlProcess(argsIn...)
	if err != nil {
		p.t.Logf(stderr)
		p.t.Fatal(fmt.Errorf("kluctl failed: %w", err))
	}
	return stdout, stderr
}

func (p *TestProject) KluctlExecute(argsIn ...string) (string, string, error) {
	if len(p.extraEnv) != 0 {
		p.t.Fatal("extraEnv is only supported in KluctlProcess(...)")
	}

	var args []string
	args = append(args, "--project-dir", p.LocalRepoDir())
	args = append(args, argsIn...)

	p.t.Logf("Runnning kluctl: %s", strings.Join(args, " "))

	var m sync.Mutex
	stdoutBuf := bytes.NewBuffer(nil)
	stdout := status.NewLineRedirector(func(line string) {
		m.Lock()
		defer m.Unlock()
		p.t.Log(line)
		stdoutBuf.WriteString(line + "\n")
	})
	stderrBuf := bytes.NewBuffer(nil)

	ctx := context.Background()
	ctx = utils.WithTmpBaseDir(ctx, p.t.TempDir())
	ctx = commands.WithStdStreams(ctx, stdout, stderrBuf)
	sh := status.NewSimpleStatusHandler(func(message string) {
		m.Lock()
		defer m.Unlock()
		p.t.Log(message)
		stderrBuf.WriteString(message + "\n")
	}, false, true)
	defer func() {
		if sh != nil {
			sh.Stop()
		}
	}()
	ctx = status.NewContext(ctx, sh)
	err := commands.Execute(ctx, args, nil)
	sh.Stop()
	sh = nil
	_ = stdout.Close()
	return stdoutBuf.String(), stderrBuf.String(), err
}

func (p *TestProject) Kluctl(argsIn ...string) (string, string, error) {
	if p.useProcess {
		return p.KluctlProcess(argsIn...)
	} else {
		return p.KluctlExecute(argsIn...)
	}
}

func (p *TestProject) KluctlMust(argsIn ...string) (string, string) {
	stdout, stderr, err := p.Kluctl(argsIn...)
	if err != nil {
		p.t.Fatal(fmt.Errorf("kluctl failed: %w", err))
	}
	return stdout, stderr
}
