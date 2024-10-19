package constant

type AppList struct {
	App        map[string]AppNet `json:"app"`
	RegionName string            `json:"region_name"`
	MqAddr     string            `json:"mq_addr"`
	GrpcAddr   string            `json:"grpc_addr"`
	GateTcp    bool              `json:"gate_tcp"`
}

type AppNet struct {
	InnerAddr string `json:"InnerAddr"`
	InnerPort string `json:"InnerPort"`
	OuterAddr string `json:"OuterAddr"`
	OuterPort string `json:"OuterPort"`
}

type MysqlConf struct {
	Dsn string `json:"dsn"`
}

type RedisConf struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}
