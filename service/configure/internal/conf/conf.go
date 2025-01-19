package conf

type Config struct {
	Author Author
	WebUI  WebUI
}

type WebUI struct {
	Enable bool
	Addr   string
	Dist   string
}

type Author struct {
	AdminUser     string
	AdminPassword string
}
