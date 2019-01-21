@echo off
setlocal enabledelayedexpansion

if {%1}=={}                         goto all
if {%1}=={help}                     goto help
if {%1}=={build}                    goto build
if {%1}=={build_tcp_server}         goto build_tcp_server
if {%1}=={test_tcp}                 goto test_tcp
if {%1}=={test_ntp}                 goto test_ntp
goto help

:all

:help
    echo.
    echo Usage: ./make [option]
    echo.
    echo Options:
    echo build                - Compile server app
    echo build_tcp_server     - Compile tcp server
    echo test_tcp             - Send test data to the TCP server
    echo test_ntp             - Send test data to the NTP server
    echo.
    exit

:build
    go build -o server.exe .\main.go
    exit

:build_tcp_server
    go build -o tcpserver.exe .\tcp\server\main\main.go
    exit

:test_tcp
    go run .\tcp\client\main.go
    exit

:test_ntp
    go run .\ntp\main.go
    exit
