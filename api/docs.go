// Package api Gofi
//
// API documentation for the Gofi.
//
//		Schemes: http
//		Host: 127.0.0.1:7677
//		Version: 1.0.0
//		Contact: rashad.ansari1996@gmail.com
//
//		Consumes:
//			- application/json
//
//		Produces:
//			- application/json
//
// swagger:meta
package api

// Error response.
// swagger:response Error
type Error struct {
	// in: body
	Body struct {
		// example: Something Went Wrong!
		// Error message.
		Message string `json:"message"`
	}
}

// No Content response.
// swagger:response NoContent
type NoContent struct{}
