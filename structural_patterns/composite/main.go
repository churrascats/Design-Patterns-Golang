package main

import "composite/component"

func main() {
	file1 := component.NewFile("myFile1")
	file2 := component.NewFile("myFile2")
	file3 := component.NewFile("myFile3")
	file4 := component.NewFile("myFile4")
	file5 := component.NewFile("myFile5")

	folder1 := component.NewFolder("myFolder1")
	folder2 := component.NewFolder("myFolder2")
	folder3 := component.NewFolder("myFolder3")

	folder1.Add(file1, file2)
	folder2.Add(file3, folder1)
	folder3.Add(file4)

	parentFolder := component.NewFolder("mainFolder")
	parentFolder.Add(folder1, folder2, folder3, file5)

	parentFolder.Search("myFile5")
}
