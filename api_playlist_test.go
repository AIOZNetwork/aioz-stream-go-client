package aiozstreamsdk

import (
	"io"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	testPlaylistID            string
	testAudioPlaylistID       string
	testDefaultTypePlaylistID string
	testPlaylistMediaAdded    bool
	testName                  = "Test Playlist"
	testCurrentId             string
	testNextId                string
	testPreviousId            string
	deletePlaylistsLater      []string
)

func TestPlaylistService_Create(t *testing.T) {
	tests := []struct {
		name     string
		request  CreatePlaylistRequest
		wantErr  bool
		storeIn  *string
		wantType string
	}{
		{
			name: "Valid Create Request",
			request: CreatePlaylistRequest{
				Name:         stringPtr(testName),
				PlaylistType: stringPtr("video"),
			},
			wantErr:  false,
			storeIn:  &testPlaylistID,
			wantType: "video",
		},
		{
			name: "Omitted Playlist Type Defaults To Video",
			request: CreatePlaylistRequest{
				Name: stringPtr(testName + " default-type"),
			},
			wantErr:  false,
			storeIn:  &testDefaultTypePlaylistID,
			wantType: "video",
		},
		{
			name: "Valid Audio Playlist Type",
			request: CreatePlaylistRequest{
				Name:         stringPtr(testName + " audio"),
				PlaylistType: stringPtr("audio"),
			},
			wantErr:  false,
			storeIn:  &testAudioPlaylistID,
			wantType: "audio",
		},
		{
			name: "Invalid Playlist Type",
			request: CreatePlaylistRequest{
				Name:         stringPtr(testName),
				PlaylistType: stringPtr("normal"),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.Playlist.Create(tt.request)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
				return
			}
			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.Data)
			require.NotNil(t, resp.Data.Playlist)
			require.NotNil(t, resp.Data.Playlist.Id)
			id := *resp.Data.Playlist.Id
			*tt.storeIn = id
			deletePlaylistsLater = append(deletePlaylistsLater, id)

			require.NotNil(t, resp.Data.Playlist.PlaylistType)
			assert.Equal(t, tt.wantType, *resp.Data.Playlist.PlaylistType)

			// The stored type, not just the create echo.
			got, err := testClient.Playlist.Get(id, PlaylistApiGetRequest{})
			require.NoError(t, err)
			require.NotNil(t, got)
			require.NotNil(t, got.Data)
			require.NotNil(t, got.Data.Playlist)
			require.NotNil(t, got.Data.Playlist.PlaylistType)
			assert.Equal(t, tt.wantType, *got.Data.Playlist.PlaylistType)
		})
	}
}

func TestPlaylistService_List(t *testing.T) {
	tests := []struct {
		name    string
		request ListPlaylistsRequest
		wantErr bool
	}{
		{
			name:    "Valid Request",
			request: ListPlaylistsRequest{},
			wantErr: false,
		},
		{
			name: "Valid Request with Filter",
			request: ListPlaylistsRequest{
				Limit:   int32Ptr(10),
				Offset:  int32Ptr(0),
				SortBy:  stringPtr("created_at"),
				OrderBy: stringPtr("desc"),
			},
			wantErr: false,
		},
		{
			name: "Valid Video Playlist Type Filter",
			request: ListPlaylistsRequest{
				PlaylistType: stringPtr("video"),
			},
			wantErr: false,
		},
		{
			name: "Valid Audio Playlist Type Filter",
			request: ListPlaylistsRequest{
				PlaylistType: stringPtr("audio"),
			},
			wantErr: false,
		},
		{
			name: "Invalid Playlist Type Filter",
			request: ListPlaylistsRequest{
				PlaylistType: stringPtr("normal"),
			},
			wantErr: true,
		},
		{
			name: "Invalid SortBy",
			request: ListPlaylistsRequest{
				SortBy: stringPtr("invalid"),
			},
			wantErr: true,
		},
		{
			name: "Invalid OrderBy",
			request: ListPlaylistsRequest{
				OrderBy: stringPtr("invalid"),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.Playlist.List(tt.request)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
			}
		})
	}
}

