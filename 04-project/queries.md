# SQL
To save some time, we are providing you the SQL queries, that you'll need.

## Required functionalities
### Retrieve all books
```sql
SELECT b.id, b.name, b.image, b.description,CONCAT(a.first_name, ' ', a.last_name) AS author, g.genre, bs.status FROM books b JOIN book_statuses bs ON bs.id = b.status_id JOIN authors a ON a.id = b.author_id JOIN genres g ON g.id = b.genre_id WHERE b.is_deleted = false ORDER BY b.id ASC;
```

### View an individual book
```sql
SELECT b.id, b.name, b.image, b.description, CONCAT(a.first_name, ' ', a.last_name) AS author, g.genre, bs.status FROM books b JOIN book_statuses bs ON bs.id = b.status_id JOIN authors a ON a.id = b.author_id
JOIN genres g ON g.id = b.genre_id WHERE b.is_deleted = false AND b.id = $1;
```

### Create a new book
```sql
INSERT INTO books (name, image, description, author_id, genre_id) VALUES ($1, $2, $3, $4, $5, $6);
```
**Note**: For the return value you can use the query for viewing a single book.

### Update an existing book
```sql
UPDATE "books" SET name = $1 image = $2 description = $3 author_id = $4 genre_id = $5 status_id = $6 WHERE book_id =$7 AND is_deleted = false
```

**Note**: For the return value you can use the query for viewing a single book.

### Delete a book
```sql
UPDATE books SET is_deleted = ture WHERE book_id = $1
```

**Note**: For the return value you can use the query for viewing a single book. Remember to fire this query, before deleting it.

## Bonus functionalities

### Use query parameters
In the query below we added a few filters - by `name`, `author` and `genre`. We also implemented pagination by using `limit` and `offset`.
```sql
SELECT b.id, b.name, b.image, b.description, CONCAT(a.first_name, ' ', a.last_name) AS author, g.genre, bs.status FROM books b JOIN book_statuses bs ON bs.id = b.status_id JOIN authors a ON a.id = b.author_id JOIN genres g ON g.id = b.genre_id WHERE b.is_deleted = false AND b.name LIKE '%$1%' AND CONCAT(a.first_name, ' ', a.last_name) LIKE '%$2%' AND g.genre LIKE '%$3%' ORDER BY b.name ASC LIMIT $4 OFFSET $5;
```

### Average rate for a book

```sql
SELECT AVG(rate) AS rate FROM book_rates WHERE book_id = $1;
```

### Authentication
#### Register
```sql
INSERT INTO users(username, password) VALUES ($1, $2);
```

#### Login
```sql
SELECT * FROM users WHERE username = $1;
```

### Book history

#### Borrow a book
```sql
INSERT INTO history(user_id, book_id, borrowed) VALUES ($1, $2, $3);

UPDATE books SET status_id = 2 WHERE id = $1 AND is_deleted = false;
```

##### Return a book
- Find the correct `history` entity:
```sql
SELECT * FROM history WHERE user_id = $1 AND book_id = $2 AND returned IS NULL;
```
- Update the `returned` field and the book status 
```sql
UPDATE history SET returned = $1 WHERE id = $2;

UPDATE books SET status_id = 1 WHERE id = $1 AND is_deleted = false;
```

### Book reviews
#### Get all reviews for a single book
```sql
SELECT r.book_id, r.id, r.text, u.id FROM reviews r JOIN users u ON u.id = r.user_id WHERE book_id = $1 AND r.is_deleted = false;
```

#### Get a single review
This is a helper query, that you might need:
```sql
SELECT r.id, u.id, u.username, r.text FROM reviews r JOIN users u ON u.id = r.user_id WHERE r.book_id = $1 AND u.id = $2 AND r.is_deleted = false;
```
#### Create a new review
```sql
INSERT INTO reviews(book_id, user_id, text) VALUES($1, $2, $3);
```
#### Update a review
```sql
UPDATE reviews SET text = $1 WHERE review_id = $2;
```
#### Delete a review
```sql
UPDATE reviews SET is_deleted = true WHERE review_id = $1;
```

### Review votes
#### Get all votes for a single review
```sql
SELECT COUNT(CASE WHEN uvr.vote_id= 1 THEN 1 END) AS likes, COUNT(CASE WHEN uvr.vote_id= 2 THEN 1 END) AS dislikes FROM user_voted_reviews uvr LEFT JOIN reviews r ON uvr.review_id = r.id WHERE uvr.review_id = $1;
```
#### Get a single user's review vote
This is a helper query, that you might need:
```sql
SELECT uvr.id, uvr.user_id, uvr.review_id, v.vote FROM user_voted_reviews uvr JOIN votes v ON v.id = uvr.vote_id WHERE uvr.review_id = $1 AND uvr.user_id = $2;
```
#### Create a new vote
```sql
INSERT INTO user_voted_reviews(user_id, review_id, vote_id) VALUES ($1, $2, $3);
```
#### Update a vote
```sql
UPDATE user_voted_reviews SET vote_id = $1 WHERE id = $2;
```
#### Delete a vote
```sql
DELETE FROM user_voted_reviews WHERE id = $1;
```
