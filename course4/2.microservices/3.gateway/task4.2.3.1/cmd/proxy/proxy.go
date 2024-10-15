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
	log.Println("start proxy server")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	service := CheckUrl(path)
	proxy := getProxy(service)
	if proxy != nil {
		proxy.ServeHTTP(w, r)
	} else {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte("Not Found"))
	}
}

func CheckUrl(url string) string {
	sl := strings.Split(url, "/")
	if len(sl) < 3 {
		return ""
	}
	if sl[1] == "api" {
		switch sl[2] {
		case "users":
			return "http://users:1080"
		case "auth":
			return "http://auth:3080"
		case "address":
			return "http://geoservice:2080"
		default:
			return ""
		}
	}

	return ""
}

func getProxy(str string) *httputil.ReverseProxy {
	target, err := url.Parse(str)
	if err != nil {
		log.Println(err)
		return nil
	}

	return httputil.NewSingleHostReverseProxy(target)
}
