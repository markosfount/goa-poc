// Code generated by goa v3.19.1, DO NOT EDIT.
//
// HTTP request path constructors for the poc service.
//
// Command:
// $ goa gen goa_poc/goa/design

package server

import (
	"fmt"
)

// GetPocPath returns the URL path to the poc service get HTTP endpoint.
func GetPocPath(id int) string {
	return fmt.Sprintf("/todos/%v", id)
}
