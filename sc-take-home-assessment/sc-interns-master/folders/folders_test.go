package folders_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folders"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_GetAllFolders(t *testing.T) {
	// Sample testing data
	sampleData := []*folders.Folder{
		{
			Id:      uuid.Must(uuid.FromString("1167c1ac-911b-4a1f-b460-a98f724c7289")),
			Name:    "heroic-bella",
			OrgId:   uuid.Must(uuid.FromString("4212d618-66ff-468a-862d-ea49fef5e183")),
			Deleted: false,
		},
		{
			Id:      uuid.Must(uuid.FromString("1167c1ac-911b-4a1f-b460-a98f724c7280")),
			Name:    "heroic-nord",
			OrgId:   uuid.Must(uuid.FromString("4212d618-66ff-468a-862d-ea49fef5e183")),
			Deleted: true,
		},
	}

	t.Run("Matching Organization ID", func(t *testing.T) {
		orgID := sampleData[0].OrgId
		result, err := folders.FetchAllFoldersByOrgID(orgID)
		// Function executes without error
		assert.NoError(t, err)
		// Expecting two results with this ID
		assert.Len(t, result, 2)
		// Assertions to check the contents of result and if they match up with the expected results in sampleData
		assert.Equal(t, sampleData[0].Name, result[0].Name)
		assert.Equal(t, sampleData[0].Deleted, result[0].Deleted)
		assert.Equal(t, sampleData[0].Id, result[0].Id)
	})

	t.Run("Non-matching organization ID", func(t *testing.T) {
		// Generate new UUID that shuld not match any in sample.json
		// 'Must' ensures an error is returned if UUID generation fails
		orgID := uuid.Must(uuid.NewV4())
		result, err := folders.FetchAllFoldersByOrgID(orgID)
		// Check if err value is nil
		assert.NoError(t, err)
		// Check because orgID provided doesn't match, the result should be an empty slice
		assert.Empty(t, result)
	})

	// t.Run("test", func(t *testing.T) {
	// 	// your test/s here
	// })

}
