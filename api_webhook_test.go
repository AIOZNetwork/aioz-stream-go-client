package aiozstreamsdk

import (
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	testWebhookForUpdateAndDelete string
	// webhookURL defaults to a webhook.site endpoint; TEST_WEBHOOK_URL
	// overrides it (applied in loadEnvVariables, after .env is loaded).
	webhookURL          = "https://webhook.site/b112b207-054e-415b-9540-1eba2bef5001"
	webhookName         = "Test Webhook"
	deleteWebhooksLater []string
)

func boolPtr(b bool) *bool {
	return &b
}

func TestWebhookService_Create(t *testing.T) {
	tests := []struct {
		name    string
		request WriteWebhookRequest
		wantErr bool
	}{
		{
			name: "Valid Create Request with All Fields",
			request: WriteWebhookRequest{
				EncodingFinished: boolPtr(true),
				EncodingStarted:  boolPtr(true),
				FileReceived:     boolPtr(true),
				Name:             stringPtr(webhookName),
				Url:              stringPtr(webhookURL),
			},
			wantErr: false,
		},
		{
			name: "Invalid Create Request Without Events",
			request: WriteWebhookRequest{
				Url:  stringPtr(webhookURL),
				Name: stringPtr(webhookName),
			},
			wantErr: true,
		},
		{
			name: "Invalid URL",
			request: WriteWebhookRequest{
				Url:  stringPtr("not-a-url"),
				Name: stringPtr(webhookName),
			},
			wantErr: true,
		},
		{
			name: "Missing URL",
			request: WriteWebhookRequest{
				Name: stringPtr(webhookName),
			},
			wantErr: true,
		},
		{
			name: "Missing Name",
			request: WriteWebhookRequest{
				Url:              stringPtr(webhookURL),
				EncodingFinished: boolPtr(true),
				EncodingStarted:  boolPtr(true),
				FileReceived:     boolPtr(true),
			},
			wantErr: false,
		},
		{
			name:    "Empty Request",
			request: WriteWebhookRequest{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.Webhook.Create(tt.request)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				require.NoError(t, err)
				require.NotNil(t, resp)
				require.NotNil(t, resp.Data)
				require.NotNil(t, resp.Data.Webhook)
				require.NotNil(t, resp.Data.Webhook.Id)
				assert.NotEmpty(t, *resp.Data.Webhook.Id)
				deleteWebhooksLater = append(
					deleteWebhooksLater,
					*resp.Data.Webhook.Id,
				)
				testWebhookForUpdateAndDelete = *resp.Data.Webhook.Id
			}
		})
	}
}

