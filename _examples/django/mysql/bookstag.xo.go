package mysql

// Code generated by xo. DO NOT EDIT.

import (
	"context"
)

// BooksTag represents a row from 'django.books_tags'.
type BooksTag struct {
	ID     int64 `json:"id"`      // id
	BookID int64 `json:"book_id"` // book_id
	TagID  int64 `json:"tag_id"`  // tag_id
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the BooksTag exists in the database.
func (bt *BooksTag) Exists() bool {
	return bt._exists
}

// Deleted returns true when the BooksTag has been marked for deletion from
// the database.
func (bt *BooksTag) Deleted() bool {
	return bt._deleted
}

// Insert inserts the BooksTag to the database.
func (bt *BooksTag) Insert(ctx context.Context, db DB) error {
	switch {
	case bt._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case bt._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO django.books_tags (` +
		`book_id, tag_id` +
		`) VALUES (` +
		`?, ?` +
		`)`
	// run
	logf(sqlstr, bt.BookID, bt.TagID)
	res, err := db.ExecContext(ctx, sqlstr, bt.BookID, bt.TagID)
	if err != nil {
		return err
	}
	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	} // set primary key
	bt.ID = int64(id)
	// set exists
	bt._exists = true
	return nil
}

// Update updates a BooksTag in the database.
func (bt *BooksTag) Update(ctx context.Context, db DB) error {
	switch {
	case !bt._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case bt._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE django.books_tags SET ` +
		`book_id = ?, tag_id = ? ` +
		`WHERE id = ?`
	// run
	logf(sqlstr, bt.BookID, bt.TagID, bt.ID)
	if _, err := db.ExecContext(ctx, sqlstr, bt.BookID, bt.TagID, bt.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the BooksTag to the database.
func (bt *BooksTag) Save(ctx context.Context, db DB) error {
	if bt.Exists() {
		return bt.Update(ctx, db)
	}
	return bt.Insert(ctx, db)
}

// Upsert performs an upsert for BooksTag.
func (bt *BooksTag) Upsert(ctx context.Context, db DB) error {
	switch {
	case bt._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO django.books_tags (` +
		`id, book_id, tag_id` +
		`) VALUES (` +
		`?, ?, ?` +
		`)` +
		` ON DUPLICATE KEY UPDATE ` +
		`book_id = VALUES(book_id), tag_id = VALUES(tag_id)`
	// run
	logf(sqlstr, bt.ID, bt.BookID, bt.TagID)
	if _, err := db.ExecContext(ctx, sqlstr, bt.ID, bt.BookID, bt.TagID); err != nil {
		return err
	}
	// set exists
	bt._exists = true
	return nil
}

// Delete deletes the BooksTag from the database.
func (bt *BooksTag) Delete(ctx context.Context, db DB) error {
	switch {
	case !bt._exists: // doesn't exist
		return nil
	case bt._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM django.books_tags ` +
		`WHERE id = ?`
	// run
	logf(sqlstr, bt.ID)
	if _, err := db.ExecContext(ctx, sqlstr, bt.ID); err != nil {
		return logerror(err)
	}
	// set deleted
	bt._deleted = true
	return nil
}

// BooksTagByBookIDTagID retrieves a row from 'django.books_tags' as a BooksTag.
//
// Generated from index 'books_tags_book_id_tag_id_29db9e39_uniq'.
func BooksTagByBookIDTagID(ctx context.Context, db DB, bookID, tagID int64) (*BooksTag, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, book_id, tag_id ` +
		`FROM django.books_tags ` +
		`WHERE book_id = ? AND tag_id = ?`
	// run
	logf(sqlstr, bookID, tagID)
	bt := BooksTag{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, bookID, tagID).Scan(&bt.ID, &bt.BookID, &bt.TagID); err != nil {
		return nil, logerror(err)
	}
	return &bt, nil
}

// BooksTagByID retrieves a row from 'django.books_tags' as a BooksTag.
//
// Generated from index 'books_tags_id_pkey'.
func BooksTagByID(ctx context.Context, db DB, id int64) (*BooksTag, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, book_id, tag_id ` +
		`FROM django.books_tags ` +
		`WHERE id = ?`
	// run
	logf(sqlstr, id)
	bt := BooksTag{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, id).Scan(&bt.ID, &bt.BookID, &bt.TagID); err != nil {
		return nil, logerror(err)
	}
	return &bt, nil
}

// BooksTagsByTagID retrieves a row from 'django.books_tags' as a BooksTag.
//
// Generated from index 'books_tags_tag_id_8d70b40a_fk_tags_tag_id'.
func BooksTagsByTagID(ctx context.Context, db DB, tagID int64) ([]*BooksTag, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, book_id, tag_id ` +
		`FROM django.books_tags ` +
		`WHERE tag_id = ?`
	// run
	logf(sqlstr, tagID)
	rows, err := db.QueryContext(ctx, sqlstr, tagID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*BooksTag
	for rows.Next() {
		bt := BooksTag{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&bt.ID, &bt.BookID, &bt.TagID); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &bt)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// Book returns the Book associated with the BooksTag's (BookID).
//
// Generated from foreign key 'books_tags_book_id_73d7d8e8_fk_books_book_id'.
func (bt *BooksTag) Book(ctx context.Context, db DB) (*Book, error) {
	return BookByBookID(ctx, db, bt.BookID)
}

// Tag returns the Tag associated with the BooksTag's (TagID).
//
// Generated from foreign key 'books_tags_tag_id_8d70b40a_fk_tags_tag_id'.
func (bt *BooksTag) Tag(ctx context.Context, db DB) (*Tag, error) {
	return TagByTagID(ctx, db, bt.TagID)
}