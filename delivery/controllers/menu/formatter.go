package menu

// ========== Menu Request =========== //

type MenuCreateRequestFormat struct {
	Menu_uid      string
	Menu_category string `json:"menu_category"`
}

type MenuUpdateRequestFormat struct {
	Menu_uid      string
	Menu_category string `json:"menu_category"`
}

// ========= Menu Response =========== //

type MenuCreateResponse struct {
	Menu_uid      string `json:"menu_uid"`
	Menu_category string `json:"menu_category"`
}

type MenuUpdateResponse struct {
	Menu_uid      string `json:"menu_uid"`
	Menu_category string `json:"menu_category"`
}

type MenuDeleteResponse struct {
	Menu_uid      string `json:"menu_uid"`
	Menu_category string `json:"menu_category"`
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
	Menu_uid      string  `json:"menu_uid"`
	Menu_category string  `json:"menu_category"`
	Foods         []Foods `json:"foods"`
}
