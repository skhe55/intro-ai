package utils

import (
	"database/sql"
	"fmt"
)

func RollbackTransaction(tx *sql.Tx, fn string, internalError error) error {
	rbErr := tx.Rollback()

	if rbErr != nil && internalError != nil {
		return fmt.Errorf("ERROR OCCURED IN [%s] RB ERROR: %v INTERNAL ERROR: %v", fn, rbErr, internalError)
	} else if rbErr == nil && internalError != nil {
		return fmt.Errorf("ERROR OCCURED IN [%s] INTERNAL ERROR: %v", fn, internalError)
	}

	return nil
}
