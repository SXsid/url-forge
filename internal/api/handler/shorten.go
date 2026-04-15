package handler

import (
	"fmt"
	"net/http"
)

func (h *URLHandler) Register(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "forge.shekharx.in/abc")
}
