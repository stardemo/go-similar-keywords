package utils

import (
	"github.com/gogf/gf/container/gvar"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/genv"
	"github.com/gogf/gf/text/gstr"
)

func FindCfg(key, defaultValue string, allowEmpty bool) *gvar.Var {
	cfg := genv.GetVar(gstr.ToUpper(gstr.Replace(key, ".", "_")))
	if cfg.IsNil() || cfg.IsEmpty() {
		cfg = g.Config().GetVar(key)
		if cfg.IsNil() || cfg.IsEmpty() {
			if !allowEmpty {
				g.Log().Panicf("Cfg :%s Cannot be empty!", key)
			}
			cfg.Set(defaultValue)
		}
		return cfg
	}
	return cfg
}
