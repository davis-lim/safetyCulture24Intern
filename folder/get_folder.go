package folder

import (
	"strings"
	"errors"
	"github.com/gofrs/uuid"
)

func GetAllFolders() []Folder {
	return GetSampleData()
}

func (f *driver) GetFoldersByOrgID(orgID uuid.UUID) []Folder {
	
	folders := f.folders

	res := []Folder{}
	for _, f := range folders {
		if f.OrgId == orgID {
			res = append(res, f)
		}
	}

	return res

}

func (f *driver) OrgIdValid(orgID uuid.UUID) bool {
	folders := f.folders
	for _, folder := range folders {
		if folder.OrgId == orgID {
			return true
		}
	}
	return false
}

func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) ([]Folder, error) {
	// Your code here...

	// Check for valid orgID
	if !f.OrgIdValid(orgID) {
		return nil, errors.New("error: orgID does not contain any folders")
	}

	// Check if folder exists
	var isFolder *Folder
	for i, folder := range f.folders {
		if folder.Name == name {
			isFolder = &f.folders[i]
		} 
	}

	if isFolder == nil {
		return nil, errors.New("error: Folder does not exist")
	}

	// Get all folders from orgID
	currOrgFolders := f.GetFoldersByOrgID(orgID)
	childFolders := []Folder{}
	
	folderExists := false
	var rootPath string
	
	// Check if folder exists within currOrgFolders by matching last segment of each path
	for _, folder:= range currOrgFolders {
		segments := strings.Split(folder.Paths, ".")
		
		lastSegment := segments[len(segments) - 1]

		if lastSegment == name {
			rootPath = folder.Paths
			folderExists = true
			break
		}
	}
	
	// Error checking for invalid path
	if !folderExists {
		return nil, errors.New("error: Folder does not exist in the specified organization")
	}
	
	//Create a prefix that all child folders should contain
	prefix := rootPath + "."

	// Checks for folders that contains the prefix and append them to childFolders
	for _, folder := range currOrgFolders {
		if strings.HasPrefix(folder.Paths, prefix) {
			childFolders = append(childFolders, folder)
		}
	}

	return childFolders, nil
}
