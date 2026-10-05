package aiozstreamsdk

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"
	"testing"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	testPublicKey          string
	testSecretKey          string
	testAnonymousSecretKey string
	testAnonymousPublicKey string

	testClient          *Client
	testAnonymousClient *Client

	testMediaID      string
	title            = "Test Video"
	description      = "Test Description"
	deleteMediaLater []string

	readyMediaOnce     sync.Once
	readyMediaIDsCache []string
	readyMediaErr      error
)

func init() {
	if err := loadEnvVariables(); err != nil {
		log.Fatalf("Failed to load environment variables: %v", err)
	}

	initializeClients()
}

func loadEnvVariables() error {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found, using environment variables")
	}

	var missingVars []string

	testPublicKey = os.Getenv("TEST_PUBLIC_KEY")
	if testPublicKey == "" {
		missingVars = append(missingVars, "TEST_PUBLIC_KEY")
	}

	testSecretKey = os.Getenv("TEST_SECRET_KEY")
	if testSecretKey == "" {
		missingVars = append(missingVars, "TEST_SECRET_KEY")
	}

	testAnonymousSecretKey = os.Getenv("TEST_ANONYMOUS_SECRET_KEY")
	if testAnonymousSecretKey == "" {
		missingVars = append(missingVars, "TEST_ANONYMOUS_SECRET_KEY")
	}

	testAnonymousPublicKey = os.Getenv("TEST_ANONYMOUS_PUBLIC_KEY")

	// Optional: the endpoint the webhook tests register. Webhook.Check makes
	// the backend deliver a test event to it, so it must be public and 2xx.
	if u := os.Getenv("TEST_WEBHOOK_URL"); u != "" {
		webhookURL = u
	}

	if len(missingVars) > 0 {
		return fmt.Errorf(
			"missing required environment variables: %v",
			missingVars,
		)
	}

	return nil
}

// initializeClients builds the authenticated and anonymous clients. Setting
// TEST_BASE_URL (e.g. http://localhost:3017/api/) points both at another
// deployment instead of the production default.
func initializeClients() {
	baseURL := os.Getenv("TEST_BASE_URL")

	b := ClientBuilder(AuthCredentials{
		PublicKey: testPublicKey,
		SecretKey: testSecretKey,
	})
	if baseURL != "" {
		b.BaseURL(baseURL)
	}
	testClient = b.Build()

	ab := ClientBuilder(AuthCredentials{
		PublicKey: testAnonymousPublicKey,
		SecretKey: testAnonymousSecretKey,
	})
	if baseURL != "" {
		ab.BaseURL(baseURL)
	}
	testAnonymousClient = ab.Build()
}

func openTestVideoFile(t *testing.T) *os.File {
	file, err := os.Open("test-assets/558k.mp4")
	if err != nil {
		t.Fatal(err)
	}
	return file
}

// openTestAsset opens a file from test-assets and closes it when the test ends.
func openTestAsset(t *testing.T, name string) *os.File {
	t.Helper()
	file, err := os.Open("test-assets/" + name)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { file.Close() })
	return file
}

