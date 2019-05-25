package transport

//RefundRequest is a model
type RefundRequest struct {
	Agency    Agency     `json:"agency"`
	Documents []Document `json:"documents"`
	Pnr       Pnr        `json:"pnr"`
}

//Document is a model
type Document struct {
	Number   string `json:"number"`
	Reason   Reason `json:"reason"`
	Status   string `json:"status"`
	TypeCode string `json:"typeCode"`
}

//Agency is a model
type Agency struct {
	Agent    Agent  `json:"agent"`
	IataCode string `json:"iataCode"`
	Name     string `json:"name"`
	TypeCode string `json:"typeCode"`
}

//Agent is a model
type Agent struct {
	Email    string `json:"email"`
	LastName string `json:"lastName"`
	Name     string `json:"name"`
}

//Reason is a model
type Reason struct {
	Comment  string `json:"comment"`
	TypeCode string `json:"typeCode"`
}

//Pnr is a model
type Pnr struct {
	RecordLocator string `json:"recordLocator"`
}

//DisplayDocumentResponse is a model
type DisplayDocumentResponse struct {
	Documents []Document `json:"documents"`
}

//RefundResponse is a model
type RefundResponse struct {
	DateCreation string `json:"dateCreation"`
	Number       string `json:"number"`
	Status       string `json:"status"`
}
