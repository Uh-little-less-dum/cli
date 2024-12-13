package git_manager

import (
	"io"
	"time"

	build_constants "github.com/Uh-little-less-dum/build/pkg/buildConstants"
	"github.com/charmbracelet/log"
	"github.com/go-git/go-git/v5"
)

type GitManager struct {
	Url        build_constants.BuildConstant
	SparsePath build_constants.BuildConstant
	Directory  string
	Timeout    time.Duration
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (g GitManager) SparseClone(targetDir string, outputManager io.Writer) {
	r, err := git.PlainClone(g.Directory, false, &git.CloneOptions{
		URL:        string(g.Url),
		NoCheckout: true,
		Progress:   outputManager,
	})
	if err == git.ErrRepositoryAlreadyExists {
		// PRIORITY: Move to a rebuild here once the build is in a working order!
		log.Fatal("repo was already cloned")
	}
	checkError(err)

	w, err := r.Worktree()
	checkError(err)

	err = w.Checkout(&git.CheckoutOptions{
		SparseCheckoutDirectories: []string{string(g.SparsePath)},
	})
	checkError(err)
}

// func (g GitManager) setHttpClient() {
// 	customClient := &http.Client{
// 		Transport: &http.Transport{
// 			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
// 		},
// 		Timeout: g.Timeout * time.Second,
// 		CheckRedirect: func(req *http.Request, via []*http.Request) error {
// 			return http.ErrUseLastResponse
// 		},
// 	}
// 	client.InstallProtocol("https", githttp.NewClient(customClient))
// }

func NewTemplateAppGitManager(targetDirectory string, timeout time.Duration) GitManager {
	return GitManager{
		Url:        build_constants.SparseCloneRepoUrl,
		SparsePath: build_constants.SparseCloneSparsePath,
		Directory:  targetDirectory,
		Timeout:    timeout,
	}
}
