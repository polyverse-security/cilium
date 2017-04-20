# js-route
This library serves to augment github.com/vulcand/vulcand by allowing Javascript-based expressions for routing.

This library seamlessly plugs into vulcand, after the pluggable routing library PR is pulled in (or you can 
revendor vulcand from the polyverse fork temporarily.)

This library allows routes to be arbitrary Javascript expressions that MUST evaluate to a boolean. Each expression
is given a well-known Request object, which is the Golang http.Request struct.

The javascript evaluator is a native Go evaluator called Otto.

Example rules can look like:

Request.Header["User-Agent"] == "something".

You can use nearly all javascript functions such as:

Request.URL.Path.contains("redflag")

Naturally this implementation is over-the-top inefficient - since the javascript is interpreted at runtime, AND 
routes have to be matched linearly.

But the library does provide a quick escape path to write rules fast.
