package flags

import (
	"fmt"
	"github.com/pesakit/pesakit/config"
	"github.com/pesakit/pesakit/mno"
	"github.com/spf13/cobra"
	"strings"
)

const (
	mnoName    = "mno"
	mnoUsage   = "mobile network operator"
	mnoDefault = ""
)

func SetMnoFlag(cmd *cobra.Command, p Type) {

	defaultMnoConf := config.DefaultMnoConf()
	switch p {
	case LOCAL:
		cmd.Flags().StringVar(&defaultMnoConf.Value, mnoName, mnoDefault, mnoUsage)
	case PERSISTENT:
		cmd.PersistentFlags().StringVar(&defaultMnoConf.Value, mnoName, mnoDefault, mnoUsage)
	}
}

func LoadMnoConfig(cmd *cobra.Command, p Type) (mno.Mno, error) {
	var (
		mnoValue mno.Mno
		err      error
		value    string
	)

	switch p {
	case LOCAL:
		value, err = cmd.Flags().GetString(mnoName)
		if err != nil {
			return "", err
		}
		if strings.TrimSpace(value) == "" {
			return "", fmt.Errorf("mno is required [tigo,airtel,vodacom]")
		}

	case PERSISTENT:
		value, err = cmd.PersistentFlags().GetString(mnoName)
		if err != nil {
			return "", err
		}
		if strings.TrimSpace(value) == "" {
			return "", fmt.Errorf("mno is required [tigo,airtel,vodacom]")
		}
	}

	mnoValue = mno.Mno(value)
	if mnoValue == mno.Unknown {
		return "", fmt.Errorf("mno is required [tigo,airtel,vodacom]")
	}

	return mnoValue, nil
}
