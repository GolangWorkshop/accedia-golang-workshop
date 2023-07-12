package webserver

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/GolangWorkshop/library/store"
)

func isBookMalformed(book *store.Book) bool {
	return book.Name == "" || book.Description == "" || book.AuthorId == "" || book.GenreId == "" || book.StatusId == ""
}

func isUserMalformed(user *store.User) bool {
	return user.Password == "" || user.Username == ""
}

func isUserAuthenticated(w http.ResponseWriter, r *http.Request) bool {
	user := r.Context().Value("username")
	if user == nil {
		log.Info().Msgf("could not authorize user")
		w.WriteHeader(http.StatusUnauthorized)
		return false
	}
	return true
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug().Msgf("unhandled method %s hit on path %s", r.Method, r.URL.Path)
	s := http.StatusNotFound
	w.WriteHeader(s)
	w.Write([]byte(http.StatusText(s)))
}

func addRoutes(sm *http.ServeMux, st store.Store) {
	addApiRoutes(sm, st)
	// add asset handlers
}

func addApiRoutes(sm *http.ServeMux, st store.Store) {
	sm.HandleFunc("/api/books", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			books, err := st.GetBooks()
			if err != nil {
				log.Error().Err(err).Msg("could not retrieve books")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(books)
		case http.MethodPost:
			check := isUserAuthenticated(w, r)
			if !check {
				return
			}

			var book store.Book
			err := json.NewDecoder(r.Body).Decode(&book)
			if err != nil {
				log.Error().Err(err).Msg("could not decode requrest body as book")
				http.Error(w, "could not decode requrest body as book", http.StatusBadRequest)
				return
			}

			if book.Id != "" || isBookMalformed(&book) {
				log.Error().Err(err).Msgf("create book failed because of malformed book object received %s", book)
				http.Error(w, "malformed book object received", http.StatusBadRequest)
				return
			}

			createdBook, err := st.CreateBook(&book)
			if err != nil || createdBook == nil {
				log.Error().Err(err).Msg("could not create book")
				w.WriteHeader(http.StatusInternalServerError)
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(createdBook)
		default:
			NotFoundHandler(w, r)
		}
	})

	sm.HandleFunc("/api/books/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			bookId := r.URL.Path[len("/api/books/"):]
			book, err := st.GetBookById(bookId)
			if err != nil {
				log.Error().Err(err).Msgf("could not retrieve book with id %s", bookId)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			if book == nil {
				log.Error().Err(err).Msgf("could not retrieve book with id %s", bookId)
				w.WriteHeader(http.StatusNotFound)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(book)

		case http.MethodPut:
			check := isUserAuthenticated(w, r)
			if !check {
				return
			}

			bookId := r.URL.Path[len("/api/books/"):]

			if bookId == "" {
				log.Error().Msgf("update book failed because of malformed book id received %s", bookId)
				http.Error(w, "book id not present", http.StatusBadRequest)
				return
			}

			var book store.Book
			err := json.NewDecoder(r.Body).Decode(&book)
			if err != nil {
				log.Error().Err(err).Msg("could not decode requrest body as book")
				http.Error(w, "could not decode requrest body as book", http.StatusBadRequest)
				return
			}

			if book.Id == "" || isBookMalformed(&book) {
				log.Error().Err(err).Msgf("update book failed because of malformed book object received %s", book)
				http.Error(w, "malformed book object received", http.StatusBadRequest)
				return
			}

			updatedBook, err := st.UpdateBookById(bookId, &book)
			if err != nil || updatedBook == nil {
				log.Error().Err(err).Msg("could not update book")
				w.WriteHeader(http.StatusInternalServerError)
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(book)

		case http.MethodDelete:
			check := isUserAuthenticated(w, r)
			if !check {
				return
			}

			bookId := r.URL.Path[len("/api/books/"):]

			if bookId == "" {
				log.Error().Msgf("delete book failed because of malformed book id received %s", bookId)
				http.Error(w, "book id not present", http.StatusBadRequest)
				return
			}

			deletedBooks, err := st.DeleteBookById(bookId)
			if err != nil || deletedBooks == 0 {
				log.Error().Err(err).Msg("could not delete book")
				w.WriteHeader(http.StatusInternalServerError)
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(deletedBooks)
		default:
			NotFoundHandler(w, r)
		}
	})

	sm.HandleFunc("/api/register", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			var user store.User
			err := json.NewDecoder(r.Body).Decode(&user)
			if err != nil {
				log.Error().Err(err).Msg("could not decode requrest body as user")
				http.Error(w, "could not decode requrest body as user", http.StatusBadRequest)
				return
			}

			if isUserMalformed(&user) {
				log.Error().Err(err).Msgf("register user failed because of malformed user object received %s", user)
				http.Error(w, "malformed user object received", http.StatusBadRequest)
				return
			}

			createdUser, err := st.Register(&user)
			if err != nil || createdUser == nil {
				log.Error().Err(err).Msg("could not register user")
				w.WriteHeader(http.StatusInternalServerError)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
		default:
			NotFoundHandler(w, r)
		}
	})

	sm.HandleFunc("/api/login", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			var user store.User
			err := json.NewDecoder(r.Body).Decode(&user)
			if err != nil {
				log.Error().Err(err).Msg("could not decode requrest body as user")
				http.Error(w, "could not decode requrest body as user", http.StatusBadRequest)
				return
			}

			if isUserMalformed(&user) {
				log.Error().Err(err).Msgf("log user in failed because of malformed user object received %s", user)
				http.Error(w, "malformed user object received", http.StatusBadRequest)
				return
			}

			tokenInfo, err := st.Login(&user)
			if err != nil {
				log.Error().Err(err).Msg("could not log user in ")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			// Set the client cookie for "token" as the JWT we just generated
			// we also set an expiry time which is the same as the token itself
			http.SetCookie(w, &http.Cookie{
				Name:    "token",
				Value:   tokenInfo.Token,
				Expires: tokenInfo.ExpiresAt,
			})
		default:
			NotFoundHandler(w, r)
		}
	})

	sm.HandleFunc("/api/logout", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodDelete:
			// Immediately clear the token cookie
			http.SetCookie(w, &http.Cookie{
				Name:    "token",
				Expires: time.Now(),
			})
		default:
			NotFoundHandler(w, r)
		}
	})
}
