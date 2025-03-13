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
	// Create a handler chain by wrapping the final handler with middlewares
	var matchedHandler http.Handler

	// Check if any route matches the request
	if _, handler := r.findMatchingRoute(req); handler != nil {
		matchedHandler = handler
	}

	// If no route matched, return 404
	if matchedHandler == nil {
		http.NotFound(w, req)
		return
	}

	// Execute the handler chain
	matchedHandler.ServeHTTP(w, req)
}

// findMatchingRoute attempts to find a route that matches the request
// It returns the matched route and an http.Handler that includes all middlewares
func (r *routerImpl) findMatchingRoute(req *http.Request) (RouteInterface, http.Handler) {
	// Check direct routes on the router
	for _, route := range r.routes {
		if r.routeMatches(route, req) {
			return route, r.wrapWithMiddlewares(route, req)
		}
	}

	// Check routes in groups
	for _, group := range r.groups {
		if route, handler := r.findMatchingRouteInGroup(group, req, ""); route != nil {
			return route, handler
		}
	}

	return nil, nil
}

// findMatchingRouteInGroup recursively searches for a matching route in a group and its subgroups
func (r *routerImpl) findMatchingRouteInGroup(group GroupInterface, req *http.Request, parentPath string) (RouteInterface, http.Handler) {
	// Combine parent path with group prefix
	groupPath := parentPath + group.GetPrefix()

	// Check routes in the current group
	for _, route := range group.GetRoutes() {
		// Create a copy of the route with adjusted path
		adjustedRoute := &routeImpl{
			method:            route.GetMethod(),
			path:              groupPath + route.GetPath(),
			handler:           route.GetHandler(),
			name:              route.GetName(),
			beforeMiddlewares: route.GetBeforeMiddlewares(),
			afterMiddlewares:  route.GetAfterMiddlewares(),
		}

		if r.routeMatches(adjustedRoute, req) {
			// Create a handler chain with group middlewares and route middlewares
			return route, r.wrapWithGroupMiddlewares(route, group, req, parentPath)
		}
	}

	// Check subgroups
	for _, subgroup := range group.GetGroups() {
		if route, handler := r.findMatchingRouteInGroup(subgroup, req, groupPath); route != nil {
			return route, handler
		}
	}

	return nil, nil
}

// routeMatches checks if a route matches the request method and path
func (r *routerImpl) routeMatches(route RouteInterface, req *http.Request) bool {
	// Check if method matches
	if route.GetMethod() != req.Method && route.GetMethod() != "" {
		return false
	}

	routePath := r.prefix + route.GetPath()
	requestPath := req.URL.Path

	// Handle catch-all routes
	if routePath == "/*" || routePath == "/**" {
		return true
	}

	// Handle wildcard patterns at the end of the path
	if len(routePath) > 2 && routePath[len(routePath)-2:] == "/*" {
		// Check if the base path matches
		basePath := routePath[:len(routePath)-2]
		return len(requestPath) >= len(basePath) && requestPath[:len(basePath)] == basePath
	}

	// For regular paths, do exact matching
	return routePath == requestPath
}

// wrapWithMiddlewares wraps a route's handler with its middlewares and the router's middlewares
func (r *routerImpl) wrapWithMiddlewares(route RouteInterface, req *http.Request) http.Handler {
	// Start with the route's handler
	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		route.GetHandler()(w, r)
	})

	// Apply route's after middlewares (in reverse order)
	for i := len(route.GetAfterMiddlewares()) - 1; i >= 0; i-- {
		handler = route.GetAfterMiddlewares()[i](handler)
	}

	// Apply router's after middlewares (in reverse order)
	for i := len(r.afterMiddlewares) - 1; i >= 0; i-- {
		handler = r.afterMiddlewares[i](handler)
	}

	// Apply route's before middlewares
	for _, middleware := range route.GetBeforeMiddlewares() {
		handler = middleware(handler)
	}

	// Apply router's before middlewares
	for _, middleware := range r.beforeMiddlewares {
		handler = middleware(handler)
	}

	return handler
}

// wrapWithGroupMiddlewares wraps a route's handler with its middlewares, the group's middlewares, and the router's middlewares
func (r *routerImpl) wrapWithGroupMiddlewares(route RouteInterface, group GroupInterface, req *http.Request, parentPath string) http.Handler {
	// Start with the route's handler
	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		route.GetHandler()(w, r)
	})

	// Apply route's after middlewares (in reverse order)
	for i := len(route.GetAfterMiddlewares()) - 1; i >= 0; i-- {
		handler = route.GetAfterMiddlewares()[i](handler)
	}

	// Apply group's after middlewares (in reverse order)
	for i := len(group.GetAfterMiddlewares()) - 1; i >= 0; i-- {
		handler = group.GetAfterMiddlewares()[i](handler)
	}

	// Apply router's after middlewares (in reverse order)
	for i := len(r.afterMiddlewares) - 1; i >= 0; i-- {
		handler = r.afterMiddlewares[i](handler)
	}

	// Apply route's before middlewares
	for _, middleware := range route.GetBeforeMiddlewares() {
		handler = middleware(handler)
	}

	// Apply group's before middlewares
	for _, middleware := range group.GetBeforeMiddlewares() {
		handler = middleware(handler)
	}

	// Apply router's before middlewares
	for _, middleware := range r.beforeMiddlewares {
		handler = middleware(handler)
	}

	return handler
}
