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
	inode := flag.Bool("i", false, "for each file, print the inode number")
	long := flag.Bool("l", false, "list in long format")
	reverse := flag.Bool("r", false, "reverse the order of the sort")
	sortSize := flag.Bool("S", false, "sort by size")
	sortTime := flag.Bool("t", false, "sort by time modified")
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

	// sort the list
	if *sortSize && !*reverse {
		sort.Sort(bySize(files))
	} else if *sortSize && *reverse {
		sort.Sort(sort.Reverse(bySize(files)))
	} else if *sortTime && !*reverse {
		sort.Sort(byTime(files))
	} else if *sortTime && *reverse {
		sort.Sort(sort.Reverse(byTime(files)))
	} else if !*reverse {
		sort.Sort(byName(files))
	} else {
		sort.Sort(sort.Reverse(byName(files)))
	}

	// print the list
	if *long {
		listFilesLong(files, *all, *inode)
	} else {
		listFiles(files, *all, *inode)
	}

}

func listFiles(fis []os.FileInfo, all, inode bool) {
	for _, fi := range fis {
		if inode {
			fmt.Printf("%8d ", getInode(fi))
		}

		if all || !strings.HasPrefix(fi.Name(), ".") {
			fmt.Println(fi.Name())
		}
	}
}

func listFilesLong(fis []os.FileInfo, all, inode bool) {
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
		if inode {
			fmt.Printf("%8d ", getInode(lfi.fileInfo))
		}

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

func getInode(fi os.FileInfo) uint64 {
	stat, ok := fi.Sys().(*syscall.Stat_t)
	if !ok {
		log.Fatalf("Not a syscall.Stat_t: %v\n", fi.Sys())
	}

	return stat.Ino
}
