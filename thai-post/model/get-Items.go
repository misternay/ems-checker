package model

type GetItems struct {
	Response ResponseModel `json:"response"`
	Message  string        `json:"message"`
	Status   bool          `json:"status"`
}

type ResponseModel struct {
	Items      map[string][]Item `json:"items"`
	TrackCount TrackModel             `json:"track_count"`
}

type TrackModel struct {
	TrackDate       string `json:"track_date"`
	CountNumber     int    `json:"count_number"`
	TrackCountLimit int    `json:"track_count_limit"`
}

type Item struct {
	Barcode             string `json:"barcode"`
	Status              string `json:"status"`
	StatusDescription   string `json:"status_description"`
	StatusDate          string `json:"status_date"`
	Location            string `json:"location"`
	Postcode            string `json:"postcode"`
	DeliveryStatus      string `json:"delivery_status"`
	DeliveryDescription string `json:"delivery_description"`
	DeliveryDatetime    string `json:"delivery_datetime"`
	ReceiverName        string `json:"receiver_name"`
	Signature           string `json:"signature"`
}

type ServiceRequest struct {
	Barcode  []string `json:"barcode"`
	Language string   `json:"language"`
	Status   string   `json:"status"`
}

type ServiceResponse struct {
	Status           string
	StatusDate       string
	DeliveryStatus   string
	DeliveryDatetime string
	ReceiverName     string
}
