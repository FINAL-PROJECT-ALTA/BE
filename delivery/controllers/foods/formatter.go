package foods

type FoodsCreateResponse struct {
	Food_uid      string `json:"food_uid"`
	Name          string `json:"name"`
	Calories      int    `json:"calories"`
	Energy        int    `json:"energy" form:"energy"`
	Carbohidrate  int    `json:"carbohidrate"`
	Protein       int    `json:"protein"`
	Unit          string `json:"unit"`
	Unit_value    int    `json:"unit_value"`
	Food_category string `json:"food_categories"`
	Image         string `json:"image"`
}

type FoodsUpdateResponse struct {
	Food_uid      string `json:"food_uid"`
	Name          string `json:"name"`
	Calories      int    `json:"calories"`
	Energy        int    `json:"energy" form:"energy"`
	Carbohidrate  int    `json:"carbohidrate"`
	Protein       int    `json:"protein"`
	Unit          string `json:"unit"`
	Unit_value    int    `json:"unit_value"`
	Food_category string `json:"food_categories"`
	Image         string `json:"image"`
}

type FoodsSearchResponse struct {
	Food_uid      string `json:"food_uid"`
	Name          string `json:"name"`
	Calories      int    `json:"calories"`
	Energy        int    `json:"energy" form:"energy"`
	Carbohidrate  int    `json:"carbohidrate"`
	Protein       int    `json:"protein"`
	Unit          string `json:"unit"`
	Unit_value    int    `json:"unit_value"`
	Food_category string `json:"food_categories"`
	Image         string `json:"image"`
}

type FoodsGetAllResponse struct {
	Food_uid      string `json:"food_uid"`
	Name          string `json:"name"`
	Calories      int    `json:"calories"`
	Energy        int    `json:"energy" form:"energy"`
	Carbohidrate  int    `json:"carbohidrate"`
	Protein       int    `json:"protein"`
	Unit          string `json:"unit"`
	Unit_value    int    `json:"unit_value"`
	Food_category string `json:"food_categories"`
	Image         string `json:"image"`
}

type FoodsCreateRequestFormat struct {
	Admin         string
	Food_uid      string
	Name          string `json:"name" validate:"required,min=2,max=20,excludesall=!@#?^#*()_+-=0123456789%&"`
	Calories      int    `json:"calories"`
	Energy        int    `json:"energy" form:"energy"`
	Carbohidrate  int    `json:"carbohidrate"`
	Protein       int    `json:"protein"`
	Unit          string `json:"unit"`
	Unit_value    int    `json:"unit_value"`
	Food_category string `json:"food_categories"`
}

type FoodsUpdateRequestFormat struct {
	Admin_uid     string
	Food_uid      string
	Name          string `json:"name" validate:"omitempty,min=2,max=20,excludesall=!@#?^#*()_+-=0123456789%&"`
	Calories      int    `json:"calories"`
	Energy        int    `json:"energy" form:"energy"`
	Carbohidrate  int    `json:"carbohidrate"`
	Protein       int    `json:"protein"`
	Unit          string `json:"unit"`
	Unit_value    int    `json:"unit_value"`
	Food_category string `json:"food_categories"`
}
