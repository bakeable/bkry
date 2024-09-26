package db

import (
	"log"

	"github.com/bakeable/bkry/fail"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func handleError(err error, message string, path *string, queries *[]string) error {
	switch status.Code(err) {
	case codes.NotFound:
		if path == nil {
			return fail.NewNotFoundError("", nil).ToError()
		} else if queries == nil {
			return fail.NewNotFoundError(*path, nil).ToError()
		} else {
			return fail.NewNotFoundError(*path, *queries).ToError()
		}
	case codes.PermissionDenied:
		return fail.NewAuthenticationError().ToError()
	case codes.Unauthenticated:
		return fail.NewAuthenticationError().ToError()
	default:
		log.Printf("Failed to write ID to document: %v", err)
		return err
	}
}

func collectQueries(queries []Query) *[]string {
	var queryStrings []string
	for _, query := range queries {
		queryStrings = append(queryStrings, query.ToString())
	}
	return &queryStrings
}