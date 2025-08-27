package exam1

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	requestsTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "myapp_requests_total",
			Help: "Количество обработанных запросов",
		},
	)

	responseTime = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "myapp_response_time_seconds",
			Help:    "Время ответа",
			Buckets: prometheus.DefBuckets,
		},
	)
)

func main() {
	// Регистрируем метрики
	prometheus.MustRegister(requestsTotal)
	prometheus.MustRegister(responseTime)

	// Твой хендлер
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		timer := prometheus.NewTimer(responseTime)
		defer timer.ObserveDuration()

		requestsTotal.Inc()
		fmt.Fprintf(w, "Hello, Prometheus!")
	})

	// Эндпоинт для метрик
	http.Handle("/metrics", promhttp.Handler())

	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
