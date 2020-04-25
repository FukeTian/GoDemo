package framework

import (
	"encoding/json"
	"io"
	"net/http"
)

// ResultOk 返回成功
func ResultOk(w http.ResponseWriter, data string) {
	io.WriteString(w, data)
}

// ResultFail 返回失败
func ResultFail(w http.ResponseWriter, err string) {
	http.Error(w, err, http.StatusBadRequest)
}

// ResultJSONOK 返回成功
func ResultJSONOK(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json, _ := json.Marshal(data)
	w.Write(json)
}
