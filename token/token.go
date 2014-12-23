/*


*/

package token

import (
	"time"
	"sync"
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
	tokens []*TokenEntry
	mutex *sync.Mutex
}


/////// TOKEN ///////
// Make new Token
func NewToken(title string, cycle int32) *Token {
	return &Token{
		0,
		0,
		title,
		cycle,
		make([]*TokenEntry, 0),
		new(sync.Mutex),
	}
}


/////// ENTRIES ///////
// Make Entry
func (t *Token) UseToken(value float32) *TokenEntry {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	newId := len(t.tokens)
	currentTime := time.UTC(time.Now())

	entry := &TokenEntry{
		newId,
		t.id,
		currentTime,
		value,
	}

	t.tokens = append(t.tokens,entry)
	return t.tokens[len(t.tokens)]
}

// Get All Token Entries
func (t *Token) GetAllTokenEntries() []*TokenEntry{

}

// Get most recent entry 
func (t *Token) GetMostRecentEntry() *TokenEntry {
	return t.tokens[len(t.tokens)]
}

// Get entry (id)
func (t *Token) getEntry(entryid int32) *TokenEntry {
	if(entryid < 0 || entryid >= len(t.tokens)) { return nil }
	return t.tokens[len(t.tokens)]
}

// Delete all Entries
func (t *Token) RemoveAllEntries() int {
	// Does this work?
	t.tokens = make([]*TokenEntry, 0)
}

// Delete Last Entry (Undo)
func (t *Token) RemoveLastEntry() int {

	return -1
}

// Delete an entry (id)
func (t *Token) RemoveEntry(entryid int32) int {

}