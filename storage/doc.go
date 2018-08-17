package storage

//go:generate go-bindata -pkg storage init.sql
//go:generate xo -x pgsql://postgres:postgres@localhost/postgres?sslmode=disable
//go:generate xo -x pgsql://postgres:postgres@localhost/postgres?sslmode=disable -N -B -T State --single-file --append -o state.xo.go -F LastState -1 -Q "SELECT * FROM state WHERE context_id = %%contextId string%% ORDER BY created_at DESC LIMIT 1"
func init() {}
