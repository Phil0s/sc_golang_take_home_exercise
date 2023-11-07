package main

import (
	"fmt"

	"github.com/georgechieng-sc/interns-2022/folders"
	"github.com/gofrs/uuid"
)

// Purpose of this code is to fetch folders from sample.json base of matchign orgID and return them

func main() {
	// Getting example OrgID
	req := &folders.FetchFolderRequest{
		OrgID: uuid.FromStringOrNil(folders.TestOrgID),
	}
	// Calling fetch folder function, returned fetchfolderresponse
	// res, err := folders.GetAllFolders(req)
	// if err != nil {
	// 	fmt.Printf("%v", err)
	// 	return
	// }
	res, err := folders.FetchAllFoldersByOrgID(req.OrgID)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	// Format fetchfolderresponse into something readable
	//fmt.Println(res)
	folders.PrettyPrint(res)
}
