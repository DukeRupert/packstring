package handlers

import (
	"fmt"
	"net/http"
)

type Contact struct{}

func NewContact() *Contact {
	return &Contact{}
}

func (c *Contact) Submit(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<p>Thank you for your inquiry. We'll be in touch soon.</p>")
}
