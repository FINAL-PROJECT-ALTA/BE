package menu

import "HealthFit/entities"

// ========== Menu Request =========== //

type MenuCreateRequestFormat struct {
	Menu_uid      string
	User_uid      string
	Menu_category string `json:"menu_category"`
	Created_by    string
	Foods         []entities.Food `json:"foods"`
}

type MenuUpdateRequestFormat struct {
	Menu_uid      string
	User_uid      string
	Menu_category string          `json:"menu_category"`
	Foods         []entities.Food `json:"foods"`
}

// ========= Menu Response =========== //

type MenuCreateResponse struct {
	Menu_uid       string `json:"menu_uid"`
	Menu_category  string `json:"menu_category"`
	Total_calories int    `json:"total_calories"`
	Created_by     string
	Foods          []entities.Food `json:"foods"`
}

type MenuUpdateResponse struct {
	Menu_uid       string `json:"menu_uid"`
	Menu_category  string `json:"menu_category"`
	Total_calories int    `json:"total_calories"`
	Created_by     string
	Foods          []entities.Food `json:"foods"`
}

type MenuDeleteResponse struct {
	Menu_uid      string `json:"menu_uid"`
	Menu_category string `json:"menu_category"`
}

type Foods struct {
	Food_uid      string `json:"food_uid"`
	Name          string `json:"name"`
	Calories      int    `json:"calories"`
	Energy        int    `json:"energy" form:"energy"`
	Carbohidrate  int    `json:"carbohidrate"`
	Protein       int    `json:"protein"`
	Unit          string `json:"unit"`
	Unit_value    int    `json:"unit_value"`
	Food_category string `json:"food_categories"`
	Images        string `json:"images"`
}

type Detail_menu struct {
	Menu_uid string `json:"menu_uid"`
	Food_uid string `json:"food_uid"`
}

type MenuGetAllResponse struct {
	Menu_uid       string          `json:"menu_uid"`
	Menu_category  string          `json:"menu_category"`
	Total_calories int             `json:"total_calories"`
	Created_by     string          `json:"created_by"`
	Foods          []entities.Food `json:"foods"`
}
