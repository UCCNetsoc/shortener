package models

// RedirectURL structure containing the url and it's slug
type RedirectURL struct {
	Slug string `gorm:"unique"`
	URL  string
}
