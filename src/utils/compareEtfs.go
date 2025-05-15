package utils

import (
	"fmt"
	"log"
	"scraper-go/src/model"
	"sort"
	"strconv"
	"strings"
)

func CompareEtf(etfs []model.EtfInfo, isJson bool) interface{} {

	var isins, tdRank, terRank, fsRank []model.RankedEtf

	for _, etf := range etfs {
		isins = append(isins, model.RankedEtf{
			IsinsRanked: etf.Isin,
		})

		tdRank = append(tdRank, model.RankedEtf{
			IsinsRanked: etf.Isin,
			Title:       etf.Title,
			Value:       parsePercent(etf.TrackingDifference),
		})
		terRank = append(terRank, model.RankedEtf{
			IsinsRanked: etf.Isin,
			Title:       etf.Title,
			Value:       parsePercent(etf.TotalExpenseRatio),
		})
		fsRank = append(fsRank, model.RankedEtf{
			IsinsRanked: etf.Isin,
			Title:       etf.Title,
			Value:       parseFundSize(etf.FundSize),
		})
	}

	sort.Slice(tdRank, func(i, j int) bool { return tdRank[i].Value < tdRank[j].Value })
	sort.Slice(terRank, func(i, j int) bool { return terRank[i].Value < terRank[j].Value })
	sort.Slice(fsRank, func(i, j int) bool { return fsRank[i].Value > fsRank[j].Value })

	if isJson {
		result := model.RankedResult{
			TrackingDifference: tdRank[:minimum(10, len(tdRank))],
			ExpenseRatio:       terRank[:minimum(10, len(terRank))],
			FundSize:           fsRank[:minimum(10, len(fsRank))],
		}
		return result
	}

	var sb strings.Builder
	sb.WriteString("Top ETFs by Tracking Difference:\n")
	for i, etf := range tdRank[:minimum(10, len(tdRank))] {
		sb.WriteString(fmt.Sprintf("%d. %s (%.2f%%)\n", i+1, etf.Title, etf.Value))
	}

	sb.WriteString("\nTop ETFs by Total Expense Ratio:\n")
	for i, etf := range terRank[:minimum(10, len(terRank))] {
		sb.WriteString(fmt.Sprintf("%d. %s (%.2f%%)\n", i+1, etf.Title, etf.Value))
	}

	sb.WriteString("\nTop ETFs by Fund Size:\n")
	for i, etf := range fsRank[:minimum(10, len(fsRank))] {
		sb.WriteString(fmt.Sprintf("%d. %s (%.2f B)\n", i+1, etf.Title, etf.Value))
	}

	return sb.String()
}

func parsePercent(s string) float64 {
	s = strings.TrimSpace(s)
	s = strings.TrimSuffix(s, "%")
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, ",", ".")
	v, err := strconv.ParseFloat(s, 64)

	if err != nil {
		log.Printf("Warning: failed to parse percent value '%s'\n", s)
		return 0
	}
	return v
}

func parseFundSize(s string) float64 {
	s = strings.ToUpper(strings.TrimSpace(s))
	s = strings.ReplaceAll(s, ",", "")
	s = strings.TrimSuffix(s, " EUR")
	s = strings.TrimSuffix(s, "MIO")
	s = strings.TrimSpace(s)

	val, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Printf("Warning: failed to parse fund size value '%s'\n", s)
		return 0
	}
	return val / 1000
}

func minimum(a, b int) int {
	if a < b {
		return a
	}
	return b
}
