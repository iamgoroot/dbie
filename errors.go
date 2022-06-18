package dbie

import (
	"database/sql"
	"errors"
	"fmt"
)

type (
	IntegrityConstraintViolationError string
	InvalidTransactionStateError      string
	InvalidCursorStateError           string
	CardinalityViolationError         string
	TransactionRollbackError          string
	Error                             string
	errWithFields                     interface {
		error
		Field(byte) string
	}
)

var ErrNoRows = Error(sql.ErrNoRows.Error())

func (e Error) Error() string {
	return fmt.Sprintf("dbie error: %s", string(e))
}
func (e IntegrityConstraintViolationError) Error() string {
	return fmt.Sprintf("integrity constraint error: %s", string(e))
}
func (e InvalidTransactionStateError) Error() string {
	return fmt.Sprintf("invalid transaction state error: %s", string(e))
}
func (e InvalidCursorStateError) Error() string {
	return fmt.Sprintf("invalid cursor state error: %s", string(e))
}
func (e CardinalityViolationError) Error() string {
	return fmt.Sprintf("cardinality violation error: %s", string(e))
}
func (e TransactionRollbackError) Error() string {
	return fmt.Sprintf("integrity constraint error: %s", string(e))
}
func Wrap(err error) error {
	if err == nil {
		return err
	}
	switch typedErr := err.(type) {
	case errWithFields:
		return getPgDescr(typedErr)
	}
	if errors.Is(err, sql.ErrNoRows) {
		return ErrNoRows
	}
	return fmt.Errorf("dbie error: %w", err)
}

func getPgDescr(err errWithFields) error {
	switch err.Field('C') {
	case "23000":
		return IntegrityConstraintViolationError("integrity_constraint_violation")
	case "23001":
		return IntegrityConstraintViolationError("restrict_violation")
	case "23502":
		return IntegrityConstraintViolationError("not_null_violation")
	case "23503":
		return IntegrityConstraintViolationError("foreign_key_violation")
	case "23505":
		return IntegrityConstraintViolationError("unique_violation")
	case "23514":
		return IntegrityConstraintViolationError("check_violation")
	case "23P01":
		return IntegrityConstraintViolationError("exclusion_violation")
	case "24000":
		return InvalidCursorStateError("invalid_cursor_state")
	case "25000":
		return InvalidTransactionStateError("invalid_transaction_state")
	case "25001":
		return InvalidTransactionStateError("active_sql_transaction")
	case "25002":
		return InvalidTransactionStateError("branch_transaction_already_active")
	case "25008":
		return InvalidTransactionStateError("held_cursor_requires_same_isolation_level")
	case "25003":
		return InvalidTransactionStateError("inappropriate_access_mode_for_branch_transaction")
	case "25004":
		return InvalidTransactionStateError("inappropriate_isolation_level_for_branch_transaction")
	case "25005":
		return InvalidTransactionStateError("no_active_sql_transaction_for_branch_transaction")
	case "25006":
		return InvalidTransactionStateError("read_only_sql_transaction")
	case "25007":
		return InvalidTransactionStateError("schema_and_data_statement_mixing_not_supported")
	case "25P01":
		return InvalidTransactionStateError("no_active_sql_transaction")
	case "25P02":
		return InvalidTransactionStateError("in_failed_sql_transaction")
	case "25P03":
		return InvalidTransactionStateError("idle_in_transaction_session_timeout")
	case "0B000":
		return InvalidTransactionStateError("invalid_transaction_initiation")
	case "21000":
		return CardinalityViolationError("cardinality_violation")
	case "40000":
		return TransactionRollbackError("transaction_rollback")
	case "40002":
		return TransactionRollbackError("transaction_integrity_constraint_violation")
	case "40001":
		return TransactionRollbackError("serialization_failure")
	case "40003":
		return TransactionRollbackError("statement_completion_unknown")
	case "40P01":
		return TransactionRollbackError("deadlock_detected")
	}
	return err
}
