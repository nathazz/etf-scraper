package utils

import (
	"bytes"
	"github.com/jung-kurt/gofpdf"
	"scraper-go/src/model"
)

func ScrapeUrls(isin string) string {
	return "https://www.trackingdifferences.com/ETF/ISIN/" + isin
}

func SaveToPDF(etfs []model.EtfInfo) ([]byte, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "ETF Report")

	pdf.Ln(12)

	pdf.SetFont("Arial", "", 12)
	for _, etf := range etfs {
		pdf.Cell(0, 10, "Title: "+etf.Title)
		pdf.Ln(8)
		pdf.Cell(0, 10, "Description: "+etf.Description)
		pdf.Ln(8)
		pdf.Cell(0, 10, "Replication: "+etf.Replication)
		pdf.Ln(8)
		pdf.Cell(0, 10, "Earnings: "+etf.Earnings)
		pdf.Ln(8)
		pdf.Cell(0, 10, "TER: "+etf.TotalExpenseRatio)
		pdf.Ln(8)
		pdf.Cell(0, 10, "TD: "+etf.TrackingDifference)
		pdf.Ln(8)
		pdf.Cell(0, 10, "Fund Size: "+etf.FundSize)
		pdf.Ln(12)
	}

	var buf bytes.Buffer
	err := pdf.Output(&buf)

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
