package version

import (
	"fmt"
	"strings"

	_ "embed"
)

//go:generate sh version.sh
//go:embed version.txt
var Version string

//func Short() string {
//	return fmt.Sprintf("atools %s", Version)
//}

func PrintWithCopyright() {
	_, _ = fmt.Println(longWithCopyright())
}

func longWithCopyright() string {
	return fmt.Sprintf("atools, version %s\nCopyright (C) 2023 Alessio Treglia <a@lessio.dev>", strings.TrimSpace(Version))
}
