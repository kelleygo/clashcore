//go:build !(android && cmfa)

package process

import "github.com/kelleygo/clashcore/constant"

func FindPackageName(metadata *constant.Metadata) (string, error) {
	return "", nil
}
