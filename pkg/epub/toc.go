// #############################################################################
// # File: toc.go                                                              #
// # Project: epub                                                             #
// # Created Date: 2023/09/11 22:07:35                                         #
// # Author: realjf                                                            #
// # -----                                                                     #
// # Last Modified: 2023/09/12 09:13:00                                        #
// # Modified By: realjf                                                       #
// # -----                                                                     #
// # Copyright (c) 2023 realjf                                                 #
// #############################################################################
package epub

import (
	"encoding/xml"
	"fmt"
	"path/filepath"
	"strconv"
)

type toc struct {
	author string // EPUB author
	title  string // EPUB title
	navXML *tocNavBody
	ncxXML *tocNcxRoot
}

// Constructor for toc
func newToc() *toc {
	t := &toc{
		navXML: newTocNavXML(),
		ncxXML: newTocNcxXML(),
	}

	return t
}

// Add a section to the TOC (navXML as well as ncxXML)
func (t *toc) addSection(index int, title string, relativePath string) {
	relativePath = filepath.ToSlash(relativePath)
	l := &tocNavItem{
		A: tocNavLink{
			Href: relativePath,
			Data: title,
		},
		Children: nil,
	}
	t.navXML.Links = append(t.navXML.Links, *l)

	np := &tocNcxNavPoint{
		ID:   "navPoint-" + strconv.Itoa(index),
		Text: title,
		Content: tocNcxContent{
			Src: relativePath,
		},
		Children: nil,
	}
	t.ncxXML.NavMap = append(t.ncxXML.NavMap, *np)
}

// Add a sub section to the TOC (navXML as well as ncxXML)
func (t *toc) addSubSection(parent string, index int, title string, relativePath string) {
	var parentNcxIndex int
	var parentNavIndex int

	relativePath = filepath.ToSlash(relativePath)
	parent = filepath.ToSlash(parent)

	for index, nav := range t.navXML.Links {
		if nav.A.Href == parent {
			parentNavIndex = index
		}
	}
	l := tocNavItem{
		A: tocNavLink{
			Href: relativePath,
			Data: title,
		},
	}
	if len(t.navXML.Links) > parentNavIndex {
		// Create a new array if none exists
		if t.navXML.Links[parentNavIndex].Children == nil {
			n := make([]tocNavItem, 0)
			t.navXML.Links[parentNavIndex].Children = &n
		}
		children := append(*t.navXML.Links[parentNavIndex].Children, l)
		t.navXML.Links[parentNavIndex].Children = &children
	} else {
		t.navXML.Links = append(t.navXML.Links, l)
	}

	// Get parent object
	for index, ncx := range t.ncxXML.NavMap {
		if ncx.Content.Src == parent {
			parentNcxIndex = index
		}
	}
	np := tocNcxNavPoint{
		ID:   "navPoint-" + strconv.Itoa(index),
		Text: title,
		Content: tocNcxContent{
			Src: relativePath,
		},
		Children: nil,
	}
	if parentNcxIndex > len(t.ncxXML.NavMap) {
		if t.ncxXML.NavMap[parentNcxIndex].Children == nil {
			n := make([]tocNcxNavPoint, 0)
			t.ncxXML.NavMap[parentNcxIndex].Children = &n
		}
		children := append(*t.ncxXML.NavMap[parentNcxIndex].Children, np)
		t.ncxXML.NavMap[parentNcxIndex].Children = &children
	} else {
		t.ncxXML.NavMap = append(t.ncxXML.NavMap, np)
	}
}

func (t *toc) setIdentifier(identifier string) {
	t.ncxXML.Meta.Content = identifier
}

func (t *toc) setTitle(title string) {
	t.title = title
}

func (t *toc) setAuthor(author string) {
	t.author = author
}

// Write the TOC files
func (t *toc) write(tempDir string) {
	t.writeNavDoc(tempDir)
	t.writeNcxDoc(tempDir)
}

// Write the the EPUB v3 TOC file (nav.xhtml) to the temporary directory
func (t *toc) writeNavDoc(tempDir string) {
	navBodyContent, err := xml.MarshalIndent(t.navXML, "    ", "  ")
	if err != nil {
		panic(fmt.Sprintf(
			"Error marshalling XML for EPUB v3 TOC file: %s\n"+
				"\tXML=%#v",
			err,
			t.navXML))
	}

	n := newXhtml(string(navBodyContent))
	n.setXmlnsEpub(xmlnsEpub)
	n.setTitle(t.title)

	navFilePath := filepath.Join(tempDir, contentFolderName, tocNavFilename)
	n.write(navFilePath)
}

// Write the EPUB v2 TOC file (toc.ncx) to the temporary directory
func (t *toc) writeNcxDoc(tempDir string) {
	t.ncxXML.Title = t.title
	t.ncxXML.Author = t.author

	ncxFileContent, err := xml.MarshalIndent(t.ncxXML, "", "  ")
	if err != nil {
		panic(fmt.Sprintf(
			"Error marshalling XML for EPUB v2 TOC file: %s\n"+
				"\tXML=%#v",
			err,
			t.ncxXML))
	}

	// Add the xml header to the output
	ncxFileContent = append([]byte(xml.Header), ncxFileContent...)
	// It's generally nice to have files end with a newline
	ncxFileContent = append(ncxFileContent, "\n"...)

	ncxFilePath := filepath.Join(tempDir, contentFolderName, tocNcxFilename)
	if err := filesystem.WriteFile(ncxFilePath, []byte(ncxFileContent), filePermissions); err != nil {
		panic(fmt.Sprintf("Error writing EPUB v2 TOC file: %s", err))
	}
}

