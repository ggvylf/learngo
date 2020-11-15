package conf

type AppConf struct {
	KafkaConf `ini:"kafka"`
	EtcdConf  `init: "etcd"`
}

type KafkaConf struct {
	Address     string `ini:"address"`
	Topic       string `ini:"topic"`
	ChanMaxSize int    `ini: "chan_max_size"`
}

type EtcdConf struct {
	Address string `ini:"address"`
	Timeout int    `ini:"timeout"`
	Key     string `ini: "collect_log_key"`
}
