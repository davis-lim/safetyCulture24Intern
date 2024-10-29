package main

import (
	"fmt"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
)

func main() {
	orgID := uuid.FromStringOrNil(folder.DefaultOrgID)

	res := folder.GetAllFolders()

	// example usage
	folderDriver := folder.NewDriver(res)
	orgFolder := folderDriver.GetFoldersByOrgID(orgID)

	folder.PrettyPrint(res)
	
	// Print all folders that belongs to orgID
	fmt.Printf("\n Folders for orgID: %s", orgID)
	folder.PrettyPrint(orgFolder)

	// Print all children folders of "stunning-horridus"
	fmt.Printf("\n Get all child folders of %s", orgFolder[4].Name)
	childFolders, err := folderDriver.GetAllChildFolders(orgID, "stunning-horridus")
	if err != nil {
		fmt.Printf("\n Error getting children folders: %v", err)
		return
	}
	fmt.Printf("\n")
	folder.PrettyPrint(childFolders)

	// Move magnetic-sinister-six and its child under hip-stingray
	fmt.Printf("\n Move magnetic-sinister-six under hip-stingray")
	movedFolder, err := folderDriver.MoveFolder("magnetic-sinister-six", "hip-stingray")
	if err != nil {
		fmt.Printf("\n Error moving folder: %v", err)
		return
	}
	folder.PrettyPrint(movedFolder)
}
