package server

import (
	"net/http"

	"github.com/Marco98/ytpodproxy/pkg/ytdlp"
	log "github.com/sirupsen/logrus"
)

func DownloadAudio(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if err := ytdlp.DownloadAudio(r.Context(), w, url, ytdlp.DownloadOptions{}); err != nil {
		log.WithError(err).Error("failed downloading audio")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
