$Repo = "TheLIama33/cforge"
$BinaryName = "cforge.exe"

$Arch = if ($env:PROCESSOR_ARCHITECTURE -eq "AMD64") { "x86_64" } else { "arm64" }
$ZipName = "cforge_Windows_$Arch.zip"
$DownloadUrl = "https://github.com/$Repo/releases/latest/download/$ZipName"

$InstallDir = "$env:LOCALAPPDATA\cforge"
$BinPath = "$InstallDir\$BinaryName"

Write-Host "Downloading cforge for Windows..." -ForegroundColor Cyan

if (!(Test-Path $InstallDir)) { New-Item -ItemType Directory -Force -Path $InstallDir | Out-Null }

$ZipPath = "$InstallDir\temp.zip"
Invoke-WebRequest -Uri $DownloadUrl -OutFile $ZipPath
Expand-Archive -Path $ZipPath -DestinationPath $InstallDir -Force
Remove-Item $ZipPath

$CurrentPath = [Environment]::GetEnvironmentVariable("Path", "User")
if ($CurrentPath -notlike "*$InstallDir*") {
    [Environment]::SetEnvironmentVariable("Path", "$CurrentPath;$InstallDir", "User")
    Write-Host "Added $InstallDir to User PATH." -ForegroundColor Green
}

Write-Host "Installation complete!" -ForegroundColor Green
Write-Host "Please restart your terminal (PowerShell/CMD) to use the command." -ForegroundColor Yellow