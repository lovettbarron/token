/*


*/

package token

import (
	"time"
)

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
	tokens []*TokenEntry
	mutex *sync.Mutex
}


/////// TOKEN ///////
// Make new Token
func NewToken() *Token {
	return &Token{
		make([]*TokenEntry, 0),
		new(sync.Mutex),
	}
}


/////// ENTRIES ///////
// Make Entry
func (t *Token) UseToken(value float32) *TokenEntry {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	newId := len(g.tokens)
	currentTime := time.UTC(time.Now())

	entry := &TokenEntry{
		newId,
		g.id,
		currentTime,
		value,
	}

	g.tokens = append(g.tokens,entry)
	return g.tokens[len(g.tokens)]
}

// Get All Token Entries
func (t *Token) GetAllTokenEntries() []*TokenEntry{

}

// Get most recent entry 
func (t *Token) GetMostRecentEntry() *TokenEntry {
	return g.tokens[len(g.tokens)]
}

// Get entry (id)
func (t *Token) getEntry(entryid int32) *TokenEntry {
	if(entryid < 0 || entryid >= len(g.tokens)) return nil
	return g.tokens[len(g.tokens)]
}

// Delete all Entries
func (t *Token) RemoveAllEntries() int {
	// Does this work?
	g.tokens := make([]*TokenEntry, 0)
}

// Delete Last Entry (Undo)
func (t *Token) RemoveLastEntry() int {

	return -1
}

// Delete an entry (id)
func (t *Token) RemoveEntry(entryid int32) int {

}