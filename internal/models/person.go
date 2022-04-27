package models


type Sex string

const (
	Male   Sex = "male"
	Female Sex = "female"
	Other  Sex = "other"
)

type Person struct {
	UUID                    string  `json:"uuid"`
	Survived                bool    `json:"survived"`
	PassengerClass          int     `json:"passengerClass"`
	Name                    string  `json:"name"`
	Sex                     Sex     `json:"sex"`
	Age                     int     `json:"age"`
	SiblingsOrSpousesAboard int     `json:"siblingsOrSpousesAboard"`
	ParentsOrChildrenAboard int     `json:"parentsOrChildrenAboard"`
	Fare                    float32 `json:"fare"`
}
