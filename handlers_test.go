package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type JSONRequest struct {
	Result uint64 `json:"result"`
}

// Fiboの正常系、異常系
func Test_app_fibo(t *testing.T) {
	var app application
	var tests = []struct {
		paramNum        string
		handler         http.HandlerFunc
		expectedStatus  int
		expectedMessage string
		expectedNumber  uint64
	}{
		// 正常系
		{"55", app.Fib, http.StatusOK, "", 139583862445},
		{"32", app.Fib, http.StatusOK, "", 2178309},

		// 異常系
		{"test", app.Fib, http.StatusBadRequest, "自然数を入力してください", 0},
		{"-1", app.Fib, http.StatusBadRequest, "自然数を入力してください", 0},
		{"0", app.Fib, http.StatusBadRequest, "1以上94未満を入力してください", 0},
		{"94", app.Fib, http.StatusBadRequest, "1以上94未満を入力してください", 0},
	}

	for _, e := range tests {
		req, _ := http.NewRequest("GET", "/fib?n=" + e.paramNum, nil)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(e.handler)
		handler.ServeHTTP(rr, req)

		// ステータスコードが一致するか検証
		if rr.Code != e.expectedStatus {
			t.Errorf("%s: wrong status returned; expected %d but got %d", e.paramNum, e.expectedStatus, rr.Code)
		}

		if rr.Code == http.StatusOK {
			var response JSONRequest
			_ = app.readJSON(rr, &response)

			// 想定の計算結果であるかどうか検証
			if response.Result != e.expectedNumber {
				t.Errorf("%s: wrong result returned; expected %d but got %d", e.paramNum, e.expectedNumber, response.Result)
			}
		} else {
			var response ErrorResponse
			_ = app.readJSON(rr, &response)

			// 想定のエラーメッセージか検証
			if response.Message != e.expectedMessage {
				t.Errorf("%s: wrong message returned; expected %s but got %s", e.paramNum, e.expectedMessage, response.Message)
			}
		}

	}
}
