very simple http call logger

build

    go build cmd/httpcalllogger/main.go
    
docker

    docker build -t test .
    docker run --name temptest --rm -p10080:10080 test 
    
use

    curl -v -d'hello' localhost:10080/somepath
    docker exec -it temptest bash
    cat /header.txt 
    cat /data.bi

you can the tool it with go get -v github.com/codecare/httpcalllogger/cmd/httpcalllogger
