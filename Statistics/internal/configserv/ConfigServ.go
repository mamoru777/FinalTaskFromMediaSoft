package configserv

type ConfigServ struct {
	GRPCAddr string `env:"USER_GRPC_ADDR" envDefault:":13997"`
}
