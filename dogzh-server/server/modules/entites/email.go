package entites

import (
	"fmt"
	"regexp"
	"strings"
)

type Email string

var emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

func (e Email) IsValid() bool {
	if len(e) < 3 || len(e) > 254 {
		return false
	}
	return emailRegex.MatchString(strings.ToLower(string(e)))
}

func (e Email) IsFromDomain(domain string) bool {
	suffix := "@" + strings.ToLower(domain)
	return strings.HasSuffix(strings.ToLower(string(e)), suffix)
}

func (e Email) GetDomain() string {
	parts := strings.Split(string(e), "@")
	if len(parts) != 2 {
		return ""
	}
	return parts[1]
}

func NewRootUser(code string) Email {
	return Email(fmt.Sprintf("admin_%s@dogzh.com", code))
}
