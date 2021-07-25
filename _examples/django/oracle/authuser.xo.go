package oracle

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
	"time"
)

// AuthUser represents a row from 'django.auth_user'.
type AuthUser struct {
	ID          int64          `json:"id"`           // id
	Password    sql.NullString `json:"password"`     // password
	LastLogin   sql.NullTime   `json:"last_login"`   // last_login
	IsSuperuser int64          `json:"is_superuser"` // is_superuser
	Username    sql.NullString `json:"username"`     // username
	FirstName   sql.NullString `json:"first_name"`   // first_name
	LastName    sql.NullString `json:"last_name"`    // last_name
	Email       sql.NullString `json:"email"`        // email
	IsStaff     int64          `json:"is_staff"`     // is_staff
	IsActive    int64          `json:"is_active"`    // is_active
	DateJoined  time.Time      `json:"date_joined"`  // date_joined
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the AuthUser exists in the database.
func (au *AuthUser) Exists() bool {
	return au._exists
}

// Deleted returns true when the AuthUser has been marked for deletion from
// the database.
func (au *AuthUser) Deleted() bool {
	return au._deleted
}

// Insert inserts the AuthUser to the database.
func (au *AuthUser) Insert(ctx context.Context, db DB) error {
	switch {
	case au._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case au._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO django.auth_user (` +
		`password, last_login, is_superuser, username, first_name, last_name, email, is_staff, is_active, date_joined` +
		`) VALUES (` +
		`:1, :2, :3, :4, :5, :6, :7, :8, :9, :10` +
		`) RETURNING id /*LASTINSERTID*/ INTO :pk`
	// run
	logf(sqlstr, au.Password, au.LastLogin, au.IsSuperuser, au.Username, au.FirstName, au.LastName, au.Email, au.IsStaff, au.IsActive, au.DateJoined)
	var id int64
	if _, err := db.ExecContext(ctx, sqlstr, au.Password, au.LastLogin, au.IsSuperuser, au.Username, au.FirstName, au.LastName, au.Email, au.IsStaff, au.IsActive, au.DateJoined, sql.Named("pk", sql.Out{Dest: &id})); err != nil {
		return err
	} // set primary key
	au.ID = int64(id)
	// set exists
	au._exists = true
	return nil
}

// Update updates a AuthUser in the database.
func (au *AuthUser) Update(ctx context.Context, db DB) error {
	switch {
	case !au._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case au._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE django.auth_user SET ` +
		`password = :1, last_login = :2, is_superuser = :3, username = :4, first_name = :5, last_name = :6, email = :7, is_staff = :8, is_active = :9, date_joined = :10 ` +
		`WHERE id = :11`
	// run
	logf(sqlstr, au.Password, au.LastLogin, au.IsSuperuser, au.Username, au.FirstName, au.LastName, au.Email, au.IsStaff, au.IsActive, au.DateJoined, au.ID)
	if _, err := db.ExecContext(ctx, sqlstr, au.Password, au.LastLogin, au.IsSuperuser, au.Username, au.FirstName, au.LastName, au.Email, au.IsStaff, au.IsActive, au.DateJoined, au.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the AuthUser to the database.
func (au *AuthUser) Save(ctx context.Context, db DB) error {
	if au.Exists() {
		return au.Update(ctx, db)
	}
	return au.Insert(ctx, db)
}

// Upsert performs an upsert for AuthUser.
func (au *AuthUser) Upsert(ctx context.Context, db DB) error {
	switch {
	case au._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `MERGE django.auth_usert ` +
		`USING (` +
		`SELECT :1 id, :2 password, :3 last_login, :4 is_superuser, :5 username, :6 first_name, :7 last_name, :8 email, :9 is_staff, :10 is_active, :11 date_joined ` +
		`FROM DUAL ) s ` +
		`ON s.id = t.id ` +
		`WHEN MATCHED THEN ` +
		`UPDATE SET ` +
		`t.password = s.password, t.last_login = s.last_login, t.is_superuser = s.is_superuser, t.username = s.username, t.first_name = s.first_name, t.last_name = s.last_name, t.email = s.email, t.is_staff = s.is_staff, t.is_active = s.is_active, t.date_joined = s.date_joined ` +
		`WHEN NOT MATCHED THEN ` +
		`INSERT (` +
		`password, last_login, is_superuser, username, first_name, last_name, email, is_staff, is_active, date_joined` +
		`) VALUES (` +
		`s.password, s.last_login, s.is_superuser, s.username, s.first_name, s.last_name, s.email, s.is_staff, s.is_active, s.date_joined` +
		`);`
	// run
	logf(sqlstr, au.ID, au.Password, au.LastLogin, au.IsSuperuser, au.Username, au.FirstName, au.LastName, au.Email, au.IsStaff, au.IsActive, au.DateJoined)
	if _, err := db.ExecContext(ctx, sqlstr, au.ID, au.Password, au.LastLogin, au.IsSuperuser, au.Username, au.FirstName, au.LastName, au.Email, au.IsStaff, au.IsActive, au.DateJoined); err != nil {
		return err
	}
	// set exists
	au._exists = true
	return nil
}

// Delete deletes the AuthUser from the database.
func (au *AuthUser) Delete(ctx context.Context, db DB) error {
	switch {
	case !au._exists: // doesn't exist
		return nil
	case au._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM django.auth_user ` +
		`WHERE id = :1`
	// run
	logf(sqlstr, au.ID)
	if _, err := db.ExecContext(ctx, sqlstr, au.ID); err != nil {
		return logerror(err)
	}
	// set deleted
	au._deleted = true
	return nil
}

// AuthUserByID retrieves a row from 'django.auth_user' as a AuthUser.
//
// Generated from index 'sys_c0012448'.
func AuthUserByID(ctx context.Context, db DB, id int64) (*AuthUser, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, password, last_login, is_superuser, username, first_name, last_name, email, is_staff, is_active, date_joined ` +
		`FROM django.auth_user ` +
		`WHERE id = :1`
	// run
	logf(sqlstr, id)
	au := AuthUser{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, id).Scan(&au.ID, &au.Password, &au.LastLogin, &au.IsSuperuser, &au.Username, &au.FirstName, &au.LastName, &au.Email, &au.IsStaff, &au.IsActive, &au.DateJoined); err != nil {
		return nil, logerror(err)
	}
	return &au, nil
}

// AuthUserByUsername retrieves a row from 'django.auth_user' as a AuthUser.
//
// Generated from index 'sys_c0012449'.
func AuthUserByUsername(ctx context.Context, db DB, username sql.NullString) (*AuthUser, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, password, last_login, is_superuser, username, first_name, last_name, email, is_staff, is_active, date_joined ` +
		`FROM django.auth_user ` +
		`WHERE username = :1`
	// run
	logf(sqlstr, username)
	au := AuthUser{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, username).Scan(&au.ID, &au.Password, &au.LastLogin, &au.IsSuperuser, &au.Username, &au.FirstName, &au.LastName, &au.Email, &au.IsStaff, &au.IsActive, &au.DateJoined); err != nil {
		return nil, logerror(err)
	}
	return &au, nil
}
