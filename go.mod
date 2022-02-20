module github.com/DymaSV/shippy-consignment-cli

go 1.16

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/DymaSV/shippy-consignment-server v0.0.0-20220202184220-9999301864ce
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.9.1
	google.golang.org/grpc v1.44.0 // indirect
)
