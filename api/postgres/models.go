package postgres

import "github.com/lib/pq"

type Product struct {
	ID   		int
	Name 		string
	Price 		int
	Description pq.StringArray
}
