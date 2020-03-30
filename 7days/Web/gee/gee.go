package gee

import (
	"html/template"
	"log"
	"net/http"
	"path"
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
	router        *router
	groups        []*RouteGroup
	htmlTemplates *template.Template
	funcMap       template.FuncMap
}

func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouteGroup = &RouteGroup{engine: engine}
	engine.groups = []*RouteGroup{engine.RouteGroup}
	return engine
}

func (g *RouteGroup) Group(prefix string) *RouteGroup {
	engine := g.engine
	newGroup := &RouteGroup{
		prefix: g.prefix + prefix,
		engine: engine,
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

func (g *RouteGroup) createStaticHandler(relativePath string, fs http.FileSystem) HandleFunc {
	absolutePath := path.Join(g.prefix, relativePath)
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))
	return func(c *Context) {
		file := c.Param("filepath")
		if _, err := fs.Open(file); err != nil {
			c.Status(http.StatusNotFound)
			return
		}

		fileServer.ServeHTTP(c.Writer, c.Req)
	}
}

func (g *RouteGroup) Static(relativePath string, root string) {
	handler := g.createStaticHandler(relativePath, http.Dir(root))
	urlPattern := path.Join(relativePath, "/*filepath")
	g.GET(urlPattern, handler)
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
	c.engine = engine
	engine.router.handle(c)
}

func (engine *Engine) SetFunMap(funcMap template.FuncMap) {
	engine.funcMap = funcMap
}

func (engine *Engine) LoadHTMLGlob(pattern string) {
	engine.htmlTemplates = template.Must(template.New("").Funcs(engine.funcMap).ParseGlob(pattern))
}