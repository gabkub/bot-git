SET GOOS=windows
SET GOARCH=amd64

cd ../main

go build -o ../bin/windows/MattermostBot.exe

SET GOOS=linux
SET GOARCH=amd64

go build -o ../bin/linux/MattermostBot