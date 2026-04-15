package handler

import (
	"fmt"
	"io"
	"net/http"

	"github/SXsid/url-forge/ui"
)

func (h *URLHandler) HomePageHandler(w http.ResponseWriter, r *http.Request) {
	f, err := ui.StaticFS.Open("index.html")
	if err != nil {
		http.Error(w, "server side error", http.StatusInternalServerError)
	}
	defer f.Close()
	if _, err := io.Copy(w, f); err != nil {
		fmt.Println(err)
		http.Error(w, "server side error", http.StatusInternalServerError)
	}
}

func (h *URLHandler) GetAnylitics(w http.ResponseWriter, r *http.Request) {
}
