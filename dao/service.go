package dao

type ServiceDetail struct {
	Info 			*ServiceInfo 	`json:"info" description:"基本信息"`
	HTTPRule 		*HttpRule 		`json:"http_rule" description:""`
	TCPRule 		*TcpRule 		`json:"tcp_rule" description:""`
	GRPCRule 		*GrpcRule 		`json:"grpc_rule" description:""`
	LoadBalance 	*LoadBalance 	`json:"load_balance" description:""`
	AccessControl 	*AccessControl	`json:"access_control" description:""`
}