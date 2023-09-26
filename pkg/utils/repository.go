package utils

import (
	"database/sql"
	"fmt"
)

func RollbackTransaction(tx *sql.Tx, internalError error) error {
	rbErr := tx.Rollback()

	if rbErr != nil && internalError != nil {
		return fmt.Errorf("RB ERROR: %v INTERNAL ERROR: %v", rbErr, internalError)
	} else if rbErr == nil && internalError != nil {
		return fmt.Errorf("INTERNAL ERROR: %v", internalError)
	}

	return nil
}
