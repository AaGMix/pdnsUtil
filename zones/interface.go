package zones

type Control interface {
	Update(ip string, domainName string) (int64, error)
	CClose() error
}
