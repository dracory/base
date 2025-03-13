package router

import "net/http"

// Handler represents the function that handles a request.
type Handler func(http.ResponseWriter, *http.Request)

// Middleware represents a middleware function.
type Middleware func(http.Handler) http.Handler

// Route represents a single route definition.
type RouteInterface interface {
	GetMethod() string
	SetMethod(method string) RouteInterface

	GetPath() string
	SetPath(path string) RouteInterface

	GetHandler() Handler
	SetHandler(handler Handler) RouteInterface

	GetName() string
	SetName(name string) RouteInterface

	AddBeforeMiddlewares(middleware []Middleware) RouteInterface
	GetBeforeMiddlewares() []Middleware

	AddAfterMiddlewares(middleware []Middleware) RouteInterface
	GetAfterMiddlewares() []Middleware
}

// Group represents a group of routes.
type GroupInterface interface {
	GetPrefix() string
	SetPrefix(prefix string) GroupInterface

	AddRoute(route RouteInterface) GroupInterface
	AddRoutes(routes []RouteInterface) GroupInterface
	GetRoutes() []RouteInterface

	AddGroup(group GroupInterface) GroupInterface
	AddGroups(groups []GroupInterface) GroupInterface
	GetGroups() []GroupInterface

	AddBeforeMiddlewares(middleware []Middleware) GroupInterface
	GetBeforeMiddlewares() []Middleware

	AddAfterMiddlewares(middleware []Middleware) GroupInterface
	GetAfterMiddlewares() []Middleware
}

// Router represents a router that can handle requests.
type RouterInterface interface {
	GetPrefix() string
	SetPrefix(prefix string) RouterInterface

	AddGroup(group GroupInterface) RouterInterface
	AddGroups(groups []GroupInterface) RouterInterface
	GetGroups() []GroupInterface

	AddRoute(route RouteInterface) RouterInterface
	AddRoutes(routes []RouteInterface) RouterInterface
	GetRoutes() []RouteInterface

	AddBeforeMiddlewares(middleware []Middleware) RouterInterface
	GetBeforeMiddlewares() []Middleware

	AddAfterMiddlewares(middleware []Middleware) RouterInterface
	GetAfterMiddlewares() []Middleware

	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

func NewRouter() RouterInterface {
	return &routerImpl{}
}

func NewRoute() RouteInterface {
	return &routeImpl{}
}

func NewGroup() GroupInterface {
	return &groupImpl{}
}

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

// GroupImpl implements the GroupInterface
// It represents a group of routes.
type groupImpl struct {
	prefix string

	routes []RouteInterface
	groups []GroupInterface

	beforeMiddlewares []Middleware
	afterMiddlewares  []Middleware
}

// Implementing GroupInterface methods for groupImpl
func (g *groupImpl) GetPrefix() string {
	return g.prefix
}

func (g *groupImpl) SetPrefix(prefix string) GroupInterface {
	g.prefix = prefix
	return g
}

func (g *groupImpl) AddRoute(route RouteInterface) GroupInterface {
	g.routes = append(g.routes, route)
	return g
}

func (g *groupImpl) AddRoutes(routes []RouteInterface) GroupInterface {
	g.routes = append(g.routes, routes...)
	return g
}

func (g *groupImpl) GetRoutes() []RouteInterface {
	return g.routes
}

func (g *groupImpl) AddGroup(group GroupInterface) GroupInterface {
	g.groups = append(g.groups, group)
	return g
}

func (g *groupImpl) AddGroups(groups []GroupInterface) GroupInterface {
	g.groups = append(g.groups, groups...)
	return g
}

func (g *groupImpl) GetGroups() []GroupInterface {
	return g.groups
}

func (g *groupImpl) AddBeforeMiddlewares(middleware []Middleware) GroupInterface {
	g.beforeMiddlewares = append(g.beforeMiddlewares, middleware...)
	return g
}

func (g *groupImpl) GetBeforeMiddlewares() []Middleware {
	return g.beforeMiddlewares
}

func (g *groupImpl) AddAfterMiddlewares(middleware []Middleware) GroupInterface {
	g.afterMiddlewares = append(g.afterMiddlewares, middleware...)
	return g
}

func (g *groupImpl) GetAfterMiddlewares() []Middleware {
	return g.afterMiddlewares
}

func (g *groupImpl) GetHandler() Handler {
	return nil // Groups do not have a single handler
}

// RouterImpl implements the RouterInterface
// It represents a router that can handle requests.
type routerImpl struct {
	prefix string

	routes []RouteInterface
	groups []GroupInterface

	beforeMiddlewares []Middleware
	afterMiddlewares  []Middleware
}

// Implementing RouterInterface methods for routerImpl
func (r *routerImpl) GetPrefix() string {
	return r.prefix
}

func (r *routerImpl) SetPrefix(prefix string) RouterInterface {
	r.prefix = prefix
	return r
}

func (r *routerImpl) AddGroup(group GroupInterface) RouterInterface {
	r.groups = append(r.groups, group)
	return r
}

func (r *routerImpl) AddGroups(groups []GroupInterface) RouterInterface {
	r.groups = append(r.groups, groups...)
	return r
}

func (r *routerImpl) GetGroups() []GroupInterface {
	return r.groups
}

func (r *routerImpl) AddRoute(route RouteInterface) RouterInterface {
	r.routes = append(r.routes, route)
	return r
}

func (r *routerImpl) AddRoutes(routes []RouteInterface) RouterInterface {
	r.routes = append(r.routes, routes...)
	return r
}

func (r *routerImpl) GetRoutes() []RouteInterface {
	return r.routes
}

func (r *routerImpl) AddBeforeMiddlewares(middleware []Middleware) RouterInterface {
	r.beforeMiddlewares = append(r.beforeMiddlewares, middleware...)
	return r
}

func (r *routerImpl) GetBeforeMiddlewares() []Middleware {
	return r.beforeMiddlewares
}

func (r *routerImpl) AddAfterMiddlewares(middleware []Middleware) RouterInterface {
	r.afterMiddlewares = append(r.afterMiddlewares, middleware...)
	return r
}

func (r *routerImpl) GetAfterMiddlewares() []Middleware {
	return r.afterMiddlewares
}

func (r *routerImpl) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// Implementation for serving HTTP requests
}
