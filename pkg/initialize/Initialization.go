package initialize

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"github.com/docker/docker/pkg/reexec"

	"github.com/sdslabs/Zeus/pkg/cgroups"
	"github.com/sdslabs/Zeus/pkg/pivotroot"
)

// runs reexec on runzInit function. Reexec in short just finds the temporary process variable and executes it as a child process.
// This makes it easier to run the same process in a new namespace and doing cgroup operations on it before it actually starts
func init() {
	reexec.Register("runz-shell", runzInit)
	if reexec.Init() {
		os.Exit(0)
	}
}

// runzInit is the function that is going to be run in the new namespace
// it is a child process that has it's own PID
// runzInit function is called the **runz-shell** while executing in the new namespace
func runzInit() {

	newrootPath := os.Args[1]
	// sets up cgroups
	if _, err := cgroups.Cgroups(os.Args[2], os.Args[3]); err != nil {
		fmt.Printf("Error setting up cgroups in runzInit- %s\n", err)
		os.Exit(1)
	}

	// mounts /proc
	if err := pivotroot.MountProc(newrootPath); err != nil {
		fmt.Printf("Error mounting /proc - %s\n", err)
		os.Exit(1)
	}

	// runs pivot_root
	if err := pivotroot.RootPivoter(newrootPath); err != nil {
		fmt.Printf("Error running pivot_root - %s\n", err)
		os.Exit(1)
	}

	// runs runz inside the container that creates a shell child process.
	runzRun()

}

func runzRun() {

	// starts the container into /bin/sh shell. Can be changed into other programs as well.
	cmd := exec.Command("/bin/sh")

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// This is the shell prompt
	PS1 := ("PS1=container@runz $PWD # ")
	cmd.Env = []string{PS1}

	// runs the /bin/sh command
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error running the /bin/sh command - %s\n", err)
		os.Exit(1)
	}

	// unmounts /proc
	if err := pivotroot.UnMountProc(cmd.Dir + "/"); err != nil {	//TODO: There is a slight chance this is wrong
		fmt.Printf("Error unmounting /proc - %s\n", err)
		os.Exit(1)
	}
}

func Begin(newrootPath string, memoryLimit string, maxPid string) {

	// reexec.Command is used to run the runzInit function in a new namespace
	cmd := reexec.Command("runz-shell", newrootPath, memoryLimit, maxPid)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// sets the new namespace flags and makes the container into root (not real root, just the permissions. This is an unprevileged container)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWNS |
			syscall.CLONE_NEWIPC |
			syscall.CLONE_NEWPID |
			syscall.CLONE_NEWNET |
			syscall.CLONE_NEWUSER |
			// syscall.CLONE_NEWCGROUP | //TODO: find out why adding this makes the program fail
			syscall.CLONE_NEWUTS,

		UidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 0,
				HostID:      os.Getuid(),
				Size:        1,
			},
		},
		GidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 0,
				HostID:      os.Getgid(),
				Size:        1,
			},
		},
	}

	// runs the reexec.Command
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error starting the reexec.Command - %s\n", err)
		os.Exit(1)
	}

	// The following code is not needed as the reexec.Command is already running the runzInit function
	// It is going to be used in the future to do operations during the runtime of the container

	// if err := cmd.Start(); err != nil {
	// 	fmt.Printf("Error starting the reexec.Command - %s\n", err)
	// 	os.Exit(1)
	// }

	// if err := cmd.Wait(); err != nil {
	// 	fmt.Printf("Error waiting for reexec.Command - %s\n", err)
	// 	os.Exit(1)
	// }
}
