
set CURDIR=`pwd`
set OLDGOPATH=%$GOPATH%
set GOPATH=%cd%

# go install ?
bin/revel run revelapp

set GOPATH=%OLDGOPATH%
