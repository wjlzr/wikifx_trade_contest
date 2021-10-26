package bcryptx

import "golang.org/x/crypto/bcrypt"

// 加密密码
func HashAndSalt(pwdStr string) (pwdHash string) {
	pwd := []byte(pwdStr)
	hash, _ := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	pwdHash = string(hash)
	return
}

// 验证密码
func ComparePasswords(hashedPwd string, plainPwd string) bool {
	byteHash := []byte(hashedPwd)
	bytePwd := []byte(plainPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePwd)
	if err != nil {
		return false
	}
	return true
}