func getFileHash(t *testing.T, file *os.File) string {
	hash := md5.New()
	_, err := io.Copy(hash, file)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := file.Seek(0, 0); err != nil {
		t.Fatal(err)
	}
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func TestMediaService_Create(t *testing.T) {
	validMetadata := []Metadata{
		{Key: stringPtr("key1"), Value: stringPtr("value1")},
		{Key: stringPtr("key2"), Value: stringPtr("value2")},
	}

	validTags := []string{"tag1", "tag2"}
	tests := []struct {
		name    string
		request CreateMediaRequest
		wantErr bool
	}{
		{
			name: "Valid Complete Request",
			request: CreateMediaRequest{
				Title:       stringPtr("Test Video"),
				Description: stringPtr("Test Description"),
				IsPublic:    boolPtr(true),
				Metadata:    &validMetadata,
				Qualities: &[]QualityConfig{
					{
						Type:          stringPtr("hls"),
						ContainerType: stringPtr("mpegts"),
						Resolution:    stringPtr("240p"),
					},
				},
				Tags: &validTags,
			},
			wantErr: false,
		},
		{
			name: "Valid Minimal Request",
			request: CreateMediaRequest{
				Title: stringPtr("Test Video"),
				Qualities: &[]QualityConfig{
					{
						Type:          stringPtr("hls"),
						ContainerType: stringPtr("mpegts"),
						Resolution:    stringPtr("240p"),
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Invalid Title - Empty",
			request: CreateMediaRequest{
				Qualities: &[]QualityConfig{
					{
						Type:          stringPtr("hls"),
						ContainerType: stringPtr("mpegts"),
						Resolution:    stringPtr("240p"),
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Invalid Title - Too Long",
			request: CreateMediaRequest{
				Title: stringPtr(strings.Repeat("a", 256)),
				Qualities: &[]QualityConfig{
					{
						Type:          stringPtr("hls"),
						ContainerType: stringPtr("mpegts"),
						Resolution:    stringPtr("240p"),
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Invalid Description - Too Long",
			request: CreateMediaRequest{
				Title:       stringPtr("Test Video"),
				Description: stringPtr(strings.Repeat("a", 1001)),
				Qualities: &[]QualityConfig{
					{
						Type:          stringPtr("hls"),
						ContainerType: stringPtr("mpegts"),
						Resolution:    stringPtr("240p"),
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Invalid Metadata - Key Too Long",
			request: CreateMediaRequest{
				Title: stringPtr("Test Video"),
				Qualities: &[]QualityConfig{
					{
						Type:          stringPtr("hls"),
						ContainerType: stringPtr("mpegts"),
						Resolution:    stringPtr("240p"),
					},
				},
				Metadata: &[]Metadata{
					{
						Key:   stringPtr(strings.Repeat("a", 256)),
						Value: stringPtr("value"),
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Invalid Qualities - Empty",
			request: CreateMediaRequest{
				Title:     stringPtr("Test Video"),
				Qualities: &[]QualityConfig{},
			},
			wantErr: false,
		},
		{
			name: "Invalid Qualities - Invalid Value",
			request: CreateMediaRequest{
				Title: stringPtr("Test Video"),
				Qualities: &[]QualityConfig{
					{
						Type:          stringPtr("video/mp4"),
						ContainerType: stringPtr("mp4"),
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Valid Qualities - All Supported",
			request: CreateMediaRequest{
				Title: stringPtr("Test Video"),
				Qualities: &[]QualityConfig{
					{
						Type:          stringPtr("hls"),
						ContainerType: stringPtr("mpegts"),
						Resolution:    stringPtr("240p"),
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Invalid Tags - Tag Too Long",
			request: CreateMediaRequest{
				Title: stringPtr("Test Video"),
				Qualities: &[]QualityConfig{
					{
						Type:          stringPtr("hls"),
						ContainerType: stringPtr("mpegts"),
						Resolution:    stringPtr("240p"),
					},
				},
				Tags: &[]string{strings.Repeat("a", 256)},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.Media.Create(tt.request)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				require.NoError(t, err)
				require.NotNil(t, resp)
				require.NotNil(t, resp.Data)
				require.NotNil(t, resp.Data.Id)
				assert.NotEmpty(t, *resp.Data.Id)
				testMediaID = *resp.Data.Id
				deleteMediaLater = append(deleteMediaLater, *resp.Data.Id)
			}
		})
	}
}

func TestMediaService_GetMediaList(t *testing.T) {
	tests := []struct {
		name    string
		request GetMediaListRequest
		wantErr bool
	}{
		{
			name:    "Valid Get Media List With No Filter",
			request: GetMediaListRequest{},
			wantErr: false,
		},
		{
			name: "Valid Get Media List With Filter",
			request: GetMediaListRequest{
				Limit:   int64Ptr(10),
				Offset:  int64Ptr(0),
				OrderBy: stringPtr("created_at"),
				SortBy:  stringPtr("desc"),
			},
			wantErr: false,
		},
		{
			name: "Valid Get Media List With Status And Type Filter",
			request: GetMediaListRequest{
				Status: &[]string{"done"},
				Type:   stringPtr("video"),
				Limit:  int64Ptr(10),
			},
			wantErr: false,
		},
		{
			name: "Invalid Status",
			request: GetMediaListRequest{
				Status: &[]string{"bogus"},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.Media.GetMediaList(tt.request)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				require.NoError(t, err)
				require.NotNil(t, resp)
				require.NotNil(t, resp.Data)
				if resp.Data.Total != nil && *resp.Data.Total > 0 {
					require.NotNil(
						t,
						resp.Data.Media,
						"total is %d but no media decoded",
						*resp.Data.Total,
					)
					assert.NotEmpty(t, *resp.Data.Media)
				}
				if tt.request.Status != nil && resp.Data.Media != nil {
					for _, m := range *resp.Data.Media {
						assert.Equal(
							t,
							stringPtr("done"),
							m.Status,
							"media %v",
							m.Id,
						)
					}
				}
			}
		})
	}
}

func TestMediaService_Update(t *testing.T) {
	requireSetupID(t, "testMediaID", testMediaID)
	notExistId := uuid.New().String()
	anonymousTest := []struct {
		name    string
		id      string
		input   UpdateMediaInfoRequest
		wantErr bool
	}{
		{
			name:    "Update other",
			id:      testMediaID,
			input:   UpdateMediaInfoRequest{Title: &title},
			wantErr: true,
		},
	}

	for _, tt := range anonymousTest {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testAnonymousClient.Media.Update(tt.id, tt.input)
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
		input   UpdateMediaInfoRequest
		wantErr bool
	}{
		{
			name: "Valid Update All Fields",
			id:   testMediaID,
			input: UpdateMediaInfoRequest{
				Title:       &title,
				Description: &description,
			},
			wantErr: false,
		},
		{
			name: "Valid Update Title Only",
			id:   testMediaID,
			input: UpdateMediaInfoRequest{
				Title: &title,
			},
			wantErr: false,
		},
		{
			name: "Valid Update Description Only",
			id:   testMediaID,
			input: UpdateMediaInfoRequest{
				Description: &description,
			},
			wantErr: false,
		},
		{
			name: "Invalid Title Length",
			id:   testMediaID,
			input: UpdateMediaInfoRequest{
				Title: stringPtr(strings.Repeat("a", 256)),
			},
			wantErr: true,
		},
		{
			name:    "Invalid Media ID",
			id:      "invalid-uuid",
			input:   UpdateMediaInfoRequest{Title: &title},
			wantErr: true,
		},
		{
			name:    "Non-existent Media ID",
			id:      "12345678-1234-1234-1234-123456789012",
			input:   UpdateMediaInfoRequest{Title: &title},
			wantErr: true,
		},
		{
			name:    "Not Exist ID",
			id:      notExistId,
			input:   UpdateMediaInfoRequest{Title: &title},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.Media.Update(tt.id, tt.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
			}
		})
	}
}

func TestMediaService_GetDetail(t *testing.T) {
	requireSetupID(t, "testMediaID", testMediaID)
	notExistId := uuid.New().String()
	anonymousTest := []struct {
		name    string
		id      string
		wantErr bool
	}{
		{
			name:    "Get other",
			id:      testMediaID,
			wantErr: true,
		},
	}
	for _, tt := range anonymousTest {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testAnonymousClient.Media.GetDetail(tt.id)
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
		checkFn func(*testing.T, *GetMediaDetailResponse)
	}{
		{
			name:    "Valid Get Detail",
			id:      testMediaID,
			wantErr: false,
			checkFn: func(t *testing.T, resp *GetMediaDetailResponse) {
				require.NotNil(t, resp.Data)
				assert.NotEmpty(t, resp.Data.Id)
				assert.NotEmpty(t, resp.Data.Title)
				assert.NotEmpty(t, resp.Status)
			},
		},
		{
			name:    "Invalid ID Format",
			id:      "invalid-uuid",
			wantErr: true,
			checkFn: nil,
		},
		{
			name:    "Non-existent ID",
			id:      "12345678-1234-1234-1234-123456789012",
			wantErr: true,
			checkFn: nil,
		},
		{
			name:    "Empty ID",
			id:      "",
			wantErr: true,
			checkFn: nil,
		},
		{
			name:    "Not Exist ID",
			id:      notExistId,
			wantErr: true,
			checkFn: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.Media.GetDetail(tt.id)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				require.NoError(t, err)
				require.NotNil(t, resp)
				if tt.checkFn != nil {
					tt.checkFn(t, resp)
				}
			}
		})
	}
}

func TestMediaService_UploadThumbnail(t *testing.T) {
	requireSetupID(t, "testMediaID", testMediaID)
	notExistId := uuid.New().String()

	// Every case opens its own file so no case reads a stream an earlier
	// case already consumed. An empty filePath sends no file at all.
	type thumbnailCase struct {
		name     string
		id       string
		fileName string
		filePath string
		wantErr  bool
	}
	openCaseFile := func(t *testing.T, filePath string) io.Reader {
		if filePath == "" {
			return nil
		}
		return openTestAsset(t, filePath)
	}

	t.Run("Anonymous Client Tests", func(t *testing.T) {
		tests := []thumbnailCase{
			{
				name:     "Unauthorized thumbnail upload",
				id:       testMediaID,
				fileName: "thumbnail.png",
				filePath: "logo.png",
				wantErr:  true,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				resp, err := testAnonymousClient.Media.UploadThumbnail(
					tt.id,
					tt.fileName,
					openCaseFile(t, tt.filePath),
				)
				assert.Equal(t, tt.wantErr, err != nil)
				if tt.wantErr {
					assert.Nil(t, resp)
				}
			})
		}
	})

	tests := []thumbnailCase{
		{
			name:     "Valid Thumbnail Upload",
			id:       testMediaID,
			fileName: "thumbnail.png",
			filePath: "logo.png",
			wantErr:  false,
		},
		{
			name:     "PNG content named .jpg",
			id:       testMediaID,
			fileName: "thumbnail.jpg",
			filePath: "logo.png",
			wantErr:  true,
		},
		{
			name:     "Invalid File Type",
			id:       testMediaID,
			fileName: "thumbnail.gif",
			filePath: "invalid-file.txt",
			wantErr:  true,
		},
		{
			name:     "Invalid Media ID",
			id:       "invalid-id",
			fileName: "thumbnail.png",
			filePath: "logo.png",
			wantErr:  true,
		},
		{
			name:     "Empty File Name",
			id:       testMediaID,
			fileName: "",
			filePath: "logo.png",
			wantErr:  true,
		},
		{
			name:     "Nil File",
			id:       testMediaID,
			fileName: "thumbnail.png",
			filePath: "",
			wantErr:  true,
		},
		{
			name:     "Not Exist ID",
			id:       notExistId,
			fileName: "thumbnail.png",
			filePath: "logo.png",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.Media.UploadThumbnail(
				tt.id,
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
}

// TestMediaService_DeleteThumbnail must run after TestMediaService_UploadThumbnail,
// which leaves a thumbnail on testMediaID.
func TestMediaService_DeleteThumbnail(t *testing.T) {
	requireSetupID(t, "testMediaID", testMediaID)
	notExistId := uuid.New().String()

	anonymousTest := []struct {
		name    string
		id      string
		wantErr bool
	}{
		{
			name:    "Delete other",
			id:      testMediaID,
			wantErr: true,
		},
	}

	for _, tt := range anonymousTest {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testAnonymousClient.Media.DeleteThumbnail(tt.id)
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
			name:    "Valid Delete Thumbnail",
			id:      testMediaID,
			wantErr: false,
		},
		{
			name:    "Invalid Media ID",
			id:      "invalid-id",
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
			resp, err := testClient.Media.DeleteThumbnail(tt.id)
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

func TestMediaService_GetCost(t *testing.T) {
	// GetCost now takes the media type; every case prices a video.
	const mediaType = "video"
	tests := []struct {
		name      string
		qualities string
		duration  float32
		wantErr   bool
	}{
		{
			name:      "Valid Single Quality",
			qualities: "720p",
			duration:  120.5,
			wantErr:   false,
		},
		{
			name:      "Valid Multiple Qualities",
			qualities: "720p,1080p",
			duration:  120.5,
			wantErr:   false,
		},
		{
			name:      "Invalid Quality",
			qualities: "invalid",
			duration:  120.5,
			wantErr:   true,
		},
		{
			name:      "Empty Quality",
			qualities: "",
			duration:  120.5,
			wantErr:   true,
		},
		{
			name:      "Negative Duration",
			qualities: "720p",
			duration:  -1,
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.Media.GetCost(
				tt.qualities,
				mediaType,
				tt.duration,
			)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				require.NoError(t, err)
				require.NotNil(t, resp)
				require.NotNil(t, resp.Data)
				assert.NotNil(t, resp.Data.Price)
			}
		})
	}
}

func TestMediaService_UploadPart(t *testing.T) {
	requireSetupID(t, "testMediaID", testMediaID)

	video := openTestVideoFile(t)
	defer video.Close()
	videoHash := getFileHash(t, video)

	index := "1"
	notExistId := uuid.New().String()

	tests := []struct {
		name    string
		id      string
		hash    *string
		index   *string
		file    *os.File
		wantErr bool
	}{
		{
			name:    "Valid File Upload",
			id:      testMediaID,
			hash:    &videoHash,
			index:   &index,
			file:    video,
			wantErr: false,
		},
		{
			name:    "Invalid Media ID",
			id:      "invalid-id",
			hash:    &videoHash,
			index:   &index,
			file:    video,
			wantErr: true,
		},
		{
			name:    "Not Exist ID",
			id:      notExistId,
			hash:    &videoHash,
			index:   &index,
			file:    video,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fileInfo, err := tt.file.Stat()
			require.NoError(t, err)
			resp, err := testClient.Media.UploadPart(
				tt.id,
				tt.hash,
				tt.index,
				tt.file.Name(),
				tt.file,
				fileInfo.Size(),
			)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
			}
		})
	}
}

func TestMediaService_UploadMediaComplete(t *testing.T) {
	requireSetupID(t, "testMediaID", testMediaID)
	notExistId := uuid.New().String()
	tests := []struct {
		name    string
		id      string
		wantErr bool
	}{
		{
			name:    "Valid Upload Media Complete",
			id:      testMediaID,
			wantErr: false,
		},
		{
			name:    "Invalid Media ID",
			id:      "invalid-id",
			wantErr: true,
		},
		{
			name:    "Empty Media ID",
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
			resp, err := testClient.Media.UploadMediaComplete(tt.id)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
			}
		})
	}

	anonymousTest := []struct {
		name    string
		id      string
		wantErr bool
	}{
		{
			name:    "Check other media complete",
			id:      testMediaID,
			wantErr: true,
		},
	}
	for _, tt := range anonymousTest {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testAnonymousClient.Media.UploadMediaComplete(tt.id)
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

func TestMediaService_GetMediaPlayerInfo(t *testing.T) {
	requireSetupID(t, "testMediaID", testMediaID)
	notExistId := uuid.New().String()
	tests := []struct {
		name    string
		id      string
		request MediaApiGetMediaPlayerInfoRequest
		wantErr bool
	}{
		{
			name:    "Valid Get Media Player Info",
			id:      testMediaID,
			request: MediaApiGetMediaPlayerInfoRequest{},
			wantErr: false,
		},
		{
			name:    "Invalid Media ID",
			id:      "invalid-id",
			request: MediaApiGetMediaPlayerInfoRequest{},
			wantErr: true,
		},
		{
			name:    "Empty Media ID",
			id:      "",
			request: MediaApiGetMediaPlayerInfoRequest{},
			wantErr: true,
		},
		{
			name:    "Not Exist ID",
			id:      notExistId,
			request: MediaApiGetMediaPlayerInfoRequest{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.Media.GetMediaPlayerInfo(tt.id, tt.request)
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

func TestMediaService_CreateCaption(t *testing.T) {
	requireSetupID(t, "testMediaID", testMediaID)
	tmpFile := createTempVTTFile(t)
	notExistId := uuid.New().String()
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	tests := []struct {
		name    string
		id      string
		lang    string
		file    *os.File
		wantErr bool
	}{
		{
			name:    "Valid Create Media Captions",
			id:      testMediaID,
			lang:    testLang,
			file:    tmpFile,
			wantErr: false,
		},
		{
			name:    "Empty Language",
			id:      testMediaID,
			lang:    "",
			file:    tmpFile,
			wantErr: true,
		},
		{
			name:    "Invalid Media ID",
			id:      "invalid-id",
			lang:    testLang,
			file:    tmpFile,
			wantErr: true,
		},
		{
			name:    "Not Exist ID",
			id:      notExistId,
			lang:    testLang,
			file:    tmpFile,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.Media.CreateCaption(
				tt.id,
				tt.lang,
				nil,
				tt.file.Name(),
				tt.file,
			)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
				return
			}
			require.NoError(t, err)
			require.NotNil(t, resp)
			assert.NotNil(
				t,
				findCaption(t, tt.id, tt.lang),
				"caption %q not listed for media %s after CreateCaption",
				tt.lang,
				tt.id,
			)
		})
	}
}

func TestMediaService_GetCaptions(t *testing.T) {
	testMediaCaptionID := readyMediaIDs(t, 1)[0]
	notExistId := uuid.New().String()
	anonymousTest := []struct {
		name    string
		id      string
		wantErr bool
	}{
		{
			name:    "Get other",
			id:      testMediaCaptionID,
			wantErr: true,
		},
	}

	for _, tt := range anonymousTest {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testAnonymousClient.Media.GetCaptions(
				tt.id,
				MediaApiGetCaptionsRequest{},
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
		request MediaApiGetCaptionsRequest
		wantErr bool
	}{
		{
			name:    "Valid Get Media Captions",
			id:      testMediaCaptionID,
			wantErr: false,
		},
		{
			name:    "Invalid Media ID",
			id:      "invalid-id",
			wantErr: true,
		},
		{
			name:    "Empty Media ID",
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
			resp, err := testClient.Media.GetCaptions(tt.id, tt.request)
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

// TestMediaService_SetDefaultCaption uses the testLang caption that
// TestMediaService_CreateCaption put on testMediaID, and must run before
// TestMediaService_DeleteCaption removes it.
func TestMediaService_SetDefaultCaption(t *testing.T) {
	requireSetupID(t, "testMediaID", testMediaID)
	notExistId := uuid.New().String()
	setTrue := SetDefaultCaptionRequest{IsDefault: boolPtr(true)}
	setFalse := SetDefaultCaptionRequest{IsDefault: boolPtr(false)}

	anonymousTest := []struct {
		name    string
		id      string
		lang    string
		wantErr bool
	}{
		{
			name:    "Set other",
			id:      testMediaID,
			lang:    testLang,
			wantErr: true,
		},
	}

	for _, tt := range anonymousTest {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testAnonymousClient.Media.SetDefaultCaption(
				tt.id,
				tt.lang,
				setTrue,
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

	// The valid cases run in order: set the caption as default, then unset it.
	tests := []struct {
		name    string
		id      string
		lang    string
		request SetDefaultCaptionRequest
		wantErr bool
	}{
		{
			name:    "Valid Set Default Caption",
			id:      testMediaID,
			lang:    testLang,
			request: setTrue,
			wantErr: false,
		},
		{
			name:    "Valid Unset Default Caption",
			id:      testMediaID,
			lang:    testLang,
			request: setFalse,
			wantErr: false,
		},
		{
			name:    "Not Exist Language",
			id:      testMediaID,
			lang:    "fr",
			request: setTrue,
			wantErr: true,
		},
		{
			name:    "Invalid Media ID",
			id:      "invalid-id",
			lang:    testLang,
			request: setTrue,
			wantErr: true,
		},
		{
			name:    "Not Exist ID",
			id:      notExistId,
			lang:    testLang,
			request: setTrue,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.Media.SetDefaultCaption(
				tt.id,
				tt.lang,
				tt.request,
			)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
				return
			}
			require.NoError(t, err)
			require.NotNil(t, resp)

			c := findCaption(t, tt.id, tt.lang)
			require.NotNil(
				t,
				c,
				"caption %q not listed for media %s",
				tt.lang,
				tt.id,
			)
			require.NotNil(t, c.IsDefault)
			assert.Equal(
				t,
				*tt.request.IsDefault,
				*c.IsDefault,
				"is_default of caption %q",
				tt.lang,
			)
		})
	}
}

func TestMediaService_DeleteCaption(t *testing.T) {
	requireSetupID(t, "testMediaID", testMediaID)
	notExistId := uuid.New().String()
	anonymousTest := []struct {
		name    string
		id      string
		lang    string
		wantErr bool
	}{
		{
			name:    "Delete other",
			id:      testMediaID,
			lang:    testLang,
			wantErr: true,
		},
	}

	for _, tt := range anonymousTest {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testAnonymousClient.Media.DeleteCaption(tt.id, tt.lang)
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
		lang    string
		wantErr bool
	}{
		{
			name:    "Valid Delete Media Captions",
			id:      testMediaID,
			lang:    testLang,
			wantErr: false,
		},
		{
			name:    "Invalid Media ID",
			id:      "invalid-id",
			lang:    testLang,
			wantErr: true,
		},
		{
			name:    "Empty Media ID",
			id:      "",
			lang:    testLang,
			wantErr: true,
		},
		{
			name:    "Invalid Language",
			id:      testMediaID,
			lang:    "invalid",
			wantErr: true,
		},
		{
			name:    "Empty Language",
			id:      testMediaID,
			lang:    "",
			wantErr: true,
		},
		{
			name:    "Not Exist ID",
			id:      notExistId,
			lang:    testLang,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.Media.DeleteCaption(tt.id, tt.lang)
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

func TestMediaService_Delete(t *testing.T) {
	notExistId := uuid.New().String()
	validMetadata := []Metadata{
		{Key: stringPtr("key1"), Value: stringPtr("value1")},
		{Key: stringPtr("key2"), Value: stringPtr("value2")},
	}

	validTags := []string{"tag1", "tag2"}
	createRequest := CreateMediaRequest{
		Title:       stringPtr("Test Video for Deletion"),
		Description: stringPtr("Test Description"),
		IsPublic:    boolPtr(true),
		Metadata:    &validMetadata,
		Qualities: &[]QualityConfig{
			{
				Type:          stringPtr("hls"),
				ContainerType: stringPtr("mpegts"),
				Resolution:    stringPtr("240p"),
			},
		},
		Tags: &validTags,
	}

	resp, err := testClient.Media.Create(createRequest)
	require.NoError(t, err, "Failed to create media")
	require.NotNil(t, resp)
	require.NotNil(t, resp.Data)
	require.NotNil(t, resp.Data.Id)
	deleteID := *resp.Data.Id

	anonymousTest := []struct {
		name    string
		id      string
		wantErr bool
	}{
		{
			name:    "Delete other",
			id:      deleteID,
			wantErr: true,
		},
	}

	for _, tt := range anonymousTest {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testAnonymousClient.Media.Delete(tt.id)
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
			name:    "Valid Delete",
			id:      deleteID,
			wantErr: false,
		},
		{
			name:    "Invalid Media ID",
			id:      "invalid-id",
			wantErr: true,
		},
		{
			name:    "Empty Media ID",
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
			resp, err := testClient.Media.Delete(tt.id)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
			}
		})
	}
	for _, id := range deleteMediaLater {
		testClient.Media.Delete(id)
	}
}

func stringPtr(s string) *string {
	return &s
}

// requireSetupID stops a test whose subject was supposed to be created by an
// earlier test in the same file (Go runs them in source order).
func requireSetupID(t *testing.T, name, id string) {
	t.Helper()
	if id == "" {
		t.Fatalf(
			"%s is not set: the Create test that produces it failed or did not run",
			name,
		)
	}
}

// readyMediaIDs returns n IDs of fully processed video media for the tests
// that need existing playable media (player themes, playlists, chapters,
// captions). TEST_MEDIA_IDS (comma-separated) takes precedence; otherwise the
// account's media list is queried once and cached for the whole run.
func readyMediaIDs(t *testing.T, n int) []string {
	t.Helper()
	readyMediaOnce.Do(func() {
		readyMediaIDsCache, readyMediaErr = loadReadyMediaIDs()
	})
	if readyMediaErr != nil {
		t.Fatalf("listing ready media: %v", readyMediaErr)
	}
	if len(readyMediaIDsCache) < n {
		t.Skipf(
			"need %d fully processed video media (status done) on the test account, found %d; upload more or set TEST_MEDIA_IDS",
			n,
			len(readyMediaIDsCache),
		)
	}
	return readyMediaIDsCache[:n]
}

// findCaption lists the media's captions through the SDK and returns the one
// in lang, or nil. It fails the test if the list is empty.
func findCaption(t *testing.T, id, lang string) *MediaCaption {
	t.Helper()
	resp, err := testClient.Media.GetCaptions(id, MediaApiGetCaptionsRequest{})
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, resp.Data)
	require.NotNil(t, resp.Data.MediaCaptions)
	require.NotEmpty(
		t,
		*resp.Data.MediaCaptions,
		"media %s has no captions",
		id,
	)
	for _, c := range *resp.Data.MediaCaptions {
		if c.Language != nil && *c.Language == lang {
			c := c
			return &c
		}
	}
	return nil
}

func loadReadyMediaIDs() ([]string, error) {
	if env := os.Getenv("TEST_MEDIA_IDS"); env != "" {
		var ids []string
		for _, id := range strings.Split(env, ",") {
			if id = strings.TrimSpace(id); id != "" {
				ids = append(ids, id)
			}
		}
		return ids, nil
	}

	// "done" is the status of a fully processed media.
	resp, err := testClient.Media.GetMediaList(GetMediaListRequest{
		Status:  &[]string{"done"},
		Type:    stringPtr("video"),
		Limit:   int64Ptr(100),
		SortBy:  stringPtr("created_at"),
		OrderBy: stringPtr("asc"),
	})
	if err != nil {
		return nil, err
	}
	if resp == nil || resp.Data == nil || resp.Data.Media == nil {
		return nil, nil
	}

	var ids []string
	for _, m := range *resp.Data.Media {
		if m.Id != nil && m.Status != nil && *m.Status == "done" {
			ids = append(ids, *m.Id)
		}
	}
	return ids, nil
}
