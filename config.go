package microfiber

type Config struct {
	AuthKeyLookup string
	Port          string
	Cache         bool
	Limitter      bool
	Logger        bool
	Metrics       bool
}
