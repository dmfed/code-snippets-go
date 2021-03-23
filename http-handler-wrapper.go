package plainauth

import "net/http"

type AuthWrapper struct {
	Token string
}

func New(token string) *AuthWrapper {
	var a AuthWrapper
	a.Token = token
	return &a
}

func (a *AuthWrapper) Valid(key string) bool {
	return key == a.Token
}

func (a *AuthWrapper) WrapHandlerFunc(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("plainauth-token")
		if !a.Valid(token) {
			http.Error(w, "not authorized", 403)
			return
		}
		handlerFunc(w, r)
	}
}

func (a *AuthWrapper) WrapHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("plainauth-token")
		if !a.Valid(token) {
			http.Error(w, "not authorized", 403)
			return
		}
		handler.ServeHTTP(w, r)
	})
}

func (a *AuthWrapper) WrapRequest(r *http.Request) *http.Request {
	r.Header.Set("plainauth-token", a.Token)
	return r
}
