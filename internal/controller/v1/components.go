package v1

import (
	"I_Dev_Kit/cmd/web/components"
	"net/http"
)

func (v *V1) GetNewProjectForm(w http.ResponseWriter, r *http.Request) {
	components.NewProjectForm().Render(r.Context(), w)
}
