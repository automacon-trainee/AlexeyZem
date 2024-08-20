package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	proxy := GetProxy("http://hugo:1313")
	if CheckURL(r.URL.Path) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("hello from API"))
	} else {
		proxy.ServeHTTP(w, r)
	}
}

func GetProxy(str string) *httputil.ReverseProxy {
	target, err := url.Parse(str)
	if err != nil {
		panic(err)
	}
	return httputil.NewSingleHostReverseProxy(target)
}

func CheckURL(str string) bool {
	sl := strings.Split(str, "/")
	if len(sl) < 2 {
		return false
	}
	return sl[1] == "api"
}
