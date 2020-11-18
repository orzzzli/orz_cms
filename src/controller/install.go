package controller

import (
	"fmt"
	"net/http"

	"github.com/orzzzli/orz_cms/src/core"
)

func InstallHandler(w http.ResponseWriter, r *http.Request) {
	tableComplete := core.CheckTableComplete(core.GlobalDB)
	fmt.Fprintf(w, core.FormatOutput(0, "success", map[string]interface{}{
		"table": core.ConvertBoolToInt(tableComplete),
	}))
}
