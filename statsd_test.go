package statsd

import (
	"reflect"
	"testing"
	"time"

	sd "github.com/cactus/go-statsd-client/statsd"
)

func TestNew(t *testing.T) {
	type args struct {
		opt *Options
	}
	tests := []struct {
		name    string
		args    args
		want    *Statsd
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.opt)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatsd_Counter(t *testing.T) {
	type fields struct {
		statsd sd.Statter
	}
	type args struct {
		key   string
		value int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Statsd{
				statsd: tt.fields.statsd,
			}
			if err := s.Counter(tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Statsd.Counter() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStatsd_Gauge(t *testing.T) {
	type fields struct {
		statsd sd.Statter
	}
	type args struct {
		key   string
		value int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Statsd{
				statsd: tt.fields.statsd,
			}
			if err := s.Gauge(tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Statsd.Gauge() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStatsd_timers(t *testing.T) {
	type fields struct {
		statsd sd.Statter
	}
	type args struct {
		key   string
		value int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Statsd{
				statsd: tt.fields.statsd,
			}
			if err := s.timers(tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Statsd.timers() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStatsd_Timing(t *testing.T) {
	type fields struct {
		statsd sd.Statter
	}
	type args struct {
		key   string
		start time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Statsd{
				statsd: tt.fields.statsd,
			}
			if err := s.Timing(tt.args.key, tt.args.start); (err != nil) != tt.wantErr {
				t.Errorf("Statsd.Timing() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStatsd_TimingDuration(t *testing.T) {
	type fields struct {
		statsd sd.Statter
	}
	type args struct {
		key      string
		duration time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Statsd{
				statsd: tt.fields.statsd,
			}
			if err := s.TimingDuration(tt.args.key, tt.args.duration); (err != nil) != tt.wantErr {
				t.Errorf("Statsd.TimingDuration() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
