package v1

import "strings"

// RemoveRepository removes repository from a list by index
//func RemoveRepository(s []Repository, i int) []Repository {
//	s[i] = s[len(s)-1]
//	return s[:len(s)-1]
//}

// RemoveApplication removes applications from a list by index
func RemoveApplication(s []Application, i int) []Application {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func GetUsernameAndRepoNameFromGithubRepositoryUrl(url string) (username string, repoName string) {
	trim := strings.TrimSuffix(url, ".git")
	urlArray := strings.Split(trim, "/")
	if len(urlArray) < 3 {
		return "", ""
	}
	repositoryName := urlArray[len(urlArray)-1]
	usernameOrorgName := urlArray[len(urlArray)-2]
	return usernameOrorgName, repositoryName
}

func GetUsernameAndRepoNameFromBitbucketRepositoryUrl(url string) (username string, repoName string) {
	trim := strings.TrimSuffix(url, ".git")
	urlArray := strings.Split(trim, "/")
	if len(urlArray) < 3 {
		return "", ""
	}
	repositoryName := urlArray[len(urlArray)-4]
	usernameOrorgName := urlArray[len(urlArray)-5]
	return usernameOrorgName, repositoryName
}
