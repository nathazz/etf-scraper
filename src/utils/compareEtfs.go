package utils

import (
	"fmt"
	"scraper-go/src/model"
	"sort"
	"strconv"
	"strings"
)

func CompareEtf(etfs []model.EtfInfo) string {
	type rankedEtf struct {
		Title string
		Value float64
	}

	var tdRank, terRank, fsRank []rankedEtf

	for _, etf := range etfs {
		tdRank = append(tdRank, rankedEtf{etf.Title, parsePercent(etf.TrackingDifference)})
		terRank = append(terRank, rankedEtf{etf.Title, parsePercent(etf.TotalExpenseRatio)})
		fsRank = append(fsRank, rankedEtf{etf.Title, parseFundSize(etf.FundSize)})
	}

	sort.Slice(tdRank, func(i, j int) bool { return tdRank[i].Value < tdRank[j].Value })
	sort.Slice(terRank, func(i, j int) bool { return terRank[i].Value < terRank[j].Value })
	sort.Slice(fsRank, func(i, j int) bool { return fsRank[i].Value > fsRank[j].Value })

	var sb strings.Builder
	sb.WriteString("Top ETFs by Tracking Difference:\n")
	for i, etf := range tdRank[:min(3, len(tdRank))] {
		sb.WriteString(fmt.Sprintf("%d. %s (%.2f%%)\n", i+1, etf.Title, etf.Value))
	}

	sb.WriteString("\nTop ETFs by Total Expense Ratio:\n")
	for i, etf := range terRank[:min(3, len(terRank))] {
		sb.WriteString(fmt.Sprintf("%d. %s (%.2f%%)\n", i+1, etf.Title, etf.Value))
	}

	sb.WriteString("\nTop ETFs by Fund Size:\n")
	for i, etf := range fsRank[:min(3, len(fsRank))] {
		sb.WriteString(fmt.Sprintf("%d. %s (%.2f B)\n", i+1, etf.Title, etf.Value))
	}

	return sb.String()
}

func parsePercent(s string) float64 {
	s = strings.TrimSuffix(s, "%")
	s = strings.ReplaceAll(s, ",", ".")
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}
	return v
}

func parseFundSize(s string) float64 {
	s = strings.ToUpper(s)
	s = strings.ReplaceAll(s, ",", "")
	if strings.HasSuffix(s, "B") {
		val, _ := strconv.ParseFloat(strings.TrimSuffix(s, "B"), 64)
		return val
	}
	if strings.HasSuffix(s, "M") {
		val, _ := strconv.ParseFloat(strings.TrimSuffix(s, "M"), 64)
		return val / 1000
	}
	val, _ := strconv.ParseFloat(s, 64)
	return val
}
