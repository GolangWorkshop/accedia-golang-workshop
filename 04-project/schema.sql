CREATE TABLE "public"."genres" (
    "id" serial NOT NULL,
    "genre" varchar NOT NULL,
    CONSTRAINT "genres_pk" PRIMARY KEY (id)
);

CREATE TABLE "public"."authors" (
    "id" serial NOT NULL,
    "first_name" varchar NULL,
    "last_name" varchar NULL,
    CONSTRAINT "authors_pk" PRIMARY KEY (id)
);

CREATE TABLE "public"."book_statuses" (
    "id" serial NOT NULL,
    "status" varchar NOT NULL,
    CONSTRAINT "book_statuses_pk" PRIMARY KEY (id)
);

CREATE TABLE "public"."books" (
  "id" serial NOT NULL,
  "name" varchar NOT NULL,
  "image" varchar NOT NULL,
  "description" varchar NOT NULL,
  "author_id" integer NOT NULL,
  "genre_id" integer NOT NULL,
  "status_id" integer NOT NULL,
  "is_deleted" bool NOT NULL DEFAULT false,
  CONSTRAINT "books_pk" PRIMARY KEY (id),
  FOREIGN KEY ("author_id") REFERENCES "public"."authors"("id") ON UPDATE restrict ON DELETE restrict,
  FOREIGN KEY ("genre_id") REFERENCES "public"."genres"("id") ON UPDATE restrict ON DELETE restrict,
  FOREIGN KEY ("status_id") REFERENCES "public"."book_statuses"("id") ON UPDATE restrict ON DELETE restrict
);

CREATE TABLE "public"."users" (
  "id" serial NOT NULL,
  "username" varchar NOT NULL,
  "password" varchar NOT null,
  CONSTRAINT "users_pk" PRIMARY KEY (id)
);

CREATE TABLE "public"."reviews" (
  "id" serial NOT NULL,
  "book_id" integer NOT NULL,
  "user_id" integer NOT NULL,
  "text" varchar NOT NULL,
  "is_deleted" bool NOT NULL DEFAULT false,
  CONSTRAINT "reviews_pk" PRIMARY KEY (id),
  FOREIGN KEY ("book_id") REFERENCES "public"."books"("id") ON UPDATE restrict ON DELETE restrict,
  FOREIGN KEY ("user_id") REFERENCES "public"."users"("id") ON UPDATE restrict ON DELETE restrict
);

CREATE TABLE "public"."votes" (
  "id" serial NOT NULL,
  "vote" varchar NOT NULL,
  CONSTRAINT "votes_pk" PRIMARY KEY (id)
);

CREATE TABLE "public"."user_voted_reviews" (
  "id" serial NOT NULL,
  "user_id" integer NOT NULL,
  "review_id" integer NOT NULL,
  "vote_id" integer NOT NULL,
  CONSTRAINT "user_voted_reviews_pk" PRIMARY KEY (id),
  FOREIGN KEY ("user_id") REFERENCES "public"."users"("id") ON UPDATE restrict ON DELETE restrict,
  FOREIGN KEY ("review_id") REFERENCES "public"."reviews"("id") ON UPDATE restrict ON DELETE restrict,
  FOREIGN KEY ("vote_id") REFERENCES "public"."votes"("id") ON UPDATE restrict ON DELETE restrict
);

CREATE TABLE "public"."history" (
  "id" serial NOT NULL,
  "user_id" integer NOT NULL,
  "book_id" integer NOT NULL,
  "borrowed" timestamp NOT NULL,
  "returned" timestamp DEFAULT NULL,
  CONSTRAINT "history_pk" PRIMARY KEY (id),
  FOREIGN KEY ("user_id") REFERENCES "public"."users"("id") ON UPDATE restrict ON DELETE restrict,
  FOREIGN KEY ("book_id") REFERENCES "public"."books"("id") ON UPDATE restrict ON DELETE restrict
);

CREATE TABLE "public"."book_rates" (
  "id" serial NOT NULL,
  "book_id" integer NOT NULL,
  "user_id" integer NOT NULL,
  "rate" integer NOT NULL DEFAULT 0,
  CONSTRAINT "book_rates_pk" PRIMARY KEY (id),
  FOREIGN KEY ("user_id") REFERENCES "public"."users"("id") ON UPDATE restrict ON DELETE restrict,
  FOREIGN KEY ("book_id") REFERENCES "public"."books"("id") ON UPDATE restrict ON DELETE restrict
);