package config

//GetDbCreds returns what you need to sign into mysql
func GetDbCreds() (string, string, string) {
	return "zane", "5245", "blog"
}
