package entity

type Animal struct {
	CommonName 		string	`json:"CommonName"`
	ScientificName 	string	`json:"ScientificName"`
	Type 			string	`json:"Type"`
	Diet 			string	`json:"Diet"`
	AverageLife 	string	`json:"AverageLife"`
	Size 			string	`json:"Size"`
	Weight 			string	`json:"Weight"`
}