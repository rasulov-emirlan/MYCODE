package storage

type Order struct {
	ID       int      `json:"id"`
	Customer string   `json:"customer"`
	Address  string   `json:"address"`
	Cost     int      `json:"cost"`
	Order    []string `json:"order"`
	Payment  int      `json:"payment"`
}


