@echo off
setlocal enabledelayedexpansion

if {%1}=={}                         goto all
if {%1}=={help}                     goto help
if {%1}=={build}                    goto build
if {%1}=={build_tcp_server}         goto build_tcp_server
if {%1}=={build_tcp_client}         goto build_tcp_client
if {%1}=={test_tcp}                 goto test_tcp
if {%1}=={test_ntp}                 goto test_ntp
goto help

:all

:help
    echo.
    echo Usage: ./make [option]
    echo.
    echo Options:
    echo build                    - build server app
    echo build_tcp_server         - build tcp server
    echo build_tcp_client         - build tcp client
    echo.
    exit

:build
    go build -o server.exe .\main.go
    exit

:build_tcp_server
    go build -o tcpserver.exe .\tcp\server\main.go
    exit

:build_tcp_client
    go build -o tcpclient.exe .\tcp\client\main.go
    exit

:test_tcp
    go run .\tcp\client\main.go
    exit

:test_ntp
    go run .\ntp\main.go
    exit
