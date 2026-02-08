# homework_golang

cd blog_golang
go mod init "blog"
go mod tidy

go test ./tests -run TestSetupDemo -v

go test ./testutil -run TestDatabaseSetup -v