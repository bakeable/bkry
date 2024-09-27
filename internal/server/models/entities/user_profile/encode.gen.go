package user_profile

// Encode converts a UserProfile struct to a map
func Encode(e UserProfile) map[string]interface{} {
    result := map[string]interface{}{
        "id": e.ID,
        "created": e.Created.Encode(),
        "modified": e.Modified.Encode(),
        "accessLevel": EncodeAccessLevel(e.AccessLevel),
        "address": e.Address,
        "avatar": e.Avatar,
        "avatarUrl": e.AvatarUrl,
        "email": e.Email,
        "firstName": e.FirstName,
        "infix": e.Infix,
        "isAdmin": e.IsAdmin,
        "_kind": e.Kind,
        "language": e.Language,
        "lastName": e.LastName,
        "prefix": e.Prefix,
    }
    return result
}
// EncodeAccessLevel converts a AccessLevel struct to a map
func EncodeAccessLevel(e AccessLevel) map[string]interface{} {
    return map[string]interface{}{
        "admin": e.Admin,
        "airlines": e.Airlines,
        "apiDocs": e.ApiDocs,
        "packingSlips": e.PackingSlips,
        "priceConfiguration": e.PriceConfiguration,
        "printingOrders": e.PrintingOrders,
        "quotation": e.Quotation,
        "sampleTest": e.SampleTest,
        "userManagement": e.UserManagement,
    }
}




