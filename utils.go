package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
)

type ErrorResponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Jsonへの変換
func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		//引数でなくてもいいようにspliteを書いているから、[0]をつける必要がある
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	//二個目以降のheaderをheaderにapeendする
	w.WriteHeader(status)
	_, err = w.Write(out)
	if err != nil {
		return err
	}

	return nil
}

// Jsonの読み取り
func (app *application) readJSON(rr *httptest.ResponseRecorder, data interface{}) error {
	dec := json.NewDecoder(rr.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(data)
	if err != nil {
		return err
	}

	return nil
}

// エラーJsonの作成
func (app *application) errorJSON(w http.ResponseWriter, message string, status ...int) error {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	var payload ErrorResponse
	payload.Error = true
	payload.Message = message

	return app.writeJSON(w, statusCode, payload)
}
