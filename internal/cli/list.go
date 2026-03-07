package cli

import (
	"fmt"
	"os"
)

func ListProfiles() error {
	files, err := os.ReadDir("/opt/stigctl/content/profiles/stig/")
	if err != nil {
		return err
	}
	for _, f := range files {
		fmt.Printf("Profile: %s\n", f.Name())
	}

	return nil
}
