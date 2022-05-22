package bootstrap

import (
	"github.com/Jeanhwea/baliqiao2/api/controllers"
	"github.com/Jeanhwea/baliqiao2/api/middlewares"
	"github.com/Jeanhwea/baliqiao2/api/routes"
	"github.com/Jeanhwea/baliqiao2/lib"
	"github.com/Jeanhwea/baliqiao2/repository"
	"github.com/Jeanhwea/baliqiao2/services"
	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	controllers.Module,
	routes.Module,
	lib.Module,
	services.Module,
	middlewares.Module,
	repository.Module,
)
