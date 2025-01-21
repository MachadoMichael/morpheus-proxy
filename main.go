package main

import (
	"net/http"

	"github.com/MachadoMichael/morpheus-proxy/config"
	"github.com/MachadoMichael/morpheus-proxy/interceptor"
)

func main() {

	err := config.Init()
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", interceptor.GetRequest)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

}

// func handleRequestAndRedirect(w http.ResponseWriter, r *http.Request) {
// 	// CORS headers
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
// 	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
//
// 	if r.Method == http.MethodOptions {
// 		w.WriteHeader(http.StatusOK)
// 		return
// 	}
//
// 	requestPath := r.URL.Query().Get("url")
// 	requestPath = strings.TrimPrefix(requestPath, Variables.BaseURL)
// 	newPath := Variables.TargetURL + requestPath
//
// 	req, err := http.NewRequest(r.Method, newPath, r.Body)
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	// Copy the headers from the original request
// 	for name, values := range r.Header {
// 		for _, value := range values {
// 			req.Header.Add(name, value)
// 		}
// 	}
//
// 	// Use http.Client to send the request
// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		http.Error(w, "Failed to send request to target", http.StatusInternalServerError)
// 		return
// 	}
// 	defer resp.Body.Close()
//
// 	// Copy the response headers to the response writer
// 	for name, values := range resp.Header {
// 		for _, value := range values {
// 			w.Header().Add(name, value)
// 		}
// 	}
//
// 	w.WriteHeader(resp.StatusCode)
// 	_, err = io.Copy(w, resp.Body)
// 	if err != nil {
// 		http.Error(w, "Failed to copy response body", http.StatusInternalServerError)
// 		return
// 	}
//
// }
