package students

import (
	"errors"
)

type Student struct {
	Id       string `json:"id"`
	Name     string `json:"Name"`
	Idnumber string `json:"Number"`
	School   string `json:"School"`
}

var students = []Student{
	{
		Id:      "1",
		Name:   "The Cinnamon Club",
		Idnumber: "The Old Westminster Library, 30-32 Great Smith St, Westminster, London SW1P 3BU",
		School:    "Indian",
	},
	{
		Id:      "2",
		Name:   "The Cinnamon Club",
		Idnumber: "The Old Westminster Library, 30-32 Great Smith St, Westminster, London SW1P 3BU",
		School:    "Indian",
	},
	{
		Id:      "3",
		Name:   "The Cinnamon Club",
		Idnumber: "The Old Westminster Library, 30-32 Great Smith St, Westminster, London SW1P 3BU",
		School:    "Indian",
	},
	{
		Id:      "4",
		Name:   "The Cinnamon Club",
		Idnumber: "The Old Westminster Library, 30-32 Great Smith St, Westminster, London SW1P 3BU",
		School:    "Indian",
	},
	{
		Id:      "5",
		Name:   "The Cinnamon Club",
		Idnumber: "The Old Westminster Library, 30-32 Great Smith St, Westminster, London SW1P 3BU",
		School:    "Indian",
	},
	
}

func GetRestaurants() []Student {
	return students
}

func GetRestaurant(id string) (*Student, error) {
	for _, r := range students {
		if r.Id == id {
			return &r, nil
		}
	}
	return nil, errors.New("Students not found")
}