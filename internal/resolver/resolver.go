package resolver

import (
	"fmt"
	"strings"
	"github.com/LeonSideln1kov/viper/internal/pypi"
	"github.com/Masterminds/semver/v3"
	"sort"
)


func parseSpec(pkg string) (string, *semver.Constraints){
	operators := []string{"==", ">=", "<=", "!=", "~=", ">", "<"}

	// Redesign cause in some cases ppl use ~= instead of ~ 
	// So its better to support both variants 
	for _, op := range operators {
		if strings.Contains(pkg, op) {
			parts := strings.SplitN(pkg, op, 2)
			if op == "~=" {
			op = "~"
			}
			constraint, err := semver.NewConstraint(fmt.Sprintf("%s %s", op, strings.TrimSpace(parts[1])))
			if err != nil {
				return pkg, nil
			}
			return strings.TrimSpace(parts[0]), constraint
		}
	}

	// If no version specified
	return strings.TrimSpace(pkg), nil
}


func ResolveVersion(pkg string) (name string, version string, err error) {
    baseName, constraint := parseSpec(pkg)
    info, err := pypi.GetPackageInfo(baseName)
    if err != nil {
        return "", "", fmt.Errorf("PyPI API error: %w", err)
    }

    var validVersions []*semver.Version
    for v := range info.Releases {
        ver, err := semver.NewVersion(v)
        if err != nil || ver.Prerelease() != "" {
            continue
        }
        if constraint == nil || constraint.Check(ver) {
            validVersions = append(validVersions, ver)
        }
    }

    if len(validVersions) == 0 {
        return "", "", fmt.Errorf("no valid versions for %s", pkg)
    }

    sort.Sort(sort.Reverse(semver.Collection(validVersions)))
    return baseName, validVersions[0].String(), nil
}