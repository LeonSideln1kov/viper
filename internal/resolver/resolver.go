package resolver

import (
	"fmt"
	"strings"
	"github.com/LeonSideln1kov/viper/internal/pypi"
	"github.com/Masterminds/semver/v3"
)


func parseSpec(pkg string) (string, *semver.Constraints){
	operators := []string{"==", ">=", "<=", "!=", "~=", ">", "<"}

	for _, op := range operators {
		if strings.Contains(pkg, op) {
			parts := strings.SplitN(pkg, op, 2)
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


func ResolveVersion(pkg string) (string, error) {
    // Parse package spec (requests>=2.0.0)
    name, constraint := parseSpec(pkg)
    
    // Get PyPI versions
    info, err := pypi.GetPackageInfo(name)
	if err != nil {
		return "", fmt.Errorf("failed to get package info: %w", err)
	}
    
    // Find latest matching version
    var latest *semver.Version
    for v := range info.Releases {
        ver, err := semver.NewVersion(v)
        if err != nil || ver.Prerelease() != "" {
            continue // Skip invalid/pre-release versions
        }
        
        if latest != nil && constraint.Check(ver) && ver.GreaterThan(latest) {
            latest = ver
        }
    }

	if latest == nil {
        return "", fmt.Errorf("no valid versions found for %s", pkg)
    }
    
    return latest.String(), nil
}