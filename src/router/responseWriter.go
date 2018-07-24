package router

import (
	"net/http"
	"net"
	"bufio"
	"encoding/json"
	"../model"
	"../common"
	"fmt"
)

type responseWriter struct {
	http.ResponseWriter
	wroteHeader bool
}

func (w *responseWriter) WriteHeader(code int) {
	if w.Header().Get("Content-Type") == "" {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
	}
	w.ResponseWriter.WriteHeader(code)
	w.wroteHeader = true
}

func (w *responseWriter) EncodeJson(v interface{}) ([]byte, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (w *responseWriter) WriteCache(s string) error {
	var data map[string]interface{}
	json.Unmarshal([]byte(s), &data)
	fmt.Println(s,data)
	jsonStr, _ := json.Marshal(data)
	fmt.Println(string(jsonStr))
	return nil
}


func (w *responseWriter) WriteJson(v interface{}) error {
	result := &model.Base{}
	result.Code = 0;
	result.Message = ""
	result.Data = v
	b, err := w.EncodeJson(result)
	if err != nil {
		return err
	}
	_, err = w.Write(b)
	if err != nil {
		return err
	}
	return nil
}


func (w *responseWriter) WriteJsonPermission(v interface{}) error {
	b, err := w.EncodeJson(v)
	if err != nil {
		return err
	}
	_, err = w.Write(b)
	if err != nil {
		return err
	}
	return nil
}


func (w *responseWriter) WriteErrorJson(code int) error {
	result := &model.Base{}
	result.Code = code;
	dict := common.ErrorDict()
	result.Message = dict(code)
	b, err := w.EncodeJson(result)
	if err != nil {
		return err
	}
	_, err = w.Write(b)
	if err != nil {
		return err
	}
	return nil
}

func (w *responseWriter) Write(b []byte) (int, error) {
	if !w.wroteHeader {
		w.WriteHeader(http.StatusOK)
	}
	return w.ResponseWriter.Write(b)
}

func (w *responseWriter) Flush() {
	if !w.wroteHeader {
		w.WriteHeader(http.StatusOK)
	}
	flusher := w.ResponseWriter.(http.Flusher)
	flusher.Flush()
}

func (w *responseWriter) CloseNotify() <-chan bool {
	notifier := w.ResponseWriter.(http.CloseNotifier)
	return notifier.CloseNotify()
}

func (w *responseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	hijacker := w.ResponseWriter.(http.Hijacker)
	return hijacker.Hijack()
}
