package analyzer

import "strings"

func BasicAnalysis(data map[string]string) map[string]string {
	results := make(map[string]string)
	for k, v := range data {
		results[k] = strings.ToUpper(v)
	}
	return results
}
