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
	Name          string `json:"name" form:"name" validate:"required,min=2,max=20,excludesall=!@#?^#*()_+-=0123456789%&"`
	Calories      int    `json:"calories" form:"calories"`
	Energy        int    `json:"energy" form:"energy"`
	Carbohidrate  int    `json:"carbohidrate" form:"carbohidrate"`
	Protein       int    `json:"protein" form:"protein"`
	Unit          string `json:"unit" form:"unit"`
	Unit_value    int    `json:"unit_value" form:"unit_value"`
	Food_category string `json:"food_categories" form:"food_categories"`
	Image         string `json:"image" form:"image"`
}
type FoodsCreateRequestFormatEdamam struct {
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
	Image         string `json:"image"`
}

type FoodsUpdateRequestFormat struct {
	Admin_uid     string
	Food_uid      string
	Name          string `json:"name" form:"name" validate:"required,min=2,max=20,excludesall=!@#?^#*()_+-=0123456789%&"`
	Calories      int    `json:"calories" form:"calories"`
	Energy        int    `json:"energy" form:"energy"`
	Carbohidrate  int    `json:"carbohidrate" form:"carbohidrate"`
	Protein       int    `json:"protein" form:"protein"`
	Unit          string `json:"unit" form:"unit"`
	Unit_value    int    `json:"unit_value" form:"unit_value"`
	Food_category string `json:"food_categories" form:"food_categories"`
	Image         string `json:"image" form:"image"`
}

// ====== API EDAMAM RESPONSE =======

type DetailNutrients struct {
	Enerc_kcal int `json:"ENERC_KCAL"`
	Procnt     int `json:"PROCNT"`
	Fat        int `json:"FAT"`
	Chocdf     int `json:"CHOCDF"`
	Fibtg      int `json:"FIBTG"`
}

type DetailFood struct {
	FoodId        string            `json:"foodId"`
	Label         string            `json:"label"`
	Nutrients     []DetailNutrients `json:"nutrients"`
	Category      string            `json:"category"`
	CategoryLabel string            `json:"categoryLabel"`
	Image         string            `json:"image"`
}

type DetailMeasures struct {
	Uri    string `json:"uri"`
	Label  string `json:"label"`
	Weight string `json:"weight"`
}

type Data struct {
	Food     []DetailFood     `json:"food"`
	Measures []DetailMeasures `json:"measures"`
}

type Response struct {
	Text   string            `json:"text"`
	Parsed map[string]string `json:"parsed"`
	Hints  []Data            `json:"hints"`
}
