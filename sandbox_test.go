package sandbox

import (
	"errors"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
)

// pledge(2) and unveil(2) last for the whole process, so the test runs
// itself again as a child and restricts only the child.
func TestRestrict(t *testing.T) {
	if runtime.GOOS != "openbsd" {
		t.Skip("pledge and unveil are OpenBSD only")
	}
	if dir := os.Getenv("SANDBOX_DIR"); dir != "" {
		restricted(t, dir)
		return
	}
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "ok"), []byte("ok\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	cmd := exec.Command(os.Args[0], "-test.run=^TestRestrict$", "-test.v")
	cmd.Env = append(os.Environ(), "SANDBOX_DIR="+dir)
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("restricted child: %v\n%s", err, out)
	}
}

func restricted(t *testing.T, dir string) {
	if err := Restrict(dir); err != nil {
		t.Fatal(err)
	}
	if _, err := os.ReadFile(filepath.Join(dir, "ok")); err != nil {
		t.Errorf("read inside dir: %v", err)
	}
	// unveil(2): outside dir, files do not exist
	if _, err := os.Open("/etc/passwd"); !errors.Is(err, fs.ErrNotExist) {
		t.Errorf("open /etc/passwd: got %v, want %v", err, fs.ErrNotExist)
	}
}
