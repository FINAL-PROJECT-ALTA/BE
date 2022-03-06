package foods

import "HealthFit/entities"

type FoodsCreateResponse struct {
	Food_uid      string `json:"food_uid"`
	Name          string `json:"name"`
	Calories      int    `json:"calories"`
	Energy        int    `json:"energy" form:"energy"`
	Carbohidrate  int    `json:"carbohidrate"`
	Protein       int    `json:"protein"`
	Food_category string `json:"food_categories"`
	Images        entities.Image
}

type FoodsUpdateResponse struct {
	Food_uid      string `json:"food_uid"`
	Name          string `json:"name"`
	Calories      int    `json:"calories"`
	Energy        int    `json:"energy" form:"energy"`
	Carbohidrate  int    `json:"carbohidrate"`
	Protein       int    `json:"protein"`
	Food_category string `json:"food_categories"`
	Images        entities.Image
}

type FoodsSearchResponse struct {
	Food_uid      string `json:"food_uid"`
	Name          string `json:"name"`
	Calories      int    `json:"calories"`
	Energy        int    `json:"energy" form:"energy"`
	Carbohidrate  int    `json:"carbohidrate"`
	Protein       int    `json:"protein"`
	Food_category string `json:"food_categories"`
	Images        entities.Image
}

type FoodsGetAllResponse struct {
	Food_uid      string `json:"food_uid"`
	Name          string `json:"name"`
	Calories      int    `json:"calories"`
	Energy        int    `json:"energy" form:"energy"`
	Carbohidrate  int    `json:"carbohidrate"`
	Protein       int    `json:"protein"`
	Food_category string `json:"food_categories"`
	Images        entities.Image
}

type Images struct {
	Url string `json:"url"`
}

type FoodsCreateRequestFormat struct {
	Admin         string
	Food_uid      string
	Name          string `json:"name"`
	Calories      int    `json:"calories"`
	Energy        int    `json:"energy" form:"energy"`
	Carbohidrate  int    `json:"carbohidrate"`
	Protein       int    `json:"protein"`
	Food_category string `json:"food_categories"`
}

type FoodsUpdateRequestFormat struct {
	Admin_uid     string
	Food_uid      string
	Name          string `json:"name"`
	Calories      int    `json:"calories"`
	Energy        int    `json:"energy" form:"energy"`
	Carbohidrate  int    `json:"carbohidrate"`
	Protein       int    `json:"protein"`
	Food_category string `json:"food_categories"`
}

type FoodsSearchRequestFormat struct {
	Name     string `json:"name"`
	Calories int    `json:"calories"`
}
