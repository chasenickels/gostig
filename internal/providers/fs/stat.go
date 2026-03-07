package fs

import (
	"fmt"
	"os"
	"syscall"
)

type StatInformation struct {
	FilePath string
}

func (s StatInformation) SyscallStat() (*syscall.Stat_t, error) {
	fileInfo, err := os.Stat(s.FilePath)
	if err != nil {
		return nil, err
	}
	syscallStat, ok := fileInfo.Sys().(*syscall.Stat_t)
	if !ok {
		return nil, fmt.Errorf("unexpected fileinfo.Sys() type")
	}

	return syscallStat, nil
}

func (s StatInformation) Gid(filePath string) int {
	return int(s.SyscallStat.Gid)
}

func (s StatInformation) Uid() int {
	return int(s.SyscallStat.Uid)
}
