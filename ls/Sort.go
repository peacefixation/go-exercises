package main

import "os"

// sort by size
type bySize []os.FileInfo

func (files bySize) Len() int {
	return len(files)
}

func (files bySize) Swap(i, j int) {
	files[i], files[j] = files[j], files[i]
}

func (files bySize) Less(i, j int) bool {
	return files[i].Size() < files[j].Size()
}

// sort by time
type byTime []os.FileInfo

func (files byTime) Len() int {
	return len(files)
}

func (files byTime) Swap(i, j int) {
	files[i], files[j] = files[j], files[i]
}

func (files byTime) Less(i, j int) bool {
	return files[i].ModTime().Before(files[j].ModTime())
}

// sort by name
type byName []os.FileInfo

func (files byName) Len() int {
	return len(files)
}

func (files byName) Swap(i, j int) {
	files[i], files[j] = files[j], files[i]
}

func (files byName) Less(i, j int) bool {
	return files[i].Name() < files[j].Name()
}