func TestPlaylistService_Update(t *testing.T) {
	requireSetupID(t, "testPlaylistID", testPlaylistID)
	notExistId := uuid.New().String()
	name := "Test Playlist"

	// Every case opens its own file so no case reads a stream an earlier
	// case already consumed. An empty filePath sends no file part at all.
	type updateCase struct {
		name     string
		id       string
		Name     *string
		fileName string
		filePath string
		wantErr  bool
		checkFn  func(*testing.T)
	}
	openCaseFile := func(t *testing.T, filePath string) io.Reader {
		if filePath == "" {
			return nil
		}
		return openTestAsset(t, filePath)
	}

	anonymousTest := []updateCase{
		{
			name:     "Update other",
			id:       testPlaylistID,
			Name:     stringPtr(name),
			fileName: "logo.png",
			filePath: "logo.png",
			wantErr:  true,
		},
	}

	for _, tt := range anonymousTest {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testAnonymousClient.Playlist.Update(
				tt.id,
				tt.Name,
				nil,
				tt.fileName,
				openCaseFile(t, tt.filePath),
			)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
			}
		})
	}

	tests := []updateCase{
		{
			name:    "Valid Update Name Only Without Thumbnail",
			id:      testPlaylistID,
			Name:    stringPtr(name),
			wantErr: false,
		},
		{
			name:     "Invalid Playlist ID",
			id:       "",
			fileName: "logo.png",
			filePath: "logo.png",
			wantErr:  true,
		},
		{
			name:     "Invalid Thumbnail",
			id:       testPlaylistID,
			fileName: "invalid-file.txt",
			filePath: "invalid-file.txt",
			wantErr:  true,
		},
		{
			name:     "PNG content named .jpg",
			id:       testPlaylistID,
			fileName: "logo.jpg",
			filePath: "logo.png",
			wantErr:  true,
		},
		{
			name:     "Not Exist ID",
			id:       notExistId,
			Name:     stringPtr(name),
			fileName: "logo.png",
			filePath: "logo.png",
			wantErr:  true,
		},
		{
			name:     "Valid Update Request With Name and Thumbnail",
			id:       testPlaylistID,
			Name:     stringPtr(name),
			fileName: "logo.png",
			filePath: "logo.png",
			wantErr:  false,
			checkFn: func(t *testing.T) {
				resp, err := testClient.Playlist.Get(
					testPlaylistID,
					PlaylistApiGetRequest{},
				)
				require.NoError(t, err)
				require.NotNil(t, resp)
				require.NotNil(t, resp.Data)
				require.NotNil(t, resp.Data.Playlist)
				require.NotNil(
					t,
					resp.Data.Playlist.ThumbnailUrl,
					"playlist should have a thumbnail after a valid upload",
				)
				assert.NotEmpty(t, *resp.Data.Playlist.ThumbnailUrl)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.Playlist.Update(
				tt.id,
				tt.Name,
				nil,
				tt.fileName,
				openCaseFile(t, tt.filePath),
			)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
				return
			}
			require.NoError(t, err)
			require.NotNil(t, resp)
			if tt.checkFn != nil {
				tt.checkFn(t)
			}
		})
	}

	t.Run("UpdateFile With Nil File", func(t *testing.T) {
		// UpdateFile is defined on *PlaylistService but missing from the
		// PlaylistServiceI interface, so reach it through the concrete type.
		svc, ok := testClient.Playlist.(*PlaylistService)
		require.True(
			t,
			ok,
			"testClient.Playlist is %T, want *PlaylistService",
			testClient.Playlist,
		)
		resp, err := svc.UpdateFile(
			testPlaylistID,
			nil,
			stringPtr(name),
			nil,
		)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
	})
}

func TestPlaylistService_GetPublic(t *testing.T) {
	requireSetupID(t, "testPlaylistID", testPlaylistID)
	notExistId := uuid.New().String()
	tests := []struct {
		name    string
		id      string
		wantErr bool
	}{
		{
			name:    "Valid Request",
			id:      testPlaylistID,
			wantErr: false,
		},
		{
			name:    "Invalid Playlist ID",
			id:      "",
			wantErr: true,
		},
		{
			name:    "Not Exist ID",
			id:      notExistId,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.Playlist.GetPublic(tt.id)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
			}
		})
	}
}

