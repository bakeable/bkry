package user_profile
import (
	"github.com/bakeable/bkry/data/entities/media"
   "github.com/bakeable/bkry/data/general/auditinfo"
)

type UserProfile struct {
	ID string `json:"id"`
	Modified auditinfo.AuditInfo `json:"modified"`
	Created auditinfo.AuditInfo `json:"created"`
	// Which tools the user has access to
	AccessLevel AccessLevel `json:"accessLevel"`
	// A user's address
	Address map[string]interface{} `json:"address"`
	// A user's avatar
	Avatar media.Media `json:"avatar"`
	// The URL of a user's avatar
	AvatarUrl string `json:"avatarUrl"`
	// A user's email address
	Email string `json:"email"`
	// A user's first name
	FirstName string `json:"firstName"`
	// A user's infix
	Infix string `json:"infix"`
	// Whether the user has administrative rights
	IsAdmin bool `json:"isAdmin"`
	// The kind of the object
	Kind string `json:"_kind"`
	// A user's language
	Language Language `json:"language"`
	// A user's last name
	LastName string `json:"lastName"`
	// A user's prefix
	Prefix string `json:"prefix"`
}



type AccessLevel struct {
	// The user's administrative access level
	Admin Permission `json:"admin"`
	// The user's access level to the airlines tool
	Airlines Permission `json:"airlines"`
	// The user's access level to the API documentation tool
	ApiDocs Permission `json:"apiDocs"`
	// The user's access level to the packing slip tool
	PackingSlips Permission `json:"packingSlips"`
	// The user's access level to the price configuration tool
	PriceConfiguration Permission `json:"priceConfiguration"`
	// The user's access level to the printing orders tool
	PrintingOrders Permission `json:"printingOrders"`
	// The user's access level to the quotation tool
	Quotation Permission `json:"quotation"`
	// The user's access level to the sample test tool
	SampleTest Permission `json:"sampleTest"`
	// The user's access level to the user management tool
	UserManagement Permission `json:"userManagement"`
}

type Permission string


type Language string




