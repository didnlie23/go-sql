package entity

type Avatar struct {
	ID       int64
	Nickname string
	Profile  *string
}

func (a *Avatar) GetProfile() string {
	if a.Profile == nil {
		return ""
	}
	return *a.Profile
}
