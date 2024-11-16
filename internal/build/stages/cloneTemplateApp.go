package clone_template_app_manager

import (
	git_manager "github.com/igloo1505/ulldCli/internal/build/gitManager"
)

func CloneTemplateApp(targetDir string) {
	gm := git_manager.NewTemplateAppGitManager(targetDir, 30)
	gm.SparseClone(targetDir)
}
