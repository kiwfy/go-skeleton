package kernel

import (
	"fmt"
	"go-skeleton/pkg"
	"time"

	"github.com/spf13/cobra"
)

type Kernel struct {
	rootCmd *cobra.Command
}

func NewKernel() *Kernel {
	k := &Kernel{}
	k.rootCmd = &cobra.Command{
		Use:   "go-skeleton",
		Short: "",
		Long:  ``,
		Run:   k.RootCmd,
	}
	return k
}

func (k *Kernel) Start() error {
	return k.rootCmd.Execute()
}

func (k *Kernel) Boot() {
	for index, dep := range pkg.KernelDependencies {
		dep.Boot()
		pkg.Logger.Info(fmt.Sprintf("[Kernel.Kernel] Booting %s", index))
	}

	pkg.Logger.Info("[Kernel.Kernel] Booting application!")
	time.Local = time.FixedZone("America/Sao_Paulo", 0)
}

func (k *Kernel) RootCmd(_ *cobra.Command, _ []string) {
	pkg.Logger.Info(fmt.Sprintf("Go Skeleton Version %v", pkg.Config.ReadConfig("VERSION")))
	pkg.Logger.Info("Use --help to check witch commands are available")
}
