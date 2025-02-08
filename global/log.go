package global

import "github.com/xiaoxlm/nccl-exporter/pkg/log"

func init() {
	(&log.Log{}).SetDefaults().Build()
}
