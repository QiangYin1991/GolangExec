package gee

import (
	"log"
	"net/http"
	"strings"
)

type HandleFunc func(*Context)

type RouteGroup struct {
	prefix      string
	middlewares []HandleFunc
	engine      *Engine
}

type Engine struct {
	*RouteGroup
	router *router
	groups []*RouteGroup
}

func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouteGroup = &RouteGroup{engine:engine}
	engine.groups = []*RouteGroup{engine.RouteGroup}
	return engine
}

func (g *RouteGroup) Group(prefix string) *RouteGroup {
	engine := g.engine
	newGroup := &RouteGroup{
		prefix:      g.prefix + prefix,
		engine:      engine,
	}

	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (g *RouteGroup) addRoute(method string, comp string, handler HandleFunc) {
	pattern := g.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	g.engine.router.addRoute(method, pattern, handler)
}

func (g *RouteGroup) GET(pattern string, handler HandleFunc) {
	g.addRoute("GET", pattern, handler)
}

func (g *RouteGroup) POST(pattern string, handler HandleFunc) {
	g.addRoute("POST", pattern, handler)
}

func (g *RouteGroup) Use(middlewares ...HandleFunc) {
	g.middlewares = append(g.middlewares, middlewares...)
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var middlewares []HandleFunc
	for _, group := range engine.groups {
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}

	c := newContext(w, req)
	c.handlers = middlewares
	engine.router.handle(c)
}