func TestPlaylistService_AddMedia(t *testing.T) {
	requireSetupID(t, "testPlaylistID", testPlaylistID)
	mediaIDs := readyMediaIDs(t, 3)
	notExistId := uuid.New().String()

	// A freshly created media has not been uploaded or processed, so the
	// backend must refuse to add it (409 media-not-ready).
	notReady, err := testClient.Media.Create(CreateMediaRequest{
		Title: stringPtr("Test Video Not Ready"),
		Qualities: &[]QualityConfig{
			{
				Type:          stringPtr("hls"),
				ContainerType: stringPtr("mpegts"),
				Resolution:    stringPtr("240p"),
			},
		},
	})
	require.NoError(t, err)
	require.NotNil(t, notReady)
	require.NotNil(t, notReady.Data)
	require.NotNil(t, notReady.Data.Id)
	notReadyID := *notReady.Data.Id
	t.Cleanup(func() { testClient.Media.Delete(notReadyID) })

	anonymousTest := []struct {
		name    string
		id      string
		request AddMediaRequest
		wantErr bool
	}{
		{
			name: "Add other",
			id:   testPlaylistID,
			request: AddMediaRequest{
				MediaId: stringPtr(mediaIDs[0]),
			},
			wantErr: true,
		},
	}

	for _, tt := range anonymousTest {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testAnonymousClient.Playlist.AddMedia(
				tt.id,
				tt.request,
			)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
			}
		})
	}

	tests := []struct {
		name    string
		id      string
		payload AddMediaRequest
		wantErr bool
	}{
		{
			name: "Valid Add First Video Request",
			id:   testPlaylistID,
			payload: AddMediaRequest{
				MediaId: stringPtr(mediaIDs[0]),
			},
			wantErr: false,
		},
		{
			name: "Valid Add Second Video Request",
			id:   testPlaylistID,
			payload: AddMediaRequest{
				MediaId: stringPtr(mediaIDs[1]),
			},
			wantErr: false,
		},
		{
			name: "Valid Add Third Video Request",
			id:   testPlaylistID,
			payload: AddMediaRequest{
				MediaId: stringPtr(mediaIDs[2]),
			},
			wantErr: false,
		},
		{
			name: "Media Not Ready",
			id:   testPlaylistID,
			payload: AddMediaRequest{
				MediaId: stringPtr(notReadyID),
			},
			wantErr: true,
		},
		{
			name: "Video Media Into Audio Playlist",
			id:   testAudioPlaylistID,
			payload: AddMediaRequest{
				MediaId: stringPtr(mediaIDs[0]),
			},
			wantErr: true,
		},
		{
			name: "Not Exist Media ID",
			id:   testPlaylistID,
			payload: AddMediaRequest{
				MediaId: stringPtr(notExistId),
			},
			wantErr: true,
		},
		{
			name: "Missing Video ID",
			id:   testPlaylistID,
			payload: AddMediaRequest{
				MediaId: stringPtr(""),
			},
			wantErr: true,
		},
		{
			name:    "Empty Request",
			id:      testPlaylistID,
			payload: AddMediaRequest{},
			wantErr: true,
		},
		{
			name: "Not Exist ID",
			id:   notExistId,
			payload: AddMediaRequest{
				MediaId: stringPtr(mediaIDs[0]),
			},
			wantErr: true,
		},
	}

	added := 0
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.Playlist.AddMedia(
				tt.id,
				tt.payload,
			)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else if assert.NoError(t, err) && assert.NotNil(t, resp) {
				added++
			}
		})
	}
	testPlaylistMediaAdded = added == 3
}

