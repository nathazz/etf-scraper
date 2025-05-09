package utils

import (
	"bytes"
	"github.com/jung-kurt/gofpdf"
	"scraper-go/src/model"
)

func SaveToPDF(etfs []model.EtfInfo) ([]byte, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetMargins(15, 20, 15)
	pdf.AddPage()

	pdf.SetFont("Arial", "B", 18)
	pdf.CellFormat(0, 12, "ETF Report", "", 1, "C", false, 0, "")
	pdf.Ln(5)

	pdf.Ln(10)

	for _, etf := range etfs {
		pdf.SetDrawColor(0, 0, 0)
		pdf.Line(15, pdf.GetY(), 195, pdf.GetY())
		pdf.Ln(6)

		pdf.SetFont("Arial", "B", 14)
		pdf.Cell(0, 10, etf.Title)
		pdf.Ln(8)

		pdf.SetFont("Arial", "I", 11)
		pdf.Cell(0, 8, "ISIN: "+etf.Isin)
		pdf.Ln(6)

		pdf.SetFont("Arial", "B", 12)
		writeKeyValue(pdf, "Replication", etf.Replication)
		writeKeyValue(pdf, "Earnings", etf.Earnings)
		writeKeyValue(pdf, "TER", etf.TotalExpenseRatio)
		writeKeyValue(pdf, "TD", etf.TrackingDifference)
		writeKeyValue(pdf, "Fund Size", etf.FundSize)
		writeKeyValue(pdf, "Description", etf.Description)

		pdf.Ln(8)
	}

	rankingText := CompareEtf(etfs)

	pdf.SetFont("Arial", "", 12)
	pdf.MultiCell(0, 7, rankingText, "", "", false)
	pdf.Ln(5)

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func writeKeyValue(pdf *gofpdf.Fpdf, key, value string) {
	pdf.SetFont("Arial", "B", 11)
	pdf.Cell(35, 8, key+":")
	pdf.SetFont("Arial", "", 11)
	pdf.MultiCell(0, 8, value, "", "", false)
}
