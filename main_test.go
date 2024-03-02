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

func TestRailFenceDecryption(t *testing.T) {
	testCases := []struct {
		cipherText string
		expected   string
		rails      int
	}{
		{
			rails:      3,
			cipherText: "EPWSAEOCN",
			expected:   "ESCAPENOW",
		},
		{
			rails:      4,
			cipherText: "ENSEOCPWA",
			expected:   "ESCAPENOW",
		},
	}
	for _, tC := range testCases {
		t.Run("Rail Fence", func(t *testing.T) {
			decoded := RailFenceDecryption(tC.rails, tC.cipherText)
			if decoded != tC.expected {
				t.Errorf("received %s, wanted %s", decoded, tC.expected)
			}
		})
	}
}

func TestStraddlingCheckerboardDecryption(t *testing.T) {
	testCases := []struct {
		straddlingKey string
		expected      string
		cipherText    string
		num1          int
		num2          int
	}{
		{
			straddlingKey: "FKMCPDYEHBIGQROSAZLUTJNWVX",
			cipherText:    "690974672309938377275387070360723094383772709",
			num1:          3,
			num2:          7,
			expected:      "DEFENDTHEEASTWALLOFTHECASTLE",
		},
	}
	for _, tC := range testCases {
		t.Run("Straddling Checkerboard", func(t *testing.T) {
			decoded := StraddlingCheckerboardDecryption(tC.straddlingKey, tC.cipherText, tC.num1, tC.num2)
			if decoded != tC.expected {
				t.Errorf("received %s, wanted %s", decoded, tC.expected)
			}
		})
	}
}

func TestQ3(t *testing.T) {
	testCases := []struct {
		straddlingKey string
		expected      string
		cipherText    string
		num1          int
		num2          int
		rails         int
	}{
		{
			straddlingKey: "XZDECAMRQKUYBLFOGVITWJHPSN",
			num1:          2, num2: 7,
			cipherText: "377767272277661122967521077712672277",
			rails:      3,
			expected:   "DONTASKTAFORSOLUTION",
		},
	}
	for _, tC := range testCases {
		t.Run("Q3", func(t *testing.T) {
			decoded := Q3(tC.straddlingKey, tC.num1, tC.num2, tC.cipherText, tC.rails)
			if decoded != tC.expected {
				t.Errorf("received %s, wanted %s", decoded, tC.expected)
			}
		})
	}
}

func TestPlayFair(t *testing.T) {
	testCases := []struct {
		key       string
		plainText string
		expected  string
	}{
		{
			key:       "PLAYFAIREXAMPLE",
			plainText: "HIDETHEGOLD",
			expected:  "BMODZBXDNAGE",
		},
		{
			key:       "PLAYFAIREXAMPLE",
			plainText: "HIDETHEGOLDINTHETREESTUMP",
			expected:  "BMODZBXDNABEKUDMUIXMMOUVIF",
		},
		{
			key:       "TRICIPHER",
			plainText: "TESTCODE",
			expected:  "RHNCRSHA",
		},
	}
	for _, tC := range testCases {
		t.Run("Playfair", func(t *testing.T) {
			encrypted := PlayFairCipher(tC.key, tC.plainText)
			if encrypted != tC.expected {
				t.Errorf("received %s, wanted %s", encrypted, tC.expected)
			}
		})
	}
}

func TestVigenere(t *testing.T) {
	testCases := []struct {
		key       string
		plainText string
		expected  string
	}{
		{
			key:       "LEMON",
			plainText: "ATTACKATDAWN",
			expected:  "LXFOPVEFRNHR",
		},
		{
			key:       "CRYPTII",
			plainText: "THEQUICKBROWNFOXJUMPSOVERLAZYDOGS",
			expected:  "VYCFNQKMSPDPVNQOHJFXAQMCGEIHAUMVL",
		},
	}
	for _, tC := range testCases {
		t.Run("Vigenere", func(t *testing.T) {
			encrypted := VigenereCipher(tC.key, tC.plainText)
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

func TestPlayFairDecryption(t *testing.T) {
	testCases := []struct {
		key        string
		cipherText string
		expected   string
	}{
		{
			key:        "PLAYFAIREXAMPLE",
			cipherText: "BMODZBXDNAGE",
			expected:   "HIDETHEGOLDX",
		},
	}
	for _, tC := range testCases {
		t.Run("Playfair Decryption", func(t *testing.T) {
			decrypted := PlayFairDecryption(tC.key, tC.cipherText)
			if decrypted != tC.expected {
				t.Errorf("received %s, wanted %s", decrypted, tC.expected)
			}
		})
	}
}

func TestVigenereDecryption(t *testing.T) {
	testCases := []struct {
		key        string
		cipherText string
		expected   string
	}{
		{
			key:        "LEMON",
			cipherText: "LXFOPVEFRNHR",
			expected:   "ATTACKATDAWN",
		},
		{
			key:        "KEY",
			cipherText: "NGMNIAKRBOGPITRFMEORCBIUSXFYYRUIW",
			expected:   "DCODECANDECRYPTVIGENEREWITHOUTKEY",
		},
	}
	for _, tC := range testCases {
		t.Run("Vigenere Decryption", func(t *testing.T) {
			decrypted := VigenereDecryption(tC.key, tC.cipherText)
			if decrypted != tC.expected {
				t.Errorf("received %s, wanted %s", decrypted, tC.expected)
			}
		})
	}
}

func TestColumnarDecryption(t *testing.T) {
	testCases := []struct {
		key        string
		cipherText string
		expected   string
	}{
		{
			key:        "YYAS",
			cipherText: "AOIXNSOXTSINRPTX",
			expected:   "TRANSPOSITION",
		},
		{
			key:        "FI",
			expected:   "WEHAVEBEENFOUNDESCAPE",
			cipherText: "WHVBEFUDSAEEAEENONECPX",
		},
		{
			key:        "FOOUR",
			expected:   "WEHAVEBEENFOUNDESCAPE",
			cipherText: "WEFEEEBOSXHEUCXVNDPXAENAX",
		},
		{
			key:        "FOURTH",
			expected:   "WEHAVEBEENFOUNDESCAPE",
			cipherText: "WBUAEOCXEENPANEXVFSXHEDE",
		},
	}
	for _, tC := range testCases {
		t.Run("Columnar Decryption", func(t *testing.T) {
			decrypted := ColumnarDecryption(tC.key, tC.cipherText)
			if decrypted != tC.expected {
				t.Errorf("received %s, wanted %s", decrypted, tC.expected)
			}
		})
	}
}

func TestCrypticConnection(t *testing.T) {
	testCases := []struct {
		playfairKey string
		vigenerKey  string
		columnarKey string
		cipherText  string
		expected    string
	}{
		{
			"TRICIPHER",
			"CODEHELP",
			"FINAL",
			"GXTWVSYXQP",
			"TESTCODE",
		},
	}
	for _, tC := range testCases {
		t.Run("Q4", func(t *testing.T) {
			decrypted := CrypticConnection(tC.playfairKey, tC.vigenerKey, tC.columnarKey, tC.cipherText)
			if decrypted != tC.expected {
				t.Errorf("received %s, wanted %s", decrypted, tC.expected)
			}
		})
	}
}
