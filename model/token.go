/*


*/

package token

type TokenEntry struct {
	entryid int32
	tokenid int32
	timestamp int32
	value float32
}

type Token struct {
	tokenid int32
	userid int32
	title string
	cycle int32
	image string
	entries []*TokenEntry
}


/////// TOKEN ///////
// Make new Token
func NewToken() *Token {

}

// Get List of All Tokens
func (t *Token) GetAllTokens() []*Token {

}

// Get most recently made Token (top)
func (t *Token) GetLastToken() *TokenEntry {

}

// Delete all Tokens
func (t *Token) RemoveAllTokens() int {

}

// Delete token
func (t *Token) RemoveToken(tokenid int32) int {

}


/////// ENTRIES ///////
// Make Entry
func (t *Token) UseToken(value float32) *TokenEntry {

}

// Get All Token Entries
func (t *Token) GetAllTokenEntries() []*TokenEntry {

}

// Get most recent entry 
func (t *Token) GetMostRecentEntry() *TokenEntry {

}

// Get entry (id)
func (t *Token) getEntry(entryid int32) *TokenEntry {

}

// Delete all Entries
func (t *Token) RemoveAllEntries() int {

}

// Delete Last Entry (Undo)
func (t *Token) RemoveLastEntry() int {

}

// Delete an entry (id)
func (t *Token) RemoveEntry(entryid int32) int {

}