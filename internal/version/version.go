package version

import (
	"context"
	"fmt"
	"runtime"
)

var (
	version  = "dev"
	template = "oidc-cli version %s (%s/%s)"
)

func String() string {
	return fmt.Sprintf(template, version, runtime.GOOS, runtime.GOARCH)
}

func CheckForUpdate(ctx context.Context) (bool, error) {
	return false, nil
}
