package config

type LogTransfer struct {
	KafkaCfg `ini:"kafka"`
	EsCfg    `ini:"es"`
}

type KafkaCfg struct {
	Address string `ini:"address"`
	Topic   string `ini:"topic"`
}

type EsCfg struct {
	Address    string `ini:"address"`
	ChanSize   int    `ini:"chansize"`
	ChanCounts int    `ini: "nums"`
}
