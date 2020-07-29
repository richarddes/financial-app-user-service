package handler

// func HandleForgotPassResetPass() http.HandlerFunc {
// 	type reqBody struct {
// 		Email             string `json:"email"`
// 		SecretToken       string `json:"secretToken"`
// 		NewPass           string `json:"newPass"`
// 		NewPassRepetition string `json:"newPassRepetition"`
// 	}

// 	return func(w http.ResponseWriter, r *http.Request) {
// 		var body reqBody
// 		err := ParseJSONBody(r.Body, &body)

// 		if err != nil {
// 			log.Panic(err)
// 			w.WriteHeader(http.StatusInternalServerError)
// 			return
// 		}

// 		if body.NewPass != body.NewPassRepetition {
// 			w.WriteHeader(http.StatusBadRequest)
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

// 		selectSecretTokenStmt := "SELECT uid,token FROM forgotPassTokens WHERE email=?;"
// 		row := db.QueryContextRowContext(selectSecretTokenStmt, body.Email)

// 		var (
// 			uid     string
// 			savedToken string
// 		)

// 		if err = row.Scan(&uid, &savedToken); err != nil {
// 			log.Panic(err)
// 			w.WriteHeader(http.StatusInternalServerError)
// 			return
// 		}

// 		err = bcrypt.CompareHashAndPassword([]byte(savedToken), []byte(body.SecretToken))
// 		if err != nil {
// 			log.Panic(err)
// 			w.WriteHeader(http.StatusUnauthorized)
// 			return
// 		}

// 		changePassStmt := `UPDATE users
// 												SET pass=?
// 												WHERE uid=?;`

// 		stmt, err := db.Prepare(changePassStmt)
// 		if err != nil {
// 			log.Panic(err)
// 			w.WriteHeader(http.StatusInternalServerError)
// 			return
// 		}
// 		_, err = stmt.Exec(body.NewPass, uid)
// 		if err != nil {
// 			log.Panic(err)
// 			w.WriteHeader(http.StatusInternalServerError)
// 			return
// 		}
// 	}
// }
