package microfiber

import "regexp"

type Config struct {
	AuthKeyLookup string
	ProtectedURLs []*regexp.Regexp
	Port          string
	Cache         bool
	Limitter      bool
	Logger        bool
	Metrics       bool
}

func (c *Config) SetProtectedUrls(paths []string) {

	for _, path := range paths {
		regx := regexp.MustCompile("^" + path + "$")
		c.ProtectedURLs = append(c.ProtectedURLs, regx)
	}
}
