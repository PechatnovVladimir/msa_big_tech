module github.com/PechatnovVladimir/msa_big_tech/gateway

go 1.25.1

require (
	github.com/PechatnovVladimir/msa_big_tech/auth v0.0.0-20250920122254-7806f6c3d5b1
	github.com/PechatnovVladimir/msa_big_tech/chat v0.0.0-20250920122254-7806f6c3d5b1
	github.com/PechatnovVladimir/msa_big_tech/social v0.0.0-20250920122254-7806f6c3d5b1
	github.com/PechatnovVladimir/msa_big_tech/users v0.0.0-20250920122254-7806f6c3d5b1
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.27.2
	google.golang.org/grpc v1.75.1
)

//replace (
//	github.com/PechatnovVladimir/msa_big_tech/auth => ../auth
//	github.com/PechatnovVladimir/msa_big_tech/chat => ../chat
//	github.com/PechatnovVladimir/msa_big_tech/social => ../social
//	github.com/PechatnovVladimir/msa_big_tech/users => ../users
//)

require (
	golang.org/x/net v0.44.0 // indirect
	golang.org/x/sys v0.36.0 // indirect
	golang.org/x/text v0.29.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20250818200422-3122310a409c // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250908214217-97024824d090 // indirect
	google.golang.org/protobuf v1.36.9 // indirect
)
