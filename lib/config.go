package mail2most

import (
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	General        General
	Logging        Logging
	Profiles       []Profile `toml:"Profile"`
	DefaultProfile Profile
	NoStateFile    bool
}
type General struct {
	File         string
	TimeInterval uint
	NoLoop       bool
}

type Logging struct {
	Loglevel string
	Logtype  string
	Logfile  string
	Output   string
}

type Profile struct {
	IgnoreDefaults bool
	Mail           maildata
	Mattermost     mattermost
	Filter         filter
}

type maildata struct {
	ImapServer, Username, Password string
	ReadOnly                       bool
	ImapTLS                        bool
	VerifyTLS                      bool
	Limit                          uint32
}

type filter struct {
	Folders, From, To, Subject []string
	Unseen                     bool
	TimeRange                  string
}

type mattermost struct {
	URL, Team, Username, Password, AccessToken string
	Channels                                   []string
	Users                                      []string
	Broadcast                                  []string
	SubjectOnly                                bool
	StripHTML                                  bool
	ConvertToMarkdown                          bool
	HideFrom                                   bool
	HideFromEmail                              bool
	HideSubject                                bool
	MailAttachments                            bool
}

func parseConfig(fileName string, conf *Config) error {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return err
	}
	_, err := toml.DecodeFile(fileName, conf)
	return err
}
