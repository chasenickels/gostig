package cli

import (
	"fmt"
	"os"
)

const PROFILE_DIR = "/opt/gostig/content/profiles/stig/"

func ListProfiles() error {
	files, err := os.ReadDir(PROFILE_DIR)
	if err != nil {
		return err
	}
	for _, f := range files {
		fmt.Printf("Profile: %s\n", f.Name())
	}

	return nil
}
