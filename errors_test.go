package dbie

import (
	"database/sql"
	"fmt"
	"testing"
)

type pgFakeError map[byte]string

func (p pgFakeError) Error() string {
	return fmt.Sprint(p.Field('C'))
}

func (p pgFakeError) Field(b byte) string {
	return p[b]
}

func TestWrap(t *testing.T) {
	tests := []struct {
		name    string
		argErr  error
		wantErr error
	}{
		{
			name:    "wrap any error",
			argErr:  fmt.Errorf("err1"),
			wantErr: fmt.Errorf("dbie error: %w", fmt.Errorf("err1")),
		},
		{
			name:    "sql.ErrNoRows",
			argErr:  sql.ErrNoRows,
			wantErr: ErrNoRows,
		},
		{
			name:    "integrity_constraint_violation",
			argErr:  pgFakeError{'C': "23000"},
			wantErr: IntegrityConstraintViolationError("integrity_constraint_violation"),
		},
		{
			name:    "restrict_violation",
			argErr:  pgFakeError{'C': "23001"},
			wantErr: IntegrityConstraintViolationError("restrict_violation"),
		},
		{
			name:    "not_null_violation",
			argErr:  pgFakeError{'C': "23502"},
			wantErr: IntegrityConstraintViolationError("not_null_violation"),
		},
		{
			name:    "foreign_key_violation",
			argErr:  pgFakeError{'C': "23503"},
			wantErr: IntegrityConstraintViolationError("foreign_key_violation"),
		},
		{
			name:    "unique_violation",
			argErr:  pgFakeError{'C': "23505"},
			wantErr: IntegrityConstraintViolationError("unique_violation"),
		},
		{
			name:    "check_violation",
			argErr:  pgFakeError{'C': "23514"},
			wantErr: IntegrityConstraintViolationError("check_violation"),
		},
		{
			name:    "exclusion_violation",
			argErr:  pgFakeError{'C': "23P01"},
			wantErr: IntegrityConstraintViolationError("exclusion_violation"),
		},
		{
			name:    "invalid_cursor_state",
			argErr:  pgFakeError{'C': "24000"},
			wantErr: InvalidCursorStateError("invalid_cursor_state"),
		},
		{
			name:    "invalid_transaction_state",
			argErr:  pgFakeError{'C': "25000"},
			wantErr: InvalidTransactionStateError("invalid_transaction_state"),
		},
		{
			name:    "active_sql_transaction",
			argErr:  pgFakeError{'C': "25001"},
			wantErr: InvalidTransactionStateError("active_sql_transaction"),
		},
		{
			name:    "branch_transaction_already_active",
			argErr:  pgFakeError{'C': "25002"},
			wantErr: InvalidTransactionStateError("branch_transaction_already_active"),
		},
		{
			name:    "held_cursor_requires_same_isolation_level",
			argErr:  pgFakeError{'C': "25008"},
			wantErr: InvalidTransactionStateError("held_cursor_requires_same_isolation_level"),
		},
		{
			name:    "inappropriate_access_mode_for_branch_transaction",
			argErr:  pgFakeError{'C': "25003"},
			wantErr: InvalidTransactionStateError("inappropriate_access_mode_for_branch_transaction"),
		},
		{
			name:    "inappropriate_isolation_level_for_branch_transaction",
			argErr:  pgFakeError{'C': "25004"},
			wantErr: InvalidTransactionStateError("inappropriate_isolation_level_for_branch_transaction"),
		},
		{
			name:    "no_active_sql_transaction_for_branch_transaction",
			argErr:  pgFakeError{'C': "25005"},
			wantErr: InvalidTransactionStateError("no_active_sql_transaction_for_branch_transaction"),
		},
		{
			name:    "read_only_sql_transaction",
			argErr:  pgFakeError{'C': "25006"},
			wantErr: InvalidTransactionStateError("read_only_sql_transaction"),
		},
		{
			name:    "schema_and_data_statement_mixing_not_supported",
			argErr:  pgFakeError{'C': "25007"},
			wantErr: InvalidTransactionStateError("schema_and_data_statement_mixing_not_supported"),
		},
		{
			name:    "no_active_sql_transaction",
			argErr:  pgFakeError{'C': "25P01"},
			wantErr: InvalidTransactionStateError("no_active_sql_transaction"),
		},
		{
			name:    "in_failed_sql_transaction",
			argErr:  pgFakeError{'C': "25P02"},
			wantErr: InvalidTransactionStateError("in_failed_sql_transaction"),
		},
		{
			name:    "idle_in_transaction_session_timeout",
			argErr:  pgFakeError{'C': "25P03"},
			wantErr: InvalidTransactionStateError("idle_in_transaction_session_timeout"),
		},
		{
			name:    "invalid_transaction_initiation",
			argErr:  pgFakeError{'C': "0B000"},
			wantErr: InvalidTransactionStateError("invalid_transaction_initiation"),
		},
		{
			name:    "cardinality_violation",
			argErr:  pgFakeError{'C': "21000"},
			wantErr: CardinalityViolationError("cardinality_violation"),
		},
		{
			name:    "transaction_rollback",
			argErr:  pgFakeError{'C': "40000"},
			wantErr: TransactionRollbackError("transaction_rollback"),
		},
		{
			name:    "transaction_integrity_constraint_violation",
			argErr:  pgFakeError{'C': "40002"},
			wantErr: TransactionRollbackError("transaction_integrity_constraint_violation"),
		},
		{
			name:    "serialization_failure",
			argErr:  pgFakeError{'C': "40001"},
			wantErr: TransactionRollbackError("serialization_failure"),
		},
		{
			name:    "statement_completion_unknown",
			argErr:  pgFakeError{'C': "40003"},
			wantErr: TransactionRollbackError("statement_completion_unknown"),
		},
		{
			name:    "deadlock_detected",
			argErr:  pgFakeError{'C': "40P01"},
			wantErr: TransactionRollbackError("deadlock_detected"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Wrap(tt.argErr)
			if tt.wantErr.Error() != err.Error() {
				t.Fatalf("want %v got %v", tt.wantErr, err)
			}
		})
	}
}
