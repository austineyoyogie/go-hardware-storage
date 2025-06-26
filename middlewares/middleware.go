package middlewares

import (
	"net/http"

	"github.com/austineyoyogie/go-hardware-store/utils"
)
// it write out every url requested on console.log
// {"method":"GET","url":"localhost:5000/search/products?q=GTX","version":"HTTP/1.1"}
// {"method":"GET","url":"localhost:5000/products/1","version":"HTTP/1.1"}
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.Debuger(struct {
			Method string `json:"method"`
			Url string `json:"url"`
			Version string `json:"version"`
		}{
			Method: r.Method,
			Url: 	r.Host + r.RequestURI,
			Version: r.Proto,	
		})
		next(w, r)
	}
}