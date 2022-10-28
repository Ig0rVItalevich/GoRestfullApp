package perfume

type Product struct {
	Id                int    `json:"id" db:"id"`
	Title             string `json:"title" binding:"required" db:"title"`
	Content           string `json:"content" binding:"required" db:"content"`
	Count             int    `json:"count" db:"count"`
	Cost              int    `json:"cost" binding:"required" db:"cost"`
	DateOfPublication string `json:"date-of-publication" db:"date_of_publication"`
	Rating            int    `json:"rating" db:"rating"`
}

type Category struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name" binding:"required"`
}

type ProductsCategories struct {
	Id         int `json:"id" db:"id"`
	ProductId  int `json:"product-id" db:"product_id"`
	CategoryId int `json:"category-id" db:"category_id"`
}

type UpdateProduct struct {
	Title   *string `json:"title"`
	Content *string `json:"content"`
	Count   *int    `json:"count"`
	Cost    *int    `json:"cost"`
	Rating  *int    `json:"rating"`
}

func (p *UpdateProduct) Validate() bool {
	if p.Title == nil && p.Content == nil && p.Count == nil && p.Cost == nil && p.Rating == nil {
		return false
	}

	return true
}
