package middleware

import "log"

func TraceLogSlice() func(c *SliceRouterContext) {
	return func(c *SliceRouterContext) {
		log.Println("TraceLogSlice before...")
		c.Next()
		log.Println("TraceLogSlice after...")
	}
}
