package main

import (
	"pocketbase/config"
	"pocketbase/flag"
	"pocketbase/superuser"

	"github.com/common-library/go/log/klog"
	"github.com/pocketbase/pocketbase"
)

func init() {
	superuser.Create("admin@admin.com", "admin123")
}

func main() {
	defer klog.Flush()

	app := pocketbase.New()

	if err := flag.Parse(app); err != nil {
		klog.ErrorS(err, "")
		return
	} else if err := config.Read(flag.Get().ConfigFile); err != nil {
		klog.ErrorS(err, "")
		return
	} else {
		klog.SetWithCallerInfo(config.Get().Log.WithCallerInfo)
		klog.Info("config", "", config.Get())
	}

	if err := app.Start(); err != nil {
		klog.ErrorS(err, "")
	}
}
