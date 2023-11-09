package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/georgechieng-sc/interns-2022/folders"
	"github.com/gofrs/uuid"
)

// // Stage one
// // Purpose of this code is to fetch folders from sample.json based of matchign orgID and return them
// func main() {
// 	// Getting example OrgID
// 	req := &folders.FetchFolderRequest{
// 		OrgID: uuid.FromStringOrNil(folders.TestOrgID),
// 	}
// 	res, err := folders.FetchAllFoldersByOrgID(req.OrgID)
// 	if err != nil {
// 		fmt.Printf("%v", err)
// 		return
// 	}
// 	// Format fetchfolderresponse into something readable
// 	folders.PrettyPrint(res)
// }

// Stage two (Pagination)
// Get desired orgID from user
func getInput(question string, r *bufio.Reader) (string, error) {
	fmt.Print(question)
	input, err := r.ReadString('\n')
	if len(input) != 38 {
		fmt.Println("input must be valid uuid format (36 characters long)")
		// Recursive loop.
		getInput(question, r)
	}
	return strings.TrimSpace(input), err
}

func askQuestion() string {
	reader := bufio.NewReader(os.Stdin)
	// Type this out for pagination test: c1556e17-b7c0-45a3-a6ae-9546248fb17a
	// or this for smaller sample size and odd number: 4212d618-66ff-468a-862d-ea49fef5e183
	// and this for small and even: 4212d618-66ff-468a-862d-ea49fef5e180
	id, _ := getInput("Specify organisation ID: ", reader)
	return id
}

func display(folders []*folders.Folder, startIdx, endIdx int) {
	fmt.Println("Displaying folders:")
	for i := startIdx; i < endIdx; i++ {
		// Incase odds number of folders and there it tries to search for the index of a sixth one, causing problems
		if i < len(folders) {
			// Spacing for easier readability.
			fmt.Printf("   ID: %s, Name: %s, Deleted: %v\n", folders[i].Id, folders[i].Name, folders[i].Deleted)
		}
	}
}

func main() {
	// Asking question
	id := askQuestion()
	orgID := uuid.FromStringOrNil(id)
	res, err := folders.FetchAllFoldersByOrgID(orgID)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	// Displaying
	pageSize := 2
	startIdx := 0
	endIdx := startIdx + pageSize
	for {
		display(res, startIdx, endIdx)

		fmt.Print("Move (0: down, 1: up, any other key: exit): ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')

		if strings.TrimSpace(input) == "1" {
			startIdx += pageSize
			if startIdx >= len(res) {
				if len(res)%2 == 0 {
					// Because when even last page will display only two
					startIdx = (len(res) - 2)
					fmt.Println("No further results. Reached last results")
				} else {
					// Because when odd last page will display only one
					startIdx = (len(res) - 1)
					fmt.Println("No further results. Reached last results")
				}
			}
			endIdx = startIdx + pageSize
		} else if strings.TrimSpace(input) == "0" {
			startIdx -= pageSize
			if startIdx < 0 {
				startIdx = 0
			}
			endIdx = startIdx + pageSize
		} else {
			// Exit
			fmt.Println("exited from search")
			break
		}
	}
}
