package main

import (
	"net/http"
	"strconv"
)

// handler

func (app *application) Fib(w http.ResponseWriter, r *http.Request) {
	convertedStrUint64, err := strconv.ParseUint(r.URL.Query().Get("n"), 10, 64)

	if err != nil {
		app.errorJSON(w, "自然数を入力してください", http.StatusBadRequest)
		return
	} else if convertedStrUint64 >= 94 || convertedStrUint64 == 0 {
		app.errorJSON(w, "1以上94未満を入力してください", http.StatusBadRequest)
		return
	}

	memo := make(map[uint64]uint64)
	ans := memorize(convertedStrUint64, memo)

	var payload = struct {
		Result uint64 `json:"result"`
	}{
		Result: ans,
	}

	err = app.writeJSON(w, http.StatusOK, payload)
	if err != nil {
		app.errorJSON(w, "jsonの生成に失敗しました。", http.StatusBadRequest)
		return
	}
}

// private

// 再帰メモ化でフィボナッチ数を割り出す
func memorize(n uint64, memo map[uint64]uint64) uint64 {
	if n < 2 {
		return n
	}
	if _, ok := memo[n]; !ok {
		memo[n] = memorize(n-2, memo) + memorize(n-1, memo)
	}
	return memo[n]
}
