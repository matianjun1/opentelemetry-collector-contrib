// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package servicegraphprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/servicegraphprocessor"

import (
	"time"
)

// Config defines the configuration options for servicegraphprocessor.
type Config struct {

	// MetricsExporter is the name of the metrics exporter to use to ship metrics.
	MetricsExporter string `mapstructure:"metrics_exporter"`

	// LatencyHistogramBuckets is the list of durations representing latency histogram buckets.
	// See defaultLatencyHistogramBucketsMs in processor.go for the default value.
	LatencyHistogramBuckets []time.Duration `mapstructure:"latency_histogram_buckets"`

	// Dimensions defines the list of additional dimensions on top of the provided:
	// - client
	// - server
	// - failed
	// - connection_type
	// The dimensions will be fetched from the span's attributes. Examples of some conventionally used attributes:
	// https://github.com/open-telemetry/opentelemetry-collector/blob/main/model/semconv/opentelemetry.go.
	Dimensions []string `mapstructure:"dimensions"`

	// Store contains the config for the in-memory store used to find requests between services by pairing spans.
	Store StoreConfig `mapstructure:"store"`
	// CacheLoop is the time to cleans the cache periodically.
	CacheLoop time.Duration `mapstructure:"cache_loop"`
	// CacheLoop is the time to expire old entries from the store periodically.
	StoreExpirationLoop time.Duration `mapstructure:"store_expiration_loop"`
}

type StoreConfig struct {
	// MaxItems is the maximum number of items to keep in the store.
	MaxItems int `mapstructure:"max_items"`
	// TTL is the time to live for items in the store.
	TTL time.Duration `mapstructure:"ttl"`
}
