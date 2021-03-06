package command

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"syscall"
)

func Init(command string)  {
	command = readFromPipe()

	pwd, err := os.Getwd()
	if err != nil {
		log.Printf("ERROR: get pwd error!\n")
		return
	}
	log.Printf("current path: %s.\n", pwd)
	pivotRoot(pwd)

	defaultMountFlags := syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV
	syscall.Mount("proc", "/proc", "proc", uintptr(defaultMountFlags), "")

	if err := syscall.Exec(command, []string{command}, os.Environ()); err != nil {
		log.Printf("syscall.Exec err: %v\n", err)
		log.Fatal(err)
	}
}

func readFromPipe() string {
	reader := os.NewFile(uintptr(3), "pipe")
	command, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Printf("reader.Read(buf) error:%v\n", err)
		return ""
	}
	return string(command)
}

func pivotRoot(root string) error {
	if err := syscall.Mount(root, root, "bind", syscall.MS_BIND|syscall.MS_REC, ""); err != nil {
		return fmt.Errorf("Mount rootfs to itself error: %v", err)
	}
	pivotDir := filepath.Join(root, ".pivot_root")
	if err := os.Mkdir(pivotDir, 0777); err != nil {
		return err
	}
	if err := syscall.PivotRoot(root, pivotDir); err != nil {
		return fmt.Errorf("pivot_root %v", err)
	}
	if err := syscall.Chdir("/"); err != nil {
		return fmt.Errorf("chdir / %v", err)
	}

	pivotDir = filepath.Join("/", ".pivot_root")
	if err := syscall.Unmount(pivotDir, syscall.MNT_DETACH); err != nil {
		return fmt.Errorf("unmount pivot_root dir %v", err)
	}
	return os.Remove(pivotDir)
}
