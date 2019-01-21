@echo off
setlocal enabledelayedexpansion

if {%1}=={}                 goto all
if {%1}=={help}             goto help
if {%1}=={run} (
    if {%2}=={server}       goto run_server
)
if {%1}=={build} (
    if {%2}=={server}       goto build_server
    if {%2}=={ntp}          goto build_ntp
    if {%2}=={tcp}          goto build_tcp_server
)
if {%1}=={test} (
    if {%2}=={ntp}          goto test_ntp
    if {%2}=={tcp}          goto test_tcp
)
goto help

:all

:help
    echo.
    echo Usage: ./make [option]
    echo.
    echo Options:
    echo run
    echo     server     - Run server app
    echo build
    echo     server     - Compile server app
    echo     tcp        - Compile TCP server
    echo     ntp        - Compile NTP client
    echo test
    echo     tcp        - Send test data to the TCP server
    echo     ntp        - Send test data to the NTP server
    echo.
    exit


:run_server
    go run .\main.go
    exit

:build_server
    go build -o server.exe .\main.go
    exit

:build_tcp_server
    go build -o tcpserver.exe .\tcp\server\main\main.go
    exit

:build_ntp
    go build -o ntpclient.exe .\ntp\main.go
    exit

:test_tcp
    go run .\tcp\client\main.go
    exit

:test_ntp
    go run .\ntp\main.go
    exit
