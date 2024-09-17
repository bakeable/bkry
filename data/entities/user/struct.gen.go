package user
import (
	"github.com/bakeable/bkry/data/entities/media"
   "github.com/bakeable/bkry/data/general/auditinfo"
)

type User struct {
	ID string `json:"id"`
	Modified auditinfo.AuditInfo `json:"modified"`
	Created auditinfo.AuditInfo `json:"created"`
	// The address of the user
	Address Address `json:"address"`
	// A user's avatar
	Avatar media.Media `json:"avatar"`
	// The actual e-mail of the user
	Email string `json:"email"`
	// The preferred locale settings of the user
	Locale Locale `json:"locale"`
	// Indicates whether the sound is muted for the user
	MuteSound bool `json:"mute_sound"`
	// The profile information of the user
	Profile Profile `json:"profile"`
	// The user's settings
	Settings string `json:"settings"`
	// A nick-name for the user to use in games
	UserName string `json:"user_name"`
}



type Address struct {
	// The city of the address
	City string `json:"city"`
	// The country code of the address
	CountryCode string `json:"country_code"`
	// The first line of the address
	Line1 string `json:"line1"`
	// The second line of the address
	Line2 string `json:"line2"`
	// The postal code of the address
	PostalCode string `json:"postal_code"`
	// The state of the address
	State string `json:"state"`
	// The street of the address
	Street string `json:"street"`
}


type Locale struct {
	// The preferred locale extended with accent
	Code string `json:"code"`
	// Fallback on this locale if the preferred accent is not available
	TwoDigitCode string `json:"two_digit_code"`
}


type Profile struct {
	// The user's date of birth
	DateOfBirth string `json:"date_of_birth"`
	// The user's first name
	FirstName string `json:"first_name"`
	// The user's gender
	Gender Gender `json:"gender"`
	// The user's infix
	Infix string `json:"infix"`
	// Indicates whether the user is an adult
	IsGrownUp bool `json:"is_grown_up"`
	// The user's last name
	LastName string `json:"last_name"`
	// The user's prefix
	Prefix string `json:"prefix"`
}

type Gender string






