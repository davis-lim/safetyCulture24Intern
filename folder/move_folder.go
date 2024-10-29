package folder

import (
	"errors"
	"fmt"
	"strings"
)

func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	// Your code here...
	var src, dest *Folder

	// Find the src and dest folders
	for i, folder := range f.folders {
		if folder.Name == name {
			src = &f.folders[i]
		} 
		if folder.Name == dst {
			dest = &f.folders[i]
		}
	}

	// Check if src exists
	if src == nil {
		return nil, errors.New("error: source folder does not exist")
	}

	// Check if dest exists
	if dest == nil {
		return nil, errors.New("error: destination folder does not exist")
	}

	// Check if src and dest shares the same orgID
	if src.OrgId != dest.OrgId {
		return nil, errors.New("error: source and destination belongs in different organisations")
	}

	// Check if dest is a child of src
	if strings.HasPrefix(dest.Paths, src.Paths + ".") {
		return nil, errors.New("error: destination is a child of source")
	}

	// Check if src and dest are the same
	if src.Name == dest.Name {
		return nil, errors.New("error: source and destination are identical folders")
	}

	originalPath := src.Paths
	newPath := dest.Paths + "." + src.Name

	//Check for source path and update all of its descendants
	for i := range f.folders {
		fmt.Printf("Checking folder at index %d with path %s\n", i, f.folders[i].Paths)
		if strings.HasPrefix(f.folders[i].Paths, originalPath) {
			f.folders[i].Paths = strings.Replace(f.folders[i].Paths, originalPath, newPath, 1)
		} else {
			// print statement for debugging
			fmt.Printf("Prefix condition failed for folder at index %d\n", i)
		}
	}

	return f.folders, nil
}
