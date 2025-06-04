package mobile

import (
	v2 "github.com/hiddify/hiddify-core/v2"
	_ "github.com/sagernet/gomobile"
)

func Setup(baseDir, workingDir, tempDir string, debug bool) error {
	return v2.Setup(baseDir, workingDir, tempDir, 0, debug)
	// return v2.Start(17078)
}
