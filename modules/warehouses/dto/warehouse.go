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
	X int `json:"x" binding:"min=0,max=25"`
	Y int `json:"y" binding:"min=0,max=25"`
}
