package pkg

import (
	"net/http"
	"os"
	"text/template"
)

type DData struct{
	Color string
	Res string
}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	filename := "markup.html"
	text, err := os.ReadFile(filename)
	webMarkup := string(text)
	if err != nil {
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}
	// var res string
	var k string
	var err2 error
	apple := &DData{Color:"#FFFFFF" , Res: "" }
	if r.Method == "POST" {
		err1 := r.ParseForm()
		if err1 != nil {
			http.Redirect(w, r, "/500", http.StatusSeeOther)
			return
		}
		inputText := r.PostFormValue("inputcontent")
		font := r.PostFormValue("fonts")
		apple.Color = r.PostFormValue("color")
		k, err2 = TextProcessor(inputText, font)
		if err2 != nil {
			http.Redirect(w, r, "/400", http.StatusSeeOther)
			return
		}
		apple.Res=k

	}
	temp, err3 := template.New("something").Parse(webMarkup)
	if err3 != nil {
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}
	err4 := temp.Execute(w, apple)
	if err4 != nil {
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
}
