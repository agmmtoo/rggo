package main

import (
	"log"
	"net/http"
)

func newMux(todoFile string) http.Handler {
	m := http.NewServeMux()
	m.HandleFunc("/", rootHandler)
	return m
}

func replyTextContent(w http.ResponseWriter, r *http.Request, code int, content string) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(code)
	w.Write([]byte(content))
}

func replyJSONContent(w http.ResponseWriter, r *http.Request, status int, resp *todoResponse) {
	body, err := resp.MarshalJSON()
	if err != nil {
		replyError(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(body)
}

func replyError(w http.ResponseWriter, r *http.Request, status int, message string) {
	log.Printf("%s %s: Error: %d %s", r.URL, r.Method, status, message)
	http.Error(w, http.StatusText(status), status)
}
