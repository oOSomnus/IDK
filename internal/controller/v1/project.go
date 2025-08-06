package v1

import (
	"I_Dev_Kit/cmd/web/components"
	"net/http"
	"strconv"
)

func (v *V1) GetProjectList(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	if page == "" {
		page = "1"
	}
	pageNum, err := strconv.Atoi(page)
	if err != nil {
		v.l.Error("invalid page number", "error", err)
		http.Error(w, "invalid page number", http.StatusBadRequest)
		return
	}
	projects, err := v.p.GetProjectsByPage(pageNum)
	if err != nil {
		v.l.Error("failed to get projects", "error", err)
		http.Error(w, "failed to get projects", http.StatusInternalServerError)
		return
	}
	components.InfiniteScrollProjectList(projects, pageNum+1).Render(r.Context(), w)
}
