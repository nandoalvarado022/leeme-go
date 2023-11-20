package shared

type Employee struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

type Test struct {
	Id int `json:"id"`
}
