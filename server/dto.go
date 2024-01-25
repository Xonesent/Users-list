package server

type Post_structure struct {
	Name string `json:"name" binding:"required"`
	Surname string `json:"surname" binding:"required"`
	Patronymic string `json:"patronymic"`
}

type Person_structure struct {
	Name        string `db:"name"`
	Surname     string `db:"surname"`
	Patronymic  string `db:"patronymic"`
	Age         uint   `db:"age"`
	Gender      string `db:"gender"`
	Nationality string `db:"nationality"`
}

type Enricher_structure struct {
	Age         int   
	Gender      string 
	Nationality string 
}

type Age_str struct {
	Age         int `json:"age"`
}

type Gender_str struct {
	Gender      string `json:"gender"`
}

type Nationality_str struct {
	CountryId   string  `json:"country_id"`
	Probability float32 `json:"probability"`
}

type Nationalities_str struct {
	Country []Nationality_str `json:"country"`
}