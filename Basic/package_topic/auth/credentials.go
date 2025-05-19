package auth

func LoginWithCredentials(username, password string) (bool, string) {
	// Simulate a login check
	if username == "admin" && password == "1234" {
		// in same directory we can get access of private function
		// but in different directory we can't access private function
		// so we need to use public function
		return true, sessionLogic()
	}
	return false, ""
}