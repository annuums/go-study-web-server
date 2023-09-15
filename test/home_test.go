package test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/annuums/go-study-web-server/app"
	"github.com/stretchr/testify/assert"
)

func TestHomeIndex(t *testing.T) {
	assert := assert.New(t)

	homeURI := "/home"
	//* NewHandler가 명시하는 Mock 서버를 생성해요.
	ts := httptest.NewServer(app.NewHandler())
	defer ts.Close()

	//* 서버 주소에 테스트할 homeURI를 붙여요
	testURL := fmt.Sprintf("%s%s", ts.URL, homeURI)

	res, err := http.Get(testURL)
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)

	data, _ := io.ReadAll(res.Body)
	assert.Equal("Hello, Home!\n", string(data))
}