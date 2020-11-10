package conf

type AppConf struct {
	KafkaConf `ini:"kafka"`
	EtcdConf  `init: "etcd"`
}

type KafkaConf struct {
	Address string `ini:"address"`
	Topic   string `ini:"topic"`
}

type EtcdConf struct {
	Address string `ini:"address"`
}
