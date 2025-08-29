package collector

import "fmt"

func CollectBasic(name string) map[string]string {
	// Placeholder: integrate APIs later
	return map[string]string{
		"Name":    name,
		"Profile": fmt.Sprintf("https://gov.example/%s", name),
		"Notes":   "Mocked government data",
	}
}
