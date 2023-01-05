package activity

type CreateActivityRequest struct {
	Title string `json:"title"`
	Email string `json:"email"`
}

type UpdateActivityRequest struct {
	Title string `json:"title"`
}
