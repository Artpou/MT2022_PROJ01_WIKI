type Article struct {
	ID        uint `gorm:"primaryKey"`
	CreatorID uint
	Titre     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}