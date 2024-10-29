package folder_test



import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

// Sample data for testing
var sampleFolders = []folder.Folder{
	{Name: "alpha", Paths: "alpha", OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111"))},
	{Name: "bravo", Paths: "alpha.bravo", OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111"))},
	{Name: "charlie", Paths: "alpha.bravo.charlie", OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111"))},
	{Name: "delta", Paths: "alpha.delta", OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111"))},
	{Name: "echo", Paths: "echo", OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111"))},
	{Name: "foxtrot", Paths: "foxtrot", OrgId: uuid.Must(uuid.FromString("22222222-2222-2222-2222-222222222222"))},
}

func TestGetFoldersByOrgID(t *testing.T) {
	

	// Define test cases
	tests := []struct {
		name     string
		orgID    uuid.UUID
		expected []folder.Folder
	}{
		{
			name:  "Valid orgID with folders",
			orgID: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
			expected: []folder.Folder{
				{Name: "alpha", Paths: "alpha", OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111"))},
				{Name: "bravo", Paths: "alpha.bravo", OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111"))},
				{Name: "charlie", Paths: "alpha.bravo.charlie", OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111"))},
				{Name: "delta", Paths: "alpha.delta", OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111"))},
				{Name: "echo", Paths: "echo", OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111"))},
			},
		},
		{
			name:     "Valid orgID with no folders",
			orgID:    uuid.Must(uuid.FromString("33333333-3333-3333-3333-333333333333")),
			expected: []folder.Folder{}, // Expect an empty slice
		},
		{
			name:     "Invalid orgID (nil UUID)",
			orgID:    uuid.Nil,
			expected: []folder.Folder{}, // Expect an empty slice
		},
	}

	// Iterate through test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			driver := folder.NewDriver(sampleFolders)
			result := driver.GetFoldersByOrgID(tt.orgID)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestGetAllChildFolders(t *testing.T) {
	orgID, _ := uuid.FromString("11111111-1111-1111-1111-111111111111")

	tests := []struct {
		name     string
		rootName string
		orgID    uuid.UUID
		expected []folder.Folder
		hasError bool
	}{
		{
			name:     "Get all child folders of 'alpha'",
			orgID:    orgID,
			rootName: "alpha",
			expected: []folder.Folder{
				{Name: "bravo", Paths: "alpha.bravo", OrgId: orgID},
				{Name: "charlie", Paths: "alpha.bravo.charlie", OrgId: orgID},
				{Name: "delta", Paths: "alpha.delta", OrgId: orgID},
			},
			hasError: false,
		},
		{
			name:     "Get all child folders of 'bravo'",
			orgID:    orgID,
			rootName: "bravo",
			expected: []folder.Folder{
				{Name: "charlie", Paths: "alpha.bravo.charlie", OrgId: orgID},
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
			name:     "Root folder does not exist",
			orgID:    orgID,
			rootName: "invalid_folder",
			expected: nil,
			hasError: true,
		},
		{
			name:     "Root folder exists but belongs to a different organization",
			orgID:    uuid.Must(uuid.FromString("22222222-2222-2222-2222-222222222222")),
			rootName: "alpha",
			expected: nil,
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			driver := folder.NewDriver(sampleFolders)
			result, error := driver.GetAllChildFolders(tt.orgID, tt.rootName)
			if tt.hasError {
				assert.Error(t, error)
			} else {
				assert.NoError(t, error)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}






// feel free to change how the unit test is structured
// func Test_folder_GetFoldersByOrgID(t *testing.T) {
// 	t.Parallel()
// 	tests := [...]struct {
// 		name    string
// 		orgID   uuid.UUID
// 		folders []folder.Folder
// 		want    []folder.Folder
// 	}{
// 		// TODO: your tests here
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			// f := folder.NewDriver(tt.folders)
// 			// get := f.GetFoldersByOrgID(tt.orgID)

// 		})
// 	}
// }

