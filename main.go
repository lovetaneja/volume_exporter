package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/lovetaneja/volume_exporter/exporter"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	listenAddress = flag.String("web.listen-address", ":9888", "Address to listen on for web interface.")
	metricPath    = flag.String("web.metrics-path", "/metrics", "Path under which to expose metrics.")
	opts          = volumeOpts{}
)

type volumePaths []string

func (v *volumePaths) Set(value string) error {
	*v = append(*v, value)
	return nil
}

func (v *volumePaths) String() string {
	return fmt.Sprint()
}

type volumeOpts struct {
	volumePath volumePaths
}

func main() {
	log.Println("Starting Volume Exporter")

	// Set up command line options
	flag.Var(&opts.volumePath, "volume-dir", "Volume to report. Format is volumneName:volumeDir")
	flag.Parse()

	if len(opts.volumePath) < 1 {
		log.Println("Missing Volume Dir")
		flag.Usage()
		os.Exit(1)
	}
	exporter.Register()

	// Start the Server
	log.Fatal(serverMetrics(*listenAddress, *metricPath))
}

func serverMetrics(listenAddress, metricsPath string) error {
	http.Handle(metricsPath, promhttp.Handler())
	return http.ListenAndServe(listenAddress, nil)
}
