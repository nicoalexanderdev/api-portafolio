package entity

type Project struct {
	Title        string   `json:"title"        binding:"required"`
	Subtitle     string   `json:"subtitle"     binding:"required"`
	Description  string   `json:"description"  binding:"required"`
	Technologies []string `json:"technologies" binding:"required"`
	URL          string   `json:"url"          binding:"required"`
	MonthYear    string   `json:"monthyear"    binding:"required"`
}
