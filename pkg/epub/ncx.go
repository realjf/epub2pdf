// #############################################################################
// # File: ncx.go                                                              #
// # Project: epub                                                             #
// # Created Date: 2023/09/12 09:11:02                                         #
// # Author: realjf                                                            #
// # -----                                                                     #
// # Last Modified: 2023/09/12 09:12:15                                        #
// # Modified By: realjf                                                       #
// # -----                                                                     #
// # Copyright (c) 2023 realjf                                                 #
// #############################################################################
package epub

import (
	"encoding/xml"
	"fmt"
)

// ============================================ toc.ncx ==============================================

// This holds the XML for the EPUB v2 TOC file (toc.ncx). This is added so the
// resulting EPUB v3 file will still work with devices that only support EPUB v2
//
// Sample: https://github.com/bmaupin/epub-samples/blob/master/minimal-v3plus2/EPUB/toc.ncx
// Spec: http://www.idpf.org/epub/20/spec/OPF_2.0.1_draft.htm#Section2.4.1
type tocNcxRoot struct {
	XMLName xml.Name         `xml:"http://www.daisy.org/z3986/2005/ncx/ ncx"`
	Version string           `xml:"version,attr"`
	Meta    tocNcxMeta       `xml:"head>meta"`
	Title   string           `xml:"docTitle>text"`
	Author  string           `xml:"docAuthor>text"`
	NavMap  []tocNcxNavPoint `xml:"navMap>navPoint"`
}

// Constructor for tocNcxRoot
func newTocNcxXML() *tocNcxRoot {
	n := &tocNcxRoot{}

	err := xml.Unmarshal([]byte(tocNcxTemplate), &n)
	if err != nil {
		panic(fmt.Sprintf(
			"Error unmarshalling tocNcxRoot: %s\n"+
				"\ttocNcxRoot=%#v\n"+
				"\ttocNcxTemplate=%s",
			err,
			*n,
			tocNcxTemplate))
	}

	return n
}

type tocNcxContent struct {
	Src string `xml:"src,attr"`
}

type tocNcxMeta struct {
	Name    string `xml:"name,attr"`
	Content string `xml:"content,attr"`
}

type tocNcxNavPoint struct {
	XMLName  xml.Name          `xml:"navPoint"`
	ID       string            `xml:"id,attr"`
	Text     string            `xml:"navLabel>text"`
	Content  tocNcxContent     `xml:"content"`
	Children *[]tocNcxNavPoint `xml:"navPoint,omitempty"`
}
