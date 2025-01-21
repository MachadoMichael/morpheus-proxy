package handler

import (
	"io"
	"net/http"
	"strings"

	"github.com/MachadoMichael/morpheus-proxy/config"
)

func HandleRequestAndRedirect(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	requestPath := getNormalizedRequestPath(r)
	newPath := getTargetPath(requestPath)
	req := createProxyRequest(r, newPath)

	resp, err := sendProxyRequest(req)
	if err != nil {
		http.Error(w, "Failed to send request to target", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	writeProxyResponse(w, resp)
}

func getNormalizedRequestPath(r *http.Request) string {
	requestPath := r.URL.Path
	return strings.TrimPrefix(requestPath, config.Variables.BaseURL)
}

func getTargetPath(requestPath string) string {
	return config.Variables.TargetURL + requestPath
}

func createProxyRequest(r *http.Request, newPath string) *http.Request {
	req, err := http.NewRequest(r.Method, newPath, r.Body)
	if err != nil {
		panic(err)
	}

	for name, values := range r.Header {
		for _, value := range values {
			req.Header.Add(name, value)
		}
	}
	return req
}

func sendProxyRequest(req *http.Request) (*http.Response, error) {
	client := &http.Client{}
	return client.Do(req)
}

func writeProxyResponse(w http.ResponseWriter, resp *http.Response) {
	for name, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(name, value)
		}
	}

	w.WriteHeader(resp.StatusCode)
	_, err := io.Copy(w, resp.Body)
	if err != nil {
		http.Error(w, "Failed to copy response body", http.StatusInternalServerError)
	}
}
