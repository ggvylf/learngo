package config

type LogTransfer struct {
	KafkaCfg
	EsCfg
}

type KafkaCfg struct {
	Address string `ini:"address"`
	Topic   string `ini:"topic"`
}

type EsCfg struct {
	Address string `ini:"address"`
}
