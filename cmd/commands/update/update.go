package update

import (
	"flag"
	"os"
	"os/exec"

	"github.com/casbin/bee/cmd/commands"
	"github.com/casbin/bee/config"
	beeLogger "github.com/casbin/bee/logger"
)

var CmdUpdate = &commands.Command{
	UsageLine: "update",
	Short:     "Update Bee",
	Long: `
Automatic run command "go get -u github.com/casbin/bee" for selfupdate
`,
	Run: updateBee,
}

func init() {
	fs := flag.NewFlagSet("update", flag.ContinueOnError)
	CmdUpdate.Flag = *fs
	commands.AvailableCommands = append(commands.AvailableCommands, CmdUpdate)
}

func updateBee(cmd *commands.Command, args []string) int {
	beeLogger.Log.Info("Updating")
	beePath := config.GitRemotePath
	cmdUp := exec.Command("go", "get", "-u", beePath)
	cmdUp.Stdout = os.Stdout
	cmdUp.Stderr = os.Stderr
	if err := cmdUp.Run(); err != nil {
		beeLogger.Log.Warnf("Run cmd err:%s", err)
	}
	return 0
}
