package model

type EtfInfo struct {
	Isin               string `json:"isin"`
	Title              string `json:"title"`
	Replication        string `json:"replication"`
	Earnings           string `json:"earnings"`
	TotalExpenseRatio  string `json:"total_expense_ratio"`
	TrackingDifference string `json:"tracking_difference"`
	FundSize           string `json:"fund_size"`
	Description        string `json:"description"`
}
