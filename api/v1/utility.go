package v1

import (
	v1 "github.com/klovercloud-ci-cd/integration-manager/core/v1"
	"github.com/klovercloud-ci-cd/integration-manager/enums"
	"log"
	"strings"
	"time"
)

func UrlFormatter(url string) string {
	if strings.Contains(url, "/src/") {
		user, repo := getUsernameAndRepoNameFromBitbucketRepositoryUrl(url)
		url = "https://bitbucket.org/" + user + "/" + repo
	} else if strings.Contains(url, "@bitbucket.org") {
		user, repo := getUsernameAndRepoNameFromGithubRepositoryUrl(url)
		url = "https://bitbucket.org/" + user + "/" + repo
	} else {
		user, repo := getUsernameAndRepoNameFromGithubRepositoryUrl(url)
		url = "https://github.com/" + user + "/" + repo
	}
	return url

}

// BranchExists returns boolean for branch existence
func BranchExists(allowedBranches []string, resourceRef string, gitType enums.REPOSITORY_TYPE) bool {
	if gitType == enums.GITHUB {
		branch := strings.Split(resourceRef, "/")[2]
		for _, each := range allowedBranches {
			if branch == each {
				return true
			}
		}
		log.Println("[Forbidden]: Branch wasn't matched!")
		return false
	} else if gitType == enums.BIT_BUCKET {
		for _, each := range allowedBranches {
			if resourceRef == each {
				return true
			}
		}
		log.Println("[Forbidden]: Branch wasn't matched!")
		return false
	}
	return false
}

func getUsernameAndRepoNameFromGithubRepositoryUrl(url string) (username string, repoName string) {
	trim := strings.TrimSuffix(url, ".git")
	urlArray := strings.Split(trim, "/")
	if len(urlArray) < 3 {
		return "", ""
	}
	repositoryName := urlArray[len(urlArray)-1]
	usernameOrorgName := urlArray[len(urlArray)-2]
	return usernameOrorgName, repositoryName
}

func getUsernameAndRepoNameFromBitbucketRepositoryUrl(url string) (username string, repoName string) {
	trim := strings.TrimSuffix(url, ".git")
	urlArray := strings.Split(trim, "/")
	if len(urlArray) < 3 {
		return "", ""
	}
	repositoryName := urlArray[len(urlArray)-1]
	usernameOrorgName := urlArray[len(urlArray)-2]
	return usernameOrorgName, repositoryName
}

func IsAppSecretValid(app v1.Application, secret string) bool {
	if secret != app.Secret {
		return false
	} else if app.SecretValidUntil.Before(time.Now().UTC()) {
		return false
	} else {
		return true
	}
}
