package entity

//LoanType ...
type LoanType struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Percent   float32 `json:"percent"`
	MaxAmount float64 `json:"max_amount"`
	MinAmount float64 `json:"min_amount"`
	CalcType  string  `json:"calc_type"`
}
