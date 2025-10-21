@echo off
:: ---------------------------------------------------------------------------
:: Call goreleaser from a container to facilitate cross-compiling
:: ---------------------------------------------------------------------------
podman image exists mycustom/goreleaser-cross:latest
if errorlevel 1 (
  echo Image not found, building...
  podman build -t mycustom/goreleaser-cross:latest -f Dockerfile.builder .
) else (
  echo Image already exists, skipping build.
)

podman run --rm ^
  -e CGO_ENABLED=1 ^
  -e GITHUB_TOKEN=%GITHUB_TOKEN% ^
  -v %cd%:/src/go/password-generator ^
  -w /src/go/password-generator ^
  mycustom/goreleaser-cross %*
