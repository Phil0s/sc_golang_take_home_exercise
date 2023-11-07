package main

import (
	"fmt"

	"github.com/georgechieng-sc/interns-2022/folders"
	"github.com/gofrs/uuid"
)

// Purpose of this code is to fetch folders from sample.json based of matchign orgID and return them
func main() {
	// Getting example OrgID
	req := &folders.FetchFolderRequest{
		OrgID: uuid.FromStringOrNil(folders.TestOrgID),
	}
	res, err := folders.FetchAllFoldersByOrgID(req.OrgID)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	// Format fetchfolderresponse into something readable
	folders.PrettyPrint(res)
}
