package metrics

import (
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

// WrapHandler wraps an http.Handler with OpenTelemetry instrumentation
func (m *AppMetrics) WrapHandler(handler http.Handler) http.Handler {
	return otelhttp.NewHandler(handler, MeterName)
}
