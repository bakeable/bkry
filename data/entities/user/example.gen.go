package user

var ExampleJSON = `{
    address: {
      city: "", // The city of the address 
      country_code: "", // The country code of the address 
      line1: "", // The first line of the address 
      line2: "", // The second line of the address 
      postal_code: "", // The postal code of the address 
      state: "", // The state of the address 
      street: "", // The street of the address 
    }, // The address of the user
    email: string, // The actual e-mail of the user
    locale: {
      code: "nl-NL", // The preferred locale extended with accent 
      two_digit_code: "nl", // Fallback on this locale if the preferred accent is not available 
    }, // The preferred locale settings of the user
    mute_sound: boolean, // Indicates whether the sound is muted for the user
    profile: {
      date_of_birth: "", // The user's date of birth 
      first_name: "", // The user's first name 
      gender: GenderValues.Unknown, // The user's gender (an enum with possible values: [Unknown Male Female])
      infix: "", // The user's infix 
      is_grown_up: false, // Indicates whether the user is an adult 
      last_name: "", // The user's last name 
      prefix: "", // The user's prefix 
    }, // The profile information of the user
    settings: string, // The user's settings
    user_name: string, // A nick-name for the user to use in games
}`