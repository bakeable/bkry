package utils

import (
	"math/rand"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func Contains(slice []string, item string) bool {
	for _, value := range slice {
		if value == item {
			return true
		}
	}
	return false
}

func NumberToFloat64(m interface{}) float64 {
    var n float64
    if v, ok := m.(float64); ok {
        n = v
    } else if v, ok := m.(int64); ok {
        n = float64(v)
    }
    return n
}

func NumberToInt64(m interface{}) int64 {
    var n int64
    if v, ok := m.(float64); ok {
        n = int64(v)
    } else if v, ok := m.(int64); ok {
        n = v
    }
    return n
}
func NumberToInt(m interface{}) int {
	var n int
	if v, ok := m.(float64); ok {
		n = int(v)
	} else if v, ok := m.(int64); ok {
		n = int(v)
	} else if v, ok := m.(int); ok {
		n = v
	}
	return n
}

func GetEditorId(c *gin.Context) *string {
	if c == nil {
		editorId := "system"
		return &editorId
	}

	editorId, ok := c.Get("userId")
	if !ok {
		editorId = "system"
	}

	// Convert editorId to string
	editorIdStr, ok := editorId.(string)
	if !ok {
		editorIdStr = "system"
	}

	return &editorIdStr
}

func MapMaterial(material string) string {
	switch material {
		case "pet":
			return "pet"
		case "rpet":
			return "pet"
		case "cardboard":
			return "cardboard"
		default:
			return "other"
	}
}


func GenerateVerificationCode() string {
	rand.Seed(time.Now().UnixNano())

	charSet := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	code := make([]byte, 6)
	for i := 0; i < 6; i++ {
		code[i] = charSet[rand.Intn(len(charSet))]
	}

	return string(code)
}

// Check if an interface{} is truly nil, even if its type is not.
func IsTrulyNil(i interface{}) bool {
    if i == nil {
        return true
    }
    v := reflect.ValueOf(i)
    return v.Kind() == reflect.Ptr && v.IsNil()
}


func GetCountryCode(country string) string {
	switch strings.ToLower(country) {
	case "engeland":
		return "en"
	case "nederland":
		return "nl"
	case "belgië":
		return "be"
	case "belgie":
		return "be"
	default:
		return "nl"
	}
}

func ExtractURLFromErrorMessage(errMsg string) string {
	re := regexp.MustCompile(`https://console\.firebase\.google\.com[^ ]+`)
	match := re.FindString(errMsg)
	return match
}