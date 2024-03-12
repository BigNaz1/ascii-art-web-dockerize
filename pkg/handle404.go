package pkg

import (
	"fmt"
	"net/http"
)

func Handle404(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>404 - Page not found</h1>")
}
