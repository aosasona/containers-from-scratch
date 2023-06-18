//go:build linux
// +build linux

package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

// docker     run image <cmd> <params> - the docker run command
// go run .   run       <cmd> <params> - the go run command for our container program

const ROOT_FS = "/home/parallels/alpine"

func main() {
	switch os.Args[1] {
	case "run", "RUN":
		run()
	case "child":
		child()
	default:
		panic("invalid command specified")
	}
}

func run() {
	fmt.Printf("[MASTER] Running %v as %d\n", os.Args[2:], os.Getpid())

	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
		Unshareflags: syscall.CLONE_NEWNS,
	}

	must(cmd.Run())
}

func child () {
	fmt.Printf("[CHILD] Running %v as %d\n", os.Args[2:], os.Getpid())

	syscall.Sethostname([]byte("container"))
	syscall.Chroot(ROOT_FS) 
	syscall.Chdir("/")
	syscall.Mount("proc", "proc", "proc", 0, "")


	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS,
	}

	must(cmd.Run())

	syscall.Unmount("proc", 0)
}
func must(err error) {
	if err != nil {
		panic(err)
	}
}
