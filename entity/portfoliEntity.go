package entity

// Hour struct to design SQL and respose
// Portfolio type
type Portfolio struct {
	ID      int      `json: ID`
	Company string   `json: Company`
	WorkOn  string   `json: WorkOn`
	Skills  []string `json: Skills`
	AppUrl  string   `json: AppUrl`
}
