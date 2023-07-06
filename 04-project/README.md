# Book Library API
In the live demo, together we are going to design and implement an API, which supports all CRUD operations for books. 
There is also a bonus section if you want to challenge yourself.
Don't worry for the SQL queries, we'll provide them to you - in the directory you can find a magical file, named `queries.md`.
During the live demo, we'll provide you a link to the DB and everything you'll need to connect to it. If you want to start earlier - there is a `diagram.png`, which is a graphical representation of the DB, a `schema.sql` file, that contains the SQL for creating the DB and of course, a `seed.sql` file with the seeds.

## Table of contents
- [**Required functionalities**](#required-functionalities)
  - [Retrieve all books](#retrieve-all-books)
  - [View an individual book](#view-an-individual-book)
  - [Create a new book](#create-a-new-book)
  - [Update an existing book](#update-an-existing-book)
  - [Delete a book](#delete-a-book)
- [**Bonus functionalities**](#bonus-functionalities)
  - [Enhance the existing features](#enhance-the-existing-features)
    - [Use search query parameters](#use-search-query-parameters)
    - [Average rate for a book](#average-rate-for-a-book)
  - [New features](#new-features)
    - [Authentication](#authentication)
      - [Register](#register)
      - [Login](#login)
      - [Logout](#logout)
    - [History of a book](#history-of-a-book)
    - [Book reviews](#book-reviews)
      - [Get all reviews for a single book](#get-all-reviews-for-a-single-book)
      - [Create a new review](#create-a-new-review)
      - [Update a review](#update-a-review)
      - [Delete a review](#delete-a-review)
    - [Review votes](#review-votes)
      - [Get all votes for a single review](#get-all-votes-for-a-single-review)
      - [Create a new vote](#create-a-new-vote)
      - [Update a vote](#update-a-vote)
      - [Delete a vote](#delete-a-vote)

## Required functionalities
### Retrieve all books
A `GET` request, that will allow you to see all the available books in the library. 

**Notes**: There are a few relations, that we'll use here:
- A one-to-many between `books` and `statuses`:
    - `free` - can be borrowed, `status_id = 1`
    - `borrowed` - `status_id = 2`
    - `unlisted` - cannot be retrieved or borrowed, `status_id = 3`
- A one-to-many between `books` and `authors`
- A one-to-many between `books` and `genres`

You should get the books from the database and return them as a response. 

- Example request: `GET/api/books`
- Example response: The books data

    ```json
    [
        {
            "id": 1,
            "name": "Pippi Longstocking",
            "image": "https://upload.wikimedia.org/wikipedia/en/7/78/L%C3%A5ngstrump_G%C3%A5r_Ombord.jpeg",
            "description": "The beloved story of a spunky young girl and her hilarious escapades. \"A rollicking story.\"--The Horn Book Tommy and his sister Annika have a new neighbor, and her name is Pippi Longstocking. She has crazy red pigtails, no parents to tell her what to do, a horse that lives on her porch, and a flair for the outrageous that seems to lead to one adventure after another!",
            "author": "Astrid Lindgren",
            "genre": "fiction",
            "status": "borrowed"
        },
        {
            "id": 2,
            "name": "Order of the Phoenix",
            "image": "https://upload.wikimedia.org/wikipedia/en/7/70/Harry_Potter_and_the_Order_of_the_Phoenix.jpg",
            "description": "Rowling and the fifth novel in the Harry Potter series. It follows Harry Potter's struggles through his fifth year at Hogwarts School of Witchcraft and Wizardry, including the surreptitious return of the antagonist Lord Voldemort, O.W.L. exams, and an obstructive Ministry of Magic.",
            "author": "Joanne Rowling",
            "genre": "fantasy",
            "status": "free"
        },
        {
            "id": 3,
            "name": "The Mysterious Affair at Styles",
            "image": "https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/lrg/9780/0084/9780008400637.jpg",
            "description": "Agatha Christieâ€™s first novel, The Mysterious Affair at Styles, was the result of a dare from her sister Madge who challenged her to write a story. The story begins when Hastings is sent back to England from the First World War due to injury and is invited to spend his sick leave at the beautiful Styles Court by his old friend John Cavendish. Here, Hastings meets Johnâ€™s step-mother, Mrs Inglethorp, and her new husband, Alfred. Despite the tranquil surroundings Hastings begins to realise that all is not right. When Mrs Inglethorp is found poisoned, suspicion falls on the family, and another old friend, Hercule Poirot, is invited to investigate.",
            "author": "Agatha Christie",
            "genre": "criminal",
            "status": "borrowed"
        },
        {
            "id": 4,
            "name": "Romeo and Juliet ",
            "image": "https://prodimage.images-bn.com/pimages/9780743477116_p0_v2_s1200x630.jpg",
            "description": "Romeo and Juliet is a tragedy written by William Shakespeare early in his career about two young star-crossed lovers whose deaths ultimately reconcile their feuding families. It was among Shakespeare's most popular plays during his lifetime and, along with Hamlet, is one of his most frequently performed plays.",
            "author": "William Shakespeare",
            "genre": "novel",
            "status": "free"
        },
        {
            "id": 5,
            "name": "Hamlet",
            "image": "https://i.pinimg.com/originals/a7/c2/5b/a7c25bb8144b76673d27593d8bf8081b.jpg",
            "description": "Hamlet is melancholy, bitter, and cynical, full of hatred for his uncle's scheming and disgust for his mother's sexuality. A reflective and thoughtful young man who has studied at the University of Wittenberg, Hamlet is often indecisive and hesitant, but at other times prone to rash and impulsive acts.\n",
            "author": "William Shakespeare",
            "genre": "tragedy",
            "status": "free"
        }
        ...
    ]
    ```

### View an individual book
Another `GET` request, but this one will allow you to see individual books. You should get the book from the database by its unique `id` and return it as a response.

- Example request: `GET/api/books/1`
- Example response: The book’s data
    ```json
    {
      "id": 1,
      "name": "Pippi Longstocking",
      "image": "https://upload.wikimedia.org/wikipedia/en/7/78/L%C3%A5ngstrump_G%C3%A5r_Ombord.jpeg",
      "description": "The beloved story of a spunky young girl and her hilarious escapades. \"A rollicking story.\"--The Horn Book Tommy and his sister Annika have a new neighbor, and her name is Pippi Longstocking. She has crazy red pigtails, no parents to tell her what to do, a horse that lives on her porch, and a flair for the outrageous that seems to lead to one adventure after another!",
      "author": "Astrid Lindgren",
      "genre": "fiction",
      "status": "free",
    }
    ```


### Create a new book
A `POST` request, that will allow you to add a new book to the system. You should send the new data to the database and return a response - the newly created book's data.

- Example request: `POST/api/books`
- Example request body:

    ```json
      {
        "name": "The Goose Egg",
        "image": "https://images-na.ssl-images-amazon.com/images/I/51A%2B8tZvLXL._SX218_BO1,204,203,200_QL40_ML2_.jpg",
        "description": "Henrietta likes her quiet life. A morning swim, a cup of tea--all is serene. But everything changes when she bumps her head and winds up with a goose egg--a REAL goose egg. Henrietta tries to return the baby goose to the nest, but her flock has flown. It's up to Henrietta to raise her.",
        "author_id": 11,
        "genre_id": 6,
      }
    ```

- Example response: The newly created book's data
    ```json
    {
        "id": 46,
        "name": "The Goose Eggs",
        "image": "https://images-na.ssl-images-amazon.com/images/I/51A%2B8tZvLXL._SX218_BO1,204,203,200_QL40_ML2_.jpg",
        "description": "Henrietta likes her quiet life. A morning swim, a cup of tea--all is serene. But everything changes when she bumps her head and winds up with a goose egg--a REAL goose egg. Henrietta tries to return the baby goose to the nest, but her flock has flown. It's up to Henrietta to raise her.",
        "author": "Liz Wong",
        "genre": "children's literature",
        "status": "free",
      }
    ```

### Update an existing book
A `PUT` request, that will allow you to update a single book. You should send the data to update and return the updated entity.
- Example request: `PUT/api/books/46`
- Example request body:
    ```json
    {
        "name": "The Goose Eggs",
        "image": "https://images-na.ssl-images-amazon.com/images/I/51A%2B8tZvLXL._SX218_BO1,204,203,200_QL40_ML2_.jpg",
        "description": "Henrietta likes her quiet life. A morning swim, a cup of tea--all is serene. But everything changes when she bumps her head and winds up with a goose egg--a REAL goose egg. Henrietta tries to return the baby goose to the nest, but her flock has flown. It's up to Henrietta to raise her.",
        "author_id": 11,
        "genre_id": 6,
        "status_id": 2
    }
    ```
- Example response: The updated book's data (here we've updated the `status`)
    ```json
    {
        "id": 46,
        "name": "The Goose Eggs",
        "image": "https://images-na.ssl-images-amazon.com/images/I/51A%2B8tZvLXL._SX218_BO1,204,203,200_QL40_ML2_.jpg",
        "description": "Henrietta likes her quiet life. A morning swim, a cup of tea--all is serene. But everything changes when she bumps her head and winds up with a goose egg--a REAL goose egg. Henrietta tries to return the baby goose to the nest, but her flock has flown. It's up to Henrietta to raise her.",
        "author": "Liz Wong",
        "genre": "children's literature",
        "status": "borrowed",
      }
    ```

## Delete a book
And the final `DELETE` request, that will allow you to delete a single book.

- Example request: `DELETE/api/books/46`
- Example response: The deleted book's data
    ```json
    {
        "id": 46,
        "name": "The Goose Eggs",
        "image": "https://images-na.ssl-images-amazon.com/images/I/51A%2B8tZvLXL._SX218_BO1,204,203,200_QL40_ML2_.jpg",
        "description": "Henrietta likes her quiet life. A morning swim, a cup of tea--all is serene. But everything changes when she bumps her head and winds up with a goose egg--a REAL goose egg. Henrietta tries to return the baby goose to the nest, but her flock has flown. It's up to Henrietta to raise her.",
        "author": "Liz Wong",
        "genre": "children's literature",
        "status": "free",
      }
    ```


## Bonus functionalities
If you wish to add some additional features, here are a few ideas.

### Enhance the existing features
#### Use search query parameters
For the feature for reading all books you can add:

- **server-side pagination**
    - instead of returning all of the books, return the data for an individual page
    - you'll need a parameter for the books per page (limit) and one for the offset
- **filtering** - by author, genre, book name
- **sorting** - ascending or descending

To achieve that use query parameters.

#### Average rate for a book
There is a table, called `book_rates`, which represents a many-to-many relation between `users`* and `books`. The table has a column `rate` - an integer. When  querying for an individual book, get all rates for that book and find the average rate. Then, in the response add an additional field, which is the calculated average rate of the book. Pick a meaningful name.

***Note**: For this feature, we do not care about `users`

### New features
#### Authentication
As mentioned above, there is a `users` table in our database with a `username` and `password` columns. Use this table and try to implement authentication - register, login and logout. Use JWT strartegy. Remember to hash all passwords. Try to protect some of the routes as well.

##### Register
- Example request: `POST/api/users`
- Example request body:
    ```json
    {
        "username": "ivan",
        "password": "1234"
    }
    ```
- Example response: Bearer token
    ```json
    {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsInVzZXJuYW1lIjoiaXZhbiIsImlhdCI6MTUxNjIzOTAyMn0.MpJjYMjJE3aDNY43QNlG4VcoIA157xtliKPVmqktdhY"
    }
    ```

##### Login
- Example request: `POST/api/session`, `POST/api/users/login`, `POST/api/login`
- Example request body:
    ```json
    {
        "username": "ivan",
        "password": "1234"
    }
    ```
- Example response: Bearer token
    ```json
    {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMSIsInVzZXJuYW1lIjoiaXZhbiIsImlhdCI6MTUxNjIzOTAyMn0.MpJjYMjJE3aDNY43QNlG4VcoIA157xtliKPVmqktdhY"
    }
    ```
##### Logout
- Example request: `DELETE/api/session`, `POST/api/users/logout`, `POST/api/logout`
- Example response: A message
    ```json
    {
        "message": "Logged out"
    }
    ```

**Note**: Try to invalidate the token.

#### History of a book
There is a table in our database, called `history`. And there is a one-to-many relation between `history` and `books`. Use that relation to implement the borrow and return functionalities. These routes should be available for authenticated users only.

##### Borrow a book
In `history` there are a `borrowed` and a `returned` field of type `datetime`. The `returned` one is nullable. When borrowing a book:
- insert a new entity in the `history` table
- use the current date and time for `borrowed`
- leave `returned` as `null`

Remember that a Bearer token should be send.

- Example request: `POST/api/books/1/history`
- Example response: The history data
    ```json
    {
        "id": 35,
        "user_id": 11,
        "borrowed": "2016-06-22 10:00:00-04",
        "returned": null
    }
    ```

##### Return a book
For returning a book, find the correct `history` entity and update only the `returend` field with the current date and time.

- Example request: `PUT/api/books/1/history`
- Example response:
    ```json
    {
        "id": 35,
        "user_id": 11,
        "borrowed": "2016-06-22 10:00:00-04",
        "returned": "2016-06-30 10:00:00-04",
    }
    ```

#### Book reviews
In our database there is a many-to-one relation between `books` and `reviews`. Each review has a `text` field and is created by a `user` (many-to-one relation).

##### Get all reviews for a single book
You can decide if you prefer to keep this route proteced. If you like, you can also add this data as a field to the response for getting an individual book. It's up to you.

- Example request: `GET/api/books/1/reviews`
- Example response: All reviews for the given book
    ```json
    {
      "reviews": [
        {
          "id": 4,
          "book_id": 1,
          "text": "This is my favourite book!",
          "user_id": 1,
        }
      ]
    }
    ```

##### Create a new review
The route should be protected.
- Example request: `POST/api/books/1/reviews`
- Example request body:
    ```json
    {
        "text": "Awesome book!"
    }
    ```
- Example response: The newly created review
    ```json
    {
        "id": 26,
        "book_id": 1,
        "user_id": 11,
        "text": "Awesome book!"
    }
    ```

##### Update a review
The route should be protected. Users should be able to update only their own reviews.
- Example request: `PUT/api/books/1/reviews`
- Example request body:
    ```json
    {
        "text": "Awesome book! I bought it as a gift to my mother."
    }
    ```
- Example response: The updated review
    ```json
    {
        "id": 26,
        "book_id": 1,
        "user_id": 11,
        "text": "Awesome book! I bought it as a gift to my mother."
    }
    ```

##### Delete a review
The route should be protected. Users should be able to delete only their own reviews.
- Example request: `DELETE/api/books/1/reviews/26`
- Example response: The updated review
    ```json
    {
        "id": 26,
        "book_id": 23,
        "user_id": 11,
        "text": "Awesome book! I bought it as a gift to my mother."
    }
    ```

#### Review votes
Each review has a number of likes and dislikes, we called these `votes`, and each vote is created by a user. So, there is a `user_voted_reviews` table, where we store the relations between these three.

##### Get all votes for a single review
Use the `user_liked_reviews` table to sum the likes and dislikes of the given review. If you wish you can protect this route for authenticated users only. You can add this data as additional field in the response for getting the reviews of a book. It's up to you.
- Example request: `GET/api/books/1/reviews/4/votes`
- Example response: The calculated votes for the given review
    ```json
    {
        "likes": 2,
        "dislikes": 1
    }
    ```
##### Create a new vote
The votes are as follows:
- like - `vote_id = 1`
- dislike - `vote_id = 2`

The route should be protected.

- Example request: `POST/api/books/1/reviews/4/votes`
- Example request body:
    ```json
    {
        "vote_id": 1
    }
    ```
- Example response: The newly created vote by the user
    ```json
    {
        "id": 16,
        "user_id": 11,
        "review_id": 4,
        "vote": "like"
    }
    ```
##### Update a vote
The route should be protected. Users should be able to update only their own votes.
- Example request: `PUT/api/books/1/reviews/4/votes/16`
- Example request body:
    ```json
    {
        "vote_id": 2
    }
    ```
- Example response: The newly created vote by the user
    ```json
    {
        "id": 16,
        "user_id": 11,
        "review_id": 4,
        "vote": "dislike"
    }
    ```
##### Delete a vote
The route should be protected. Users should be able to delete only their own votes.
- Example request: `DELETE/api/books/1/reviews/4/votes/16`
- Example response: The deleted vote
    ```json
    {
        "id": 16,
        "user_id": 11,
        "review_id": 4,
        "vote": "dislike"
    }
    ```

