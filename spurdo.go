// a spurdo dranzlator paggage.
package spurdo

import (
	"errors"
	"math/rand"
	"regexp"
	"strings"
)

// tagez a sdrign to be "dranzlated" idno spurdo
func Translate(String string) (string, error) {

	if String == "" {
		return "", errors.New("spurdo: string can't be empty")
	}

	replacements := map[string]string{
		"america": "clapistan",
		"meme":    "maymay",
		"some":    "sum",
		"epic":    "ebin",
		"kek":     "geg",
		"right":   "rite",
		"ng":      "nk",
		"ic":      "ig",
		"ing":     "ign",
		"alk":     "olk",
		"ys":      "yz",
		"ws":      "wz",
		"us":      "uz",
		"ts":      "tz",
		"ss":      "sz",
		"rs":      "rz",
		"ns":      "nz",
		"ms":      "mz",
		"ls":      "lz",
		"is":      "iz",
		"gs":      "gz",
		"fs":      "fz",
		"es":      "ez",
		"ds":      "dz",
		"bs":      "bz",
		"tr":      "dr",
		"pr":      "br",
		"nt":      "dn",
		"mm":      "m",
		"lt":      "ld",
		"kn":      "gn",
		"cr":      "gr",
		"ck":      "gg",
		"va":      "ba",
		"up":      "ub",
		"pi":      "bi",
		"pe":      "be",
		"po":      "bo",
		"ot":      "od",
		"op":      "ob",
		"ke":      "ge",
		"it":      "id",
		"iv":      "ib",
		"et":      "ed",
		"ex":      "egz",
		"ev":      "eb",
		"co":      "go",
		"cx":      "gg",
		"ca":      "ga",
		"ap":      "ab",
		"af":      "ab",
		"th":      "d",
		"wh":      "w",
	}

	ebinfaces := map[int]string{
		0:  ":D",
		1:  ":DD",
		2:  ":DDD",
		3:  ":-D",
		4:  "XD",
		5:  "XXD",
		6:  "XDD",
		7:  "XXDD",
		8:  "xD",
		9:  "xDD",
		10: ":dd",
		11: ":d",
	}

	String = strings.ToLower(String)

	for search, replace := range replacements {
		String = strings.ReplaceAll(String, search, replace)
	}
	r := regexp.MustCompile(",|\n.")
	matches := r.FindAllString(String, -1)
	for _, match := range matches {
		String = strings.ReplaceAll(String, match, ebinfaces[rand.Intn(len(ebinfaces))])
	}

	return String, nil
}
