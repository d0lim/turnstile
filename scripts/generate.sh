root_dir=$(git rev-parse --show-toplevel)

go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/execquery "$root_dir/ent/schema"
wire "$root_dir/internal/di/wire.go"
# swag init -d "$root_dir/cmd/server,$root_dir/internal/" -o "$root_dir/docs"

echo "[INFO] go generate success"