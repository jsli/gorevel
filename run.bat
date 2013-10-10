@echo off

set OLDGOPATH=%GOPATH%
set GOPATH=%cd%

bin\revel run revelapp

set GOPATH=%OLDGOPATH%
