package perfume

type Order struct {
	Id                int    `json:"id" db:"id"`
	Status            string `json:"status" db:"status" binding:"required"`
	DateOfPublication string `json:"date-of-publication" db:"date_of_publication"`
	Comment           string `json:"comment" db:"comment"`
	UserId            int    `json:"user-id" db:"user_id"`
	DateOfCompletion  string `json:"date-of-completion" db:"date_of_completion"`
}

type productsIdSlice struct {
	Products []int
}

type UpdateOrder struct {
	Status              *string         `json:"status"`
	ProductsToBeAdded   productsIdSlice `json:"products-to-be-added"`
	ProductsToBeDeleted productsIdSlice `json:"products-to-be-deleted"`
}

func (o *UpdateOrder) Validate() bool {
	if o.Status == nil && len(o.ProductsToBeAdded.Products) == 0 && len(o.ProductsToBeDeleted.Products) == 0 {
		return false
	}

	return true
}

type OrdersProducts struct {
	Id        int `json:"id" db:"id"`
	Count     int `json:"count" db:"count"`
	OrderId   int `json:"order-id" db:"order_id"`
	ProductId int `json:"product-id" db:"product_id"`
}
