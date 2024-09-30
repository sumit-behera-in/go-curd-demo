# starting commands
go mod init "name of project"
go get github.com/gin-gonic/gin
go get go.mongodb.org/mongo-driver


# for any dependency problem use

go get -u
go get github.com/golang/snappy
go mod tidy
go build