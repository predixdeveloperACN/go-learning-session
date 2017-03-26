package model

type Patient struct {
	Id   int `json:"id,omitempty"  db:"id"`
	Name string `json:"name,omitempty"  db:"name,"`
	Weight int `json:"weight,omitempty"  db:"weight"`
	Height int `json:"height,omitempty"  db:"height"`
	Age int `json:"age,omitempty"  db:"age"`
	Gender string `json:"gender"  db:"gender"`
	BustSize int `json:"bustSize,omitempty"  db:"bustsize"`
	WaistSize int `json:"waistSize,omitempty"  db:"waistsize"`
	HipSize int `json:"hipSize,omitempty"  db:"hipsize"`
}
