package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

func ChainMiddleware(h http.Handler, m ...Middleware) http.Handler {
	//кто прочитал, тот лох
	//здесь соединяются все мидлвари, начиная с последнего потому что так нужно
	middlewareAmount := len(m)
	if middlewareAmount == 0 {
		return h
	}

	for i := middlewareAmount - 1; i >= 0; i-- {
		h = m[i](h)
	}

	return h
}
