package base

import (
	"github.com/vitesales/cyrus-pool/h"
	"github.com/vitesales/cyrus/models"
	"github.com/vitesales/cyrus/models/security"
	"github.com/vitesales/cyrus/server"
	"github.com/vitesales/cyrus/tools/logging"
)

const (
	// MODULE_NAME is the name of this module
	MODULE_NAME = "base"
)

var log logging.Logger

func init() {
	log = logging.GetLogger("base")
	server.RegisterModule(&server.Module{
		Name: MODULE_NAME,
		PostInit: func() {
			err := models.ExecuteInNewEnvironment(security.SuperUserID, func(env models.Environment) {
				h.Group().NewSet(env).ReloadGroups()
			})
			if err != nil {
				log.Panic("Error while initializing", "error", err)
			}
		},
	})
}
