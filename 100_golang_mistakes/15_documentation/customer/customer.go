// Package customer provides basic customer model with methods.
package customer

// Customer represents customer model
type Customer struct {
	id string
}

// New returns new customer
func New(id string) *Customer {
	return &Customer{id: id}
}

// NewCustomer returns new customer
// Deprecated: Wrong naming. Use New instead
func NewCustomer(id string) *Customer {
	return &Customer{id: id}
}

// ID returns Customer's id
func (c Customer) ID() string {
	return c.id
}
