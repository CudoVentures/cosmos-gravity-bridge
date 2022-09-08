module github.com/althea-net/cosmos-gravity-bridge/module

go 1.15

require (
	github.com/confio/ics23/go v0.7.0 // indirect
	github.com/cosmos/cosmos-sdk v0.45.3
	github.com/ethereum/go-ethereum v1.10.3
	github.com/gin-gonic/gin v1.7.0 // indirect
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.5.2
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/pkg/errors v0.9.1
	github.com/regen-network/cosmos-proto v0.3.1
	github.com/spf13/cobra v1.4.0
	github.com/stretchr/objx v0.3.0 // indirect
	github.com/stretchr/testify v1.7.1
	github.com/tendermint/tendermint v0.34.19
	github.com/tendermint/tm-db v0.6.6
	google.golang.org/genproto v0.0.0-20211223182754-3ac035c7e7cb
	google.golang.org/grpc v1.45.0
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4

replace google.golang.org/grpc => google.golang.org/grpc v1.33.2

// replace github.com/cosmos/cosmos-sdk => ../../cosmos-sdk

replace github.com/cosmos/cosmos-sdk => github.com/CudoVentures/cosmos-sdk v0.0.0-20220816082327-65532d606824
