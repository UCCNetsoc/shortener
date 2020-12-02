package models

// Request contains the domain, slug and target domain
type Request struct {
	Domain string `json:"domain"`
	Slug   string `json:"slug"`
	Target string `json:"target"`
}
