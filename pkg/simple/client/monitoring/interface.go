/*
Copyright 2020 KubeSphere Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package monitoring

import (
	"time"
)

type Interface interface {
	GetMetric(expr string, time time.Time) Metric
	GetMetricOverTime(expr string, start, end time.Time, step time.Duration) Metric
	GetNamedMetrics(metrics []string, time time.Time, opt QueryOption) []Metric
	GetNamedMetricsOverTime(metrics []string, start, end time.Time, step time.Duration, opt QueryOption) []Metric
	GetMetadata(namespace string) []Metadata
	GetLabelValues(label string, matches []string, start, end time.Time) []string
	GetMetricLabelSet(expr string, start, end time.Time) []map[string]string

	// meter
	GetNamedMeters(meters []string, time time.Time, opts []QueryOption) []Metric
	GetNamedMetersOverTime(metrics []string, start, end time.Time, step time.Duration, opts []QueryOption) []Metric
}
