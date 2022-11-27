package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("NOPE")
	}
}

func run() {
	fmt.Printf("Executing %v as PID %d\n", os.Args[2:], os.Getgid())
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Create a sep proc -> new Namespace
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
	}

	execWithoutFail(cmd.Run())
}

func child() {
	fmt.Printf("Running %v as PID = %d\n", os.Args[2:], os.Getpid())

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	execWithoutFail(syscall.Sethostname([]byte("Container")))
	execWithoutFail(cmd.Run())
}

func execWithoutFail(err error) {
	if err != nil {
		panic(err)
	}
}
