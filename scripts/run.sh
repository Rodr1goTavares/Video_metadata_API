script_dir=$(dirname "$0")
cd "$script_dir"

cd ../

go run ./cmd/videomt/main.go
