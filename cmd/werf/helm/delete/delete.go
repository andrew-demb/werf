package delete

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/flant/logboek"
	"github.com/flant/shluz"
	"github.com/flant/werf/cmd/werf/common"
	"github.com/flant/werf/pkg/deploy"
	"github.com/flant/werf/pkg/deploy/helm"
	"github.com/flant/werf/pkg/true_git"
	"github.com/flant/werf/pkg/werf"
	"github.com/spf13/cobra"
)

var CommonCmdData common.CmdData

var CmdData struct {
	helm.DeleteOptions
}

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "delete RELEASE_NAME",
		Short:                 "Delete release from Kubernetes with all resources associated with the last release revision",
		Aliases:               []string{"del", "remove", "rm"},
		DisableFlagsInUseLine: true,
		Annotations: map[string]string{
			common.CmdEnvAnno: common.EnvsDescription(),
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := common.ValidateMinimumNArgs(1, args, cmd); err != nil {
				return err
			}
			return runDelete(args)
		},
	}

	common.SetupTmpDir(&CommonCmdData, cmd)
	common.SetupHomeDir(&CommonCmdData, cmd)

	common.SetupKubeConfig(&CommonCmdData, cmd)
	common.SetupKubeContext(&CommonCmdData, cmd)
	common.SetupHelmReleaseStorageNamespace(&CommonCmdData, cmd)
	common.SetupHelmReleaseStorageType(&CommonCmdData, cmd)

	cmd.Flags().BoolVar(&CmdData.DisableHooks, "no-hooks", false, "Prevent hooks from running during deletion")
	cmd.Flags().BoolVar(&CmdData.Purge, "purge", false, "Remove the release from the store and make its name free for later use")
	cmd.Flags().Int64Var(&CmdData.Timeout, "timeout", 300, "Time in seconds to wait for any individual Kubernetes operation (like Jobs for hooks)")

	return cmd
}

func runDelete(releaseNames []string) error {
	if err := werf.Init(*CommonCmdData.TmpDir, *CommonCmdData.HomeDir); err != nil {
		return fmt.Errorf("initialization error: %s", err)
	}

	if err := shluz.Init(filepath.Join(werf.GetServiceDir(), "locks")); err != nil {
		return err
	}

	if err := true_git.Init(true_git.Options{Out: logboek.GetOutStream(), Err: logboek.GetErrStream()}); err != nil {
		return err
	}

	helmReleaseStorageType, err := common.GetHelmReleaseStorageType(*CommonCmdData.HelmReleaseStorageType)
	if err != nil {
		return err
	}

	deployInitOptions := deploy.InitOptions{
		HelmInitOptions: helm.InitOptions{
			KubeConfig:                  *CommonCmdData.KubeConfig,
			KubeContext:                 *CommonCmdData.KubeContext,
			HelmReleaseStorageNamespace: *CommonCmdData.HelmReleaseStorageNamespace,
			HelmReleaseStorageType:      helmReleaseStorageType,
			ReleasesMaxHistory:          0,
		},
	}
	if err := deploy.Init(deployInitOptions); err != nil {
		return err
	}

	errors := []string{}
	for _, releaseName := range releaseNames {
		if err := helm.Delete(releaseName, CmdData.DeleteOptions); err != nil {
			errors = append(errors, err.Error())
		}
	}

	if len(errors) > 0 {
		return fmt.Errorf("following errors have occured during removal of specified releases: %s", strings.Join(errors, "; "))
	}

	return nil
}
