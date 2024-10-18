package httpserver

import (
	"fmt"
	"net/http"
)

func getVersion(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Classly-v0.1.0")
}
