package repo

import (
	"MrHour/entity"
)

// FindPortfolios to list the messages
func FindPortfolios() ([]entity.Portfolio, error) {
	portfolios := []entity.Portfolio{}
	portfolio1 := entity.Portfolio{ID: 12, Company: "Emirates Transport", Skills: []string{
		"VueJS",
		"Element UI",
		"Google Map",
		"Javascript",
		"SCSS",
		"CSS",
		"HTML"}, WorkOn: "Fleet Management App"}
	portfolio2 := entity.Portfolio{ID: 13, Company: "DOT Transport", Skills: []string{
		"VueJS",
		"Element UI",
		"Google Map",
		"Javascript",
		"SCSS",
		"CSS",
		"HTML"}, WorkOn: "Fleet Management App"}
	portfolios = append(portfolios, portfolio1)
	portfolios = append(portfolios, portfolio2)
	return portfolios, nil
}
