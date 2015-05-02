package gnhandlers

import (
	"net/http"
)

func PostItem(r *http.Request) (int, error) {
	r.Form.Encode()
	return http.StatusOK, nil
}
