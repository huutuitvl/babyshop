// File: internal/module/invoice/pdf/generator.go
package pdf

import (
	"bytes"
	"html/template"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

type InvoiceViewData struct {
	OrderCode     string
	CustomerName  string
	Address       string
	Phone         string
	TotalAmount   float64
	Items         []InvoiceItem
}

type InvoiceItem struct {
	Name     string
	Price    float64
	Quantity int
	Subtotal float64
}

// RenderHTML renders the invoice HTML and returns the file path
func RenderHTML(data InvoiceViewData) (string, error) {
	tmpl, err := template.ParseFiles("template/invoice.html")
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}

	// Ensure output folder exists
	os.MkdirAll("output", os.ModePerm)

	filename := "invoice_" + time.Now().Format("20060102150405") + ".html"
	htmlPath := filepath.Join("output", filename)
	if err := os.WriteFile(htmlPath, buf.Bytes(), 0644); err != nil {
		return "", err
	}

	return htmlPath, nil
}

// GeneratePDF calls puppeteer to render PDF from html
func GeneratePDF(htmlPath, pdfPath string) error {
	cmd := exec.Command("node", "render-pdf.js", htmlPath, pdfPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// ExportInvoice generates both html and pdf from invoice data
func ExportInvoice(data InvoiceViewData) (string, error) {
	htmlPath, err := RenderHTML(data)
	if err != nil {
		return "", err
	}

	pdfPath := htmlPath[:len(htmlPath)-5] + "pdf"
	if err := GeneratePDF(htmlPath, pdfPath); err != nil {
		return "", err
	}

	return pdfPath, nil
}
