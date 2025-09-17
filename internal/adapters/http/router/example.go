package router
import "net/http"
func Routes() http.Handler { return http.NewServeMux() }
