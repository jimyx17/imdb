package imdb

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// A Name represents an IMDb name (actor, director, writer, etc.).
type Name struct {
	ID        string `json:",omitempty"`
	URL       string `json:",omitempty"`
	FullName  string `json:",omitempty"`
	Profile   string `json:",omitempty"`
	Biography string `json:",omitempty"`
	Birthday  string `json:",omitempty"`
}

// String formats a Name.
func (n *Name) String() string {
	return fmt.Sprintf("IMDb %s: %s", n.ID, n.FullName)
}

var nmRE = regexp.MustCompile(`^nm\d+$`)

const nameURL = "https://www.imdb.com/name/%s"
const bioURL = "https://www.imdb.com/name/%s/bio"

// NewName gets, parses and returns a Name by its ID.
func NewName(c *http.Client, id string) (*Name, error) {
	if !nmRE.MatchString(id) {
		return nil, ErrInvalidID
	}
	resp, err := c.Get(fmt.Sprintf(nameURL, id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	page, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	bioresp, err := c.Get(fmt.Sprintf(bioURL, id))
	if err != nil {
		return nil, err
	}
	defer bioresp.Body.Close()
	biopage, err := ioutil.ReadAll(bioresp.Body)
	if err != nil {
		return nil, err
	}

	n := Name{}
	if err := n.Parse(page, biopage); err != nil {
		return nil, err
	}
	return &n, nil
}

// Regular expressions to parse a Name.
var (
	nameIDRE       = regexp.MustCompile(`<link rel="canonical" href="https://www.imdb.com/name/(nm\d+)/"`)
	nameFullNameRE = regexp.MustCompile(`<meta property='og:title' content="(.*?)"`)
	profileRE      = regexp.MustCompile(`<meta property='og:image' content="([\w\.\/\:\-]+@@).*.jpg"`)
	bioRE          = regexp.MustCompile(`<h4 class="li_group">Mini Bio \(\d\)</h4>\n +<div [^>]+>\n +<p>\n +(.*)\n +</p>`)
	birthRE        = regexp.MustCompile(`(?s) +<div id="name-born-info" class="txt-block">\n.*<time datetime="([\d-]+)">\n`)
	tagsRE         = regexp.MustCompile(`(<[^>]+>)`)
)

// Parse parses a Name from its page.
func (n *Name) Parse(page, biopage []byte) error {
	// ID, URL
	s := nameIDRE.FindSubmatch(page)
	if s == nil {
		return NewErrParse("id")
	}
	n.ID = string(s[1])
	n.URL = fmt.Sprintf(nameURL, n.ID)

	// FullName
	s = nameFullNameRE.FindSubmatch(page)
	if s == nil {
		return NewErrParse("full name")
	}
	if len(s[1]) == 0 {
		return NewErrParse("full name empty")
	}
	n.FullName = decode(string(s[1]))
	s = profileRE.FindSubmatch(page)
	if s != nil && len(s[1]) > 0 {
		n.Profile = string(s[1]) + ".jpg"
	}

	s = birthRE.FindSubmatch(page)
	if s != nil && len(s[1]) > 0 {
		n.Birthday = string(s[1])
	}

	s = bioRE.FindSubmatch(biopage)
	if s == nil || len(s[1]) == 0 {
		return nil
	}
	rs := tagsRE.ReplaceAll(s[1], []byte{})
	n.Biography = string(rs)

	return nil
}
