package exporter

import (
	"github.com/prometheus/client_golang/prometheus"
)

type volumeCollector struct {
	volumeBytesTotal *prometheus.Desc
	volumeBytesFree  *prometheus.Desc
	volumeBytesUsed  *prometheus.Desc
}

func newVolumeCollector() *volumeCollector {
	namespace := "volume"
	return &volumeCollector{
		volumeBytesTotal: prometheus.NewDesc(prometheus.BuildFQName(namespace, "", "bytes_total"), "Total Size", []string{"name", "path"}, nil),
		volumeBytesFree:  prometheus.NewDesc(prometheus.BuildFQName(namespace, "", "bytes_free"), "Free Size", []string{"name", "path"}, nil),
		volumeBytesUsed:  prometheus.NewDesc(prometheus.BuildFQName(namespace, "", "bytes_used"), "Used Size", []string{"name", "path"}, nil),
	}
}

func (collector *volumeCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.volumeBytesTotal
	ch <- collector.volumeBytesFree
	ch <- collector.volumeBytesUsed
}

func (collector *volumeCollector) Collect(ch chan<- prometheus.Metric) {
	var metricValue float64
	if 1 == 1 {
		metricValue = 20
	}
	ch <- prometheus.MustNewConstMetric(collector.volumeBytesTotal, prometheus.GaugeValue, metricValue, "log", "path")
	ch <- prometheus.MustNewConstMetric(collector.volumeBytesFree, prometheus.GaugeValue, metricValue-8, "log", "path")
	ch <- prometheus.MustNewConstMetric(collector.volumeBytesUsed, prometheus.GaugeValue, metricValue-12, "log", "path")
}

func Register() {
	collector := newVolumeCollector()
	prometheus.MustRegister(collector)
}
