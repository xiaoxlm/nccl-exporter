package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"github.com/xiaoxlm/nccl-exporter/global"
	"github.com/xiaoxlm/nccl-exporter/pkg/nccl"
	pkgPrometheus "github.com/xiaoxlm/nccl-exporter/pkg/prometheus"
	"log"
	"net/http"
	"time"
)

var ncclMetrics *nccl.NCCL

func init() {
	ncclMetrics = nccl.NewNCCL()
	gauge := ncclMetrics.GetGauge()
	http.Handle("/metrics", promhttp.HandlerFor(pkgPrometheus.NewMetricsRegistry(map[string]string{
		"service":    "nccl-exporter",
		"ai_metrics": global.NCCLMetricsLabel,
	}, gauge), promhttp.HandlerOpts{}))
}

func main() {
	go func() {
		for {
			if err := ncclMetrics.SetGaugeValue(); err != nil {
				logrus.Errorf("setGaugeValue error. err:%v ", err)
			}
			time.Sleep(2 * time.Second)
		}
	}()

	fmt.Println("Starting server at :9134")
	if err := http.ListenAndServe(":9134", nil); err != nil {
		log.Fatalln(err)
	}
}
