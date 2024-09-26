package user

import (
	"github.com/bakeable/bkry/tools"
	"github.com/bakeable/bkry/data/entities/media"
	"github.com/bakeable/bkry/data/general/auditinfo"
)

// Decode converts a map to User struct
func Decode(m interface{}) User {
	if m, ok := m.(map[string]interface{}); ok {
		return User{
			ID: utils.TryDecodeString(m["id"]),
			Created: auditinfo.Decode(m["created"]),
			Modified: auditinfo.Decode(m["modified"]),
			Address: DecodeAddress(m["address"]),
			Avatar: media.Decode(m["avatar"]),
			Email: utils.DecodeString(m["email"], ""),
			Locale: DecodeLocale(m["locale"]),
			MuteSound: utils.DecodeBool(m["mute_sound"], false),
			Profile: DecodeProfile(m["profile"]),
			Settings: utils.DecodeString(m["settings"], ""),
			UserName: utils.DecodeString(m["user_name"], ""),
		}
	}

	return User{}
}

// DecodeAll converts a slice of maps to a slice of User struct
func DecodeAll(ms interface{}) []User {
	var entities []User

	if arr, ok := ms.([]map[string]interface{}); ok {
		for _, m := range arr {
			entities = append(entities, Decode(m))
		}
		return entities
	}

	if arr, ok := ms.([]interface{}); ok {
		for _, m := range arr {
			entities = append(entities, Decode(m))
		}
		return entities
	}

	return entities
}

// DecodeAddress converts a map to Address struct
func DecodeAddress(m interface{}) Address {
	if m == nil {
		return Address{}
	}
	if val, ok := m.(map[string]interface{}); ok {
		return Address{
			City: utils.DecodeString(val["city"], ""),
			CountryCode: utils.DecodeString(val["country_code"], ""),
			Line1: utils.DecodeString(val["line1"], ""),
			Line2: utils.DecodeString(val["line2"], ""),
			PostalCode: utils.DecodeString(val["postal_code"], ""),
			State: utils.DecodeString(val["state"], ""),
			Street: utils.DecodeString(val["street"], ""),
		}
	}
	return Address{}
}


// DecodeLocale converts a map to Locale struct
func DecodeLocale(m interface{}) Locale {
	if m == nil {
		return Locale{}
	}
	if val, ok := m.(map[string]interface{}); ok {
		return Locale{
			Code: utils.DecodeString(val["code"], "nl-NL"),
			TwoDigitCode: utils.DecodeString(val["two_digit_code"], "nl"),
		}
	}
	return Locale{}
}


// DecodeProfile converts a map to Profile struct
func DecodeProfile(m interface{}) Profile {
	if m == nil {
		return Profile{}
	}
	if val, ok := m.(map[string]interface{}); ok {
		return Profile{
			DateOfBirth: utils.DecodeString(val["date_of_birth"], ""),
			FirstName: utils.DecodeString(val["first_name"], ""),
			Gender: Gender(utils.DecodeString(val["gender"], "Unknown")),
			Infix: utils.DecodeString(val["infix"], ""),
			IsGrownUp: utils.DecodeBool(val["is_grown_up"], false),
			LastName: utils.DecodeString(val["last_name"], ""),
			Prefix: utils.DecodeString(val["prefix"], ""),
		}
	}
	return Profile{}
}


