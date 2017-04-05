@echo off

if not exist nogit (
	mkdir nogit
) 

echo setup tolua
if not exist u3d\Assets\Bsn\ThirdPart\tolua (
	pushd nogit
	if not exist tolua (
		echo clone tolua from https://github.com/topameng/tolua.git
		git clone https://github.com/topameng/tolua.git
	)
	popd

	xcopy nogit\tolua u3d\Assets\Bsn\ThirdPart\tolua /E /Y /I
	pushd u3d\Assets\Bsn\ThirdPart\tolua
		rmdir Luajit /S /Q
		rmdir ProjectSettings /S /Q
		rmdir Unity5.x /S /Q
		pushd Assets
			rmdir Editor /S /Q
			del Editor.meta /Q
			rmdir Lua /S /Q
			del Lua.meta /Q
			pushd Source
				del LuaConst.* /Q
			popd
			pushd ToLua
				rmdir Examples /S /Q
				del Examples.meta /Q
				pushd Editor
					del ToLuaExport.* /Q
					del ToLuaMenu.* /Q
				popd
				pushd Core
					del ToLua.cs /Q
					del ToLua.cs.meta /Q
				popd
			popd
		popd
	popd
)

if exist u3d\Assets\ABOut (
	rmdir u3d\Assets\ABOut /S /Q
)
rem GOBIN=E:\github\bsn\gopath\bin
rem GOPATH=E:\github\bsn\gopath
rem GOROOT=D:\Software\Go17
rem 
rem go get github.com/astaxie/beego
pause