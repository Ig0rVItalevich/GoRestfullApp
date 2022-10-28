package perfume

type LikeProduct struct {
	Id                int    `json:"id" db:"id"`
	Mark              int    `json:"mark" db:"mark"`
	DateOfPublication string `json:"date-of-publication" db:"date_of_publication"`
	ProductId         int    `json:"product-id" db:"product_id"`
	UserId            int    `json:"user-id" db:"user_id"`
}

type UpdateLikeProduct struct {
	Mark *int `json:"mark" db:"mark"`
}

func (l *UpdateLikeProduct) Validate() bool {
	if l.Mark == nil {
		return false
	}

	return true
}
