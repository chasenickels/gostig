package fs

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
)

type AtomicEditor struct {
	BackupDir string
}

func dataContains(data string, expectedString string) bool {
	return strings.Contains(data, expectedString)
}

func atomicEdit(filePath string, expectedText string) (err error) {

	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

}

func atomicWrite(filename string, data string, mode os.FileMode) (err error) {
	dir := filepath.Dir(filename)
	tmpFile, err := os.CreateTemp(dir, filepath.Base(filename)+".tmp")
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			os.Remove(tmpFile.Name())
		}
	}()

	if _, err = tmpFile.WriteString(data); err != nil {
		return err
	}
	return nil
}

func ensurePermission(filePath string, mode os.FileMode) (err error) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return err
	}
	if fileInfo.Mode().Perm() == mode {
		return nil
	}

	err = os.Chmod(filePath, mode)
	if err != nil {
		return err
	}
	fmt.Println("Mode change successful")
	return nil
}

func ensureUserOwner(filePath string, owningUser string) (err error) {
	fileInfo, err := os.Stat(filePath)
	syscallStat := fileInfo.Sys().(*syscall.Stat_t)
	u, err := user.Lookup(owningUser)
	if err != nil {
		return err
	}

	userInt, err := strconv.Atoi(u.Uid)
	if err != nil {
		return err
	}

	if userInt == int(syscallStat.Uid) {
		return nil
	}

	err = os.Chown(filePath, userInt, int(syscallStat.Gid))

	return nil
}

func ensureGroupOwner(filePath string, owningGroup string) (err error) {
	fileInfo, err := os.Stat(filePath)
	syscallStat := fileInfo.Sys().(*syscall.Stat_t)
	g, err := user.LookupGroup(owningGroup)
	if err != nil {
		return err
	}

	groupInt, err := strconv.Atoi(g.Gid)
	if err != nil {
		return err
	}

	if groupInt == int(syscallStat.Uid) {
		return nil
	}

	err = os.Chown(filePath, int(syscallStat.Uid), groupInt)
	if err != nil {
		return nil
	}
	return nil
}
