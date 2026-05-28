package utils

import "os"

// Exists reports whether anything exists at path (symlink, file, or dir).
func Exists(path string) bool {
	_, err := os.Lstat(path)
	return err == nil
}
