package data

import (
	"encoding/json"
	"io"
)

type Courses struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (c *Courses) FromJSON(r io.Reader) error {
	decode := json.NewDecoder(r)
	return decode.Decode(c)
}

func GetCourse() []*Courses {
	return CourseList
}

var CourseList = []*Courses{
	{
		ID:          1,
		Name:        "HTML",
		Description: "Basics of HTML",
	},
	{
		ID:          2,
		Name:        "GOLANG",
		Description: "Basics of GOLANG",
	},
}
