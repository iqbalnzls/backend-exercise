package domain

type Listings struct {
	Id          int64
	UserId      int64
	Price       int64
	ListingType string
	CreatedAt   int64
	UpdatedAt   int64
	Users       Users `gorm:"foreignKey:user_id;"`
}
