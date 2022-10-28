package perfume

type Review struct {
	Id                int    `json:"id" db:"id"`
	Content           string `json:"content" db:"content" binding:"required"`
	Rating            int    `json:"rating" db:"rating"`
	DateOfPublication string `json:"date-of-publication" db:"date_of_publication"`
	UserId            int    `json:"user-id" db:"user_id"`
	ProductId         int    `json:"product-id" db:"product_id"`
}

type UpdateReview struct {
	Content *string `json:"content"`
	Rating  *int    `json:"rating"`
}

func (r *UpdateReview) Validate() bool {
	if r.Content == nil && r.Rating == nil {
		return false
	}

	return true
}
