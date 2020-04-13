package conf

type LogTransferCfg struct {
	KafkaCfg	`ini:"kafka"`
	ESCfg		`ini:"es"`

}

// Kafka 配置
type KafkaCfg struct {
	Address string `ini:"address"`
	Topic string `ini:"topic"`
}

// es 配置
type ESCfg struct {
	Address string `ini:"address"`
}