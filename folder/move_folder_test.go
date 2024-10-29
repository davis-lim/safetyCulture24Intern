package folder_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

// Sample data for testing
var moveSampleFolders = []folder.Folder{
	{Name: "alpha", Paths: "alpha", OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111"))},
	{Name: "bravo", Paths: "alpha.bravo", OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111"))},
	{Name: "charlie", Paths: "alpha.bravo.charlie", OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111"))},
	{Name: "delta", Paths: "alpha.delta", OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111"))},
	{Name: "echo", Paths: "alpha.delta.echo", OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111"))},
	{Name: "foxtrot", Paths: "foxtrot", OrgId: uuid.Must(uuid.FromString("22222222-2222-2222-2222-222222222222"))},
	{Name: "golf", Paths: "golf", OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111"))},
}

func Test_folder_MoveFolder(t *testing.T) {
	// TODO: your tests here
	driver := folder.NewDriver(sampleFolders)

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
				{Name: "alpha", Paths: "alpha", OrgId: moveSampleFolders[0].OrgId},
				{Name: "bravo", Paths: "alpha.delta.bravo", OrgId: moveSampleFolders[1].OrgId},
				{Name: "charlie", Paths: "alpha.delta.bravo.charlie", OrgId: moveSampleFolders[2].OrgId},
				{Name: "delta", Paths: "alpha.delta", OrgId: moveSampleFolders[3].OrgId},
				{Name: "echo", Paths: "alpha.delta.echo", OrgId: moveSampleFolders[4].OrgId},
				{Name: "foxtrot", Paths: "foxtrot", OrgId: moveSampleFolders[5].OrgId},
				{Name: "golf", Paths: "golf", OrgId: moveSampleFolders[6].OrgId},
			},
		},
		{
			name:        "Source folder does not exist",
			src:         "invalid_folder",
			dst:         "delta",
			hasError: true,
			expectedErr: "error: source folder does not exist",
		},
		{
			name:        "Destination folder does not exist",
			src:         "bravo",
			dst:         "invalid_folder",
			hasError: true,
			expectedErr: "error: destination folder does not exist",
		},
		{
			name:        "Move bravo to a folder in a different org (foxtrot)",
			src:         "bravo",
			dst:         "foxtrot",
			hasError: true,
			expectedErr: "error: source and destination belongs in different organisations",
		}, 
		{
			name:        "Move bravo to itself",
			src:         "bravo",
			dst:         "bravo",
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
			result, err := driver.MoveFolder(tt.src, tt.dst)

			if tt.expectedErr != "" {
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
