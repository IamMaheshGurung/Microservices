package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"xample.com/m/data"
)

type CourseHandler struct {
	l *log.Logger
}

func NewCourse(l *log.Logger) *CourseHandler {
	return &CourseHandler{l}
}

func (c CourseHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		c.getCourse(w, r)
	case http.MethodPost:
		c.addCourse(w, r)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}

func (c CourseHandler) getCourse(w http.ResponseWriter, r *http.Request) {
	cl := data.GetCourse()
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(cl)
	if err != nil {
		http.Error(w, "Unable to Marshall JSON", http.StatusInternalServerError)
	}

}

func (c CourseHandler) addCourse(w http.ResponseWriter, r *http.Request) {
	c.l.Println("HTTP POST Method Handle")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Cannot read input", http.StatusBadRequest)
		return
	}

	c.l.Println("The Body input is :", string(body))

	bodyCourse := &data.Courses{}
	err = bodyCourse.FromJSON(bytes.NewReader(body))
	if err != nil {
		http.Error(w, "Unable to read the error", http.StatusBadRequest)
	}
	c.l.Printf("The Courses are %#v", bodyCourse)

}
