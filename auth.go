package main

func tokenOk(token string, acl map[string]string) bool {
	found := false
	for _, value := range acl {
		if value == token {
			found = true
		}
	}
	return found
}
