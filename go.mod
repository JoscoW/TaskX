module taskx

go 1.22.0

replace taskx-server => ./server

require (
	github.com/go-sql-driver/mysql v1.8.1
	github.com/spf13/cobra v1.8.0
	github.com/stretchr/testify v1.9.0
	google.golang.org/grpc v1.63.2
	google.golang.org/protobuf v1.34.0
	taskx-server v0.0.0-00010101000000-000000000000
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/net v0.21.0 // indirect
	golang.org/x/sys v0.17.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240227224415-6ceb2ff114de // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
