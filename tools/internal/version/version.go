package version

import "runtime/debug"

func Version() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "UNKNOWN"
	}

	return info.Main.Version
}
