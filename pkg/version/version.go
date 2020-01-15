package version

import (
	"encoding/json"
	"fmt"
	"runtime"
	"strings"
)

var (
	// TODO: Deprecate gitMajor and gitMinor, use only gitVersion
	// instead. First step in deprecation, keep the fields but make
	// them irrelevant. (Next we'll take it out, which may muck with
	// scripts consuming the kubectl version output - but most of
	// these should be looking at gitVersion already anyways.)
	gitMajor string = "1" // major version, always numeric
	gitMinor string = "6" // minor version, numeric possibly followed by "+"

	// semantic version, derived by build scripts (see
	// https://github.com/kubernetes/kubernetes/blob/master/docs/design/versioning.md
	// for a detailed discussion of this field)
	//
	// TODO: This field is still called "gitVersion" for legacy
	// reasons. For prerelease versions, the build metadata on the
	// semantic version is a git hash, but the version itself is no
	// longer the direct output of "git describe", but a slight
	// translation to be semver compliant.
	gitVersion   string = "v1.6.2+$Format:%h$"
	gitCommit    string = "$Format:%H$"    // sha1 from git, output of $(git rev-parse HEAD)
	gitTreeState string = "not a git tree" // state of git tree, either "clean" or "dirty"
	gitBranch    string = "unknown"

	buildDate string = "1970-01-01T00:00:00Z" // build date in ISO8601 format, output of $(date -u +'%Y-%m-%dT%H:%M:%SZ')
)

// Info contains versioning information.
// TODO: Add []string of api versions supported? It's still unclear
// how we'll want to distribute that information.
type Info struct {
	// Major        string `json:"major"`
	// Minor        string `json:"minor"`
	GitVersion   string `json:"gitVersion"`
	GitBranch    string `json:"gitBranch"`
	GitCommit    string `json:"gitCommit"`
	GitTreeState string `json:"gitTreeState"`
	BuildDate    string `json:"buildDate"`
	GoVersion    string `json:"goVersion"`
	Compiler     string `json:"compiler"`
	Platform     string `json:"platform"`
}

// String returns info as a human-friendly version string.
func (info Info) String() string {
	return info.GitVersion
}

// Get returns the overall codebase version. It's for detecting
// what code a binary was built from.
func Get() Info {
	// These variables typically come from -ldflags settings and in
	// their absence fallback to the settings in pkg/version/base.go
	return Info{
		// Major:        gitMajor,
		// Minor:        gitMinor,
		GitVersion:   gitVersion,
		GitBranch:    gitBranch,
		GitCommit:    gitCommit,
		GitTreeState: gitTreeState,
		BuildDate:    buildDate,
		GoVersion:    runtime.Version(),
		Compiler:     runtime.Compiler,
		Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}

func GetJsonString() string {
	bs, _ := json.MarshalIndent(Get(), "", "  ")
	return string(bs)
}

func shortDate(dateStr string) string {
	var buf strings.Builder
	for _, c := range dateStr {
		if c >= '0' && c <= '9' {
			buf.WriteRune(c)
		}
	}
	dateStr = buf.String()
	if strings.HasPrefix(dateStr, "20") {
		dateStr = dateStr[2:]
	}
	if len(dateStr) > 8 {
		dateStr = dateStr[:8]
	}
	return dateStr
}

func GetShortString() string {
	v := Get()
	return fmt.Sprintf("%s(%s%s)", v.GitBranch, v.GitCommit, shortDate(v.BuildDate))
}
