package jsroute

import "net/http"

const NOT_FOUND_HANDLER_RESP_CODE = 900

type defaultNotFoundHandler struct {
}

func (n *defaultNotFoundHandler) ServeHTTP(resp http.ResponseWriter, r *http.Request) {
	resp.WriteHeader(http.StatusNotFound)
}
