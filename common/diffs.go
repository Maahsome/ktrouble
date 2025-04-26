package common

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"ktrouble/internal"

	"github.com/pmezard/go-difflib/difflib"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func OutputContextDiff(local []byte, remote []byte, contextLines int) {

	text := ""

	localReader := bytes.NewReader(local)
	localBuffer := new(bytes.Buffer)
	remoteReader := bytes.NewReader(remote)
	remoteBuffer := new(bytes.Buffer)

	serr := internal.SortYAML(localReader, localBuffer, 2)
	if serr != nil {
		Logger.WithError(serr).Error("Error sorting original yaml")
	}
	serr = internal.SortYAML(remoteReader, remoteBuffer, 2)
	if serr != nil {
		Logger.WithError(serr).Error("Error sorting edit yaml")
	}

	sortedLocal := localBuffer.Bytes()
	sortedRemote := remoteBuffer.Bytes()

	switch getDiffStyle() {
	case "unified":
		diff := difflib.UnifiedDiff{
			A:        difflib.SplitLines(string(sortedRemote)),
			B:        difflib.SplitLines(string(sortedLocal)),
			FromFile: "Upstream Repository",
			ToFile:   "Local Repository",
			Context:  contextLines,
		}
		text, _ = difflib.GetUnifiedDiffString(diff)
	case "context":
		diff := difflib.ContextDiff{
			A:        difflib.SplitLines(string(sortedRemote)),
			B:        difflib.SplitLines(string(sortedLocal)),
			FromFile: "Upstream Repository",
			ToFile:   "Local Repository",
			Context:  contextLines,
		}
		text, _ = difflib.GetContextDiffString(diff)

	}

	customPager := getPager()
	if len(customPager) > 0 {
		cmd := exec.Command(customPager)
		// cmd := exec.Command("/usr/local/bin/delta")

		// Feed it with the string you want to display.
		cmd.Stdin = strings.NewReader(text)

		// This is crucial - otherwise it will write to a null device.
		cmd.Stdout = os.Stdout

		// Fork off a process and wait for it to terminate.
		pageerr := cmd.Run()
		if pageerr != nil {
			logrus.WithError(pageerr).Error("Error calling to pager")
		}
	} else {
		fmt.Println(text)
	}
}

func getPager() string {

	customPager := os.Getenv("KTROUBLE_PAGER")
	if len(customPager) > 0 {
		return customPager
	}
	customPager = viper.GetString("Pager")
	if len(customPager) > 0 {
		return customPager
	}
	customPager = os.Getenv("PAGER")
	if len(customPager) > 0 {
		return customPager
	}
	return ""
}

func getDiffStyle() string {
	diffStyle := viper.GetString("diffStyle")
	if len(diffStyle) > 0 {
		return diffStyle
	}
	return "unified"
}
