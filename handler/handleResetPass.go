package handler

// func HandleResetPass(env *config.Env) http.HandlerFunc {
// 	type reqBody struct {
// 		CurrentPass       string `json:"currentPass"`
// 		NewPass           string `json:"newPass"`
// 		NewPassRepetition string `json:"newPassRepetition"`
// 	}

// 	return func(w http.ResponseWriter, r *http.Request) {
// 		uid, err := auth.ValidateAuthCookie(r)
// 		if err != nil {
// 			status := helper.CheckForUnauthorizedHeader(err)
// 			w.WriteHeader(status)
// 			return
// 		}

// 		var body reqBody
// 		err = helper.ParseJSONBody(r.Body, &body)

// 		if err != nil {
// 			log.Panic(err)
// 			w.WriteHeader(http.StatusInternalServerError)
// 			return
// 		}

// 		if body.NewPass != body.NewPassRepetition {
// 			w.WriteHeader(http.StatusBadRequest)
// 			return
// 		}

// 		selectUserPassStmt := `SELECT pass
// 														FROM users
// 														WHERE uid=?;`

// 		row := env.db.QueryContextRowContext(selectUserPassStmt, uid)

// 		var savedPass string

// 		if err = row.Scan(&savedPass); err != nil {
// 			log.Panic(err)
// 			w.WriteHeader(http.StatusInternalServerError)
// 			return
// 		}

// 		err = bcrypt.CompareHashAndPassword([]byte(savedPass), []byte(body.CurrentPass))
// 		if err != nil {
// 			log.Panic(err)
// 			w.WriteHeader(http.StatusUnauthorized)
// 			return
// 		}

// 		changePassStmt := `UPDATE users
// 												SET pass=?
// 												WHERE uid=?;`

// 		stmt, err := env.DB.Prepare(changePassStmt)
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
