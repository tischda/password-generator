@echo off

gcc --version >nul 2>&1
if errorlevel 1 (
	echo Installing WinLibs...
	winget install -e --id BrechtSanders.WinLibs.POSIX.UCRT --silent --accept-source-agreements --accept-package-agreements
	echo.
	echo Adding gcc to PATH...
 	set "PATH=%PATH%;%LOCALAPPDATA%\Microsoft\WinGet\Packages\BrechtSanders.WinLibs.POSIX.UCRT_Microsoft.Winget.Source_8wekyb3d8bbwe\mingw64\bin"
)

gcc --version | findstr "gcc"

REM Mind trailing spaces!
set CGO_ENABLED=1
set GOOS=windows
set GOARCH=amd64

make build
