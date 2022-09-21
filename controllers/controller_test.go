package controllers

import (
	"github.com/gavv/httpexpect/v2"
	"github.com/go-playground/assert/v2"
	"net/http"
	"testing"
)

func TestStage1(t *testing.T) {
	r := InitRouter()
	e := httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(r),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true),
		},
	})
	t.Run("stage1 test", func(t *testing.T) {
		resStr := httpRequestExpectBodyStr(e, "/stage1", "abcd")
		assert.Equal(t, resStr, "bcde")
	})
	t.Run("stage2 test", func(t *testing.T) {
		resStr := httpRequestExpectBodyStr(e, "/stage2", "abcd\nabcd")
		assert.Equal(t, resStr, "bcde\nbcde")
	})
	t.Run("edge case test 1", func(t *testing.T) {
		resStr := httpRequestExpectBodyStr(e, "/stage1", "abcd./\nd")
		assert.Equal(t, resStr, "bcde./\ne")
	})
	t.Run("edge case test 2", func(t *testing.T) {
		bodyStr := "`10-=~!@#$%^&*()_+qp[]QP{}al;'AL:'zm,./ZM<>?"
		resStr := httpRequestExpectBodyStr(e, "/stage2", bodyStr)
		t.Log(resStr)
	})
}

func httpRequestExpectBodyStr(e *httpexpect.Expect, path string, bodyText string) string {
	return e.POST(path).
		WithText(bodyText).
		Expect().
		Status(http.StatusOK).
		Text().
		Raw()
}
