package api

import (
	"net/http"
	"testing"
	"time"

	"github.com/gavv/httpexpect"
	"github.com/gin-gonic/gin"
)

var httpExpectConf *httpexpect.Config

func HttpExpectConf(t *testing.T) *httpexpect.Config {
	if httpExpectConf != nil {
		s, _ := EngineServer(gin.ReleaseMode)
		return &httpexpect.Config{
			Client: &http.Client{
				Transport: httpexpect.NewBinder(s.Engine()),
				Jar:       httpexpect.NewJar(),
				Timeout:   time.Second * 30,
			},
			Reporter: httpexpect.NewRequireReporter(t),
		}
	}
	return httpExpectConf
}
func HttpExpectConfDebug(t *testing.T) *httpexpect.Config {
	if httpExpectConf != nil {
		c := HttpExpectConf(t)
		c.Printers = []httpexpect.Printer{
			httpexpect.NewCurlPrinter(t),
			httpexpect.NewDebugPrinter(t, true),
		}
	}
	return httpExpectConf
}
func TestDBFillTestData(t *testing.T) {
	httpexpect.WithConfig(*HttpExpectConf(t)).
		GET("/db/fill-test-data").
		WithHeader("tag", "sqlite1").
		Expect().
		Status(http.StatusOK).
		JSON().Object().Keys().ContainsOnly("size", "table", "tag")
}
