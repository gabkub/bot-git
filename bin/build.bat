SET CURRENT_DIR=%~dp0
SET GOOS=windows
SET GOARCH=amd64
cd ../main

go build -o ../bin/windows/MattermostBot.exe
IF %ERRORLEVEL% NEQ 0 (
   exit %ERRORLEVEL%
)
SET GOOS=linux
SET GOARCH=amd64

go build -o ../bin/linux/MattermostBot
IF %ERRORLEVEL% NEQ 0 (
   exit %ERRORLEVEL%
)
SET GOOS=windows
cd %CURRENT_DIR%