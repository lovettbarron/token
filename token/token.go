/*


*/

package token

import (
	"time"
	"sync"
)

type TokenEntry struct {
	entryid int64
	tokenid int64
	timestamp int64
	value float32
}

type Token struct {
	tokenid int64
	userid int64
	title string
	cycle Cycle
	tokens []*TokenEntry
	mutex *sync.Mutex
}


/////// TOKEN ///////
// Make new Token
func NewToken(title string, cycle *Cycle) *Token {
	return &Token{
		0,
		0,
		title,
		*cycle,
		make([]*TokenEntry, 0),
		new(sync.Mutex),
	}
}


/////// ENTRIES ///////
// Make Entry
func (t *Token) UseToken(user User, value float32) *TokenEntry {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	newId := int64(len(user.tokens))
	currentTime := time.Now().Unix()

	entry := &TokenEntry{
		newId,
		user.userid,
		currentTime,
		value,
	}

	t.tokens = append(t.tokens,entry)
	return t.tokens[len(t.tokens)]
}

// Get All Token Entries
func (t *Token) GetAllTokenEntries() []*TokenEntry{
	return nil
}

// Get most recent entry 
func (t *Token) GetMostRecentEntry() *TokenEntry {
	return t.tokens[len(t.tokens)]
}

// Get entry (id)
func (t *Token) GetEntry(entryid int64) *TokenEntry {
	if(entryid < 0 || entryid >= int64(len(t.tokens))) { return nil }
	return t.tokens[len(t.tokens)-1]
}

// Delete all Entries
func (t *Token) RemoveAllEntries() int {
	// Does this work?
	t.tokens = make([]*TokenEntry, 0)
	return 0
}

// Delete Last Entry (Undo)
func (t *Token) RemoveLastEntry() int {

	return 0
}

// Delete an entry (id)
func (t *Token) RemoveEntry(entryid int32) int {
	return 0
}