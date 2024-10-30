package folder_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

// Returns Sample Folders
func getMoveSampleFolders() []folder.Folder {
	orgID1, _ := uuid.FromString("11111111-1111-1111-1111-111111111111")
	orgID2, _ := uuid.FromString("22222222-2222-2222-2222-222222222222")
    return []folder.Folder{
        {Name: "alpha", Paths: "alpha", OrgId: orgID1},
        {Name: "bravo", Paths: "alpha.bravo", OrgId: orgID1},
        {Name: "charlie", Paths: "alpha.bravo.charlie", OrgId: orgID1},
        {Name: "delta", Paths: "alpha.delta", OrgId: orgID1},
        {Name: "echo", Paths: "alpha.delta.echo", OrgId: orgID1},
        {Name: "foxtrot", Paths: "alpha.delta.echo.foxtrot", OrgId: orgID1},
        {Name: "golf", Paths: "golf", OrgId: orgID2},
    }
}

func Test_folder_MoveFolder(t *testing.T) {
	orgID1, _ := uuid.FromString("11111111-1111-1111-1111-111111111111")
	orgID2, _ := uuid.FromString("22222222-2222-2222-2222-222222222222")

	tests := []struct {
		name        string
		src         string
		dst         string
		expected    []folder.Folder
		hasError bool
		expectedErr string
	}{
		{
			name: "Move bravo under delta",
			src:  "bravo",
			dst:  "delta",
			expected: []folder.Folder{
				{Name: "alpha", Paths: "alpha", OrgId: orgID1},
				{Name: "bravo", Paths: "alpha.delta.bravo", OrgId: orgID1},
				{Name: "charlie", Paths: "alpha.delta.bravo.charlie", OrgId: orgID1},
				{Name: "delta", Paths: "alpha.delta", OrgId: orgID1},
				{Name: "echo", Paths: "alpha.delta.echo", OrgId: orgID1},
				{Name: "foxtrot", Paths: "alpha.delta.echo.foxtrot", OrgId: orgID1},
				{Name: "golf", Paths: "golf", OrgId: orgID2},
			},
		},
		{
			name: "Move delta under bravo",
			src:  "delta",
			dst:  "bravo",
			expected: []folder.Folder{
				{Name: "alpha", Paths: "alpha", OrgId: orgID1},
				{Name: "bravo", Paths: "alpha.bravo", OrgId: orgID1},
				{Name: "charlie", Paths: "alpha.bravo.charlie", OrgId: orgID1},
				{Name: "delta", Paths: "alpha.bravo.delta", OrgId: orgID1},
				{Name: "echo", Paths: "alpha.bravo.delta.echo", OrgId: orgID1},
				{Name: "foxtrot", Paths: "alpha.bravo.delta.echo.foxtrot", OrgId: orgID1},
				{Name: "golf", Paths: "golf", OrgId: orgID2},
			},
		},
		{
			name:        "Source folder does not exist",
			src:         "testingFolder",
			dst:         "alpha",
			hasError: true,
			expectedErr: "error: source folder does not exist",
		},
		{
			name:        "Destination folder does not exist",
			src:         "echo",
			dst:         "testingFolder",
			hasError: true,
			expectedErr: "error: destination folder does not exist",
		},
		{
			name:        "Source and destination belongs in different organisations",
			src:         "alpha",
			dst:         "golf",
			hasError: true,
			expectedErr: "error: source and destination belongs in different organisations",
		}, 
		{
			name:        "Move folder to itself",
			src:         "echo",
			dst:         "echo",
			hasError: true,
			expectedErr: "error: source and destination are identical folders",
		},
		{
			name:        "Move bravo under its descendant charlie",
			src:         "bravo",
			dst:         "charlie",
			hasError: true,
			expectedErr: "error: destination is a child of source",
		},
		
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Uses fresh sampleFolders for each test (assumed no persistance)
			driver := folder.NewDriver(getMoveSampleFolders())
			result, err := driver.MoveFolder(tt.src, tt.dst)

			if tt.hasError {
				// Test for expected errors
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectedErr)
			} else {
				// Test for expected folder structure
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
	
}
