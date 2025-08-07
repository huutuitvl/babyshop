// internal/pdf/invoice_data.go
package pdf

type InvoiceViewData struct {
	OrderCode    string
	CustomerName string
	Address      string
	Phone        string
	TotalAmount  float64
	Items        []InvoiceItem
}

type InvoiceItem struct {
	Name     string
	Price    float64
	Quantity int
	Subtotal float64
}
