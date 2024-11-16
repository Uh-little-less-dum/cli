package git_manager

import (
	"time"

	"github.com/charmbracelet/log"
	"github.com/go-git/go-git/v5"
	"github.com/igloo1505/ulldCli/internal/build/constants"
)

type GitManager struct {
	Url        string
	SparsePath string
	Directory  string
	Timeout    time.Duration
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (g GitManager) SparseClone(targetDir string) {
	r, err := git.PlainClone(g.Directory, false, &git.CloneOptions{
		URL:        g.Url,
		NoCheckout: true,
		// Progress:   progressManager,
	})
	if err == git.ErrRepositoryAlreadyExists {
		// PRIORITY: Move to a rebuild here once the build is in a working order!
		log.Fatal("repo was already cloned")
	}
	checkError(err)

	w, err := r.Worktree()
	checkError(err)

	err = w.Checkout(&git.CheckoutOptions{
		SparseCheckoutDirectories: []string{g.SparsePath},
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
		Url:        constants.SparseCloneRepoUrl,
		SparsePath: constants.SparseCloneSparsePath,
		Directory:  targetDirectory,
		Timeout:    timeout,
	}
}
