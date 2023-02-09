package global

type App struct {
	Template      string
	PageSize      int
	JwtSecret     string
	JwtExpiresAt  string
	SigningMethod string
	PidPath       string
}
