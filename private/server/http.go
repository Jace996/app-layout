package server

import (
	_ "embed"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	khttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/jace996/app-layout/api"
	api2 "github.com/jace996/multiapp/pkg/api"
	sapi "github.com/jace996/multiapp/pkg/api"
	"github.com/jace996/multiapp/pkg/authn/jwt"
	conf2 "github.com/jace996/multiapp/pkg/conf"
	"github.com/jace996/multiapp/pkg/localize"
	"github.com/jace996/multiapp/pkg/server"
	"github.com/jace996/multiapp/pkg/server/common"
	kithttp "github.com/jace996/multiapp/pkg/server/http"
	"github.com/jace996/saas"
	shttp "github.com/jace996/saas/http"
	uow2 "github.com/jace996/uow"
	kuow "github.com/jace996/uow/kratos"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf2.Services,
	sCfg *conf2.Security,
	tokenizer jwt.Tokenizer,
	ts saas.TenantStore,
	uowMgr uow2.Manager,
	reqDecoder khttp.DecodeRequestFunc,
	resEncoder khttp.EncodeResponseFunc,
	errEncoder khttp.EncodeErrorFunc,
	mOpt *shttp.WebMultiTenancyOption,
	apiOpt *api2.Option,
	validator sapi.TrustedContextValidator,
	register []kithttp.ServiceRegister,
	logger log.Logger) *khttp.Server {
	var opts []khttp.ServerOption
	cfg := common.GetConf(c, api.ServiceName)
	opts = kithttp.PatchOpts(logger, opts, cfg, sCfg, reqDecoder, resEncoder, errEncoder)
	m := []middleware.Middleware{
		recovery.Recovery(),
		tracing.Server(),
		logging.Server(logger),
		metrics.Server(),
		localize.I18N(),
		validate.Validator(),
		jwt.ServerExtractAndAuth(tokenizer, logger),
		sapi.ServerPropagation(apiOpt, validator, logger),
		server.Saas(mOpt, ts, validator),
		kuow.Uow(uowMgr),
	}
	opts = append(opts, []khttp.ServerOption{
		khttp.Middleware(
			m...,
		),
	}...)
	srv := khttp.NewServer(opts...)
	kithttp.ChainServiceRegister(register...).Register(srv, m...)
	return srv
}
