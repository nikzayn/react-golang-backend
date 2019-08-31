package user

//If the username and password is valid or not
func UserIsValid(uName, pwd string) bool {
	_uName, _pwd, _isValid := "nikzayn", "nikhil123", false

	if uName == _uName && pwd == _pwd {
		_isValid = true
	} else {
		_isValid = false
	}
	return _isValid
}
