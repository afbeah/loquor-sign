package models

type Symbol struct{
	ID string `json: "id"`
	Name string `json: "name"`
	Image string `json: "image"`
	Category_ID string `json: "category_id"`
}