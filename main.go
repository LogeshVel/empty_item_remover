package main

import (
	"flag"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

type pathStatus []string
type emptyFileMap map[string]pathStatus

// emptyFileMap will be the map with the key as the filename and the value as the string slice
// the string slice contains the file's Path as the first element and the Status as the 2nd ele
// Status can be of Deleted or Not deleted depends on the flag we set.

func main() {
	flagGivenPath := flag.String("p", "", "Path, from where the empty files need to find and proceed for deletion")
	flagDeleteFile := flag.Bool("f", false, "Provide this flag to delete the empty files")
	flagRecursiveDelete := flag.Bool("r", false, "Provide this flag to delete the empty Items recursively")
	// -p "path/to/chk" can be "." if the path is current dir.
	// if only the flag p is set then this program will look for the empty files in the currnt dir and prints in the output screen with status as Not deleted since the -f flag is not set
	//
	// to delete the empty files just set -f thats it. if p and f is set then the program will deleted the empty files in  the curnt dir
	// to delete the empty files recursively starting from the given path then set -r
	// -p "." -f -r this will make the program to search for the empty files in all the dir from thje curnt path and deletes it.
	// also the list of files and the path will be printed in the output screen.
	flag.Parse()

	givenPath, deleteFile, recursiveDelete := *flagGivenPath, *flagDeleteFile, *flagRecursiveDelete

	fmt.Printf("Delete File : %v\n", deleteFile)
	fmt.Printf("Recursive Delete : %v\n", recursiveDelete)
	if givenPath == "" {
		fmt.Println("Provide the Path to find the empty files")
		return
	}
	fmt.Printf("The Provided path to find the empty Files %q\n\n", givenPath)

	var emptyFileMapList = []emptyFileMap{}
	findAndRemoveEmptyItems(givenPath, &emptyFileMapList, deleteFile, recursiveDelete)

	// Print table
	checkLenAndPrintTable(emptyFileMapList, "File Name")
}

func checkLenAndPrintTable(mapToChkNPrint []emptyFileMap, fileOrDir string) {
	if len(mapToChkNPrint) == 0 {
		fmt.Printf("No empty %s is found\n", fileOrDir[:4])
	} else {
		printMapInTable(mapToChkNPrint, fileOrDir)
	}
}

func printMapInTable(mapToPrint []emptyFileMap, fileOrDir string) {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("S.No", fileOrDir, "Path", "Status")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	for item, listItem := range mapToPrint {
		for key, value := range listItem {
			tbl.AddRow(item+1, key, value[0]+"/", value[1])

		}
	}

	tbl.Print()
	fmt.Println()
}

func currentPathItemList(curntPath string) []fs.FileInfo {
	items, err := ioutil.ReadDir(curntPath)
	if err != nil {
		log.Fatalln(err)
	}
	return items
}

// gets the list of items in the current path and loop thorugh that list to check if the item size has 0 bytes and not a dir then del
// if the recursive del is enabled by setting -r flag then this func will check each dir for the same
func findAndRemoveEmptyItems(cpath string, fileMap *[]emptyFileMap, deleteF bool, recurDel bool) {
	files := currentPathItemList(cpath)
	for _, f := range files {
		if f.Size() == 0 && !f.IsDir() {
			fmt.Println(f.Name())
			emptyRemover(f, cpath, deleteF, fileMap)
		} else if recurDel && f.IsDir() {
			nonEmptyDirPath := path.Join([]string{cpath, f.Name()}...)
			findAndRemoveEmptyItems(nonEmptyDirPath, fileMap, deleteF, recurDel)
		}
	}
}

// calls the deleteItem func if the -d flag is set and updates the list with the Deleted status
// else it will just update the list with Not deleted status
func emptyRemover(itemObj fs.FileInfo, curntPath string, deleteFile bool, updateEmtFileMap *[]emptyFileMap) {
	*updateEmtFileMap = append(*updateEmtFileMap, emptyFileMap{itemObj.Name(): pathStatus{curntPath}})
	if deleteFile {
		deleteItem(itemObj, curntPath)
		(*updateEmtFileMap)[len(*updateEmtFileMap)-1][itemObj.Name()] = append((*updateEmtFileMap)[len(*updateEmtFileMap)-1][itemObj.Name()], "Deleted")
	} else {
		(*updateEmtFileMap)[len((*updateEmtFileMap))-1][itemObj.Name()] = append((*updateEmtFileMap)[len((*updateEmtFileMap))-1][itemObj.Name()], "Not Deleted")
	}

}

// this is the core function which will delete the item
func deleteItem(itemObject fs.FileInfo, givenPath string) {
	deletionPath := path.Join([]string{givenPath, itemObject.Name()}...)
	fmt.Printf("\nDeleting the File %q under the path %q\n", itemObject.Name(), givenPath)

	err := os.RemoveAll(deletionPath)
	if err != nil {
		log.Fatalf("errored while deleting the File %q under the path %q\n %v\n", itemObject.Name(), givenPath, err)
	}
	fmt.Printf("Deleted the File %q under the path %q\n\n", itemObject.Name(), givenPath)
}
