package textnormalize



import "bytes"



func normalizeLines(s []rune) *bytes.Buffer {

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



	// Lexer
		lex := NewLineLex(ch)



	// Write the string of characters -- the rune array -- into the Lexer.
		for i := 0 ; i < len(s) ; i++ {

			lex.WriteRune(s[i])

		} // for

		lex.WriteEof()

		// Wait until the writing has fully propagated through.
		<-doneCh



	// Return
		return outputBuffer
}

// NormalizeLineSeparatorsString normalizes the *line separators* in UNICODE text,
// by changing all the different ways of specifying a *line separation* into '\u2028'.
// This includes:
// "\r\n"   → "\u2028",
// "\r"     → "\u2028",
// "\n"     → "\u2028",
// "\u0085" → "\u2028".
func NormalizeLineSeparatorsString(s string) string {

	outputBuffer := normalizeLines(  []rune(s)  )

	return outputBuffer.String()
}

// NormalizeLineSeparators normalizes the *line separators* in UNICODE text,
// by changing all the different ways of specifying a *line separation* into '\u2028'.
// This includes:
// "\r\n"   → "\u2028",
// "\r"     → "\u2028",
// "\n"     → "\u2028",
// "\u0085" → "\u2028".
func NormalizeLineSeparators(s []rune) []rune {

	outputBuffer := normalizeLines( s )

	return []rune(outputBuffer.String())
}



// A LineLex is a very simple *lexical analyzer*, that will normalize *line separation* in UNICODE text,
// by changing all the different ways of specifying a *line separation* into '\u2028'.
// This includes:
// "\r\n"   → "\u2028",
// "\r"     → "\u2028",
// "\n"     → "\u2028",
// "\u0085" → "\u2028".
type LineLex struct {
	canonicalLineSeparator rune
	previousRuneWasCarriageReturn bool
	outputChannel chan<- rune
}



// NewLineLex returns a LineNex, that will normalize the *line separation* in UNICODE text,
// by changing all the different ways of specifying a *line separation* into '\u2028'.
// This includes:
// "\r\n"   → "\u2028",
// "\r"     → "\u2028",
// "\n"     → "\u2028",
// "\u0085" → "\u2028".
func NewLineLex(c chan<- rune) (*LineLex) {

	me := LineLex{
		canonicalLineSeparator        : '\u2028',
		previousRuneWasCarriageReturn : false,
		outputChannel                 : c,
	}

	return &me
}



func (me *LineLex) escape(r rune) rune {
	result := r



	switch r {
		case
			'\r',
			'\n',
			'\u0085':
			result = me.canonicalLineSeparator
	}



	return result
}

// WriteRune passes a rune to a LineLex.
func (me *LineLex) WriteRune(r rune) {

	previousRuneWasCarriageReturn := me.previousRuneWasCarriageReturn



	me.previousRuneWasCarriageReturn = ('\r' == r)



	switch {
		case !previousRuneWasCarriageReturn && '\r' == r:
			// Nothing here.

		case !previousRuneWasCarriageReturn && '\r' != r:
			me.outputChannel <- me.escape(r)

		case previousRuneWasCarriageReturn && '\r' == r:
			me.outputChannel <- me.canonicalLineSeparator

		case previousRuneWasCarriageReturn && '\r' != r && '\n' == r:
			me.outputChannel <- me.canonicalLineSeparator

		case previousRuneWasCarriageReturn && '\r' != r && '\n' != r:
			me.outputChannel <- me.canonicalLineSeparator
			me.outputChannel <- me.escape(r)

		default:
			me.outputChannel <- me.escape(r)
	}

}

// WriteEof closes a LineLex.
func (me *LineLex) WriteEof() {

	if me.previousRuneWasCarriageReturn {
		me.outputChannel <- me.canonicalLineSeparator
	}

	me.previousRuneWasCarriageReturn = false

	close(me.outputChannel)

}
