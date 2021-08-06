// Package handlers provides gin handlers for all business logic.
//
// A handler should be provided in this form:
// ```
// func XxxHandler(dep Dependency) gin.HandlerFunc {
//     return func(c *gin.Context) {
//         /* biz logic */
//     }
// }
// ```
// All dependent interfaces should be injected as arguments to factory functions.
package handlers
