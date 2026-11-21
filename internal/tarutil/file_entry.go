package tarutil

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var _ Entry = (*FileEntry)(nil)

type FileEntry struct {
	Path string
}

func NewEntryFromFilePath(path string) Entry {
	return FileEntry{
		Path: path,
	}
}

func NewEntryFromFilePaths(paths ...string) []Entry {
	var entries []Entry
	for _, path := range paths {
		entries = append(entries, NewEntryFromFilePath(path))
	}
	return entries
}

func (t FileEntry) writeEntry(tw lowLevelWriter) error {
	for _, seg := range strings.Split(filepath.ToSlash(t.Path), "/") {
		if seg == ".." {
			return fmt.Errorf("invalid file path")
		}
	}
	fi, err := os.Lstat(t.Path)
	if err != nil {
		return fmt.Errorf("unable to stat file %q: %w", t.Path, err)
	}
	return writeEntry(tw, t.Path, fi, func() (io.Reader, error) {
		return os.Open(t.Path)
	})
}
