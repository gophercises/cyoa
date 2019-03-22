package settings

type Settings struct {
	FilePath string
	ListenPort string
	TemplatePath string
}

func (conf *Settings) GetFilePath() string {
	return conf.FilePath
}

func (conf *Settings) GetListenPort() string {
	return conf.ListenPort
}

func (conf *Settings) GetTemplatePath() string {
	return conf.TemplatePath
}