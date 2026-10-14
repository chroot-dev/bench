package sandbox

import "golang.org/x/sys/unix"

// Restrict leaves dir readable, hides the rest of the file system and
// drops every system call outside the stdio and rpath promises.
func Restrict(dir string) error {
	if err := unix.Unveil(dir, "r"); err != nil {
		return err
	}
	if err := unix.UnveilBlock(); err != nil {
		return err
	}
	return unix.Pledge("stdio rpath", "")
}
