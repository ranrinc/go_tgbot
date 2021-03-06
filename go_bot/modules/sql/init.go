package sql

import (
	"github.com/atechnohazard/ginko/go_bot"
	"github.com/atechnohazard/ginko/go_bot/modules/utils/error_handling"
	"github.com/go-pg/pg"
)

var SESSION *pg.DB

func init() {
	opt, err := pg.ParseURL(go_bot.BotConfig.SqlUri)
	error_handling.HandleErrorAndExit(err)

	SESSION = pg.Connect(opt)
}
