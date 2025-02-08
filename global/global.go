package global

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

var (
	NCCLMetricsLabel = os.Getenv("NCCL_METRICS_LABEL")
	LokiURL          = os.Getenv("LOKI_URL")
)

func init() {
	fmt.Printf(`
====env var====
NCCL_METRICS_LABEL=%s
LOKI_URL=%s
====env var====

`, NCCLMetricsLabel, LokiURL)
	if NCCLMetricsLabel == "" {
		logrus.Fatalln("env var NCCLMetricsLabel is empty")
	}

	if LokiURL == "" {
		logrus.Fatalln("env var LOKI_URL is empty")
	}

}
