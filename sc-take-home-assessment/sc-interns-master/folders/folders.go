package folders

import (
	"github.com/gofrs/uuid"
)

/** Function seems redundant/repeating Fetchallfolders function
func GetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {
	// Removed undeclared variables

	// Removed old function and updated with a new one
	// Calling fetch folder function to get a slice of all the matching folders to specified ID
	// Note: r now holds a COPY of resFolder.
	r, err := FetchAllFoldersByOrgID(req.OrgID)
	if err != nil {
		return nil, err
	}
	// Slice folder
	var fp []*Folder
	// Putting the values of r into a new slice of folders
	for _, v := range r {
		// Create a new folder object to save the values from v.Prevents duplicate references
		NewFolder := Folder{
			Id:      v.Id,
			Name:    v.Name,
			OrgId:   v.OrgId,
			Deleted: v.Deleted,
		}
		// Then add this with a pointer to the new folder object to slice fp.
		fp = append(fp, &NewFolder)
		// Repeating this process gives me a full slice
	}

	return &FetchFolderResponse{Folders: fp}, nil
} **/

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
