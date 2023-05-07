To init dependencies

go mod int example/Car_API

To download gin framework

go get github.com/gin-gonic/gin

curl localhost:8080/cars

curl localhost:8080/cars --include --header "Content-Type: application/json" -d @body.json --request "POST"

if needed to stop server, run the following command on terminal,
lsof -n -i4TCP:8080
kill -9 <PID>

curl localhost:8080/buycar?id=2 --request "PUT"
curl localhost:8080/sellcar?id=4 --request "PUT"