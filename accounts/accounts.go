package accounts

type User struct {
	Username    string `json:"username,omitempty"`
	Password    string `json:"password,omitempty"`
	PhotoURL    string `json:"photo_url,omitempty"`
	TotalRating int    `json:"total_rating,omitempty"`
	NumRatings  int    `json:"num_ratings,omitempty"`
	UserDID     string
}
