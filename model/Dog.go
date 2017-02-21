package model

type Dog struct {
    Id int `json:"id,omitempty"  db:"id"`
    Breed string `json:"breed,omitempty"  db:"breed"`
    Weight int `json:"weight,omitempty"  db:"weight"`
    Height int `json:"height,omitempty"  db:"height"`
}
