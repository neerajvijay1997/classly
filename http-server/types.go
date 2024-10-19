package httpserver

type SignUpRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type SignUpResponse struct {
	Message  string `json:"message"`
	UserName string `json:"user_name"`
}

type GetVersionResponse struct {
	Version string `json:"version"`
}

type GetUserDetailsResponse struct {
	UserName string `json:"user_name"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type CreateClassRequest struct {
	UserName  string `json:"user_name"`
	ClassName string `json:"class_name"`
	StartDate string `json:"start_date"` // Expected format: "2006-01-02"
	EndDate   string `json:"end_date"`   // Expected format: "2006-01-02"
	Capacity  uint32 `json:"capacity"`
}

type CreateClassResponse struct {
	Message string `json:"message"`
	ClassID string `json:"class_id"`
}

type BookClassRequest struct {
	UserName    string `json:"user_name"`
	ClassID     string `json:"class_id"`
	BookingDate string `json:"booking_date"` // Expected format: "2006-01-02"
}

type BookClassResponse struct {
	Message        string `json:"message"`
	ClassSessionId string `json:"class_session_id"`
}
