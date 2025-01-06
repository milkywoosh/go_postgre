package utils

import "golang.org/x/crypto/bcrypt"

// reference: https://medium.com/@jcox250/password-hash-salt-using-golang-b041dc94cb72
func HashPasswordUser(password string) (string, error) {
	hashed_pass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return string(""), nil
	}
	return string(hashed_pass), nil
}

func DecryptPasswordUser(hashed_password string, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashed_password), []byte(password))
	if err != nil {
		return false, err
	}
	// if error nil
	return true, nil
}
