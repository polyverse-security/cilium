package jsroute

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/robertkrimen/otto"
	"github.com/robertkrimen/otto/parser"
	"github.com/vulcand/vulcand/router"
	"net/http"
	"strings"
	"sync"
)

func NewMux() router.Router {
	log.Info("New Javascript router created.")
	return &jsRouter{
		notFound: &defaultNotFoundHandler{},
		routes:   map[string]http.Handler{},
		vm:       *otto.New(),
	}
}

type jsRouter struct {
	routes   map[string]http.Handler
	notFound http.Handler
	vm       otto.Otto
	lock     sync.Mutex
}

func (r *jsRouter) SetNotFound(nf http.Handler) error {
	r.notFound = nf
	return nil
}

func (r *jsRouter) GetNotFound() http.Handler {
	return r.notFound
}

func (r *jsRouter) IsValid(routedef string) bool {
	routedef = r.sanitizeRoute(routedef)

	//validate routedef is valid
	if _, err := parser.ParseFile(nil, "", routedef, 0); err != nil {
		log.WithFields(log.Fields{"Error": err, "RouteDefinition": routedef}).Error("Error parsing route definition in javascript.")
		return false
	}
	log.WithFields(log.Fields{"RouteDef": routedef}).Debug("Route valid for javascript router.")
	return true
}

func (r *jsRouter) Handle(routedef string, handler http.Handler) error {
	if handler == nil {
		log.WithField("RouteDefinition", routedef).Error("Attempt to set a nil handler against route.")
		return fmt.Errorf("Handler cannot be nil.")
	}

	if !r.IsValid(routedef) {
		return fmt.Errorf("Route Definition %v not found valid for Javascript expression router.", routedef)
	}

	if _, ok := r.routes[routedef]; ok {
		log.WithFields(log.Fields{"RouteDefinition": routedef}).Warning("Route definition already exists in this router, and has a handler assigned to it. Overriding the handler but this is highly unusual.")
	}
	r.routes[routedef] = handler
	return nil
}

func (r *jsRouter) Remove(routedef string) error {
	if _, ok := r.routes[routedef]; !ok {
		log.WithField("RouteDefinition", routedef).Error("No route definition/handler exists. Unable to remove.")
		return fmt.Errorf("A handler did not exist for route %v. Didn't remove any handlers", routedef)
	}
	delete(r.routes, routedef)
	return nil
}

func (r *jsRouter) ServeHTTP(resp http.ResponseWriter, req *http.Request) {

	log.WithField("Request", req).Debug("Javascript router attempting to match request against known routes.")

	//iterate through all routes linearly - the gift of a powerful evaluator and the curse
	for route, handler := range r.routes {
		if r.matches(route, req) {
			handler.ServeHTTP(resp, req)
			return //Exit when match!
		}
	}

	log.WithField("Request", req).Warning("This request did not match any registered javascript routes. Sending to default not found handler.")
	//finally fall back to not found handler of this combined router.
	r.notFound.ServeHTTP(resp, req)
}

func (r *jsRouter) matches(routedef string, req *http.Request) bool {
	r.lock.Lock()
	defer r.lock.Unlock()

	routedef = r.sanitizeRoute(routedef)
	r.vm.Set("Request", req)
	if value, err := r.vm.Run(routedef); err != nil {
		log.WithFields(log.Fields{"Error": err, "Request": req, "RouteDef": routedef}).Error("Error when evaluating request against route definition")
		return false
	} else if match, err := value.ToBoolean(); err != nil {
		log.WithFields(log.Fields{"Error": err, "Value": value}).Error("Expression result not a boolean; not convertable to a boolean")
		return false
	} else if match {
		log.WithFields(log.Fields{"Request": req, "RouteDef": routedef}).Debug("Request matched route definition!")
		return true
	} else {
		log.WithFields(log.Fields{"Request": req, "RouteDef": routedef}).Debug("Request did not match route definition!")
		return false
	}
}

func (r *jsRouter) sanitizeRoute(routedef string) string {
	if strings.HasPrefix(routedef, "javascript:") {
		sanitizedRoutedef := routedef[len("javascript:"):]
		log.WithFields(log.Fields{"Original": routedef, "Sanitized": sanitizedRoutedef}).Debug("Sanitized route definition by removing the javascript prefix.")
		return sanitizedRoutedef
	}
	return routedef
}
