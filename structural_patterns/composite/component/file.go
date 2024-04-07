package component

import (
	"fmt"
	"strings"
)

type File struct {
	name       string
	folderName string
}

func NewFile(name string) *File {
	return &File{
		name: name,
	}
}

func (f *File) Search(keyword string) {
	if strings.ToLower(keyword) == strings.ToLower(f.name) {
		fmt.Printf("Searched for keyword %s. File %s Found in %s Folder.\n", keyword, f.name, f.folderName)
	}
}

func (f *File) ComponentedAdded(parentFolderName string) {
	f.folderName = parentFolderName
	fmt.Printf("File %s added into %s parent folder\n", f.name, parentFolderName)
}
