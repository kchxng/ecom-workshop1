package middlewares

import ginprometheus "github.com/zsais/go-gin-prometheus"

func NewCustomPrometheus(subsystem string) *ginprometheus.Prometheus {
	return ginprometheus.NewPrometheus(subsystem)
}
