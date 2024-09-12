package handlers

import (
	"net/http"
	"strings"
)

type Message struct {
	Message string `json:"message"`
}

func GetMessage(w http.ResponseWriter, r *http.Request) {
	msg := Message{
		Message: "HI I am from Golang, Handler,",
	}
	cleanMessage := strings.ReplaceAll(msg.Message, "\"\"", "")
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(`<h2 id=\"message\">` + cleanMessage + `</h2>
	<form hx-post="/submit" hx-target="#response">
            <input type="text" class="border p-2" name="inputField" placeholder="Enter text">
            <button type="submit" class="px-4 py-2 bg-blue-500 text-white rounded">Submit</button>
        </form>
        <div id="response"></div>`))

}

func Submithandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Unable to parseForm ", http.StatusBadRequest)
		return
	}
	input := r.FormValue("inputField")
	response := "<div id=\"response\"> You submitted:" + input + "</div>"
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(response))
}
