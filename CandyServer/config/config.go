package config

type Order struct {
	Money      int    `json:"money"`
	CandyType  string `json:"candyType"`
	CandyCount int    `json:"candyCount"`
}

type Response struct {
	Change   int    `json:"change,omitempty"`
	Thanks   string `json:"thanks,omitempty"`
	ErrorMsg string `json:"errorMsg,omitempty"`
}

var CandyPrice = map[string]int{
	"CE": 10,
	"AA": 15,
	"NT": 17,
	"DE": 21,
	"YR": 23,
}
