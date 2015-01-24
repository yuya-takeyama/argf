package argf

import (
	"bytes"
	"os/exec"
	"testing"
)

func TestArgf(t *testing.T) {
	cmd := exec.Command("go", "run", "testcmd/main.go")
	cmd.Stdin = bytes.NewBufferString("foo\nbar\nbaz\n")
	stdout := new(bytes.Buffer)
	cmd.Stdout = stdout

	err := cmd.Run()
	if err != nil {
		t.Fatal("Failed to run command to test")
	}

	if stdout.String() != "foo\nbar\nbaz\n" {
		t.Fatal("Not matched")
	}
}

func TestWithOneFile(t *testing.T) {
	cmd := exec.Command("go", "run", "testcmd/main.go", "testcmd/file1")
	stdout := new(bytes.Buffer)
	cmd.Stdout = stdout

	err := cmd.Run()
	if err != nil {
		t.Fatal("Failed to run command to test")
	}

	if stdout.String() != "foo\nbar\nbaz\n" {
		t.Fatal("Not matched")
	}
}

func TestWithTwoFiles(t *testing.T) {
	cmd := exec.Command("go", "run", "testcmd/main.go", "testcmd/file1", "testcmd/file2")
	stdout := new(bytes.Buffer)
	cmd.Stdout = stdout

	err := cmd.Run()
	if err != nil {
		t.Fatal("Failed to run command to test")
	}

	if stdout.String() != "foo\nbar\nbaz\nhoge\nfuga\npiyo\n" {
		t.Fatal("Not matched")
	}
}

func TestWithTwoFilesAndStdin(t *testing.T) {
	cmd := exec.Command("go", "run", "testcmd/main.go", "testcmd/file1", "testcmd/file2")
	cmd.Stdin = bytes.NewBufferString("this will be not read\n")
	stdout := new(bytes.Buffer)
	cmd.Stdout = stdout

	err := cmd.Run()
	if err != nil {
		t.Fatal("Failed to run command to test")
	}

	if stdout.String() != "foo\nbar\nbaz\nhoge\nfuga\npiyo\n" {
		t.Fatal("Not matched")
	}
}
