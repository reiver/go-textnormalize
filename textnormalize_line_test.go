package textnormalize



import (
	"bytes"
	"strings"
    "testing"
)



func es(s string) string {

	ss := s

	ss = strings.Replace(ss, "\t",     "\\t",     -1)
	ss = strings.Replace(ss, "\r",     "\\r",     -1)
	ss = strings.Replace(ss, "\n",     "\\n",     -1)
	ss = strings.Replace(ss, "\u0085", "\\u0085", -1) // next line
	ss = strings.Replace(ss, "\u2028", "\\u2028", -1) // line separator
	ss = strings.Replace(ss, "\u2029", "\\u2029", -1) // paragraph separator

	return ss
}

func TestEverything(t *testing.T) {

	tests := []struct {
		S []rune
		Expected []rune
	}{
		{
			[]rune("apple"),
			[]rune("apple"),
		},
		{
			[]rune(" apple"),
			[]rune(" apple"),
		},
		{
			[]rune("apple"),
			[]rune("apple"),
		},
		{
			[]rune(" apple "),
			[]rune(" apple "),
		},
		{
			[]rune("apple banana"),
			[]rune("apple banana"),
		},
		{
			[]rune("apple banana  cherry"),
			[]rune("apple banana  cherry"),
		},
		{
			[]rune("apple++banana ++ cherry"),
			[]rune("apple++banana ++ cherry"),
		},





		{
			[]rune("\n"),
			[]rune("\u2028"),
		},
		{
			[]rune(" \n"),
			[]rune(" \u2028"),
		},
		{
			[]rune("\n "),
			[]rune("\u2028 "),
		},
		{
			[]rune(" \n "),
			[]rune(" \u2028 "),
		},





		{
			[]rune("\n\n"),
			[]rune("\u2028\u2028"),
		},
		{
			[]rune(" \n\n"),
			[]rune(" \u2028\u2028"),
		},
		{
			[]rune("\n\n "),
			[]rune("\u2028\u2028 "),
		},
		{
			[]rune(" \n\n "),
			[]rune(" \u2028\u2028 "),
		},





		{
			[]rune("\n\n\n"),
			[]rune("\u2028\u2028\u2028"),
		},
		{
			[]rune(" \n\n\n"),
			[]rune(" \u2028\u2028\u2028"),
		},
		{
			[]rune("\n\n\n "),
			[]rune("\u2028\u2028\u2028 "),
		},
		{
			[]rune(" \n\n\n "),
			[]rune(" \u2028\u2028\u2028 "),
		},





		{
			[]rune("\n\n\n\n"),
			[]rune("\u2028\u2028\u2028\u2028"),
		},
		{
			[]rune(" \n\n\n\n"),
			[]rune(" \u2028\u2028\u2028\u2028"),
		},
		{
			[]rune("\n\n\n\n "),
			[]rune("\u2028\u2028\u2028\u2028 "),
		},
		{
			[]rune(" \n\n\n\n "),
			[]rune(" \u2028\u2028\u2028\u2028 "),
		},





		{
			[]rune("\u0085"),
			[]rune("\u2028"),
		},
		{
			[]rune(" \u0085"),
			[]rune(" \u2028"),
		},
		{
			[]rune("\u0085 "),
			[]rune("\u2028 "),
		},
		{
			[]rune(" \u0085 "),
			[]rune(" \u2028 "),
		},





		{
			[]rune("\u0085\u0085"),
			[]rune("\u2028\u2028"),
		},
		{
			[]rune(" \u0085\u0085"),
			[]rune(" \u2028\u2028"),
		},
		{
			[]rune("\u0085\u0085 "),
			[]rune("\u2028\u2028 "),
		},
		{
			[]rune(" \u0085\u0085 "),
			[]rune(" \u2028\u2028 "),
		},





		{
			[]rune("\u0085\u0085\u0085"),
			[]rune("\u2028\u2028\u2028"),
		},
		{
			[]rune(" \u0085\u0085\u0085"),
			[]rune(" \u2028\u2028\u2028"),
		},
		{
			[]rune("\u0085\u0085\u0085 "),
			[]rune("\u2028\u2028\u2028 "),
		},
		{
			[]rune(" \u0085\u0085\u0085 "),
			[]rune(" \u2028\u2028\u2028 "),
		},





		{
			[]rune("\u0085\u0085\u0085\u0085"),
			[]rune("\u2028\u2028\u2028\u2028"),
		},
		{
			[]rune(" \u0085\u0085\u0085\u0085"),
			[]rune(" \u2028\u2028\u2028\u2028"),
		},
		{
			[]rune("\u0085\u0085\u0085\u0085 "),
			[]rune("\u2028\u2028\u2028\u2028 "),
		},
		{
			[]rune(" \u0085\u0085\u0085\u0085 "),
			[]rune(" \u2028\u2028\u2028\u2028 "),
		},





		{
			[]rune("\r"),
			[]rune("\u2028"),
		},
		{
			[]rune(" \r"),
			[]rune(" \u2028"),
		},
		{
			[]rune("\r "),
			[]rune("\u2028 "),
		},
		{
			[]rune(" \r "),
			[]rune(" \u2028 "),
		},





		{
			[]rune("\r\r"),
			[]rune("\u2028\u2028"),
		},
		{
			[]rune(" \r\r"),
			[]rune(" \u2028\u2028"),
		},
		{
			[]rune("\r\r "),
			[]rune("\u2028\u2028 "),
		},
		{
			[]rune(" \r\r "),
			[]rune(" \u2028\u2028 "),
		},





		{
			[]rune("\r\r\r"),
			[]rune("\u2028\u2028\u2028"),
		},
		{
			[]rune(" \r\r\r"),
			[]rune(" \u2028\u2028\u2028"),
		},
		{
			[]rune("\r\r\r "),
			[]rune("\u2028\u2028\u2028 "),
		},
		{
			[]rune(" \r\r\r "),
			[]rune(" \u2028\u2028\u2028 "),
		},





		{
			[]rune("\r\r\r\r"),
			[]rune("\u2028\u2028\u2028\u2028"),
		},
		{
			[]rune(" \r\r\r\r"),
			[]rune(" \u2028\u2028\u2028\u2028"),
		},
		{
			[]rune("\r\r\r\r "),
			[]rune("\u2028\u2028\u2028\u2028 "),
		},
		{
			[]rune(" \r\r\r\r "),
			[]rune(" \u2028\u2028\u2028\u2028 "),
		},





		{
			[]rune("\r\n"),
			[]rune("\u2028"),
		},
		{
			[]rune(" \r\n"),
			[]rune(" \u2028"),
		},
		{
			[]rune("\r\n "),
			[]rune("\u2028 "),
		},
		{
			[]rune(" \r\n "),
			[]rune(" \u2028 "),
		},





		{
			[]rune("\r\n\r\n"),
			[]rune("\u2028\u2028"),
		},
		{
			[]rune(" \r\n\r\n"),
			[]rune(" \u2028\u2028"),
		},
		{
			[]rune("\r\n\r\n "),
			[]rune("\u2028\u2028 "),
		},
		{
			[]rune(" \r\n\r\n "),
			[]rune(" \u2028\u2028 "),
		},





		{
			[]rune("\r\n\r\n\r\n"),
			[]rune("\u2028\u2028\u2028"),
		},
		{
			[]rune(" \r\n\r\n\r\n"),
			[]rune(" \u2028\u2028\u2028"),
		},
		{
			[]rune("\r\n\r\n\r\n "),
			[]rune("\u2028\u2028\u2028 "),
		},
		{
			[]rune(" \r\n\r\n\r\n "),
			[]rune(" \u2028\u2028\u2028 "),
		},





		{
			[]rune("\r\n\r\n\r\n\r\n"),
			[]rune("\u2028\u2028\u2028\u2028"),
		},
		{
			[]rune(" \r\n\r\n\r\n\r\n"),
			[]rune(" \u2028\u2028\u2028\u2028"),
		},
		{
			[]rune("\r\n\r\n\r\n\r\n "),
			[]rune("\u2028\u2028\u2028\u2028 "),
		},
		{
			[]rune(" \r\n\r\n\r\n\r\n "),
			[]rune(" \u2028\u2028\u2028\u2028 "),
		},





		{
			[]rune("\r\r\n\r\n\r\u0085\u0085\r\r"),
			[]rune("\u2028\u2028\u2028\u2028\u2028\u2028\u2028\u2028"),
		},
		{
			[]rune(" \r\r\n\r\n\r\u0085\u0085\r\r"),
			[]rune(" \u2028\u2028\u2028\u2028\u2028\u2028\u2028\u2028"),
		},
		{
			[]rune("\r\r\n\r\n\r\u0085\u0085\r\r "),
			[]rune("\u2028\u2028\u2028\u2028\u2028\u2028\u2028\u2028 "),
		},
		{
			[]rune(" \r\r\n\r\n\r\u0085\u0085\r\r "),
			[]rune(" \u2028\u2028\u2028\u2028\u2028\u2028\u2028\u2028 "),
		},





		{
			[]rune("\n\r\r\n\r\n\r\u0085\u0085\r\r"),
			[]rune("\u2028\u2028\u2028\u2028\u2028\u2028\u2028\u2028\u2028"),
		},
		{
			[]rune(" \n\r\r\n\r\n\r\u0085\u0085\r\r"),
			[]rune(" \u2028\u2028\u2028\u2028\u2028\u2028\u2028\u2028\u2028"),
		},
		{
			[]rune("\n\r\r\n\r\n\r\u0085\u0085\r\r "),
			[]rune("\u2028\u2028\u2028\u2028\u2028\u2028\u2028\u2028\u2028 "),
		},
		{
			[]rune(" \n\r\r\n\r\n\r\u0085\u0085\r\r "),
			[]rune(" \u2028\u2028\u2028\u2028\u2028\u2028\u2028\u2028\u2028 "),
		},
	}



	for iii := 0 ; iii < len(tests) ; iii++ {


		// Store the output of the Lexer in 'outputBuffer'.
			doneCh := make(chan bool)

			outputBuffer := bytes.NewBuffer(nil)

			ch := make(chan rune)

			go func() {
				for r,ok := <-ch ; ok ; r,ok = <-ch {

					outputBuffer.WriteRune(r)
				}

				doneCh <- true
			}()



		// Expected
			datum := tests[iii]

			expected := datum.Expected



		// Lexer
			lex := NewLineLex(ch)



		// Write the string of characters -- the rune array -- into the Lexer.
			for i := 0 ; i < len(datum.S) ; i++ {

				lex.WriteRune(datum.S[i])

			} // for

			lex.WriteEof()

			// Wait until the writing has fully propagated through.
			<-doneCh



		// Verify that the length of the tokens we got from the Lexer is the number we expect.
			if len(expected) != len([]rune(outputBuffer.String())) {
				t.Errorf("Expected [%v] rune(s), but got [%v] instead, for [%v]. Namely: [%v].", len(expected), outputBuffer.Len(), es(string(datum.S)), es(outputBuffer.String()))
	/////////// CONTINUE
				continue
			}



		// Verify that the the output from the Lexer is what we expect.
			if string(expected) != outputBuffer.String() {
				t.Errorf("Expected [%v], but got [%v] instead.", es(string(expected)), es(outputBuffer.String()))
	/////////// CONTINUE
				continue
			}



		// Verify that get expected results using the NormalizeLineSeparators() helper function.
			normalizedRunes := NormalizeLineSeparators(datum.S)

			if string(expected) != string(normalizedRunes) {
				t.Errorf("With NormalizeLineSeparators(), expected [%v], but got [%v] instead.", es(string(expected)), es(string(normalizedRunes)))
	/////////// CONTINUE
				continue
			}



		// Verify that get expected results using the NormalizeLineSeparatorsString() helper function.
			normalizedString := NormalizeLineSeparatorsString( string(datum.S) )

			if string(expected) != normalizedString {
				t.Errorf("With NormalizeLineSeparatorsString(), expected [%v], but got [%v] instead.", es(string(expected)), es(normalizedString))
	/////////// CONTINUE
				continue
			}


	} // for
}

