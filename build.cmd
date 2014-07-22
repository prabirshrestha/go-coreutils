@pushd %~dp0
@if not exist bin md bin
@cd bin
@del * /q
go build ..\src\base64
go build ..\src\env
go build ..\src\false
go build ..\src\logname
go build ..\src\ls
go build ..\src\mkdir
go build ..\src\mktemp
go build ..\src\pwd
go build ..\src\realpath
go build ..\src\seq
go build ..\src\sleep
go build ..\src\true
go build ..\src\uptime
go build ..\src\whoami
@popd