func TestPlaylistService_Get(t *testing.T) {
	requireSetupID(t, "testPlaylistID", testPlaylistID)
	notExistId := uuid.New().String()
	anonymousTest := []struct {
		name    string
		id      string
		wantErr bool
	}{
		{
			name:    "Get other",
			id:      testPlaylistID,
			wantErr: true,
		},
	}

	for _, tt := range anonymousTest {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testAnonymousClient.Playlist.Get(
				tt.id,
				PlaylistApiGetRequest{},
			)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
			}
		})
	}

	tests := []struct {
		name    string
		id      string
		request PlaylistApiGetRequest
		wantErr bool
	}{
		{
			name:    "Valid Playlist ID",
			id:      testPlaylistID,
			wantErr: false,
		},
		{
			name: "Valid Playlist ID with sortBy and orderBy",
			id:   testPlaylistID,
			request: PlaylistApiGetRequest{}.
				SortBy("created_at").
				OrderBy("desc"),
			wantErr: false,
		},
		{
			name:    "Invalid Playlist ID",
			id:      "",
			wantErr: true,
		},
		{
			name:    "Empty Request",
			id:      testPlaylistID,
			request: PlaylistApiGetRequest{},
			wantErr: false,
		},
		{
			name:    "Invalid SortBy",
			id:      testPlaylistID,
			request: PlaylistApiGetRequest{}.SortBy("invalid"),
			wantErr: true,
		},
		{
			name:    "Invalid orderBy",
			id:      testPlaylistID,
			request: PlaylistApiGetRequest{}.OrderBy("invalid"),
			wantErr: true,
		},
		{
			name:    "Not Exist ID",
			id:      notExistId,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.Playlist.Get(tt.id, tt.request)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
				return
			}
			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.Data)
			require.NotNil(t, resp.Data.Playlist)

			if !testPlaylistMediaAdded {
				t.Log(
					"AddMedia did not add all three media; skipping item checks",
				)
				return
			}
			require.NotNil(t, resp.Data.Playlist.Items)
			items := *resp.Data.Playlist.Items
			require.Len(t, items, 3)
			var middleFound bool
			for _, item := range items {
				assert.NotNil(
					t,
					item.Media,
					"playlist item %v has no media",
					item.Id,
				)
				// The middle item has both neighbours, which the move test needs.
				if item.Id != nil && item.NextId != nil &&
					item.PreviousId != nil {
					middleFound = true
					testCurrentId = *item.Id
					testNextId = *item.NextId
					testPreviousId = *item.PreviousId
				}
			}
			require.True(
				t,
				middleFound,
				"no playlist item has both a next and a previous item",
			)
		})
	}
}

// requirePlaylistItems skips the item tests when the playlist items they act
// on were never recorded (AddMedia skipped or failed).
func requirePlaylistItems(t *testing.T) {
	t.Helper()
	if testCurrentId == "" || testNextId == "" || testPreviousId == "" {
		t.Skip(
			"playlist items not available: AddMedia/Get did not record them",
		)
	}
}

func TestPlaylistService_MoveItem(t *testing.T) {
	requireSetupID(t, "testPlaylistID", testPlaylistID)
	requirePlaylistItems(t)
	anonymousTest := []struct {
		name    string
		id      string
		payload MoveItemRequest
		wantErr bool
	}{
		{
			name: "Move other",
			id:   testPlaylistID,
			payload: MoveItemRequest{
				CurrentId:  stringPtr(testNextId),
				NextId:     stringPtr(testCurrentId),
				PreviousId: stringPtr(testPreviousId),
			},
			wantErr: true,
		},
	}

	for _, tt := range anonymousTest {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testAnonymousClient.Playlist.MoveItem(
				tt.id,
				tt.payload,
			)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
			}
		})
	}

	tests := []struct {
		name    string
		id      string
		payload MoveItemRequest
		wantErr bool
	}{
		{
			name: "Valid Move Video Request",
			id:   testPlaylistID,
			payload: MoveItemRequest{
				CurrentId:  stringPtr(testNextId),
				NextId:     stringPtr(testCurrentId),
				PreviousId: stringPtr(testPreviousId),
			},
			wantErr: false,
		},
		{
			name:    "Invalid Playlist ID",
			id:      "",
			wantErr: true,
		},
		{
			name:    "Missing Current ID",
			id:      testPlaylistID,
			wantErr: true,
		},
		{
			name: "Missing Next ID",
			id:   testPlaylistID,
			payload: MoveItemRequest{
				CurrentId: stringPtr(testCurrentId),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.Playlist.MoveItem(
				tt.id,
				tt.payload,
			)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
			}
		})
	}
}

