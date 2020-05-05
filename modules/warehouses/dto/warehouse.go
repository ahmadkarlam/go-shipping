package dto

type Warehouse struct {
	Code  string `json:"code"`
	Stock int    `json:"stock"`
	X     int    `json:"x"`
	Y     int    `json:"y"`
}

type SendingCost struct {
	From     Warehouse
	Distance int
	Cost     int
	Day      int
	Hour     int
}

type SendVaccineToLocationRequest struct {
	X int `json:"x"`
	Y int `json:"y"`
}
