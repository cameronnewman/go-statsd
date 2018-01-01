package statsd

import (
	"fmt"
	"time"

	sd "github.com/cactus/go-statsd-client/statsd"
)

const sampleRate = 1.0

//Stats interface
type Stats interface {
	StatsCounter(key string, value int64)
	StatsGauge(key string, value int64)
	Timing(key string, start time.Time)
	TimingDuration(key string, duration time.Duration)
}

//Statsd used to monitor and measure performance
type Statsd struct {
	sd sd.Statter
}

//New creates a statsd client instance
func New(cfg *Config) (*Statsd, error) {
	conn := fmt.Sprintf("%s:%d", cfg.StatsdHost, cfg.StatsdPort)

	statsd := &Statsd{}
	stats, err := sd.NewBufferedClient(conn, cfg.StatsdNamespace, 0, 0)
	if nil != err {
		return nil, err
	}
	statsd.sd = stats
	return statsd, nil

}

//StatsCounter increase counter for key
func (stsd *Statsd) StatsCounter(key string, value int64) {
	if stsd != nil && stsd.sd != nil {
		stsd.sd.Inc(key, value, sampleRate)
	}
}

//StatsGauge gauge for key
func (stsd *Statsd) StatsGauge(key string, value int64) {
	if stsd != nil && stsd.sd != nil {
		stsd.sd.Gauge(key, value, sampleRate)

	}
}

func (stsd *Statsd) statsTimers(key string, value int64) {
	if stsd != nil && stsd.sd != nil {
		stsd.sd.Timing(key, value, sampleRate)
	}
}

//Timing measure for key
func (stsd *Statsd) Timing(key string, start time.Time) {
	elapsed := time.Since(start)

	stsd.statsTimers(key, elapsed.Nanoseconds()/1000000)
}

//TimingDuration measures duration for key
func (stsd *Statsd) TimingDuration(key string, duration time.Duration) {
	if stsd != nil && stsd.sd != nil {
		stsd.sd.TimingDuration(key, duration, sampleRate)
	}
}
