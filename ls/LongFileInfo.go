package main

/*
This is not used. Need to learn how to perform uid/gid lookups cross platform.
*/

import (
	"fmt"
	"os"
	"os/user"
)

// LongFileInfo for listing files with extra information
type LongFileInfo struct {
	fileInfo os.FileInfo
	uid      uint32
	gid      uint32
}

// GetUsername looks up the user by uid and returns the username
func (l *LongFileInfo) GetUsername() string {
	usr, err := user.LookupId(string(l.uid))
	fmt.Println(err)
	if err != nil {
		return ""
	}
	return usr.Username
}

// GetGroupName looks up the Group by gid and returns the group name
func (l *LongFileInfo) GetGroupName() string {
	group, err := user.LookupGroupId(string(l.gid))
	if err != nil {
		return ""
	}
	return group.Name
}
