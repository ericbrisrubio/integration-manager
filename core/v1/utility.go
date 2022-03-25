package v1

// RemoveRepository removes repository from a list by index
func RemoveRepository(s []Repository, i int) []Repository {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

// RemoveApplication removes applications from a list by index
func RemoveApplication(s []Application, i int) []Application {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}