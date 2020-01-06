package set

import (
	"github.com/urfave/cli"

	"bygui86/konf/commons"
	"bygui86/konf/kubeconfig"
	"bygui86/konf/logger"
	"bygui86/konf/utils"
)

func BuildCommand() *cli.Command {
	logger.Logger.Debug("🐛 Create set command")
	home := utils.GetHomeDirOrExit("set")
	return &cli.Command{
		Name:  "set",
		Usage: "Set local or global Kubernetes context",
		Subcommands: cli.Commands{
			{
				Name:   "local",
				Usage:  "Set local Kubernetes context (current shell)",
				Action: setLocal,
			},
			{
				Name:   "global",
				Usage:  "Set global Kubernetes context",
				Action: setGlobal,
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:     utils.GetUrfaveFlagName(commons.CustomKubeConfigFlagName, commons.CustomKubeConfigFlagShort),
						Usage:    commons.CustomKubeConfigFlagDescription,
						EnvVar:   commons.CustomKubeConfigPathEnvVar,
						Value:    kubeconfig.GetCustomKubeConfigPathDefault(home),
						Required: false,
					},
				},
			},
		},
	}
}