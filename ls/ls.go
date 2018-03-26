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

func main() {
	all := flag.Bool("a", false, "list all files including files that start with '.'")
	noSort := flag.Bool("f", false, "output is not sorted")
	human := flag.Bool("h", false, "show file size with unit suffixes (for byte, kilobyte, etc) to reduce the number of digits")
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

	if !*noSort {
		// sort the list
		if *sortSize {
			sortList(bySize(files), *reverse)
		} else if *sortTime {
			sortList(byTime(files), *reverse)
		} else {
			// default sort by name
			sortList(byName(files), *reverse)
		}
	}

	// print the list
	if *long {
		listFilesLong(files, *all, *inode, *human)
	} else {
		listFiles(files, *all, *inode)
	}

}

// sort an array that implements the sort.Interface
func sortList(toSort sort.Interface, reverse bool) {
	if reverse {
		sort.Sort(sort.Reverse(toSort))
	} else {
		sort.Sort(toSort)
	}
}

// print the list of os.FileInfo
func listFiles(fis []os.FileInfo, all, showInode bool) {
	for _, fi := range fis {
		if showInode {
			fmt.Printf("%8d ", getInode(fi))
		}

		if all || !strings.HasPrefix(fi.Name(), ".") {
			fmt.Println(fi.Name())
		}
	}
}

// print the list of os.FileInfo and show extra information
func listFilesLong(fis []os.FileInfo, all, showInode, human bool) {
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
		if showInode {
			fmt.Printf("%8d ", getInode(lfi.fileInfo))
		}

		fmt.Printf("%s  ", lfi.fileInfo.Mode())
		fmt.Printf("%10d  ", lfi.uid)
		fmt.Printf("%10d  ", lfi.gid)
		if human {
			fmt.Printf("%10s  ", convertUnits(lfi.fileInfo.Size()))
		} else {
			fmt.Printf("%10d  ", lfi.fileInfo.Size())
		}
		fmt.Printf("%s  ", formatDate(lfi.fileInfo.ModTime()))
		fmt.Printf("%s", lfi.fileInfo.Name())
		fmt.Println()
	}
}

// format a time.Time, show the year if it's more than 1 year old
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

// get the inode for an os.FileInfo
func getInode(fi os.FileInfo) uint64 {
	stat, ok := fi.Sys().(*syscall.Stat_t)
	if !ok {
		log.Fatalf("Not a syscall.Stat_t: %v\n", fi.Sys())
	}

	return stat.Ino
}

// convert a file size to larger units to reduce the number of digits
func convertUnits(size int64) string {
	magnitude := 0
	units := []string{"B", "K", "M", "G", "T"}
	for size > 1024 && magnitude < len(units)-1 {
		size /= 1024
		magnitude++
	}

	return fmt.Sprintf("%d%s", size, units[magnitude])
}
