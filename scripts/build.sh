script_dir=$(dirname "$0")
cd "$script_dir"

cd ../

go build -o ./build/videomt ./cmd/videomt/main.go
