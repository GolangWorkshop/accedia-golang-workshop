package store

type Book struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Image       string `json:"image"`
	Description string `json:"description"`
	AuthorId    string `json:"author_id"`
	GenreId     string `json:"genre_id"`
	StatusId    string `json:"status_id"`
}

func (st *store) GetBooks() (*[]Book, error) {
	rows, err := st.db.Query("SELECT id, name, image, description, author_id, genre_id, status_id FROM books WHERE is_deleted = false")
	if err != nil {
		log.Error().Err(err).Msg("could not retrieve books")
		return nil, err
	}
	defer rows.Close()

	books := []Book{}
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.Id, &book.Name, &book.Image, &book.Description, &book.AuthorId, &book.GenreId, &book.StatusId)
		if err != nil {
			log.Error().Err(err).Msg("could not retrieve books")
			return nil, err
		}
		books = append(books, book)
	}

	return &books, nil
}

func (st *store) GetBookById(id string) (*Book, error) {
	rows, err := st.db.Query("SELECT id, name, image, description, author_id, genre_id, status_id FROM books WHERE id = $1 AND is_deleted = false", id)
	if err != nil {
		log.Error().Err(err).Msgf("could not retrieve book with id %s", id)
		return nil, err
	}
	defer rows.Close()

	books := []Book{}
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.Id, &book.Name, &book.Image, &book.Description, &book.AuthorId, &book.GenreId, &book.StatusId)
		if err != nil {
			log.Error().Err(err).Msgf("could not retrieve book with id %s", id)
			return nil, err
		}
		books = append(books, book)
	}

	if len(books) == 0 {
		return nil, nil
	}

	return &books[0], nil
}

func (st *store) CreateBook(book *Book) (b *Book, err error) {
	row := st.db.QueryRow("INSERT INTO books (name, image, description, author_id, genre_id, status_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		book.Name, book.Image, book.Description, book.AuthorId, book.GenreId, book.StatusId)

	err = row.Scan(&book.Id)
	if err != nil {
		log.Error().Err(err).Msgf("could not create book %s", book)
		return nil, err
	}

	return book, nil
}

func (st *store) UpdateBookById(id string, book *Book) (*Book, error) {
	_, err := st.db.Exec("UPDATE books SET name = $1, image = $2, description = $3, author_id = $4, genre_id = $5, status_id = $6 WHERE id = $7 AND is_deleted = false",
		book.Name, book.Image, book.Description, book.AuthorId, book.GenreId, book.StatusId, id)
	if err != nil {
		log.Error().Err(err).Msgf("failed to update book %s", book)
		return nil, err
	}

	return book, err
}

func (st *store) DeleteBookById(id string) (affectedRows int64, err error) {
	res, err := st.db.Exec("UPDATE books SET is_deleted = true WHERE id = $1", id)
	if err != nil {
		log.Error().Err(err).Msgf("failed to delete book %s", id)
		return 0, err
	}

	return res.RowsAffected()
}
