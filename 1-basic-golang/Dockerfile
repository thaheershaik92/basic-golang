###########################
#Base Image
###########################

From ubuntu As Build

Run apt-get update && apt-get install -y golang-go

Env GO111MODULE=off

Copy . .

Run CGO_ENABLED=0 go build -0 variable.go

ENTRYPOINT ["VARIABLE.GO"]
