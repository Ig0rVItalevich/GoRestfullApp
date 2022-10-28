package perfume

type LikeReview struct {
	Id                int    `json:"id" db:"id"`
	Mark              int    `json:"mark" db:"mark" binding:"required"`
	DateOfPublication string `json:"date-of-publication" db:"date_of_publication"`
	ReviewId          int    `json:"review-id" db:"review_id"`
	UserId            int    `json:"user-id" db:"user_id"`
}

type UpdateLikeReview struct {
	Mark *int `json:"mark"`
}

func (l *UpdateLikeReview) Validate() bool {
	if l.Mark == nil {
		return false
	}

	return true
}
