// Copyright (c) 2019 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package m3msg

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/m3db/m3/src/aggregator/aggregator"
	xm3msg "github.com/m3db/m3/src/cmd/services/m3coordinator/server/m3msg"
	"github.com/m3db/m3/src/metrics/aggregation"
	"github.com/m3db/m3/src/metrics/metadata"
	"github.com/m3db/m3/src/metrics/metric"
	"github.com/m3db/m3/src/metrics/metric/aggregated"
	"github.com/m3db/m3/src/metrics/metric/unaggregated"
	"github.com/m3db/m3/src/metrics/policy"
	xtime "github.com/m3db/m3/src/x/time"

	"github.com/stretchr/testify/require"
)

var (
	defaultStoragePolicy = policy.NewStoragePolicy(10*time.Second, xtime.Second, 48*time.Hour)

	expectedMetric = aggregated.Metric{
		Type:      metric.GaugeType,
		ID:        []byte("FakeMetricID"),
		TimeNanos: 0,
		Value:     2.333,
	}

	expectedMetadata = metadata.TimedMetadata{
		AggregationID: aggregation.MustCompressTypes(aggregation.Last),
		StoragePolicy: defaultStoragePolicy,
	}
)

func TestGetWriteFn(t *testing.T) {
	aggregator := &mockAggregator{}
	writeFn := GetWriteFn(aggregator, nil, nil)

	callback := &mockCallback{
		record: make(map[xm3msg.CallbackType]int),
	}

	writeFn(
		context.TODO(),
		expectedMetric.ID,
		expectedMetric.TimeNanos,
		expectedMetric.TimeNanos,
		expectedMetric.Value,
		defaultStoragePolicy,
		callback,
	)

	require.Equal(t, map[xm3msg.CallbackType]int{xm3msg.OnSuccess: 1}, callback.record)
	require.Equal(t, 1, aggregator.count)
	require.Equal(t, expectedMetric, *aggregator.receivedMetrics[0])
	require.Equal(t, expectedMetadata, *aggregator.receivedMetadata[0])
}

type mockAggregator struct {
	sync.RWMutex

	count            int
	receivedMetrics  []*aggregated.Metric
	receivedMetadata []*metadata.TimedMetadata
}

func (agg *mockAggregator) Open() error {
	return nil
}

func (agg *mockAggregator) AddUntimed(metric unaggregated.MetricUnion, metas metadata.StagedMetadatas) error {
	return nil
}

func (agg *mockAggregator) AddTimed(metric aggregated.Metric, metadata metadata.TimedMetadata) error {
	return nil
}

func (agg *mockAggregator) AddForwarded(metric aggregated.ForwardedMetric, metadata metadata.ForwardMetadata) error {
	return nil
}

func (agg *mockAggregator) AddPassThrough(metric aggregated.Metric, metadata metadata.TimedMetadata) error {
	agg.Lock()
	defer agg.Unlock()

	agg.count++
	agg.receivedMetrics = append(agg.receivedMetrics, &metric)
	agg.receivedMetadata = append(agg.receivedMetadata, &metadata)
	return nil
}

func (agg *mockAggregator) Resign() error {
	return nil
}

func (agg *mockAggregator) Status() aggregator.RuntimeStatus {
	return aggregator.RuntimeStatus{
		FlushStatus: aggregator.FlushStatus{
			ElectionState: aggregator.UnknownState,
			CanLead:       true,
		},
	}
}

func (*mockAggregator) Close() error {
	return nil
}

type mockCallback struct {
	record map[xm3msg.CallbackType]int
}

func (c *mockCallback) Callback(t xm3msg.CallbackType) {
	c.record[t]++
}