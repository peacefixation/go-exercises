package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"syscall"
	"time"
)

func usage() {
	fmt.Println("Usage: ls")
	os.Exit(1)
}

func main() {
	all := flag.Bool("a", false, "list all files including files that start with '.'")
	long := flag.Bool("l", false, "list in long format")
	flag.Parse()

	filepath := "."
	if flag.NArg() > 0 {
		filepath = flag.Arg(0)
	}

	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Open file error: %v\n", err)
	}

	files, err := file.Readdir(-1)
	file.Close()
	if err != nil {
		log.Fatalf("Read directory error: %v\n", err)
	}

	sort.Slice(files, func(i, j int) bool { return files[i].Name() < files[j].Name() })
	if *long {
		listFilesLong(files, *all)
	} else {
		listFiles(files, *all)
	}

}

func listFiles(fis []os.FileInfo, all bool) {
	for _, fi := range fis {
		if all || !strings.HasPrefix(fi.Name(), ".") {
			fmt.Println(fi.Name())
		}
	}
}

func listFilesLong(fis []os.FileInfo, all bool) {
	// copy the file info into a new array of long file info with some extra attributes
	// check the length of some columns as we go so we can space them correctly later
	// turns out I can't get usernames and group names working (from uid / gid) so this is moot
	var lfis []LongFileInfo
	for _, fi := range fis {
		if all || !strings.HasPrefix(fi.Name(), ".") {
			longFileInfo := new(LongFileInfo)
			longFileInfo.fileInfo = fi
			longFileInfo.uid = fi.Sys().(*syscall.Stat_t).Uid
			longFileInfo.gid = fi.Sys().(*syscall.Stat_t).Gid

			lfis = append(lfis, *longFileInfo)
		}
	}

	for _, lfi := range lfis {
		fmt.Printf("%s  ", lfi.fileInfo.Mode())
		fmt.Printf("%10d  ", lfi.uid)
		fmt.Printf("%10d  ", lfi.gid)
		fmt.Printf("%10d  ", lfi.fileInfo.Size())
		fmt.Printf("%s  ", formatDate(lfi.fileInfo.ModTime()))
		fmt.Printf("%s", lfi.fileInfo.Name())
		fmt.Println()
	}
}

func formatDate(t time.Time) string {
	oneYearAgo := time.Now().AddDate(-1, 0, 0)
	var formatted string
	if t.After(oneYearAgo) {
		formatted = t.Format("Jan _2 15:04")
	} else {
		formatted = t.Format("Jan _2  2006")
	}
	return formatted
}
