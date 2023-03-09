package http

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	go_gin "golang.source-fellows.com/seminar/go-gin/v2"
	"golang.source-fellows.com/seminar/go-gin/v2/mocks"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleAllAuto_success(t *testing.T) {
	controller := gomock.NewController(t)
	repository := mocks.NewMockAutoRepository(controller)

	var autos []go_gin.Auto
	repository.EXPECT().GetAllAuto().Return(autos, nil)

	r := gin.Default()
	r.GET("/", handleGetAutos(repository))

	tserver := httptest.NewServer(r)
	defer tserver.Close()

	response, err := http.Get(tserver.URL)

	if response.StatusCode != http.StatusOK {
		t.Errorf("wrong status code: %v", response.StatusCode)
		return
	}

	if err != nil {
		t.Errorf("should not be an error: %v", err)
		return
	}
	if response == nil {
		t.Error("response should be successful but is nil")
	}
}
