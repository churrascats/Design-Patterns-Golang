package component

import "fmt"

type Folder struct {
	name       string
	components []Component
}

func NewFolder(name string) *Folder {
	return &Folder{
		name: name,
	}
}

func (f *Folder) Search(keyword string) {
	fmt.Printf("Serching recursively for keyword %s in folder %s\n", keyword, f.name)

	for _, composite := range f.components {
		composite.Search(keyword)
	}
}

func (f *Folder) ComponentedAdded(parentFolderName string) {
	fmt.Printf("Folder %s added into %s parent folder\n", f.name, parentFolderName)
}

func (f *Folder) Add(components ...Component) {

	for _, component := range components {
		component.ComponentedAdded(f.name)
	}

	f.components = append(f.components, components...)
}
