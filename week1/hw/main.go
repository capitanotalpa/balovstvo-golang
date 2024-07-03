package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

func dirTree(out io.Writer, path string, printFiles bool) error {
	return printDirRecursive(path, printFiles, out, "")
}

/*
Modifies array of dir entries, sorts it by filename
*/
func sortDirByNames(dir []os.DirEntry) {
	sort.Slice(dir, func(i, j int) bool {
		return dir[i].Name() < dir[j].Name()
	})
}

/*
Returns true or false whether dir is empty or not
*/

func isDirEmpty(dirName string) (bool, error) {
	dirFile, err := os.Open(dirName)
	if err != nil {
		return true, err
	}
	dir, err := dirFile.ReadDir(0)
	if err != nil {
		return true, err
	}
	return len(dir) == 0, nil
}

/*
Returns prefix for console prefix by cgiven riterias
*/
func getPrefix(isDir, isEmpty, isLast bool) string {
	if isLast {
		return "└───"
	}
	if isDir && isEmpty {
		return "├───"
	}
	return "├───"
}

/*
Filters directory from files
*/
func filterDir(dir []os.DirEntry) (ret []os.DirEntry) {
	for _, d := range dir {
		if d.IsDir() {
			ret = append(ret, d)
		}
	}
	return
}

func getFileSizeString(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		return "", err
	}
	if size := stat.Size(); size > 0 {
		return fmt.Sprintf("(%vb)", size), nil
	} else {
		return "(empty)", nil
	}
}

/*
Recursively prints directories and files if printFiles is true
*/
func printDirRecursive(dirName string, printFiles bool, out io.Writer, stringPrefix string) error {
	dirFile, err := os.Open(dirName)
	if err != nil {
		return err
	}
	defer dirFile.Close()

	dir, err := dirFile.ReadDir(0)
	if err != nil {
		return err
	}
	sortDirByNames(dir)
	if !printFiles {
		dir = filterDir(dir)
	}

	for i, dirEntry := range dir {
		path := dirName + "/" + dirEntry.Name()
		isLast := i == len(dir)-1
		if dirEntry.IsDir() {
			var isEmpty bool
			if !printFiles {
				isEmpty = true
			} else {
				isEmpty, err = isDirEmpty(path)
				if err != nil {
					return err
				}
			}

			pref := getPrefix(true, isEmpty, isLast)
			fmt.Fprintln(out, stringPrefix+pref+dirEntry.Name())
			var margin string
			if isLast {
				margin = "\t"
			} else {
				margin = "│\t"
			}
			printDirRecursive(path, printFiles, out, stringPrefix+margin)
		} else if printFiles {
			pref := getPrefix(false, false, isLast)
			sizeString, err := getFileSizeString(path)
			if err != nil {
				return err
			}
			fmt.Fprintln(out, stringPrefix+pref+dirEntry.Name()+" "+sizeString)
		}
	}

	return nil
}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}
