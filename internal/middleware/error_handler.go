package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Error   string            `json:"error"`
	Details map[string]string `json:"details,omitempty"`
}

func ErrorFormatterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Check if there are any errors
		errs := c.Errors
		if len(errs) > 0 {
			// Default error response
			resp := ErrorResponse{
				Error: errs[0].Err.Error(),
			}

			// Check for validation errors
			if ve, ok := errs[0].Err.(validator.ValidationErrors); ok {
				details := make(map[string]string)
				for _, fe := range ve {
					details[fe.Field()] = validationErrorMessage(fe)
				}
				resp.Details = details
			}

			c.JSON(http.StatusBadRequest, resp)
			c.Abort()
			return
		}
	}
}

// Optional: Custom message per validation tag
func validationErrorMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "gt":
		return "Must be greater than " + fe.Param()
	case "oneof":
		return "Must be one of: " + fe.Param()
	case "datetime":
		return "Invalid date format, use YYYY-MM-DD"
	default:
		return fe.Error()
	}
}
