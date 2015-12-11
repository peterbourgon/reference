## Motivation

This is a list of high level features and functionality a reference
Go kit REST-ish app should have.  Such an application might be
presented to developers as a way to structure new Go kit applications,
serving as a common organizational point of departure for new
applications.

* The app should expose URLs that are familiar to programmers with
experience with REST APIs.  This means different things to different
people, but generally URLs should telegraph the thing being talked
about.  Ideally, URLs may contain path tokens (/users/:id) that can
be easily evaluated in the proximity of the Go kit Endpoint.

* HTTP verbs should be part of the request routing.  The use of a
well known mux such as [Gorilla Mux](https://github.com/gorilla/mux)
or [httprouter](https://github.com/julienschmidt/httprouter) is an
acceptable way to satisfy this requirement, together with the
previous requirement around URL tokens.

* The application should demonstrate how to use a configurable
security middleware.  OAuth2 would be a reasonable choice to
demonstrate.  If HTTP verbs and URL paths both determine routing
per above, demonstrating how OAuth2 scopes can be applied to specific
route is desirable.

* The application should have the ability to produce log records
that are consistent with an organization's existing or expected
formats.

* Go kit already uses the Go context.Context object.  An example
of how to use this for request timeout and cancellation, and request
context values read/write is desirable.

* The app should expose _internal_ management endpoints that expose
general info, health indicators, service dependencies, metrics, and
other documentation references.  These endpoints typically reside
at /internal/info, /internal/health, etc.  The /internal/info
endpoint should contain enough information for a proxy to register
this application in service discovery on behalf of the application.
The proxy might periodically poll the /internal/health endpoint to
maintain confidence that the app should remain registered and
in-service.

* The app should demonstrate how to use Go kit circuitbreaker
functionality when calling remote services.

* The app should demonstrate how to use Go kit request tracing.

* The app should demonstrate recommended ways to connect to common
data services, including MySQL, RabbitMQ, Redis, and Zookeeper.

* The application should be capable of pushing metrics to its
destination every minute, or other configurable time period.

* A demonstration of graceful application restart and shutdown would
be a nice to have. Graceful meaning take no new requests, but allow
in-flight requests to complete.

* A demonstration of how to publish Swagger docs for the app's
public API is desirable.




