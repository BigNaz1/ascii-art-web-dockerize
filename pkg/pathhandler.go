package pkg

import (
	"net/http"
)

func PathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/500":
		Handle500(w, r)
	case "/400":
		Handle400(w, r)
	default:
		Handle404(w, r)
	}

}
