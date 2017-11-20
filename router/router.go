package router
import (
	"fmt"
	"net/http"
	"path"
)
//错误类型的定义
const (
	errNotFound = "not found"
	errUnsupported = "http unsupported"
	RouterKey = "%s-%s"
)
//路由方法的定义

type (
	//路由对象
	Router struct {
		//全局hander
		globalHander []HanderFunc
		//基础路径
		basePath string
		//路由分发
		routers map[string]*route
	}
	//方法分发
	route struct {
		//http
		method string
		//handler
		handlers []HanderFunc
	}
	//上下文对象
	Context struct {
		Requset *http.Request
		Writer http.ResponseWriter
		handlers []HanderFunc
		index int8
	}
	HanderFunc func(*Context)
)
//实例
 func New() *Router {
	 return &Router {
		routers:  make(map[string]*route),
		 basePath: "/",
	 }
 }
 //全局处理
 func (r *Router) Use(handlers ...HanderFunc) {
	 r.globalHander = append(r.globalHander, handlers...)
 }
 //分组路由
 func (r *Router) Group(partPath string,fn func(),handlers ...HanderFunc) {
	 rootBasePath := r.basePath
	 rootHandlers := r.globalHander
	 r.basePath = path.Join(r.basePath, partPath)
	 r.globalHander = r.combineHandlers(handlers)
	 fn()
	 r.basePath = rootBasePath
	 r.globalHander = rootHandlers
 }
//  type opt struct {
// 	 mode string
// 	 partPath string
// 	 handlers ...HanderFunc
// 	 queryFunc func
//  }

 //get
 func (r *Router) GET(partPath string, handlers ...HanderFunc) {
	 	path := path.Join(r.basePath, partPath)
		handlers = r.combineHandlers(handlers)
		r.addRoute(http.MethodGet, path, handlers)
 }
 //post
 func (r *Router) POST(partPath string, handlers ...HanderFunc) {
	 	path := path.Join(r.basePath, partPath)
		handlers = r.combineHandlers(handlers)
	 	r.addRoute(http.MethodPost, path, handlers)
 }
 //run
 func (r *Router) Run(port string) error {
	 return http.ListenAndServe(port, r)
 }
 //请求合并
 func (r *Router) combineHandlers(handlers []HanderFunc)[]HanderFunc {
	 combineLength := len(r.globalHander) + len(handlers)
	 combineHandler := make([]HanderFunc, combineLength)
	 copy(combineHandler, r.globalHander)
	 copy(combineHandler[len(r.globalHander):], handlers)
	 return combineHandler
 }
 //add
 func (r *Router) addRoute(method, path string,handlers []HanderFunc) {
	 route := &route{
		 method: method,
		 handlers: handlers,
	 }
	 r.routers[fmt.Sprintf(RouterKey, path, method)] = route
 }
 //server
 func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	 httpMethod :=  req.Method
	 path := req.URL.Path
	 route, ok := r.routers[fmt.Sprintf(RouterKey, path, httpMethod)]
	 if !ok {
		 w.WriteHeader(http.StatusNotFound)
		 fmt.Fprintf(w,errNotFound)
		 return
	 }
	 c := &Context{
		 Requset: req,
		 Writer: w,
		 handlers: route.handlers,
		 index: -1,
	 }
	 c.Next()
 }

func (c *Context) Next() {
	c.index++
	if n := int8(len(c.handlers)); c.index < n {
		c.handlers[c.index](c)
	}
}
