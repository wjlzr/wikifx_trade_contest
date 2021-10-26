package middleware

import (
	"net/http"

	"github.com/k0kubun/pp"
)

type Customintercept struct {
}

func NewCustominterceptMiddleware() *Customintercept {
	return &Customintercept{}
}

func (m *Customintercept) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		lang := r.Header.Get("language")
		if lang == "" {
			r.Header.Set("language", "en")
		}
		pp.Println(r.Header.Get("language"))
		next(w, r)
	}
}
