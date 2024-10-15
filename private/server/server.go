package server

import (
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/goava/di"
	api2 "github.com/jace996/app-layout/api"
	"github.com/jace996/app-layout/private/biz"
	"github.com/jace996/app-layout/private/data"
	"github.com/jace996/multiapp/pkg/api"
	"github.com/jace996/multiapp/pkg/authz/authz"
	kitdi "github.com/jace996/multiapp/pkg/di"
	"github.com/jace996/multiapp/pkg/server"
	"github.com/jace996/saas"
	"github.com/jace996/saas/seed"
	"github.com/jace996/uow"
)

// ProviderSet is server providers.
var ProviderSet = kitdi.NewSet(
	NewAuthorizationOption,
	kitdi.NewProvider(NewHTTPServer, di.As(new(transport.Server))),
	kitdi.NewProvider(NewGRPCServer, di.As(new(transport.Server))),
	NewSeeder,
	NewSeeding,
	kitdi.Value(ClientName))

var ClientName api.ClientName = api2.ServiceName

func NewSeeding(uow uow.Manager, migrate *data.Migrate, post *biz.PostSeeder) seed.Contrib {
	return seed.Chain(server.NewUowContrib(uow, seed.Chain(migrate, post)))
}

func NewSeeder(ts saas.TenantStore, seeds []seed.Contrib) seed.Seeder {
	res := seed.NewDefaultSeeder(server.NewTraceContrib(server.SeedChangeTenant(ts, seeds...)))
	return res
}

func NewAuthorizationOption() *authz.Option {
	return authz.NewAuthorizationOption()
}
