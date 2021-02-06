package models

// Link contains the domain, slug and target domain
type Link struct {
	Slug string `gorm:"primaryKey" json:"slug"`
	URL  string `json:"url"`
}
