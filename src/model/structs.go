package model

type EtfInfo struct {
	Title              string `json:"title"`
	Replication        string `json:"replication"`
	Earnings           string `json:"earnings"`
	TotalExpenseRatio  string `json:"total_expense_ratio"`
	Description        string `json:"description"`
	TrackingDifference string `json:"tracking_difference"`
	FundSize           string `json:"fund_size"`
}
