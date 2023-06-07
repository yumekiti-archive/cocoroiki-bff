package domain

type Family struct {
	ID       int
	Name     string
	CreateAt string
	UpdateAt string
}

type FamilyUser struct {
	FamilyID int
	UserID   int
	CreateAt string
	UpdateAt string
}