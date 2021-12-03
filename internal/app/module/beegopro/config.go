package beegopro

import (
	"io/ioutil"

	"github.com/casbin/bee/internal/pkg/utils"
	beeLogger "github.com/casbin/bee/logger"
)
var CompareExcept = []string{"@BeeGenerateTime"}
func (c *Container) GenConfig() {
	if utils.IsExist(c.BeegoProFile) {
		beeLogger.Log.Fatalf("beego pro toml exist")
		return
	}

	err := ioutil.WriteFile("beegopro.toml", []byte(BeegoToml), 0644)
	if err != nil {
		beeLogger.Log.Fatalf("write beego pro toml err: %s", err)
		return
	}
}
