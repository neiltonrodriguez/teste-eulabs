package domain

type Product struct {
	Id  int `json:"id"`
	Name  string `json:"name"`
	Description string `json:"description"`
	Value float64 	`json:"value"`

}

type ProductDTO struct {
	Id  int `json:"id"`
	Name  string `json:"name"`
	Description string `json:"description"`
	Value float64 	`json:"value"`
	
}