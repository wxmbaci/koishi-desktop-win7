//go:build !windows

package pathutil

import (
	"fmt"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

func UserDataDir() (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", fmt.Errorf("failed to get home directory: %w", err)
	}

	return filepath.Join(home, "koishijs/koishi-desktop"), nil
}
