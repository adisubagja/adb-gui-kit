#!/usr/bin/env bash

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
# shellcheck source=./common.sh
source "$SCRIPT_DIR/common.sh"

main() {
  require_cmd go
  require_cmd pnpm
  require_cmd wails

  require_file "$ROOT_DIR/bin/linux/adb"
  require_file "$ROOT_DIR/bin/linux/fastboot"

  local version
  version="${ADBKIT_VERSION:-$(cd "$ROOT_DIR" && project_version)}"
  log "Linux builder ready for version ${version}"

  ensure_frontend_dependencies
  build_frontend_assets

  log "Building Linux binary"
  (cd "$ROOT_DIR" && wails build -clean -upx)

  log "Linux base build complete"
}

main "$@"
