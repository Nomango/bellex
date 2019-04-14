@echo off
setlocal enabledelayedexpansion

if {%1}=={}                 goto all
if {%1}=={help}             goto help
if {%1}=={run} (
    if {%2}=={server}       goto run_server
)
if {%1}=={build} (
    if {%2}=={server}       goto build_server
    if {%2}=={ntp}          goto build_ntp_client
    if {%2}=={tcp}          goto build_tcp_client
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
    echo test
    echo     tcp        - Send test data to the TCP server
    echo     ntp        - Send test data to the NTP server
    echo.
    exit


:run_server
    go run .\server\main.go
    exit

:build_server
    go build -o server.exe .\server\main.go
    exit

:build_ntp_client
    go build -o tcp_client.exe .\client\ntp\
    exit

:build_tcp_client
    go build -o tcp_client.exe .\client\tcp\
    exit

:test_tcp
    go run .\client\tcp\
    exit

:test_ntp
    go run .\client\ntp\
    exit
