package controller

import (
	"fmt"
	"net/http"

	error2 "github.com/orzzzli/orz_cms/src/error"

	"github.com/orzzzli/orz_cms/src/core"
)

func InstallHandler(w http.ResponseWriter, r *http.Request) {
	installed := core.IsInstalled(core.GlobalDB)
	//已安装
	if installed {
		fmt.Fprintf(w, core.FormatOutput(0, "success", map[string]interface{}{
			"installed": core.ConvertBoolToInt(installed),
		}))
		return
	}
	//进行安装
	tableComplete := core.CheckTableComplete(core.GlobalDB)
	if !tableComplete {
		res := core.InstallTable(core.GlobalDB)
		if !res {
			fmt.Fprintf(w, core.FormatOutput(error2.InstallDefaultTableError, "error", nil))
			return
		}
	}
	res := core.InstallDefaultData(core.GlobalDB)
	if !res {
		fmt.Fprintf(w, core.FormatOutput(error2.InstallDefaultDataError, "error", nil))
		return
	}
	res = core.ConfirmInstall(core.GlobalDB)
	if !res {
		fmt.Fprintf(w, core.FormatOutput(error2.ConfirmInstallError, "error", nil))
		return
	}
	fmt.Fprintf(w, core.FormatOutput(0, "success", nil))
}

func InstallInfoHandler(w http.ResponseWriter, r *http.Request) {
	installed := core.IsInstalled(core.GlobalDB)
	tableComplete := core.CheckTableComplete(core.GlobalDB)
	fmt.Fprintf(w, core.FormatOutput(0, "success", map[string]interface{}{
		"installed": core.ConvertBoolToInt(installed),
		"table":     core.ConvertBoolToInt(tableComplete),
	}))
}

func ResetHandler(w http.ResponseWriter, r *http.Request) {
	tableComplete := core.CheckTableComplete(core.GlobalDB)
	fmt.Fprintf(w, core.FormatOutput(0, "success", map[string]interface{}{
		"table": core.ConvertBoolToInt(tableComplete),
	}))
}
