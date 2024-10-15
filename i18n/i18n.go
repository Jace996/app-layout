package i18n

import (
	"embed"
	"github.com/jace996/multiapp/pkg/localize"
)

var (
	//go:embed  embed
	f embed.FS
)

func init() {
	localize.RegisterFileBundle(localize.FileBundle{Fs: f})
}
