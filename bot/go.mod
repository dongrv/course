module bot

go 1.21.0

require (
	casino v0.0.0-incompatible
	github.com/sirupsen/logrus v1.8.1
	google.golang.org/protobuf v1.28.1
)

require golang.org/x/sys v0.0.0-20220728004956-3c1f35247d10 // indirect

replace (
	casino => ./../../../betaServer // TODO 业务服路径
	github.com/aaagame/common => ./../../../go_common
)
