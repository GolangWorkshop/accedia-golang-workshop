package store

import (
	"time"

	"github.com/GolangWorkshop/library/util"
	"github.com/golang-jwt/jwt/v4"
)

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (st *store) Register(user *User) (u *User, err error) {
	hash, err := util.HashPassword(user.Password)

	if err != nil {
		log.Error().Err(err).Msgf("could not hash password")
		return nil, err
	}

	row := st.db.QueryRow("INSERT INTO users(username, password) VALUES ($1, $2) RETURNING id",
		user.Username, hash)

	err = row.Scan(&user.Id)
	if err != nil {
		log.Error().Err(err).Msgf("could not create user %s", user.Username)
		return nil, err
	}

	return user, nil
}

func (st *store) GetUserByUsername(username string) (*User, error) {
	rows, err := st.db.Query("SELECT * FROM users WHERE username = $1", username)
	if err != nil {
		log.Error().Err(err).Msgf("could not retrieve user with username %s", username)
		return nil, err
	}
	defer rows.Close()

	users := []User{}
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Username, &user.Password)
		if err != nil {
			log.Error().Err(err).Msgf("could not retrieve user with username %s", username)
			return nil, err
		}
		users = append(users, user)
	}

	if len(users) == 0 {
		return nil, nil
	}

	return &users[0], nil
}

func (st *store) Login(user *User) (token *util.JwtInfo, err error) {
	dbUser, err := st.GetUserByUsername(user.Username)
	if err != nil {
		log.Error().Err(err).Msgf("could not retrieve user %s", user.Username)
		return nil, err
	}

	match := util.CheckPasswordHash(user.Password, dbUser.Password)

	if match == false {
		log.Error().Err(err).Msgf("incorrect password")
		return nil, err
	}

	// Declare the expiration time of the token
	// here, we have kept it as 30 minutes
	expirationTime := time.Now().Add(30 * time.Minute)

	// Create the JWT claims, which includes the username and expiry time
	claims := &util.JwtClaims{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string
	tknStr, err := jwtToken.SignedString(util.JwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		log.Error().Err(err).Msgf("could not generate jwt token")
		return nil, err
	}

	jwtInfo := util.JwtInfo{
		Token:     tknStr,
		ExpiresAt: expirationTime,
	}

	return &jwtInfo, nil
}
