package router

import (
	"net/http"
	"strings"
	"text/template"
)

func chatRoomHandler(rw http.ResponseWriter, r *http.Request) {
	nome := strings.TrimSpace(r.FormValue("name"))
	confirm := r.FormValue("confirm")
	if r.Method != "POST" || nome == "" || confirm != "on" {
		http.Redirect(rw, r, "/", http.StatusSeeOther)
		return
	}

	data := map[string]interface{}{
		"UserName": nome,
	}

	tmpl, err := template.ParseFiles("templates/chat.html")
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	if err = tmpl.Execute(rw, data); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

}
