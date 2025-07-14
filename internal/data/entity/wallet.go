package entity

type Wallet struct {
	Model
	UserID  int   `json:"user_id"`
	Balance int64 `json:"balance"`
}
