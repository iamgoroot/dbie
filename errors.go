package dbie

import (
	"database/sql"
	"fmt"
	pgErr "github.com/uptrace/bun/driver/pgdriver"
)

type (
	ErrIntegrityConstraintViolation Err
	ErrInvalidTransactionState      Err
	ErrInvalidCursorState           Err
	ErrCardinalityViolation         Err
	ErrTransactionRollback          Err
	ErrNoRows                       error
	Err                             struct {
		error
		desc string
	}
)

var NoRows = ErrNoRows(sql.ErrNoRows)

func (e Err) Error() string {
	return fmt.Sprintf(e.error.Error(), e.desc)
}

func Wrap(err error) error {
	switch typedErr := err.(type) {
	case pgErr.Error:
		return getPgDescr(typedErr)
	}
	switch err {
	case sql.ErrNoRows:
		return NoRows
	}
	return err
}

func getPgDescr(err pgErr.Error) error {
	switch err.Field('C') {
	case "23000":
		return ErrIntegrityConstraintViolation{err, "integrity_constraint_violation"}
	case "23001":
		return ErrIntegrityConstraintViolation{err, "restrict_violation"}
	case "23502":
		return ErrIntegrityConstraintViolation{err, "not_null_violation"}
	case "23503":
		return ErrIntegrityConstraintViolation{err, "foreign_key_violation"}
	case "23505":
		return ErrIntegrityConstraintViolation{err, "unique_violation"}
	case "23514":
		return ErrIntegrityConstraintViolation{err, "check_violation"}
	case "23P01":
		return ErrIntegrityConstraintViolation{err, "exclusion_violation"}
	case "24000":
		return ErrInvalidCursorState{err, "invalid_cursor_state"}
	case "25000":
		return ErrInvalidTransactionState{err, "invalid_transaction_state"}
	case "25001":
		return ErrInvalidTransactionState{err, "active_sql_transaction"}
	case "25002":
		return ErrInvalidTransactionState{err, "branch_transaction_already_active"}
	case "25008":
		return ErrInvalidTransactionState{err, "held_cursor_requires_same_isolation_level"}
	case "25003":
		return ErrInvalidTransactionState{err, "inappropriate_access_mode_for_branch_transaction"}
	case "25004":
		return ErrInvalidTransactionState{err, "inappropriate_isolation_level_for_branch_transaction"}
	case "25005":
		return ErrInvalidTransactionState{err, "no_active_sql_transaction_for_branch_transaction"}
	case "25006":
		return ErrInvalidTransactionState{err, "read_only_sql_transaction"}
	case "25007":
		return ErrInvalidTransactionState{err, "schema_and_data_statement_mixing_not_supported"}
	case "25P01":
		return ErrInvalidTransactionState{err, "no_active_sql_transaction"}
	case "25P02":
		return ErrInvalidTransactionState{err, "in_failed_sql_transaction"}
	case "25P03":
		return ErrInvalidTransactionState{err, "idle_in_transaction_session_timeout"}
	case "0B000":
		return ErrInvalidTransactionState{err, "invalid_transaction_initiation"}
	case "21000":
		return ErrCardinalityViolation{err, "cardinality_violation"}
	case "40000":
		return ErrTransactionRollback{err, "transaction_rollback"}
	case "40002":
		return ErrTransactionRollback{err, "transaction_integrity_constraint_violation"}
	case "40001":
		return ErrTransactionRollback{err, "serialization_failure"}
	case "40003":
		return ErrTransactionRollback{err, "statement_completion_unknown"}
	case "40P01":
		return ErrTransactionRollback{err, "deadlock_detected"}
	}
	return err
}
