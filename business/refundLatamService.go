package business

import (
	transport "refund-latam-mock/transport"
	"time"
)

const dateCreationFormat = "02-01-2006"

//RequestRefund is a func
func RequestRefund(request transport.RefundRequest) transport.RefundResponse {
	return transport.RefundResponse{Number: "2801974",
		Status:       "Ingresado a la WEB",
		DateCreation: time.Now().Format(dateCreationFormat),
	}
}

//DisplayDocumentRefund is a func
func DisplayDocumentRefund(document string, pnr string) transport.DisplayDocumentResponse {
	documents := []transport.Document{
		transport.Document{
			Number:   document,
			Status:   "Disponible",
			TypeCode: "TKT",
		},
	}
	return transport.DisplayDocumentResponse{Documents: documents}
}
