module github.com/sageflow/sageapi

go 1.15

require (
	github.com/99designs/gqlgen v0.13.0
	github.com/gin-gonic/contrib v0.0.0-20201101042839-6a891bf89f19
	github.com/gin-gonic/gin v1.6.3
	github.com/golang/protobuf v1.4.1
	github.com/mwitkow/go-proto-validators v0.3.2
	github.com/sageflow/sageflow v0.0.0-20210108233356-e663f3625227
	github.com/soheilhy/cmux v0.1.4
	github.com/stretchr/testify v1.6.1 // indirect
	github.com/vektah/gqlparser/v2 v2.1.0
	golang.org/x/sync v0.0.0-20190423024810-112230192c58
	google.golang.org/grpc v1.27.0
	google.golang.org/protobuf v1.25.0
)

replace github.com/sageflow/sageflow v0.0.0-20210108233356-e663f3625227 => ../sageflow
