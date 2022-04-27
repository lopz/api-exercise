package metric

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/middleware"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	httpDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name: "http_duration_seconds",
		Help: "Duration of HTTP requests.",
		//Buckets: prometheus.DefBuckets,
	}, []string{"code", "bytes", "method", "path"})

	httpRequest = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_request_total",
			Help: "How many HTTP requests processed",
		},
		[]string{"code", "method", "path", "remote_addr"},
	)
)

func init() {
	prometheus.Register(httpDuration)
	prometheus.MustRegister(httpRequest)
}

func MiddlewarePrometheus(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		httpRequest.WithLabelValues(strconv.Itoa(ww.Status()), r.Method, r.URL.Path, r.RemoteAddr).Inc()
		timer := prometheus.NewTimer(httpDuration.WithLabelValues(strconv.Itoa(ww.Status()), strconv.Itoa(ww.BytesWritten()), r.Method, r.URL.Path))
		next.ServeHTTP(ww, r)
		timer.ObserveDuration()
	})

}
