$ErrorActionPreference = "Stop"

function Write-BuildLog {
    param([string]$Message)
    Write-Host "[build] $Message"
}

$RootDir = (Resolve-Path (Join-Path $PSScriptRoot "../..")).Path

$AdbPath = Join-Path $RootDir "bin/windows/adb.exe"
$FastbootPath = Join-Path $RootDir "bin/windows/fastboot.exe"

if (-not (Test-Path $AdbPath)) {
    throw "Required file not found: $AdbPath"
}

if (-not (Test-Path $FastbootPath)) {
    throw "Required file not found: $FastbootPath"
}

$WailsJsonPath = Join-Path $RootDir "wails.json"
$WailsConfig = Get-Content -Raw $WailsJsonPath | ConvertFrom-Json
$Version = if ($env:ADBKIT_VERSION) { $env:ADBKIT_VERSION } else { $WailsConfig.info.productVersion }

Write-BuildLog "Windows builder ready for version $Version"

Write-BuildLog "Installing frontend dependencies"
pnpm install --frozen-lockfile --dir "$RootDir/frontend"

Write-BuildLog "Building frontend assets"
pnpm --dir "$RootDir/frontend" build

Write-BuildLog "Building NSIS installer"
Push-Location $RootDir
wails build -clean -upx -nsis
Pop-Location

Write-BuildLog "Windows base build complete"
