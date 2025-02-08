package nccl

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
	"github.com/xiaoxlm/nccl-exporter/global"
	"github.com/xiaoxlm/nccl-exporter/pkg/loki"
	"strings"
	"time"
)

type NCCL struct {
	query string
	gauge prometheus.Gauge
}

func NewNCCL() *NCCL {
	return &NCCL{
		query: `{app="dlrover"} |= "RuntimeError:"`, // RuntimeError: NCCL error
		gauge: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "nccl_error", // 指标名称
			Help: "NCCL error", // 帮助信息
		}),
	}
}

func (nccl *NCCL) GetGauge() prometheus.Gauge {
	return nccl.gauge
}

func (nccl *NCCL) SetGaugeValue() error {
	now := time.Now()
	var start int64 = now.Add(-10 * time.Second).UnixNano()
	var end int64 = now.UnixNano()
	resp, err := nccl.queryLoki(global.LokiURL, start, end)
	if err != nil {
		return err
	}

	if !resp.GetRecords { // 只有没有数据的时候才设置成0
		nccl.gauge.Set(0)
		return nil
	}

	if resp.Value != 0 { // 有数据，只是日志不对
		nccl.gauge.Set(resp.Value)
	}

	return nil
}

func (nccl *NCCL) queryLoki(lokiURL string, start, end int64) (value *RESP, err error) {
	resp, err := loki.QueryLoki(lokiURL, nccl.query, start, end)
	if err != nil {
		return nil, err
	}

	if len(resp.Data.Result) == 0 { // 查询不到日志数据
		logrus.Warningf("no records from loki. start: %d, end: %d", start, end)
		return &RESP{}, nil
	}

	var (
		runtimeErrorCount float64
		Labels            = make(map[string]string)
	)
	for _, res := range resp.Data.Result {
		Labels = res.Stream

		if len(res.Values) < 1 {
			continue
		}

		for _, values := range res.Values {
			val := values.([]interface{})[1].(string)
			if !strings.Contains(val, "RuntimeError") {
				continue
			}
			//if !strings.Contains(val, "NCCL") {
			//	continue
			//}
			runtimeErrorCount++
		}

	}

	if runtimeErrorCount == 0 {
		return &RESP{GetRecords: true, Labels: Labels}, nil
	}

	return &RESP{
		GetRecords: true,
		Value:      runtimeErrorCount,
		Labels:     Labels,
	}, nil
}

type RESP struct {
	GetRecords bool
	Value      float64
	Labels     map[string]string
}
