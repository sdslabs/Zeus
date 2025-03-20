package pivotroot

import (
	"os"
	"path/filepath"
	"syscall"
)

// mounts /proc to new root
//TODO: find out where it's source is, because I have absolutely no idea where this proc is coming from and I am 100% sure it carries my data from my device.
//TODO: ISSUE: hostname visible inside container, but root filesystem isn't, I am pinning this on the rexec package, but I am not sure.
func MountProc(newroot string) error {
	source := "proc"
	target := filepath.Join(newroot, "/proc")
	fstype := "proc"
	flags := 0
	data := ""

	os.MkdirAll(target, 0755)
	if err := syscall.Mount(
		source,
		target,
		fstype,
		uintptr(flags),
		data,
	); err != nil {
		return err
	}

	return nil
}

// unmounts /proc
func UnMountProc(newroot string) error {
	target := filepath.Join(newroot, "/proc")

	if err := syscall.Unmount(
		target,
		0,
	); err != nil {
		return err
	}

	return nil
}