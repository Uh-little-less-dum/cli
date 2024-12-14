package stage_clone_template_app

import (
	"io"

	git_manager "github.com/Uh-little-less-dum/cli/internal/build/gitManager"
)

func Run(targetDir string, outputManager io.Writer) {
	gm := git_manager.NewTemplateAppGitManager(targetDir, 30)
	gm.SparseClone(targetDir, outputManager)
}
