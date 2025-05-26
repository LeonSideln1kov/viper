package pypi


import (
	"fmt"
	"encoding/json"
	"net/http"
)


type PackageInfo struct {
    Releases map[string][]struct {
        URL string `json:"url"`
    } `json:"releases"`
}

func GetPackageInfo(name string) (*PackageInfo, error) {
	resp, err := http.Get(fmt.Sprintf("https://pypi.org/pypi/%s/json", name))
	if err != nil {
		fmt.Printf("package %s not found", name)
		return nil, err
	}
	defer resp.Body.Close()

	var info PackageInfo
	if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
		return nil, err
	}
	return &info, nil
}
