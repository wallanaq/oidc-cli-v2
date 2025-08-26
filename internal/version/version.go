package version

import (
	"fmt"
	"runtime"
)

var (
	version  = "v0.0.0"
	template = "oidc-cli version %s (%s/%s)"
)

func Current() string {
	return version
}

func String() string {
	return fmt.Sprintf(template, version, runtime.GOOS, runtime.GOARCH)
}
