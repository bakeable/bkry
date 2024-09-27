package middleware

import (
	"github.com/bakeable/bkry/fail"

	"github.com/gin-gonic/gin"
)

func RecoveryWithWriter() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				// Check if r is error interface
				if e, ok := r.(fail.IError); ok {
					// r implements IError interface
					// Handle the error accordingly
					c.JSON(e.GetResponseCode(), gin.H{
						"error": e.GetError(),
						"info": e.GetData(),
					})
				} else if err, ok := r.(error); ok {
                    if err.Error()[5:] == "fail:" {
                        // Convert error to IError
                        e, err := fail.FromError(err)
                        if err != nil {
                            panic(err)
                        }
                        
                        // r implements IError interface
                        // Handle the error accordingly
                        c.JSON(e.GetResponseCode(), gin.H{
                            "error": e.GetError(),
                            "info": e.GetData(),
                        })
                    } else {
                        panic(err)
                    }
				} else {
					// r does not implement IError interface
					// Handle the error as a panic
					panic(r)
				}
			}
		}()
		c.Next()
	}
}