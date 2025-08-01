package tests

import (
	"net/http"
	"testing"

	"github.com/juanMaAV92/go-server-template/cmd/handlers/health"
	healthService "github.com/juanMaAV92/go-server-template/internal/services/health"
	"github.com/juanMaAV92/go-server-template/tests/helpers"
	"github.com/juanMaAV92/go-utils/pointers"
	"github.com/juanMaAV92/go-utils/testhelpers"
	"github.com/stretchr/testify/assert"
)

func Test_HealthCheck(t *testing.T) {

	path := "/"
	cases := []testhelpers.HttpTestCase{
		{
			TestName: "HealthCheck",
			Request: testhelpers.TestRequest{
				Method:    "GET",
				Url:       path,
				PathParam: []testhelpers.TestPathParam{},
				Header:    map[string]string{},
			},
			Response: testhelpers.ExpectedResponse{
				Status: http.StatusOK,
				Body:   pointers.Pointer(`{"status":"OK"}`),
			},
		},
	}

	app := helpers.NewTestServer()
	service := healthService.NewService()
	handler := health.NewHandler(service)

	for _, test := range cases {
		t.Run(test.TestName, func(t *testing.T) {
			ctx, recorder := testhelpers.PrepareContextFormTestCase(app.Server.Echo, test)
			err := handler.Check(ctx)
			assert.NoError(t, err)
			assert.Equal(t, test.Response.Status, recorder.Code)
			assert.JSONEq(t, *test.Response.Body, recorder.Body.String())
		})
	}
}
