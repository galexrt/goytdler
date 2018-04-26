package main

import (
	"context"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/coreos/pkg/capnslog"
	"github.com/coreos/pkg/flagutil"
	"github.com/galexrt/goytdler/data"
	"github.com/galexrt/goytdler/pkg/options"
	"github.com/galexrt/goytdler/pkg/routes"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/common/version"
)

var (
	logger        = capnslog.NewPackageLogger("github.com/galexrt/goytdler", "main")
	goytdlerFlags = flag.NewFlagSet("goytdler", flag.ExitOnError)
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [OPTION]...\n", os.Args[0])
	goytdlerFlags.PrintDefaults()
	os.Exit(0)
}

func setLogLevel() {
	// parse given log level string then set up corresponding global logging level
	ll, err := capnslog.ParseLevel(options.Opts.LogLevel)
	if err != nil {
		logger.Warningf("failed to set log level %s. %+v", options.Opts.LogLevel, err)
	}
	capnslog.SetGlobalLogLevel(ll)
}

// loadTemplate loads templates embedded by go-assets-builder
func loadTemplate() (*template.Template, error) {
	t := template.New("")
	for name, file := range data.Assets.Files {
		if file.IsDir() || !strings.HasSuffix(name, ".tmpl") {
			continue
		}
		h, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, err
		}
		t, err = t.New(name).Parse(string(h))
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

func init() {
	goytdlerFlags.StringVar(&options.Opts.LogLevel, "log-level", "INFO", "log level")
	goytdlerFlags.StringVar(&options.Opts.ListenAddress, "listen-address", ":1454", "IP:PORT leave IP empty for wildcard listen")
	goytdlerFlags.StringVar(&options.Opts.YoutubeDLPath, "youtube-dl-path", "/usr/bin/youtube-dl", "Path to the youtube-dl executable")
	goytdlerFlags.StringVar(&options.Opts.OutputPath, "output-path", "./goytdler-output", "Output path for youtube-dl downloads")

	goytdlerFlags.Parse(os.Args[1:])
}

func main() {
	flagutil.SetFlagsFromEnv(goytdlerFlags, "GOYTDLER")

	setLogLevel()

	logger.Infof("starting goytdler %s %s", version.Info(), version.BuildContext())

	r := gin.New()

	t, err := loadTemplate()
	if err != nil {
		panic(err)
	}
	r.SetHTMLTemplate(t)

	r.GET("/", routes.Index)
	r.POST("/download", routes.Download)

	srv := &http.Server{
		Addr:    options.Opts.ListenAddress,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logger.Error(err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logger.Info("shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	srv.Shutdown(ctx)
	cancel()

	logger.Infof("exiting")
}
