// #############################################################################
// # File: navigation.go                                                       #
// # Project: epub                                                             #
// # Created Date: 2023/09/12 09:12:44                                         #
// # Author: realjf                                                            #
// # -----                                                                     #
// # Last Modified: 2023/09/12 09:13:03                                        #
// # Modified By: realjf                                                       #
// # -----                                                                     #
// # Copyright (c) 2023 realjf                                                 #
// #############################################################################
package epub

import (
	"encoding/xml"
	"fmt"
)

// ============================================ nav.xhtml ==============================================

// This holds the body XML for the EPUB v3 TOC file (nav.xhtml). Since this is
// an XHTML file, the rest of the structure is handled by the xhtml type
//
// Sample: https://github.com/bmaupin/epub-samples/blob/master/minimal-v3plus2/EPUB/nav.xhtml
// Spec: http://www.idpf.org/epub/301/spec/epub-contentdocs.html#sec-xhtml-nav
type tocNavBody struct {
	XMLName  xml.Name     `xml:"nav"`
	EpubType string       `xml:"epub:type,attr"`
	H1       string       `xml:"h1"`
	Links    []tocNavItem `xml:"ol>li"`
}

// Constructor for tocNavBody
func newTocNavXML() *tocNavBody {
	b := &tocNavBody{
		EpubType: tocNavEpubType,
	}
	err := xml.Unmarshal([]byte(tocNavBodyTemplate), &b)
	if err != nil {
		panic(fmt.Sprintf(
			"Error unmarshalling tocNavBody: %s\n"+
				"\ttocNavBody=%#v\n"+
				"\ttocNavBodyTemplate=%s",
			err,
			*b,
			tocNavBodyTemplate))
	}

	return b
}

type tocNavItem struct {
	A        tocNavLink    `xml:"a"`
	Children *[]tocNavItem `xml:"ol>li,omitempty"`
}

type tocNavLink struct {
	XMLName xml.Name `xml:"a"`
	Href    string   `xml:"href,attr"`
	Data    string   `xml:",chardata"`
}
