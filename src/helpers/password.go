package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

func GetPwdHash(pwd string, salt string) (string,error) {
	password := []byte(pwd + salt)
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hash),nil
}

func ComparePassword(inputPwd string,userPwd string, salt string) error {
	pwd,_ := GetPwdHash(inputPwd,salt)
	return bcrypt.CompareHashAndPassword([]byte(pwd),[]byte(userPwd))
}
