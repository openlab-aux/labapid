package main

func tokenOk(token string, acl map[string]string) bool {

	for _, v := range acl {
		if v == token {
			return true
		}
	}
	return false
}
