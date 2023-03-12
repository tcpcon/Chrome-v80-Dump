package chromium

type (
	Browser struct {
		Accounts      []Account
		Cookies       []Cookie
		WebHistory    []WebHistory
		SearchHistory []SearchHistory
		CreditCards   []CreditCard
		Autofill      []Autofill
	}

	Account struct {
		Origin, Action, User, Pass, Date, LastUsed string
	}

	Cookie struct {
		Host, Name, Path, Value, Expires, IsSecure string
	}

	WebHistory struct {
		Url, Title, VisitCount, LastVisit string
	}

	SearchHistory struct {
		Term string
	}

	CreditCard struct {
		Number, Name, ExpirationMonth, ExpirationYear string
	}

	Autofill struct {
		Name, Value, Created, LastUsed string
	}
)
