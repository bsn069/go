@echo off
echo run in anywhere to config env

echo check gopath
if not exist %GOPATH% (
	echo not found %GOPATH%
	exit
)
cd \d %GOPATH% 

if not exist src (
	mkdir src
)

if not exist bin (
	mkdir bin
)

echo check bsn069\go
pushd src
	if not exist github.com\bsn069 (
		mkdir github.com\bsn069
	)

	pushd github.com\bsn069
		if not exist go (
			echo clone https://github.com/bsn069/go.git 
			git clone https://github.com/bsn069/go.git
		)

		pushd go
			if not exist nogit (
				mkdir nogit
			)
		popd
	popd
popd

echo check third part

pushd src
	echo check golang.org\x
	if not exist golang.org\x (
		mkdir golang.org\x
	)

	if not exist github.com\golang\crypto (
		echo download golang.org\x\crypto
		go get github.com/golang/crypto
	)
	if not exist golang.org\x\crypto (
		ln -s %cd%\github.com\golang\crypto %cd%\golang.org\x\crypto
	)

	if not exist github.com\golang\net (
		echo download golang.org\x\net
		go get github.com/golang/net
	)
	if not exist golang.org\x\net (
		ln -s %cd%\github.com\golang\net %cd%\golang.org\x\net
	)

	if not exist github.com\golang\tools (
		echo download golang.org\x\tools
		go get github.com/golang/tools
	)
	if not exist golang.org\x\tools (
		ln -s %cd%\github.com\golang\tools %cd%\golang.org\x\tools
	)
popd

pushd src
	if not exist github.com\golang\protobuf\protoc-gen-go (
		echo download github.com\golang\protobuf\protoc-gen-go
		go get github.com/golang/protobuf/protoc-gen-go
	)

	if not exist github.com\ssdb\gossdb (
		echo download github.com\ssdb\gossdb
		go get github.com/ssdb/gossdb
	)

	if not exist github.com\syndtr\goleveldb\leveldb (
		echo download github.com\syndtr\goleveldb\leveldb
		go get github.com/syndtr/goleveldb/leveldb
	)

	if not exist github.com\xtaci\kcp-go (
		echo download github.com\xtaci\kcp-go
		go get github.com/xtaci/kcp-go
	)
popd
 

if not exist bin\protoc-gen-go.exe (
	echo install protoc-gen-go
	pushd src\github.com\golang\protobuf\protoc-gen-go
		go install
	popd
)

pushd src\github.com\bsn069\go
echo %cd%
	if not exist nogit (
echo dsfds22
		mkdir nogit
	)
echo dsfds11

	pushd nogit
		if not exist "ssdb-bin" (
			echo download ssdb-bin
			echo git clone https:\github.com\ideawu\ssdb-bin.git
		)
	popd
popd
echo ds33fds

cd src\github.com\bsn069\go
pause