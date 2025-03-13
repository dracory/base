package router

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
