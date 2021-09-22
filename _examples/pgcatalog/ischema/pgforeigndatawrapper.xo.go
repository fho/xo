package ischema

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"

	"github.com/fho/xo/_examples/pgcatalog/pgtypes"
)

// PgForeignDataWrapper represents a row from 'information_schema._pg_foreign_data_wrappers'.
type PgForeignDataWrapper struct {
	Oid                        pgtypes.NullOid  `json:"oid"`                           // oid
	Fdwowner                   pgtypes.NullOid  `json:"fdwowner"`                      // fdwowner
	Fdwoptions                 []sql.NullString `json:"fdwoptions"`                    // fdwoptions
	ForeignDataWrapperCatalog  sql.NullString   `json:"foreign_data_wrapper_catalog"`  // foreign_data_wrapper_catalog
	ForeignDataWrapperName     sql.NullString   `json:"foreign_data_wrapper_name"`     // foreign_data_wrapper_name
	AuthorizationIdentifier    sql.NullString   `json:"authorization_identifier"`      // authorization_identifier
	ForeignDataWrapperLanguage sql.NullString   `json:"foreign_data_wrapper_language"` // foreign_data_wrapper_language
}
