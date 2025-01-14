package helpers

const (
	StatusBadRequest          = 400 // error validate
	StatusUnauthorized        = 401 // error verify session token
	StatusForbidden           = 403
	StatusNotFound            = 404
	StatusInternalServerError = 500 // error server
	StatusOK                  = 200 // success
)
