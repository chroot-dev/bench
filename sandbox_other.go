//go:build !openbsd

package sandbox

import "errors"

// Restrict needs pledge(2) and unveil(2), which only OpenBSD has.
func Restrict(dir string) error {
	return errors.ErrUnsupported
}
