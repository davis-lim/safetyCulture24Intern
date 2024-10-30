package folder_test



import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

// Returns Sample Folders
func GetSampleFolders() []folder.Folder {
	orgID1, _ := uuid.FromString("11111111-1111-1111-1111-111111111111")
	orgID2, _ := uuid.FromString("22222222-2222-2222-2222-222222222222")
	orgID3, _ := uuid.FromString("33333333-3333-3333-3333-333333333333")
	return []folder.Folder {
		{Name: "alpha", Paths: "alpha", OrgId: orgID1},
		{Name: "alpha", Paths: "alpha", OrgId: orgID3},
        {Name: "bravo", Paths: "alpha.bravo", OrgId: orgID1},
        {Name: "charlie", Paths: "alpha.bravo.charlie", OrgId: orgID1},
        {Name: "delta", Paths: "alpha.delta", OrgId: orgID1},
        {Name: "echo", Paths: "alpha.delta.echo", OrgId: orgID1},
        {Name: "foxtrot", Paths: "foxtrot", OrgId: orgID2},
        {Name: "golf", Paths: "golf", OrgId: orgID3},
	}
}

func TestGetFoldersByOrgID(t *testing.T) {
	orgID1, _ := uuid.FromString("11111111-1111-1111-1111-111111111111")
	orgID2, _ := uuid.FromString("22222222-2222-2222-2222-222222222222")
	orgID3, _ := uuid.FromString("33333333-3333-3333-3333-333333333333")
	
	// Define test cases
	tests := []struct {
		name     string
		orgID    uuid.UUID
		expected []folder.Folder
	}{
		{
			name:  "Valid orgID with folders",
			orgID: orgID1,
			expected: []folder.Folder{
				{Name: "alpha", Paths: "alpha", OrgId: orgID1},
				{Name: "bravo", Paths: "alpha.bravo", OrgId: orgID1},
				{Name: "charlie", Paths: "alpha.bravo.charlie", OrgId: orgID1},
				{Name: "delta", Paths: "alpha.delta", OrgId: orgID1},
				{Name: "echo", Paths: "alpha.delta.echo", OrgId: orgID1},
			},
		},
		{
			name:  "Valid orgID with overlapping names",
			orgID: orgID3,
			expected: []folder.Folder{
				{Name: "alpha", Paths: "alpha", OrgId: orgID3},
				{Name: "golf", Paths: "golf", OrgId: orgID3},
			},
		},
		{
			name:  "Valid orgID with only one folder",
			orgID: orgID2,
			expected: []folder.Folder{
				{Name: "foxtrot", Paths: "foxtrot", OrgId: orgID2},
			},
		},
		{
			name:     "Valid orgID with no folders",
			orgID:    uuid.Must(uuid.FromString("22222222-1111-1111-1111-222222222222")),
			expected: []folder.Folder{}, // Expect an empty slice
		},
		{
			name:     "Invalid orgID",
			orgID:    uuid.Nil,
			expected: []folder.Folder{}, // Expect an empty slice
		},
	}

	// Iterate through test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			driver := folder.NewDriver(GetSampleFolders())
			result := driver.GetFoldersByOrgID(tt.orgID)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestGetAllChildFolders(t *testing.T) {
	orgID, _ := uuid.FromString("11111111-1111-1111-1111-111111111111")
	orgID2, _ := uuid.FromString("22222222-2222-2222-2222-222222222222")

	// Define test cases
	tests := []struct {
		name     string
		rootName string
		orgID    uuid.UUID
		expected []folder.Folder
		hasError bool
		expectedErr string
	}{
		{
			name:     "Get all child folders of 'alpha'",
			orgID:    orgID,
			rootName: "alpha",
			expected: []folder.Folder{
				{Name: "bravo", Paths: "alpha.bravo", OrgId: orgID},
				{Name: "charlie", Paths: "alpha.bravo.charlie", OrgId: orgID},
				{Name: "delta", Paths: "alpha.delta", OrgId: orgID},
				{Name: "echo", Paths: "alpha.delta.echo", OrgId: orgID},
			},
			hasError: false,
		},
		{
			name:     "Get all child folders of 'delta'",
			orgID:    orgID,
			rootName: "delta",
			expected: []folder.Folder{
				{Name: "echo", Paths: "alpha.delta.echo", OrgId: orgID},
			},
			hasError: false,
		},
		{
			name:     "Get all child folders of 'echo' (no descendants)",
			orgID:    orgID,
			rootName: "echo",
			expected: []folder.Folder{}, // Expect an empty slice
			hasError: false,
		},
		{
			name:     "Invalid orgID",
			orgID:    uuid.Must(uuid.FromString("11111111-3333-1111-1111-111111111111")),
			rootName: "alpha",
			expected: nil,
			hasError: true,
			expectedErr: "error: orgID does not contain any folders",
		},
		{
			name:     "Folder does not exist",
			orgID:    orgID,
			rootName: "invalid_folder",
			expected: nil,
			hasError: true,
			expectedErr: "error: Folder does not exist",
		},
		{
			name:     "Root folder exists but belongs to a different organization",
			orgID:    orgID2,
			rootName: "alpha",
			expected: nil,
			hasError: true,
			expectedErr: "error: Folder does not exist in the specified organization",
		},
	}

	// Iterate through test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			driver := folder.NewDriver(GetSampleFolders())
			result, error := driver.GetAllChildFolders(tt.orgID, tt.rootName)
			if tt.hasError {
				assert.Error(t, error)
				assert.EqualError(t, error, tt.expectedErr)
			} else {
				assert.NoError(t, error)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

