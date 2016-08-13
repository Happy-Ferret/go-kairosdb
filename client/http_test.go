package client

import (
    "github.com/dan-compton/go-kairosdb/kairosdb"
    "testing"
    "time"
)

var testHttpClient = New("http://localhost:8080")

// Test Query for built-in kairosdb.jvm.free_memory
func TestQueryMetrics(t *testing.T) {
    req := &kairosdb.QueryMetricsRequest{
        Metrics:[]*kairosdb.QueryMetric{
            &kairosdb.QueryMetric{
                Name: "kairosdb.jvm.free_memory",
            },
         },
         CacheTime: 0,
        StartAbsolute: (time.Now().Add(-5*time.Minute)).UnixNano()/int64(time.Millisecond),
        EndAbsolute: time.Now().UnixNano()/int64(time.Millisecond),
    }

    resp, err := testHttpClient.Query(req)
    if err != nil {
        t.Fatalf("client return err%s", err.Error())
    }
    if len(resp.Errors) > 0{
        t.Fatalf("kairosdb return err: %v", resp.Errors)
    }
}