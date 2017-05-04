@echo off
set ssdbbinname=ssdb-server-1.9.4.exe

cd ..

if not exist nogit\ssdb-bin (
    echo nof found nogit\ssdb-bin
    pushd nogit
        echo download nogit\ssdb-bin
        git clone https://github.com/ideawu/ssdb-bin.git
    popd
)
set ssdbbin=nogit\ssdb-bin\%ssdbbinname%

if not exist %ssdbbin% (
    echo not found %ssdbbin%
    goto Exit0
)    

set id=10010
echo listen 127.0.0.1:%id%
if not exist ssdb\data\%id%\var (
    mkdir ssdb\data\%id%\var
)
start %ssdbbin% .\ssdb\config\%id%.conf

set id=10011
echo listen 127.0.0.1:%id%
if not exist ssdb\data\%id%\var (
    mkdir ssdb\data\%id%\var
)
start %ssdbbin% .\ssdb\config\%id%.conf

set id=10020
echo listen 127.0.0.1:%id%
if not exist ssdb\data\%id%\var (
    mkdir ssdb\data\%id%\var
)
start %ssdbbin% .\ssdb\config\%id%.conf

set id=10021
echo listen 127.0.0.1:%id%
if not exist ssdb\data\%id%\var (
    mkdir ssdb\data\%id%\var
)
start %ssdbbin% .\ssdb\config\%id%.conf
 
cd ssdb
:Exit0