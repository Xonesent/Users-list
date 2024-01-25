package server

type Post_structure struct {
	Name       string `json:"name" binding:"required"`
	Surname    string `json:"surname" binding:"required"`
	Patronymic string `json:"patronymic"`
}

type Patch_structure struct {
	Id          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Surname     string `json:"surname" db:"surname"`
	Patronymic  string `json:"patronymic" db:"patronymic"`
	Age         int    `json:"age" db:"age"`
	Gender      string `json:"gender" db:"gender"`
	Nationality string `json:"nationality" db:"nationality"`
}

type Get_structure struct {
	Id          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Surname     string `json:"surname" db:"surname"`
	Patronymic  string `json:"patronymic" db:"patronymic"`
	Age         int    `json:"age" db:"age"`
	Gender      string `json:"gender" db:"gender"`
	Nationality string `json:"nationality" db:"nationality"`
	Limit       int    `json:"limit" db:"limit"`
	Offset      int    `json:"offset" db:"offset"`
}

type Person_structure struct {
	Name        string `json:"name" db:"name"`
	Surname     string `json:"surname" db:"surname"`
	Patronymic  string `json:"patronymic" db:"patronymic"`
	Age         int    `json:"age" db:"age"`
	Gender      string `json:"gender" db:"gender"`
	Nationality string `json:"nationality" db:"nationality"`
}

type Enricher_structure struct {
	Age         int
	Gender      string
	Nationality string
}

type Age_str struct {
	Age int `json:"age"`
}

type Gender_str struct {
	Gender string `json:"gender"`
}

type Nationality_str struct {
	CountryId   string  `json:"country_id"`
	Probability float32 `json:"probability"`
}

type Nationalities_str struct {
	Country []Nationality_str `json:"country"`
}
