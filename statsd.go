package statsd

import (
	"fmt"
	"time"

	sd "github.com/cactus/go-statsd-client/statsd"
)

const sampleRate = 1.0

//Metrics interface
type Metrics interface {
	Counter(key string, value int64)
	Gauge(key string, value int64)
	Timing(key string, start time.Time)
	TimingDuration(key string, duration time.Duration)
}

//Statsd used to monitor and measure performance
type Statsd struct {
	statsd sd.Statter
}

//Options for statsd
type Options struct {
	Host      string `json:"host"`
	Port      int    `json:"port"`
	Namespace string `json:"namespace"`
}

//New creates a statsd client instance
func New(opt *Options) (*Statsd, error) {
	conn := fmt.Sprintf("%s:%d", opt.Host, opt.Port)
	statter, err := sd.NewClientWithConfig(&sd.ClientConfig{
		Address: conn,
		Prefix:  opt.Namespace,
	})
	if nil != err {
		return nil, err
	}

	return &Statsd{
		statsd: statter,
	}, nil
}

//Counter increase counter for key
func (s *Statsd) Counter(key string, value int64) error {
	if s != nil && s.statsd != nil {
		err := s.statsd.Inc(key, value, sampleRate)
		if err != nil {
			return err
		}
	}
	return nil
}

//Gauge gauge for key
func (s *Statsd) Gauge(key string, value int64) error {
	if s != nil && s.statsd != nil {
		err := s.statsd.Gauge(key, value, sampleRate)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Statsd) timers(key string, value int64) error {
	if s != nil && s.statsd != nil {
		err := s.statsd.Timing(key, value, sampleRate)
		if err != nil {
			return err
		}
	}
	return nil
}

//Timing measure for key
func (s *Statsd) Timing(key string, start time.Time) error {
	elapsed := time.Since(start)
	err := s.timers(key, elapsed.Nanoseconds()/1000000)
	if err != nil {
		return err
	}
	return nil
}

//TimingDuration measures duration for key
func (s *Statsd) TimingDuration(key string, duration time.Duration) error {
	if s != nil && s.statsd != nil {
		err := s.statsd.TimingDuration(key, duration, sampleRate)
		if err != nil {
			return err
		}
	}
	return nil
}
