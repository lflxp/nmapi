package pkg

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
)

func ExecCommand(cmd string) (string, error) {
	pipeline := exec.Command("/bin/sh", "-c", cmd)
	var out bytes.Buffer
	var stderr bytes.Buffer
	pipeline.Stdout = &out
	pipeline.Stderr = &stderr
	err := pipeline.Run()
	if err != nil {
		return stderr.String(), err
	}
	// fmt.Println(stderr.String())
	return out.String(), nil
}

func ReadFile(path string) ([]byte, error) {
	bytes, err := ioutil.ReadFile(path)
	defer DeleteFile(path)
	return bytes, err
}

func DeleteFile(path string) error {
	err := os.Remove(path)
	return err
}
