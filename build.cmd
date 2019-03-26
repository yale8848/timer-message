set GOOS=windows
set GOARCH=amd64
set CGO_ENABLED=1
set CC=E:\go-vgo\Mingw\bin\gcc.exe
set CXX=E:\go-vgo\Mingw\bin\g++.exe
go build  -ldflags "-H windowsgui"  -o timer-message(0b5_a_a_a_a_a).exe -x ./main/main.go
rem go build  -ldflags "-H windowsgui"  -o timer-message(@hourly).exe -x ./main/main.go
