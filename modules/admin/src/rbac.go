package admin

import (
	."github.com/casbin/casbin"
	_"github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/casbin/gorm-adapter"
	"../../../framework"
)

var db = gormadapter.NewAdapter("sqlite3", xf.GetDataDir("database/rbac.db"), true)
var Rbac = NewEnforcer(xf.GetDataDir("casbin/rbac_model.conf"), db, false)