package handler

import (
	"errors"
	"net/http"

	"github/SXsid/url-forge/internal/domain"
)

func (h *URLHandler) Redirect(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")
	if code == "" {
		return
	}
	ctx := r.Context()
	url, err := h.repo.GetURL(ctx, code)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, domain.ErrCodeNotFound) {
			status = http.StatusNotFound
		}
		http.Error(w, err.Error(), status)
		return
	}

	http.Redirect(w, r, url, http.StatusMovedPermanently)
}
