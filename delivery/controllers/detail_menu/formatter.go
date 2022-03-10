package detailmenu

type DetailMenuCreateRequestFormat struct {
	Menu_uid string `json:"menu_uid"`
	Food_uid string `json:"food_uid"`
}

type DetailMenuCreateResponse struct {
	Detail_menu_uid string `json:"detail_menu_uid"`
	Menu_uid        string `json:"menu_uid"`
	Food_uid        string `json:"food_uid"`
}

type DetailMenuGetResponse struct {
	Detail_menu_uid string `json:"detail_menu_uid"`
	Menu_uid        string `json:"menu_uid"`
	Food_uid        string `json:"food_uid"`
}
