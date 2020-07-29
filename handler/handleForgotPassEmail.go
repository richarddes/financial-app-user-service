package handler

// func HandleForgotPassEmail() http.HandlerFunc {
// 	type reqBody struct {
// 		Email string `json:"email"`
// 	}

// 	return func(w http.ResponseWriter, r *http.Request) {
// 		var body reqBody
// 		err := ParseJSONBody(r.Body, &body)

// 		if err != nil {
// 			log.Panic(err)
// 			w.WriteHeader(http.StatusInternalServerError)
// 			return
// 		}

// 		emailValid, err := ValidateEmail(body.Email)
// 		if err != nil {
// 			log.Panic(err)
// 			w.WriteHeader(http.StatusInternalServerError)
// 			return
// 		}

// 		if !emailValid {
// 			w.WriteHeader(http.StatusUnauthorized)
// 			return
// 		}

// 		addTokenStmt := `INSERT INTO forgotPassTokens
// 											VALUES(SELECT uid  WHERE email=?, ?)`

// 		stmt, err := db.Prepare(addTokenStmt)
// 		if err != nil {
// 			log.Panic(err)
// 			w.WriteHeader(http.StatusInternalServerError)
// 			return
// 		}

// 		secretToken := GenRandomInt64()

// 		passHash, err := bcrypt.GenerateFromPassword([]byte(secretToken), bcrypt.DefaultCost)
// 		if err != nil {
// 			log.Panic(err)
// 			w.WriteHeader(http.StatusInternalServerError)
// 			return
// 		}

// 		_, err = stmt.Exec(body.Email, string(passHash))
// 		if err != nil {
// 			log.Panic(err)
// 			w.WriteHeader(http.StatusInternalServerError)
// 			return
// 		}
// 	}
// }
