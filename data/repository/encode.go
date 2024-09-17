package repository

import (
	"fmt"

	utils "github.com/bakeable/bkry/tools"
)

func tryEncode(foreignKeyMap *EntityForeignKeyMap) map[string]interface{} {
	if utils.IsTrulyNil(foreignKeyMap) {
		return nil
	}
	// Catch the panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic")
		}
	}()

	// Try to encode
	encodedForeignKeyMap := (*foreignKeyMap).Encode()

	// Return the encoded foreign key map
	return encodedForeignKeyMap
}