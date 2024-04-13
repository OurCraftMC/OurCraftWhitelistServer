package Config

type ConfigStruct struct {
	WhitelistFilePath string `yaml:"whitelist_file_path"`
	HTMLPath          string `yaml:"html_path"`
	ListenAddr        string `yaml:"listen_addr"`
	SecretKey         string `yaml:"secret_key"`
}
