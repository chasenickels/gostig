package cli

import (
	"fmt"
	"gostig/internal/platform"
)

func determineProfile() (string, error) {

	osrelease, err := platform.ReadFile("/etc/os-release")
	if err != nil {
		return "", err
	}
	switch osrelease["ID"] {
	case "ubuntu18.04":
		return PROFILE_DIR + "/ubuntu18.yml", nil
	case "ubuntu20.04":
		return PROFILE_DIR + "/ubuntu20.yml", nil
	case "ubuntu22.04":
		return PROFILE_DIR + "/ubuntu22.yml", nil
	case "ubuntu24.04":
		return PROFILE_DIR + "/ubuntu24.yml", nil
	}
	return "", err
}

func ApplyProfile() {
	fmt.Println("whatever for now")
}

// TODO: Implement apply command
