package handler

import (
	"fmt"
	"net/http"
)

func (h *URLHandler) Redirect(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusPermanentRedirect)
	fmt.Fprint(w, r.PathValue("code"))
}
