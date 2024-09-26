package user_profile

import (
	"github.com/bakeable/bkry/tools"
	"github.com/bakeable/bkry/data/entities/media"
	"github.com/bakeable/bkry/data/general/auditinfo"
)

// Decode converts a map to UserProfile struct
func Decode(m interface{}) UserProfile {
	if m, ok := m.(map[string]interface{}); ok {
		return UserProfile{
			ID: utils.TryDecodeString(m["id"]),
			Created: auditinfo.Decode(m["created"]),
			Modified: auditinfo.Decode(m["modified"]),
			AccessLevel: DecodeAccessLevel(m["accessLevel"]),
			Address: utils.DecodeMap(m["address"], nil),
			Avatar: media.Decode(m["avatar"]),
			AvatarUrl: utils.DecodeString(m["avatarUrl"], ""),
			Email: utils.DecodeString(m["email"], ""),
			FirstName: utils.DecodeString(m["firstName"], ""),
			Infix: utils.DecodeString(m["infix"], ""),
			IsAdmin: utils.DecodeBool(m["isAdmin"], false),
			Kind: utils.DecodeString(m["_kind"], "UserProfile"),
			Language: Language(utils.DecodeString(m["language"], "nl")),
			LastName: utils.DecodeString(m["lastName"], ""),
			Prefix: utils.DecodeString(m["prefix"], ""),
		}
	}

	return UserProfile{}
}

// DecodeAll converts a slice of maps to a slice of UserProfile struct
func DecodeAll(ms interface{}) []UserProfile {
	var entities []UserProfile

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

// DecodeAccessLevel converts a map to AccessLevel struct
func DecodeAccessLevel(m interface{}) AccessLevel {
	if m == nil {
		return AccessLevel{}
	}
	if val, ok := m.(map[string]interface{}); ok {
		return AccessLevel{
			Admin: Permission(utils.DecodeString(val["admin"], "none")),
			Airlines: Permission(utils.DecodeString(val["airlines"], "none")),
			ApiDocs: Permission(utils.DecodeString(val["apiDocs"], "none")),
			PackingSlips: Permission(utils.DecodeString(val["packingSlips"], "none")),
			PriceConfiguration: Permission(utils.DecodeString(val["priceConfiguration"], "none")),
			PrintingOrders: Permission(utils.DecodeString(val["printingOrders"], "none")),
			Quotation: Permission(utils.DecodeString(val["quotation"], "none")),
			SampleTest: Permission(utils.DecodeString(val["sampleTest"], "none")),
			UserManagement: Permission(utils.DecodeString(val["userManagement"], "none")),
		}
	}
	return AccessLevel{}
}


