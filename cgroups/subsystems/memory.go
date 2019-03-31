package subsystems

import (
	"fmt"
	"io/ioutil"
	"log"
	"path"
)

func Set(content string) error {
	absolutePath := ""
	if absolutePath = FindAbsolutePath("memory"); absolutePath == "" {
		log.Printf("ERROR: absoutePath is empty!\n")
		return fmt.Errorf("ERROR: absoutePath is empty!\n")
	}
	if err := ioutil.WriteFile(path.Join(absolutePath, "memory.limit_in_bytes"), []byte(content),0644); err != nil {
		log.Printf("ERROR write content:%s.\n", content)
		return fmt.Errorf("ERROR write content:%s.\n", content)
	}
	return nil
}

func Apply(pid string) error {
	absolutePath := ""
	if absolutePath = FindAbsolutePath("memory"); absolutePath == "" {
		log.Printf("ERROR: absoutePath is empty!\n")
		return fmt.Errorf("ERROR: absoutePath is empty!\n")
	}
	log.Printf("Apply absolutePath:%s, taskPath:%s\n", absolutePath, path.Join(absolutePath, "tasks"))
	if err := ioutil.WriteFile(path.Join(absolutePath, "tasks"), []byte(pid),0644); err != nil {
		log.Printf("ERROR write pid:%s.\n", pid)
		return fmt.Errorf("ERROR write pid:%s.\n", pid)
	} else {
		log.Printf("err : %v\n", err)
	}
	return nil
}
