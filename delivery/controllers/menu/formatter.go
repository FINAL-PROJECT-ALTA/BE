package menu

// ========== Menu Request =========== //

type MenuCreateRequestFormat struct {
	Menu_uid       string
	Menu_category  string `json:"menu_category"`
	Total_calories int    `json:"total_calories"`
}

type MenuUpdateRequestFormat struct {
	Menu_uid       string
	Menu_category  string `json:"menu_category"`
	Total_calories int    `json:"total_calories"`
}

// ========= Menu Response =========== //

type MenuCreateResponse struct {
	Menu_uid       string `json:"menu_uid"`
	Menu_category  string `json:"menu_category"`
	Total_calories int    `json:"total_calories"`
}

type MenuUpdateResponse struct {
	Menu_uid       string `json:"menu_uid"`
	Menu_category  string `json:"menu_category"`
	Total_calories int    `json:"total_calories"`
}

type MenuDeleteResponse struct {
	Menu_uid       string `json:"menu_uid"`
	Menu_category  string `json:"menu_category"`
	Total_calories int    `json:"total_calories"`
}

type Images struct {
	Url string `json:"url"`
}

type Foods struct {
	Food_uid      string   `json:"food_uid"`
	Name          string   `json:"name"`
	Calories      int      `json:"calories"`
	Energy        int      `json:"energy" form:"energy"`
	Carbohidrate  int      `json:"carbohidrate"`
	Protein       int      `json:"protein"`
	Food_category string   `json:"food_categories"`
	Images        []Images `json:"images"`
}

type MenuGetResponse struct {
	Menu_uid       string  `json:"menu_uid"`
	Menu_category  string  `json:"menu_category"`
	Total_calories int     `json:"total_calories"`
	Foods          []Foods `json:"foods"`
}
