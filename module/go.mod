module github.com/althea-net/cosmos-gravity-bridge/module

go 1.15

require (
	github.com/cometbft/cometbft v0.37.1
	github.com/cometbft/cometbft-db v0.7.0
	github.com/cosmos/cosmos-sdk v0.47.0-rc2.0.20230220103612-f094a0c33410
	github.com/cosmos/gogoproto v1.4.10
	github.com/ethereum/go-ethereum v1.11.6
	github.com/gogo/protobuf v1.3.2
	github.com/golang/protobuf v1.5.3
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/pkg/errors v0.9.1
	github.com/regen-network/cosmos-proto v0.3.1
	github.com/spf13/cobra v1.6.1
	github.com/stretchr/testify v1.8.2
	google.golang.org/genproto v0.0.0-20230306155012-7f2fa6fef1f4
	google.golang.org/grpc v1.55.0
)

// replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4

// replace github.com/cosmos/cosmos-sdk => ../../cosmos-sdk

replace github.com/cosmos/cosmos-sdk => github.com/CudoVentures/cosmos-sdk v0.0.0-20230628122035-744b6e4f35cb
