package domain

type Post struct {
	ID       int
	UserID   int
	Content  string
	ImageURL string
	CreatedAt string
	UpdatedAt string

	User 	 User
}