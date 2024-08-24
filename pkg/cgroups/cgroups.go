package cgroups

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/sdslabs/Zeus/pkg/utils"
)


const (
	cgroupPath  = "/sys/fs/cgroup"
)

// Cgroups sets up cgroups for the container
func Cgroups(memoryLimit string, maxpid string) (string, error) {

	groupName := fmt.Sprintf("runz_%s", utils.GenerateUUID())
	// fmt.Println("UUID: ", groupName)

	cgroupDir := filepath.Join(cgroupPath, groupName)

	if _, err := os.Stat(cgroupDir); os.IsNotExist(err) {
		if err := os.Mkdir(cgroupDir, 0755); err != nil {
			return "", fmt.Errorf("failed to create cgroup: %v", err)
		}
	}

	// sets maximum memory
	if err := os.WriteFile(filepath.Join(cgroupDir, "memory.max"), []byte(memoryLimit), 0644); err != nil {
		return "", fmt.Errorf("failed to set memory limit: %v", err)
	}

	// sets maximum pid
	if err := os.WriteFile(filepath.Join(cgroupDir, "pids.max"), []byte(maxpid), 0644); err != nil {
		return "", fmt.Errorf("failed to set pid limit: %v", err)
	}

	// sets PID of the process
	if err := os.WriteFile(filepath.Join(cgroupDir, "cgroup.procs"), []byte(strconv.Itoa(os.Getpid())), 0644); err != nil {
		return "", fmt.Errorf("failed to add process to cgroup: %v", err)
	}

	return cgroupDir, nil
}
