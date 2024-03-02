package main

import "testing"

func TestAdfgvx(t *testing.T) {
	testCases := []struct {
		key1             string
		key2             string
		plainText        string
		expected         string
		permutation      string
		transpositionKey string
	}{
		{
			key1:             "GRAPHY",
			key2:             "CRYPTO",
			permutation:      "PH0QG64MEA1YL2NOFDXKR3CVS5ZW7BJ9UTI8",
			plainText:        "DEFENDTHEEASTWALLOFTHECASTLE",
			expected:         "YYRTRYCRYXYAPAPGHAGAYROROYYRTRAARPPGPHPXHYAGPPRPGXRYCRTYOPOX",
			transpositionKey: "GERMAN",
		},
		{
			key1:             "GRAPHY",
			key2:             "CRYPTO",
			permutation:      "J9I1E8D2Z6Y3M0C7K5LQOX4SFTUAGHNRWPBV",
			plainText:        "TEXTTOENCRYPT",
			expected:         "TTCOTPPYOXHRGHXRRHRRCTORXPAAPX",
			transpositionKey: "CIPHER",
		},
	}
	for _, tC := range testCases {
		t.Run("ADFGVX", func(t *testing.T) {
			encrypted := modifiedAdfgvxCipher(tC.key1, tC.key2, tC.permutation, tC.plainText, tC.transpositionKey)
			if encrypted != tC.expected {
				t.Errorf("received %s, wanted %s", encrypted, tC.expected)
			}
		})
	}
}


func TestColumnarTransposition(t *testing.T) {
	testCases := []struct {
		key       string
		plainText string
		expected  string
	}{
		{
			key:       "FOURTH",
			plainText: "WEHAVEBEENFOUNDESCAPE",
			expected:  "WBUAEOCXEENPANEXVFSXHEDE",
		},
		{
			key:       "FI",
			plainText: "WEHAVEBEENFOUNDESCAPE",
			expected:  "WHVBEFUDSAEEAEENONECPX",
		},
		{
			key:       "FOOUR",
			plainText: "WEHAVEBEENFOUNDESCAPE",
			expected:  "WEFEEEBOSXHEUCXVNDPXAENAX",
		},
	}
	for _, tC := range testCases {
		t.Run("Columnar", func(t *testing.T) {
			encrypted := ColumnarTransposition(tC.key, tC.plainText)
			if encrypted != tC.expected {
				t.Errorf("received %s, wanted %s", encrypted, tC.expected)
			}
		})
	}
}
