package util

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os/exec"
)

// ensureLF replaces "\r\n" with "\n"
func ensureLF(content []byte) []byte {
	return bytes.Replace(content, []byte("\r\n"), []byte("\n"), -1)
}

// comapreFiles comapres two files at given paths
func CompareFiles(pathA string, pathB string) error {
	contentA, err := ioutil.ReadFile(pathA)
	if err != nil {
		return err
	}
	contentB, err := ioutil.ReadFile(pathB)
	if err != nil {
		return err
	}
	if !bytes.Equal(ensureLF(contentA), ensureLF(contentB)) {
		output := new(bytes.Buffer)
		output.WriteString(fmt.Sprintf("%q and %q are not equal!\n\n", pathA, pathB))

		diffPath, err := exec.LookPath("diff")
		if err != nil {
			return nil
		}
		diffCmd := exec.Command(diffPath, "-u", pathA, pathB)
		diffCmd.Stdout = output
		diffCmd.Stderr = output

		output.WriteString("$ diff -u " + pathA + " " + pathB + "\n")
		if err := diffCmd.Run(); err != nil {
			output.WriteString("\n" + err.Error())
		}
		return errors.New(output.String())
	}
	return nil
}
