package router

// RouteImpl implements the RouteInterface
// It represents a single route definition.
type routeImpl struct {
	method  string
	path    string
	handler Handler
	name    string

	beforeMiddlewares []Middleware
	afterMiddlewares  []Middleware
}

func (r *routeImpl) GetMethod() string {
	return r.method
}

func (r *routeImpl) SetMethod(method string) RouteInterface {
	r.method = method
	return r
}

func (r *routeImpl) GetPath() string {
	return r.path
}

func (r *routeImpl) SetPath(path string) RouteInterface {
	r.path = path
	return r
}

func (r *routeImpl) GetHandler() Handler {
	return r.handler
}

func (r *routeImpl) SetHandler(handler Handler) RouteInterface {
	r.handler = handler
	return r
}

func (r *routeImpl) GetName() string {
	return r.name
}

func (r *routeImpl) SetName(name string) RouteInterface {
	r.name = name
	return r
}

func (r *routeImpl) AddBeforeMiddlewares(middleware []Middleware) RouteInterface {
	r.beforeMiddlewares = append(r.beforeMiddlewares, middleware...)
	return r
}

func (r *routeImpl) GetBeforeMiddlewares() []Middleware {
	return r.beforeMiddlewares
}

func (r *routeImpl) AddAfterMiddlewares(middleware []Middleware) RouteInterface {
	r.afterMiddlewares = append(r.afterMiddlewares, middleware...)
	return r
}

func (r *routeImpl) GetAfterMiddlewares() []Middleware {
	return r.afterMiddlewares
}
