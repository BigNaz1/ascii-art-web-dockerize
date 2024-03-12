package pkg

import (
	"fmt"
	"net/http"
)

func Handle400(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>400 - Bad Request Error</h1>")
}
