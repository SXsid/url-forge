package migrations

import "embed"

//go:embed *.sql
var MigratonFs embed.FS
