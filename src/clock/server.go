package clock

import (
	"net/http"
	"strings"
	"time"
)

func handler(write http.ResponseWriter, read *http.Request) {
	write.Header().Set("content-type", "text/plain; charset=utf-8")
	path := strings.Trim(read.URL.Path, "/")
	if path == "" {
		path = "local"
	}
	text, err := Render(path, time.Now())
	if err != nil {
		write.WriteHeader(http.StatusNotFound)
		_, _ = write.Write([]byte(err.Error() + "\n" + Help() + "\n"))
		return
	}
	_, _ = write.Write([]byte(text + "\n"))
}

func Serve(addr string) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	return http.ListenAndServe(addr, mux)
}
