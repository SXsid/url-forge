package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github/SXsid/url-forge/internal/domain"
)

func (h *URLHandler) Redirect(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")
	if code == "" {
		return
	}
	ctx := r.Context()
	rUrl, err := h.cahche.Get(ctx, code)
	if err != nil {
		fmt.Printf("error whiel fetching form cache for %s", code)
	}
	if rUrl != "" {
		http.Redirect(w, r, rUrl, http.StatusMovedPermanently)
		return
	}
	url, err := h.repo.GetURL(ctx, code)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, domain.ErrCodeNotFound) {
			status = http.StatusNotFound
		}
		http.Error(w, err.Error(), status)
		return
	}
	if err := h.cahche.Set(ctx, code, url); err != nil {
		fmt.Printf("error while updating cahce for %s", code)
	}

	http.Redirect(w, r, url, http.StatusMovedPermanently)
}
