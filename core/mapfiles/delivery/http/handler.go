package mapFilesHttp

import (
	"GOSecretProject/core/mapfiles/model/base"
	"encoding/json"
	"net/http"
	"time"
)

type Handler struct {

}

func NewHandler() *Handler {
	return &Handler{

	}
}

func (h *Handler) GetMapFiles(w http.ResponseWriter, r *http.Request) {
	mapFiles := []base.MapFile{
		base.MapFile{
			ID:      1,
			Name:    "Карта",
			Tag:     "ddwada",
			JSON:    "{0: { lines: [{432, 240, 528, 240}, {528, 128, 432, 128}, {528, 240, 528, 128}, {432, 128, 432, 240}]}, 1: { lines: [{144, 64, 384, 224}, {560, 96, 560, 192}, {496, 96, 432, 192}, {672, 96, 672, 192}]}, 2: { lines: [{544, 112, 448, 224}, {448, 224, 544, 224}, {448, 112, 544, 112}]}, 4: { lines: [{432, 160, 592, 240}]}}",
			Changed: time.Now(),
		},
		base.MapFile{
			ID:      2,
			Name:    "Карта два",
			Tag:     "adwww",
			JSON:    "{1: { lines: [{160, 192, 288, 16}, {128, 64, 304, 176}, {128, 272, 336, 240}]}}",
			Changed: time.Now(),
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mapFiles)
}