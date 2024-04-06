package prototype

import "fmt"

type Folder struct {
	Children []Cloneable
	Name     string
}

func (f *Folder) Print(indentation string) {
	fmt.Println(indentation + f.Name)
	for _, child := range f.Children {
		child.Print(indentation + indentation)
	}
}

func (f *Folder) Clone() Cloneable {
	cloneFolder := &Folder{Name: f.Name + "_clone"}
	var tempChildren []Cloneable

	for _, i := range f.Children {
		copy := i.Clone()
		tempChildren = append(tempChildren, copy)
	}

	cloneFolder.Children = tempChildren
	return cloneFolder
}
