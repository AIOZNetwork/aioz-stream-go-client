package aiozstreamsdk

import (
	"io"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	testLang       = "en"
	chapterContent = `WEBVTT

00:00:00.000 --> 00:01:00.000
Chapter 1

00:01:00.000 --> 00:02:00.000
Chapter 2`
)

func createTempVTTFile(t *testing.T) *os.File {
	tmpFile, err := os.CreateTemp("", "test-*.vtt")
	if err != nil {
		t.Fatal(err)
	}

	_, err = tmpFile.WriteString(chapterContent)
	if err != nil {
		t.Fatal(err)
	}

	tmpFile.Seek(0, 0)
	return tmpFile
}

func TestMediaChapterService_Create(t *testing.T) {
	testMediaIDForChapter := readyMediaIDs(t, 1)[0]
	notExistId := uuid.New().String()
	tmpFile := createTempVTTFile(t)
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	tests := []struct {
		name    string
		mediaID string
		lang    string
		file    *os.File
		wantErr bool
	}{
		{
			name:    "Valid Create",
			mediaID: testMediaIDForChapter,
			lang:    testLang,
			file:    tmpFile,
			wantErr: false,
		},
		{
			name:    "Invalid Media ID",
			mediaID: "invalid-id",
			lang:    testLang,
			file:    tmpFile,
			wantErr: true,
		},
		{
			name:    "Invalid Language",
			mediaID: testMediaIDForChapter,
			lang:    "invalid",
			file:    tmpFile,
			wantErr: true,
		},
		{
			name:    "Not Exist ID",
			mediaID: notExistId,
			lang:    testLang,
			file:    tmpFile,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var reader io.Reader
			var fileName string
			if tt.file != nil {
				reader = tt.file
				fileName = tt.file.Name()
			}

			resp, err := testClient.MediaChapter.Create(
				tt.mediaID,
				tt.lang,
				fileName,
				reader,
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

func TestMediaChapterService_Get(t *testing.T) {
	testMediaIDForChapter := readyMediaIDs(t, 1)[0]
	notExistId := uuid.New().String()
	anonymousTest := []struct {
		name    string
		mediaID string
		wantErr bool
	}{
		{
			name:    "Get other",
			mediaID: testMediaIDForChapter,
			wantErr: true,
		},
	}

	for _, tt := range anonymousTest {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testAnonymousClient.MediaChapter.Get(
				tt.mediaID,
				MediaChapterApiGetRequest{},
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
		mediaID string
		request MediaChapterApiGetRequest
		wantErr bool
		checkFn func(*testing.T, *GetMediaChaptersResponse)
	}{
		{
			name:    "Valid Get",
			mediaID: testMediaIDForChapter,
			request: MediaChapterApiGetRequest{}.
				Limit(10).
				Offset(0),
			wantErr: false,
			checkFn: func(t *testing.T, resp *GetMediaChaptersResponse) {
				assert.NotNil(t, resp.Data)
			},
		},
		{
			name:    "Invalid Media ID",
			mediaID: "invalid-id",
			request: MediaChapterApiGetRequest{},
			wantErr: true,
		},
		{
			name:    "Not Exist ID",
			mediaID: notExistId,
			request: MediaChapterApiGetRequest{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.MediaChapter.Get(tt.mediaID, tt.request)
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

func TestMediaChapterService_Delete(t *testing.T) {
	testMediaIDForChapter := readyMediaIDs(t, 1)[0]
	notExistId := uuid.New().String()
	anonymousTest := []struct {
		name    string
		mediaID string
		lang    string
		wantErr bool
	}{
		{
			name:    "Delete other",
			mediaID: testMediaIDForChapter,
			lang:    testLang,
			wantErr: true,
		},
	}

	for _, tt := range anonymousTest {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testAnonymousClient.MediaChapter.Delete(
				tt.mediaID,
				tt.lang,
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
		mediaID string
		lang    string
		wantErr bool
	}{
		{
			name:    "Valid Delete",
			mediaID: testMediaIDForChapter,
			lang:    testLang,
			wantErr: false,
		},
		{
			name:    "Invalid Media ID",
			mediaID: "invalid-id",
			lang:    testLang,
			wantErr: true,
		},
		{
			name:    "Invalid Language",
			mediaID: testMediaIDForChapter,
			lang:    "invalid",
			wantErr: true,
		},
		{
			name:    "Empty Media ID",
			mediaID: "",
			lang:    testLang,
			wantErr: true,
		},
		{
			name:    "Empty Language",
			mediaID: testMediaIDForChapter,
			lang:    "",
			wantErr: true,
		},
		{
			name:    "Not Exist ID",
			mediaID: notExistId,
			lang:    testLang,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.MediaChapter.Delete(tt.mediaID, tt.lang)
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
