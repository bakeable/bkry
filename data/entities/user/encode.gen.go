package user

// Encode converts a User struct to a map
func Encode(e User) map[string]interface{} {
    result := map[string]interface{}{
        "id": e.ID,
        "created": e.Created.Encode(),
        "modified": e.Modified.Encode(),
        "address": EncodeAddress(e.Address),
        "avatar": e.Avatar,
        "email": e.Email,
        "locale": EncodeLocale(e.Locale),
        "mute_sound": e.MuteSound,
        "profile": EncodeProfile(e.Profile),
        "settings": e.Settings,
        "user_name": e.UserName,
    }
    return result
}
// EncodeAddress converts a Address struct to a map
func EncodeAddress(e Address) map[string]interface{} {
    return map[string]interface{}{
        "city": e.City,
        "country_code": e.CountryCode,
        "line1": e.Line1,
        "line2": e.Line2,
        "postal_code": e.PostalCode,
        "state": e.State,
        "street": e.Street,
    }
}

// EncodeLocale converts a Locale struct to a map
func EncodeLocale(e Locale) map[string]interface{} {
    return map[string]interface{}{
        "code": e.Code,
        "two_digit_code": e.TwoDigitCode,
    }
}

// EncodeProfile converts a Profile struct to a map
func EncodeProfile(e Profile) map[string]interface{} {
    return map[string]interface{}{
        "date_of_birth": e.DateOfBirth,
        "first_name": e.FirstName,
        "gender": e.Gender,
        "infix": e.Infix,
        "is_grown_up": e.IsGrownUp,
        "last_name": e.LastName,
        "prefix": e.Prefix,
    }
}




