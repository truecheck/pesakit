package flags

import (
	"github.com/pesakit/pesakit/config"
	"github.com/spf13/cobra"
)

const (
	mnoName    = "mno"
	mnoUsage   = "mobile network operator"
	mnoDefault = ""
)

func SetMnoFlag(cmd *cobra.Command, p Type, mno *config.Mno) {
	if mno == nil {
		mno = config.DefaultMnoConf()
	}
	switch p {
	case LOCAL:
		cmd.Flags().StringVar(&mno.Value, mnoName, mnoDefault, mnoUsage)
	case PERSISTENT:
		cmd.PersistentFlags().StringVar(&mno.Value, mnoName, mnoDefault, mnoUsage)
	}
}
