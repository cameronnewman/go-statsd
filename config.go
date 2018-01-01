package statsd

//Config for statsd
type Config struct {
	StatsdHost      string `json:"statsd_host"`
	StatsdPort      int    `json:"statsd_port"`
	StatsdNamespace string `json:"statsd_namespace"`
}