func TestWebhookService_Update(t *testing.T) {
	requireSetupID(
		t,
		"testWebhookForUpdateAndDelete",
		testWebhookForUpdateAndDelete,
	)
	notExistId := uuid.New().String()
	anonymousTest := []struct {
		name    string
		id      string
		request WriteWebhookRequest
		wantErr bool
	}{
		{
			name: "Update other",
			id:   testWebhookForUpdateAndDelete,
			request: WriteWebhookRequest{
				Name: stringPtr("Updated Webhook"),
			},
			wantErr: true,
		},
	}

	for _, tt := range anonymousTest {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testAnonymousClient.Webhook.Update(tt.id, tt.request)
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
		request WriteWebhookRequest
		wantErr bool
	}{
		{
			name: "Valid Update All Fields",
			id:   testWebhookForUpdateAndDelete,
			request: WriteWebhookRequest{
				EncodingFinished: boolPtr(true),
				EncodingStarted:  boolPtr(false),
				FileReceived:     boolPtr(true),
				Name:             stringPtr("Updated Webhook"),
				Url:              stringPtr(webhookURL),
			},
			wantErr: false,
		},
		{
			name: "Update Partial Fields, only Name",
			id:   testWebhookForUpdateAndDelete,
			request: WriteWebhookRequest{
				Name: stringPtr("Updated Name Only"),
			},
			wantErr: false,
		},
		{
			name: "Empty Name",
			id:   testWebhookForUpdateAndDelete,
			request: WriteWebhookRequest{
				Name: stringPtr(""),
			},
			wantErr: true,
		},
		{
			name: "Blank Name",
			id:   testWebhookForUpdateAndDelete,
			request: WriteWebhookRequest{
				Name: stringPtr("   "),
			},
			wantErr: true,
		},
		{
			name: "Invalid URL",
			id:   testWebhookForUpdateAndDelete,
			request: WriteWebhookRequest{
				Url: stringPtr("not-a-url"),
			},
			wantErr: true,
		},
		{
			name:    "Invalid ID",
			id:      "invalid-id",
			request: WriteWebhookRequest{},
			wantErr: true,
		},
		{
			name:    "Not Exist ID",
			id:      notExistId,
			request: WriteWebhookRequest{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.Webhook.Update(tt.id, tt.request)
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

func TestWebhookService_List(t *testing.T) {
	tests := []struct {
		name    string
		request WebhookApiListRequest
		wantErr bool
		checkFn func(*testing.T, *ListWebhooksResponse)
	}{
		{
			name:    "List All Webhooks",
			request: WebhookApiListRequest{},
			wantErr: false,
			checkFn: func(t *testing.T, resp *ListWebhooksResponse) {
				assert.NotNil(t, resp.Data)
			},
		},
		{
			name: "List with Pagination",
			request: WebhookApiListRequest{}.
				Limit(10).
				Offset(0),
			wantErr: false,
			checkFn: func(t *testing.T, resp *ListWebhooksResponse) {
				require.NotNil(t, resp.Data)
				require.NotNil(t, resp.Data.Webhooks)
				assert.LessOrEqual(t, len(*resp.Data.Webhooks), 10)
			},
		},
		{
			name: "List with Search",
			request: WebhookApiListRequest{}.
				Search("test"),
			wantErr: false,
		},
		{
			name: "List with Event Filters",
			request: WebhookApiListRequest{}.
				EncodingFinished(true).
				EncodingStarted(false).
				FileReceived(true),
			wantErr: false,
		},
		{
			name: "Invalid Offset",
			request: WebhookApiListRequest{}.
				Offset(-1),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testClient.Webhook.List(tt.request)
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

func TestWebhookService_Get(t *testing.T) {
	requireSetupID(
		t,
		"testWebhookForUpdateAndDelete",
		testWebhookForUpdateAndDelete,
	)
	notExistId := uuid.New().String()
	anonymousTest := []struct {
		name    string
		id      string
		wantErr bool
	}{
		{
			name:    "Get other",
			id:      testWebhookForUpdateAndDelete,
			wantErr: true,
		},
	}

	for _, tt := range anonymousTest {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testAnonymousClient.Webhook.Get(tt.id)
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
		checkFn func(*testing.T, *WebhookResponse)
	}{
		{
			name:    "Valid Get",
			id:      testWebhookForUpdateAndDelete,
			wantErr: false,
			checkFn: func(t *testing.T, resp *WebhookResponse) {
				require.NotNil(t, resp.Data)
				require.NotNil(t, resp.Data.Webhook)
				assert.NotEmpty(t, resp.Data.Webhook.Id)
				assert.NotEmpty(t, resp.Data.Webhook.Url)
				assert.NotEmpty(t, resp.Data.Webhook.Name)
			},
		},
		{
			name:    "Invalid ID",
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
			resp, err := testClient.Webhook.Get(tt.id)
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

// TestWebhookService_Check must run before TestWebhookService_Delete removes
// testWebhookForUpdateAndDelete. The backend delivers a real test event to the
// webhook's URL, so the valid case also depends on outbound delivery.
func TestWebhookService_Check(t *testing.T) {
	requireSetupID(
		t,
		"testWebhookForUpdateAndDelete",
		testWebhookForUpdateAndDelete,
	)
	notExistId := uuid.New().String()
	anonymousTest := []struct {
		name    string
		id      string
		wantErr bool
	}{
		{
			name:    "Check other",
			id:      testWebhookForUpdateAndDelete,
			wantErr: true,
		},
	}

	for _, tt := range anonymousTest {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testAnonymousClient.Webhook.Check(tt.id)
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
		name          string
		id            string
		wantErr       bool
		needsDelivery bool
	}{
		{
			name:          "Valid Check",
			id:            testWebhookForUpdateAndDelete,
			wantErr:       false,
			needsDelivery: true,
		},
		{
			name:    "Invalid ID",
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
			if tt.needsDelivery && os.Getenv("TEST_WEBHOOK_URL") == "" {
				t.Skip(
					"Webhook.Check delivers a real test event; set TEST_WEBHOOK_URL to a public endpoint that answers 2xx (the default webhook.site URL returns 404, so the backend answers 502 webhook-check-failed)",
				)
			}
			resp, err := testClient.Webhook.Check(tt.id)
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

func TestWebhookService_Delete(t *testing.T) {
	requireSetupID(
		t,
		"testWebhookForUpdateAndDelete",
		testWebhookForUpdateAndDelete,
	)
	notExistId := uuid.New().String()
	anonymousTest := []struct {
		name    string
		id      string
		wantErr bool
	}{
		{
			name:    "Delete other",
			id:      testWebhookForUpdateAndDelete,
			wantErr: true,
		},
	}

	for _, tt := range anonymousTest {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := testAnonymousClient.Webhook.Delete(tt.id)
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
			id:      testWebhookForUpdateAndDelete,
			wantErr: false,
		},
		{
			name:    "Invalid ID",
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
			resp, err := testClient.Webhook.Delete(tt.id)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
			}
		})
	}

	for _, id := range deleteWebhooksLater {
		testClient.Webhook.Delete(id)
	}
}
