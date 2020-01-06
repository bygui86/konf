package split

import (
	"bygui86/kubeconfigurator/commons"
	"bygui86/kubeconfigurator/kubeconfig"
	"bygui86/kubeconfigurator/utils"

	"github.com/urfave/cli"
)

func BuildCommand() *cli.Command {
	home := utils.GetHomeDirOrExit("split")
	return &cli.Command{
		Name:   "split",
		Usage:  "Split kube-config into multiple single Kubernetes configurations based on the context",
		Action: split,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:     utils.GetUrfaveFlagName(commons.CustomKubeConfigFlagName, commons.CustomKubeConfigFlagShort),
				Usage:    commons.CustomKubeConfigFlagDescription,
				EnvVar:   commons.CustomKubeConfigPathEnvVar,
				Value:    kubeconfig.GetCustomKubeConfigPathDefault(home),
				Required: false,
			},
			cli.StringFlag{
				Name:     utils.GetUrfaveFlagName(commons.SingleConfigsFlagName, commons.SingleConfigsFlagShort),
				Usage:    commons.SingleConfigsFlagDescription,
				EnvVar:   commons.SingleConfigsPathEnvVar,
				Value:    kubeconfig.GetSingleConfigsPathDefault(home),
				Required: false,
			},
		},
	}
}
