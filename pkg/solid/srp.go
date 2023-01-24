package solid

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
)

// 文字列helloを返します。
func Hello() string {
	return "hello"
}

/*

*/

var entryCount = 0

type Journal struct {
	entries []string
}

// Jounalの構造体のポインタを返却
func NewJounal(entiry []string) *Journal {
	return &Journal{
		entries: entiry,
	}
}
// 出力: Jounalのentriesを改行して表示
func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

// 構造体のJounalにカウントと文字列を入れて、カウントを戻り値で返す
func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s",
		entryCount,
		text)
	j.entries = append(j.entries, entry)
	return entryCount
}

func (j *Journal) RemoveEntry(index int) {
	// ...
}

// 
func (j *Journal) Save(filename string) {
	_ = ioutil.WriteFile(filename,
		[]byte(j.String()), 0644)
}

func (j *Journal) Load(filename string) {

}

func (j *Journal) LoadFromWeb(url *url.URL) {

}

var lineSeparator = "\n"
func SaveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename,
		[]byte(strings.Join(j.entries, lineSeparator)), 0644)
}

type Persistence struct {
	lineSeparator string
}

func (p *Persistence) saveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename,
		[]byte(strings.Join(j.entries, p.lineSeparator)), 0644)
}


