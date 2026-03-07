package content

import (
	"os"

	"gopkg.in/yaml.v3"
)

func loadYaml(profileFile string) (Profile, error) {
	yamlFile, err := os.ReadFile(profileFile)
	if err != nil {
		return Profile{}, err
	}
	var profile Profile

	err = yaml.Unmarshal(yamlFile, &profile)
	if err != nil {
		return Profile{}, err
	}
	return profile, nil
}
