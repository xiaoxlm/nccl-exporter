package nccl

import (
	"fmt"
	"testing"
)

func TestQueryMFU(t *testing.T) {
	lokiURL := "http://10.129.60.70:3100"

	var start int64 = 1738981383000000000
	var end int64 = 1738981637000000000
	gotMfuValue, err := NewNCCL().queryLoki(lokiURL, start, end)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(gotMfuValue)
}
