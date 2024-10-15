package api

import (
	_ "embed"

	"github.com/jace996/multiapp/pkg/authz/authz"
)

const (
	ResourcePost = "github.com/jace996/app-layout.post"
)

//go:embed permission.yaml
var permission []byte

func init() {
	authz.LoadFromYaml(permission)
}
