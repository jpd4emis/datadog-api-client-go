package test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestTelemetryHeaders(t *testing.T) {
	ctx, stop := WithClient(WithFakeAuth(context.Background()), t)
	defer stop()

	// Mock a random endpoint and make sure we send the operation id header. Return an arbitrary success response code.
	gock.New("https://api.datadoghq.com").
		Get("dashboard/lists/manual/1234/dashboards").
		MatchHeader("DD-OPERATION-ID", "GetDashboardListItems").
		MatchHeader("User-Agent", "^datadog-api-client-go/\\d\\.\\d\\.\\d.*? \\(go .*?; os .*?; arch .*?\\)$").
		Reply(299)
	defer gock.Off()

	_, httpresp, err := Client(ctx).DashboardListsApi.GetDashboardListItems(ctx, 1234).Execute()
	assert.Nil(t, err)
	assert.Equal(t, 299, httpresp.StatusCode)
}