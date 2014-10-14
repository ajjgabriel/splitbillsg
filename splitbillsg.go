package splitbillsg

import (
    "html/template"
    "net/http"
)



func init() {
    http.HandleFunc("/", root)
	http.Handle("/css/", http.FileServer(http.Dir(".")))
	http.Handle("/js/", http.FileServer(http.Dir(".")))
}


func root(w http.ResponseWriter, r *http.Request) {
	
	whatsYourTimeNowForm.ExecuteTemplate(w, "splitbillsg.htm", "");
}


var whatsYourTimeNowForm = template.Must(template.New("").ParseFiles("splitbillsg.htm"))