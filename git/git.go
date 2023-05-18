package git

import (
	"errors"
	"fmt"
	"log"
	"path"

	at "github.com/deepsourcelabs/artifacts/types"
	git "github.com/libgit2/git2go/v34"
)

// GitOperator represents the Git operations.
type GitOperator interface {
	Commit() error
	Push() error
	ConfigureAuthor() error
	ConfigureRemote() error
	CheckoutToNewBranch() error
}

// GitClient is an implementation of the GitOperator interface that uses the libgit2 library.
type GitClient struct {
	Repository *git.Repository
	// CommitData *at.PatchCommit
}

// NewGitClient initializes a GitClient with a repository and commit data.
func NewGitClient(repositoryPath string, commitData *at.PatchCommit) (*GitClient, error) {
	repo, err := git.OpenRepository(repositoryPath)
	if err != nil {
		return nil, err
	}
	return &GitClient{
		Repository: repo,
		// CommitData: commitData,
	}, nil
}

// ConfigureAuthor configures the user name and email for the local Git repository.
func (g *GitClient) ConfigureAuthor() error {
	config, err := g.Repository.Config()
	if err != nil {
		return err
	}

	err = config.SetString("user.name", g.CommitData.Author.Name)
	if err != nil {
		return err
	}

	err = config.SetString("user.email", g.CommitData.Author.Email)
	return err
}

// Commit adds and commits the changes with the given message.
func (g *GitClient) Commit() error {
	index, err := g.Repository.Index()
	if err != nil {
		return err
	}

	err = index.AddAll([]string{}, git.IndexAddDefault, nil)
	if err != nil {
		return err
	}

	treeID, err := index.WriteTree()
	if err != nil {
		return err
	}

	tree, err := g.Repository.LookupTree(treeID)
	if err != nil {
		return err
	}

	headRef, err := g.Repository.Head()
	if err != nil {
		return err
	}

	headCommit, err := g.Repository.LookupCommit(headRef.Target())
	if err != nil {
		return err
	}

	signature, err := g.Repository.DefaultSignature()
	if err != nil {
		return err
	}

	// The refname parameter in the git.Repository.CreateCommit() function is optional and can
	// be an empty string. It is typically used when creating new references or updating existing references.
	// In the context of the Commit() function, if we want to create a new commit without associating
	// it with a specific reference, we can pass an empty string as the refname parameter.
	_, err = g.Repository.CreateCommit("", signature, signature, g.CommitData.Commit.Title, tree, headCommit)
	return err
}

// ConfigureRemote adds a new remote to the local Git repository.
func (g *GitClient) ConfigureRemote() error {
	_, err := g.Repository.Remotes.Create(g.CommitData.Remote.Name, g.CommitData.Remote.URL)
	return err
}

// CheckoutToNewBranch creates and checks out a new branch in the local Git repository.
func (g *GitClient) CheckoutToNewBranch() (err error) {
	branchRef, err := g.Repository.Head()
	if err != nil {
		return
	}

	commit, err := g.Repository.LookupCommit(branchRef.Target())
	if err != nil {
		return
	}

	branchName := "refs/heads/" + g.CommitData.Commit.DestinationBranch
	_, err = g.Repository.LookupBranch(branchName, git.BranchLocal)
	if err != nil {
		// Branch doesn't exist, create it
		_, err = g.Repository.CreateBranch(g.CommitData.Commit.DestinationBranch, commit, false)
		if err != nil {
			return
		}
	}
	return
}

// Push pushes the local branch to the remote repository.
func (g *GitClient) Push() error {
	remote, err := g.Repository.Remotes.Lookup(g.CommitData.Remote.Name)
	if err != nil {
		log.Println(err)
		return err
	}

	// Get the current branch reference
	branchRef, err := g.Repository.Head()
	if err != nil {
		log.Println(err)
		return err
	}

	// Check if the HEAD reference is pointing to a branch
	if branchRef.Type() != git.ReferenceSymbolic {
		err := errors.New("HEAD reference is not pointing to a branch")
		log.Println(err)
		return err
	}

	// Get the current branch name
	branchName, err := branchRef.Branch().Name()
	if err != nil {
		log.Println(err)
	}
	branchShortName := path.Base(branchName)

	// Push the branch to the remote repository
	remoteBranchName := fmt.Sprintf("refs/heads/%s", branchShortName)
	log.Println(remoteBranchName)
	pushSpec := fmt.Sprintf("%s:%s", branchRef.Target().String(), remoteBranchName)

	pushOptions := &git.PushOptions{
		RemoteCallbacks: git.RemoteCallbacks{
			CredentialsCallback: func(url string, username string, allowedTypes git.CredentialType) (*git.Credential, error) {
				if allowedTypes&(git.CredentialTypeUserpassPlaintext|git.CredentialTypeDefault) != 0 {
					cred, err := git.NewCredentialUserpassPlaintext("auth-token", "x-oauth-basic")
					if err != nil {
						return nil, err
					}
					return cred, nil
				}
				return nil, nil
			},
		},
	}

	err = remote.Push([]string{pushSpec}, pushOptions)
	if err != nil {
		return err
	}

	return nil
}
