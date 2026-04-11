package api

import (
	"fmt"
	"io"
	"net/http"

	"github/SXsid/url-forge/ui"
)

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
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

func Redirect(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusPermanentRedirect)
	fmt.Fprint(w, r.PathValue("code"))
}

func Register(w http.ResponseWriter, r *http.Request) {
}
