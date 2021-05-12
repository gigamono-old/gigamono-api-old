module github.com/gigamono/gigamono-api

go 1.15

require (
	github.com/99designs/gqlgen v0.13.0
	github.com/gigamono/gigamono v0.0.0-20210510234020-a879327fc4e6
	github.com/gin-gonic/contrib v0.0.0-20201101042839-6a891bf89f19
	github.com/gin-gonic/gin v1.7.1
	github.com/go-playground/validator/v10 v10.4.1
	github.com/soheilhy/cmux v0.1.4
	github.com/vektah/gqlparser/v2 v2.1.0
	golang.org/x/sync v0.0.0-20190423024810-112230192c58
	google.golang.org/grpc v1.27.0
)

replace github.com/gigamono/gigamono v0.0.0-20210510234020-a879327fc4e6 => ../gigamono
