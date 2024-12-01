package models

import (
	"fmt"
	"regexp"
	"strings"
)

// ValidationError represents a validation error
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

// ValidationErrors represents multiple validation errors
type ValidationErrors []ValidationError

func (e ValidationErrors) Error() string {
	if len(e) == 0 {
		return ""
	}

	var msgs []string
	for _, err := range e {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// Validator interface for input validation
type Validator interface {
	Validate() error
}

var slugRegex = regexp.MustCompile(`^[-a-zA-Z0-9_]+$`)

// ValidateSlug validates a slug string
func ValidateSlug(slug string) error {
	if slug == "" {
		return &ValidationError{
			Field:   "slug",
			Message: "cannot be empty",
		}
	}
	if !slugRegex.MatchString(slug) {
		return &ValidationError{
			Field:   "slug",
			Message: "must contain only alphanumeric characters, hyphens, and underscores",
		}
	}
	return nil
}

// ValidateRequired validates that a string is not empty
func ValidateRequired(field, value string) error {
	if value == "" {
		return &ValidationError{
			Field:   field,
			Message: "cannot be empty",
		}
	}
	return nil
}
