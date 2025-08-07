// internal/pdf/html.go
package pdf

import (
	"bytes"
	"html/template"
	"os"
)

func RenderHTML(data InvoiceViewData) (string, error) {
	tmpl, err := template.ParseFiles("template/invoice.html")
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}
	// Lưu ra file tạm
	tmpFile := "output/invoice.html"
	if err := os.WriteFile(tmpFile, buf.Bytes(), 0644); err != nil {
		return "", err
	}
	return tmpFile, nil
}
