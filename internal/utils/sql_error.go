package utils

import (
	"fmt"

	"github.com/lib/pq"
)

func MapSQLError(err error) error {
	if err == nil {
		return nil
	}

	if pqErr, ok := err.(*pq.Error); ok {
		switch pqErr.Code {
		case "23505":
			message := fmt.Sprintf(" This %s is already in use.", pqErr.Constraint)
			return fmt.Errorf("%s", message)
		case "23503":
			return fmt.Errorf("The related resource was not found.")
		case "23502":
			return fmt.Errorf("A required field is missing.")
		default:
			return fmt.Errorf("An unexpected error occurred: %s", pqErr.Message)
		}
	}

	return fmt.Errorf("An unexpected error occurred: %v", err)
}
