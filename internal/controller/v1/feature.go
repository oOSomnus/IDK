package v1

import (
	"I_Dev_Kit/cmd/web/components"
	"net/http"
)

func (v *V1) GetStats(w http.ResponseWriter, r *http.Request) {
	stats := v.f.GetStats()
	components.MainQuickStats(stats).Render(r.Context(), w)
}
