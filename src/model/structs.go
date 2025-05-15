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

type EtfRequest struct {
	Isins []string `json:"isins"`
}

type RankedEtf struct {
	IsinsRanked string  `json:"isins"`
	Title       string  `json:"title"`
	Value       float64 `json:"value"`
}

type RankedResult struct {
	TrackingDifference []RankedEtf `json:"tracking_difference"`
	ExpenseRatio       []RankedEtf `json:"total_expense_ratio"`
	FundSize           []RankedEtf `json:"fund_size"`
}
