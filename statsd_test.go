package statsd

import (
	"reflect"
	"testing"
	"time"

	sd "github.com/cactus/go-statsd-client/statsd"
)

func TestNew(t *testing.T) {
	type args struct {
		cfg *Config
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
			got, err := New(tt.args.cfg)
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

func TestStatsd_StatsCounter(t *testing.T) {
	type fields struct {
		sd sd.Statter
	}
	type args struct {
		key   string
		value int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stsd := &Statsd{
				sd: tt.fields.sd,
			}
			stsd.StatsCounter(tt.args.key, tt.args.value)
		})
	}
}

func TestStatsd_StatsGauge(t *testing.T) {
	type fields struct {
		sd sd.Statter
	}
	type args struct {
		key   string
		value int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stsd := &Statsd{
				sd: tt.fields.sd,
			}
			stsd.StatsGauge(tt.args.key, tt.args.value)
		})
	}
}

func TestStatsd_statsTimers(t *testing.T) {
	type fields struct {
		sd sd.Statter
	}
	type args struct {
		key   string
		value int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stsd := &Statsd{
				sd: tt.fields.sd,
			}
			stsd.statsTimers(tt.args.key, tt.args.value)
		})
	}
}

func TestStatsd_Timing(t *testing.T) {
	type fields struct {
		sd sd.Statter
	}
	type args struct {
		key   string
		start time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stsd := &Statsd{
				sd: tt.fields.sd,
			}
			stsd.Timing(tt.args.key, tt.args.start)
		})
	}
}
