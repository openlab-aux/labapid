package main

func tokenOk(token string, acl map[string]string) bool {
	_, ok := acl[token]
	return ok
}
