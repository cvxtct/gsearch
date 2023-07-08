coverage https://stackoverflow.com/questions/10516662/how-to-measure-test-coverage-in-go
https://go.dev/doc/diagnostics
https://medium.com/compass-true-north/memory-profiling-a-go-service-cd62b90619f9
https://github.com/DataDog/go-profiler-notes/blob/main/guide/README.md

go test -v . -coverprofile=c.out
go tool cover -html=c.out
go run -race .

$ go test -race mypkg    // to test the package
$ go run -race mysrc.go  // to run the source file
$ go build -race mycmd   // to build the command
$ go install -race mypkg // to install the package