func TestPlaylistService_RemoveMedia(t *testing.T) {
	requireSetupID(t, "testPlaylistID", testPlaylistID)
	requirePlaylistItems(t)
	notExistId := uuid.New().String()
	anonymousTest := []struct {
		name    string
		id      string
		itemId  string
		wantErr bool
	}{
		{
			name:    "Remove other",
			id:      testPlaylistID,
			itemId:  testPreviousId,
			wantErr: true,
		},
	}

	for _, tt := range anonymousTest {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testAnonymousClient.Playlist.RemoveMedia(
				tt.id,
				tt.itemId,
				RemoveMediaRequest{},
			)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
			}
		})
	}

	tests := []struct {
		name    string
		id      string
		itemId  string
		wantErr bool
	}{
		{
			name:    "Valid Remove First Video Request",
			id:      testPlaylistID,
			itemId:  testPreviousId,
			wantErr: false,
		},
		{
			name:    "Valid Remove Second Video Request",
			id:      testPlaylistID,
			itemId:  testCurrentId,
			wantErr: false,
		},
		{
			name:    "Valid Remove Third Video Request",
			id:      testPlaylistID,
			itemId:  testNextId,
			wantErr: false,
		},
		{
			name:    "Invalid Playlist ID",
			id:      "",
			wantErr: true,
		},
		{
			name:    "Missing Item ID",
			id:      testPlaylistID,
			wantErr: true,
		},
		{
			name:    "Not Exist ID",
			id:      notExistId,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.Playlist.RemoveMedia(
				tt.id,
				tt.itemId,
				RemoveMediaRequest{},
			)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
			}
		})
	}
}

// TestPlaylistService_DeleteThumbnail deletes the thumbnail that the valid
// thumbnail case in TestPlaylistService_Update uploaded.
func TestPlaylistService_DeleteThumbnail(t *testing.T) {
	requireSetupID(t, "testPlaylistID", testPlaylistID)
	notExistId := uuid.New().String()
	anonymousTest := []struct {
		name    string
		id      string
		wantErr bool
	}{
		{
			name:    "Delete other",
			id:      testPlaylistID,
			wantErr: true,
		},
	}

	for _, tt := range anonymousTest {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testAnonymousClient.Playlist.DeleteThumbnail(
				tt.id,
			)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
			}
		})
	}

	tests := []struct {
		name    string
		id      string
		wantErr bool
	}{
		{
			name:    "Valid Playlist ID",
			id:      testPlaylistID,
			wantErr: false,
		},
		{
			name:    "Invalid Playlist ID",
			id:      "",
			wantErr: true,
		},
		{
			name:    "Not Exist ID",
			id:      notExistId,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.Playlist.DeleteThumbnail(tt.id)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
			}
		})
	}
}

func TestPlaylistService_Delete(t *testing.T) {
	requireSetupID(t, "testPlaylistID", testPlaylistID)
	notExistId := uuid.New().String()
	anonymousTest := []struct {
		name    string
		id      string
		wantErr bool
	}{
		{
			name:    "Delete other",
			id:      testPlaylistID,
			wantErr: true,
		},
	}

	for _, tt := range anonymousTest {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testAnonymousClient.Playlist.Delete(tt.id)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
			}
		})
	}

	tests := []struct {
		name    string
		id      string
		wantErr bool
	}{
		{
			name:    "Valid Playlist ID",
			id:      testPlaylistID,
			wantErr: false,
		},
		{
			name:    "Invalid Playlist ID",
			id:      "",
			wantErr: true,
		},
		{
			name:    "Not Exist ID",
			id:      notExistId,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.Playlist.Delete(tt.id)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
			}
		})
	}

	for _, id := range deletePlaylistsLater {
		testClient.Playlist.Delete(id)
	}
}

func int32Ptr(i int32) *int32 {
	return &i
}
