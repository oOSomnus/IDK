package v1

import (
	"I_Dev_Kit/cmd/web/components"
	"net/http"
)

func (v *V1) GetStats(w http.ResponseWriter, r *http.Request) {
	stats, err := v.f.GetStats()
	if err != nil {
		v.l.Error("failed to get feature stats", "error", err)
		http.Error(w, "failed to get feature stats", http.StatusInternalServerError)
		return
	}
	components.MainQuickStats(stats).Render(r.Context(), w)
}
