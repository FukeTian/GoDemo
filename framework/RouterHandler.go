package framework

import (
	"fmt"
	"net/http"
)

// Router newRouterHandler
var Router *RouterHandler = new(RouterHandler)

// RouterHandler struct
type RouterHandler struct {
}

var mux = make(map[string]func(http.ResponseWriter, *http.Request))

func (p *RouterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	if fun, ok := mux[r.URL.Path]; ok {
		fun(w, r)
		return
	}

	http.Error(w, "error URL:"+r.URL.String(), http.StatusBadRequest)
}

// Router 路由方法
func (p *RouterHandler) Router(relativePath string, handler func(http.ResponseWriter, *http.Request)) {
	mux[relativePath] = handler
}
