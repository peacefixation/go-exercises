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
	recursive := flag.Bool("R", false, "recursively list subdirectories")
	sortSize := flag.Bool("S", false, "sort by size")
	sortTime := flag.Bool("t", false, "sort by time modified")
	flag.Parse()

	path := "."
	if flag.NArg() > 0 {
		path = strings.TrimSuffix(flag.Arg(0), "/")
	}

	files := readPath(path)

	// print the list
	listFiles(path, files, *all, *long, *inode, *human, *noSort, *sortSize, *sortTime, *reverse, *recursive)
}

// read a directory path and return the files it contains
func readPath(path string) []os.FileInfo {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Open file error: %v\n", err)
	}

	files, err := file.Readdir(-1)
	file.Close()
	if err != nil {
		log.Fatalf("Read directory error: %v\n", err)
	}

	return files
}

// sort files according to sort type
func sortFiles(files []os.FileInfo, sortSize, sortTime, reverse bool) {
	// sort the list
	if sortSize {
		sortList(bySize(files), reverse)
	} else if sortTime {
		sortList(byTime(files), reverse)
	} else {
		// default sort by name
		sortList(byName(files), reverse)
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

// print the list of files
func listFiles(basepath string, files []os.FileInfo, all, long, inode, human, noSort, sortSize, sortTime,
	reverse, recursive bool) {

	if !noSort {
		sortFiles(files, sortSize, sortTime, reverse)
	}

	// print the files
	for _, file := range files {
		if all || !strings.HasPrefix(file.Name(), ".") {
			printFile(file, long, inode, human)
		}
	}

	if recursive {
		// print the files in each subdirectory
		for _, file := range files {
			if file.Mode().IsDir() {
				dirpath := basepath + "/" + file.Name()
				subDirFiles := readPath(dirpath)
				printPath(dirpath)
				listFiles(dirpath, subDirFiles, all, long, inode, human, noSort, sortSize, sortTime, reverse, recursive)
			}
		}
	}
}

// print a file
func printFile(file os.FileInfo, long, inode, human bool) {
	if inode {
		fmt.Printf("%8d ", getInode(file))
	}

	if long {
		fmt.Printf("%s  ", file.Mode())
		fmt.Printf("%10d  ", file.Sys().(*syscall.Stat_t).Uid)
		fmt.Printf("%10d  ", file.Sys().(*syscall.Stat_t).Gid)
		if human {
			fmt.Printf("%10s  ", convertUnits(file.Size()))
		} else {
			fmt.Printf("%10d  ", file.Size())
		}
		fmt.Printf("%s  ", formatDate(file.ModTime()))
	}

	fmt.Printf("%s", file.Name())
	fmt.Println()
}

// print a path
func printPath(path string) {
	fmt.Println("\n" + path + ":")
}

// format the date, show the year if it's more than 1 year old
func formatDate(date time.Time) string {
	oneYearAgo := time.Now().AddDate(-1, 0, 0)
	var formatted string
	if date.After(oneYearAgo) {
		formatted = date.Format("Jan _2 15:04")
	} else {
		formatted = date.Format("Jan _2  2006")
	}
	return formatted
}

// get the inode for a file
func getInode(file os.FileInfo) uint64 {
	stat, ok := file.Sys().(*syscall.Stat_t)
	if !ok {
		log.Fatalf("Not a syscall.Stat_t: %v\n", file.Sys())
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
