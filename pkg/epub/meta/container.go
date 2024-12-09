// #############################################################################
// # File: container.go                                                        #
// # Project: meta                                                             #
// # Created Date: 2024/12/09 23:05:47                                         #
// # Author: realjf                                                            #
// # -----                                                                     #
// # Last Modified: 2024/12/09 23:20:54                                        #
// # Modified By: realjf                                                       #
// # -----                                                                     #
// #                                                                           #
// #############################################################################
package meta

// ====================================== Container container.xml ===================================
type Container struct {
	Rootfiles Rootfiles `xml:"rootfiles"`
}

type Rootfiles struct {
	Rootfile []Rootfile `xml:"rootfile"`
}

type Rootfile struct {
	FullPath  string `xml:"full-path,attr"`
	MediaType string `xml:"media-type,attr"`
}

// ====================================== Package content.opf ===================================

type Package struct {
	Metadata Metadata `xml:"metadata"`
	Manifest Manifest `xml:"manifest"`
	Spine    Spine    `xml:"spine"`
}

type Metadata struct {
	Title       string `xml:"dc:title"`
	Creator     string `xml:"dc:creator"`
	Language    string `xml:"dc:language"`
	Identifier  string `xml:"dc:identifier"`
	Publisher   string `xml:"dc:publisher"`
	PublishedAt string `xml:"dc:date"`
}

type Manifest struct {
	Items []ManifestItem `xml:"item"`
}

type ManifestItem struct {
	ID        string `xml:"id,attr"`
	Href      string `xml:"href,attr"`
	MediaType string `xml:"media-type,attr"`
}

type Spine struct {
	Itemrefs []Itemref `xml:"itemref"`
}

type Itemref struct {
	IDRef string `xml:"idref,attr"`
}
