package model

type (
	Proponents struct {
		Id   int    `json:"id" gorm:"primary key"`
		Name string `json:"name"`
		Term string `json:"term"`
	}

	Committees struct {
		Id   int    `json:"id" gorm:"primary key"`
		Name string `json:"name"`
	}

	Departments struct {
		Id   int    `json:"id" gorm:"primary key"`
		Name string `json:"name"`
	}

	Divisions struct {
		Name string `json:"name"`
		Code string `json:"code"`
	}
)
