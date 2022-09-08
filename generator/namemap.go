package generator

import (
	"encoding/json"
	"net/http"
)

func GetNameMap(url string) (map[string]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var nameMap map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&nameMap); err != nil {
		return nil, err
	}
	return nameMap, nil
}
