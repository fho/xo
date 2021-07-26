-- Generated by xo for the booktest schema.

-- table authors
CREATE TABLE authors (
  author_id NUMBER GENERATED ALWAYS AS IDENTITY,
  name NVARCHAR2 NOT NULL,
  CONSTRAINT authors_pkey UNIQUE (author_id)
);

-- index authors_name_idx
CREATE INDEX authors_name_idx ON authors (name);

-- table books
CREATE TABLE books (
  book_id NUMBER GENERATED ALWAYS AS IDENTITY,
  author_id NUMBER NOT NULL CONSTRAINT books_author_id_fkey REFERENCES authors (author_id),
  isbn NVARCHAR2 NOT NULL,
  title NVARCHAR2 NOT NULL,
  year NUMBER NOT NULL,
  available TIMESTAMP WITH TIME ZONE(6) NOT NULL,
  description NCLOB,
  tags NCLOB NOT NULL,
  CONSTRAINT books_isbn_key UNIQUE (isbn),
  CONSTRAINT books_pkey UNIQUE (book_id)
);

-- index books_title_idx
CREATE INDEX books_title_idx ON books (title, year);

-- function say_hello
CREATE FUNCTION say_hello(name IN NVARCHAR2) RETURN NVARCHAR2 AS
BEGIN
  RETURN 'hello ' || name\;
END\;;