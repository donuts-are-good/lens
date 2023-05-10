package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/tabwriter"
	"time"
)

func main() {
	var dir string
	var err error

	if len(os.Args) > 1 {
		dir = os.Args[1]
	} else {
		dir, err = os.Getwd()
		if err != nil {
			fmt.Println("Error getting current directory:", err)
			return
		}
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error reading directory contents:", err)
		return
	}

	sort.Slice(entries, func(i, j int) bool {
		return strings.ToLower(entries[i].Name()) < strings.ToLower(entries[j].Name())
	})

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	fmt.Printf("dir: %s\n", dir)

	for _, entry := range entries {
		info, _ := entry.Info()
		fileType := ""
		fileExt := strings.ToUpper(filepath.Ext(info.Name()))

		isExecutable := info.Mode().Perm()&0111 != 0

		switch {
		case entry.IsDir():
			fileType = "📁 DIR"
		case info.Mode()&os.ModeSymlink != 0:
			fileType = "🔗 LINK"
		case info.Mode().IsRegular() && isExecutable:
			fileType = "⚙️ EXE"
			entry = &customDirEntry{
				entry: entry,
				name:  "*" + entry.Name(),
			}
		default:
			if len(fileExt) > 0 {
				fileType = "📄 " + fileExt[1:]
			} else {
				fileType = "❓ ..?"
			}
		}

		fmt.Fprintf(w, "%s\t%s\t%d\t%s\n", entry.Name(), fileType, info.Size(), info.ModTime().Format(time.RFC1123))
	}

	w.Flush()
}

type customDirEntry struct {
	entry fs.DirEntry
	name  string
}

func (c *customDirEntry) Name() string {
	return c.name
}

func (c *customDirEntry) IsDir() bool {
	return c.entry.IsDir()
}

func (c *customDirEntry) Type() fs.FileMode {
	return c.entry.Type()
}

func (c *customDirEntry) Info() (fs.FileInfo, error) {
	return c.entry.Info()
}
