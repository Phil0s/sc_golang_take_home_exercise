package folders

import (
	"github.com/gofrs/uuid"
)

//Getallfolder function seems redundant. Just reassigning the folders to a new slice.

// This function sorts out all the matching folders from the entire sample data
func FetchAllFoldersByOrgID(orgID uuid.UUID) ([]*Folder, error) {
	// Removed old function and added updated one that creates a new folder object when appending to prevent duplicate ref
	// Recall GetSampleData simply returns a slice of all folder objects in sample.json, does not sort yet.
	folders, err := GetSampleData()
	if err != nil {
		return nil, err
	}
	// New slice to house all the matching folders after sorting logic (below)
	resFolder := []*Folder{}
	// Sorting commences
	for _, folder := range folders {
		if folder.OrgId == orgID {
			// Create a new Folder object to save the values from the original. Prevents duplicate references
			NewFolder := Folder{
				Id:      folder.Id,
				Name:    folder.Name,
				OrgId:   folder.OrgId,
				Deleted: folder.Deleted,
			}
			resFolder = append(resFolder, &NewFolder)
		}
	}
	return resFolder, nil
}
