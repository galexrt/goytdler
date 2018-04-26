package routes

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"strings"

	"github.com/galexrt/goytdler/pkg/options"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "/templates/index.tmpl", nil)
}

func Download(c *gin.Context) {
	paramURL := c.PostForm("url")
	if paramURL == "" {
		c.HTML(http.StatusBadRequest, "/templates/error.tmpl", nil)
		return
	}

	cmdName := options.Opts.YoutubeDLPath
	// TODO use flag values

	cmdArgs := []string{"-x", "--audio-format=mp3", "--format=bestvideo+bestaudio", "--audio-quality=0", paramURL}

	cmd := exec.Command(cmdName, cmdArgs...)
	cmd.Dir = options.Opts.OutputPath
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		c.HTML(http.StatusBadRequest, "/templates/error.tmpl", fmt.Errorf("error creating StdoutPipe for Cmd. %+v", err))
		return
	}

	scanner := bufio.NewScanner(cmdReader)

	cmdStarted := false
	c.Stream(func(w io.Writer) bool {
		if !cmdStarted {
			err = cmd.Start()
			if err != nil {
				w.Write([]byte(fmt.Sprintf("Error starting Cmd %+v", err)))
				return false
			}
		}
		w.Write([]byte(fmt.Sprintf("Running command: %s %s\n\n", cmdName, strings.Join(cmdArgs, " "))))
		for scanner.Scan() {
			w.Write(scanner.Bytes())
			w.Write([]byte("\n"))
		}
		err = cmd.Wait()
		if err != nil {
			w.Write([]byte(fmt.Sprintf("Error waiting for Cmd %+v", err)))
			return false
		}
		return true
	})
}
