package shortener

import (
	"testing"
)

func TestGenerateShortLink(t *testing.T) {
	testTable := []map[string]string{{
		"link":   "https://www.linkedin.com/in/afistapratama",
		"expect": "BxjspBB8",
	}, {
		"link":   "https://www.google.com/search?q=impact+byte&rlz=1C1GCEA_enID937ID937&oq=impact+byte&aqs=chrome..69i57j35i39j0j0i22i30l7.1454j1j15&sourceid=chrome&ie=UTF-8",
		"expect": "ZgqphFEB",
	}, {
		"link":   "https://www.google.com/search?q=afista+pratama&rlz=1C1GCEA_enID937ID937&oq=afista+pratama&aqs=chrome..69i57j46i13i175i199j0i8i13i30j69i60l3.2077j1j15&sourceid=chrome&ie=UTF-8",
		"expect": "bcxL8FzG",
	}}

	for _, test := range testTable {
		result := GenerateShortLink(test["link"], "")

		if result != test["expect"] {
			t.Fail()
		}
	}
}

func TestCheckingValidLink(t *testing.T) {
	testTable := []map[string]string{{
		"link":   "google.com",
		"expect": "http://www.google.com",
	}, {
		"link":   "www.google.com",
		"expect": "http://www.google.com",
	}, {
		"link":   "https://google.com",
		"expect": "https://www.google.com",
	}, {
		"link":   "http://google.com",
		"expect": "http://www.google.com",
	}}

	for _, test := range testTable {
		result := CheckingValidLink(test["link"])

		if result != test["expect"] {
			t.Fail()
		}
	}
}
