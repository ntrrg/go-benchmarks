package pb_test

import (
	"os/exec"
	"path/filepath"
	"testing"
)

func TestProtoGen(t *testing.T) {
	cmd := exec.Command("make", "proto")
	cmd.Dir = ".."
	if err := cmd.Run(); err != nil {
		t.Fatal(err)
	}

	pbfiles, err := filepath.Glob("*.pb.go")
	if err != nil {
		t.Fatal(err)
	}

	args := []string{"diff", "--quiet", "--"}
	cmd = exec.Command("git", append(args, pbfiles...)...)
	if err := cmd.Run(); err != nil {
		t.Fatal(err)
	}
}
