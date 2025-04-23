package request

import "net/http"

func HandleBody[T any](w *http.ResponseWriter, r *http.Request) (*T, error) {
	body, err := decode[T](r.Body)
	if err != nil {
		//TODO response err
		return nil, err
	}
	err = isValid[T](body)
	if err != nil {
		//TODO response err
		return nil, err
	}
	return &body, err
}
