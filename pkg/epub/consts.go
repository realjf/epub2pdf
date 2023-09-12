// #############################################################################
// # File: consts.go                                                           #
// # Project: epub                                                             #
// # Created Date: 2023/09/11 21:59:59                                         #
// # Author: realjf                                                            #
// # -----                                                                     #
// # Last Modified: 2023/09/12 08:07:43                                        #
// # Modified By: realjf                                                       #
// # -----                                                                     #
// # Copyright (c) 2023 realjf                                                 #
// #############################################################################
package epub

const (
	CSSFolderName   = "css"
	FontFolderName  = "fonts"
	ImageFolderName = "images"
	VideoFolderName = "videos"
	AudioFolderName = "audio"
)

const (
	cssFileFormat          = "css%04d%s"
	defaultCoverBody       = `<img src="%s" alt="Cover Image" />`
	defaultCoverCSSContent = `body {
  background-color: #FFFFFF;
  margin-bottom: 0px;
  margin-left: 0px;
  margin-right: 0px;
  margin-top: 0px;
  text-align: center;
}
img {
  max-height: 100%;
  max-width: 100%;
}
`
	defaultCoverCSSFilename   = "cover.css"
	defaultCoverCSSSource     = "cover.css"
	defaultCoverImgFormat     = "cover%s"
	defaultCoverXhtmlFilename = "cover.xhtml"
	defaultEpubLang           = "en"
	fontFileFormat            = "font%04d%s"
	imageFileFormat           = "image%04d%s"
	videoFileFormat           = "video%04d%s"
	sectionFileFormat         = "section%04d.xhtml"
	urnUUIDPrefix             = "urn:uuid:"
	audioFileFormat           = "audio%04d%s"
)

// ============================================ toc ==============================================

const (
	tocNavBodyTemplate = `
	<nav epub:type="toc">
      <h1>Table of Contents</h1>
      <ol>
      </ol>
    </nav>
`
	tocNavFilename       = "nav.xhtml"
	tocNavItemID         = "nav"
	tocNavItemProperties = "nav"
	tocNavEpubType       = "toc"

	tocNcxFilename = "toc.ncx"
	tocNcxItemID   = "ncx"
	tocNcxTemplate = `
<ncx xmlns="http://www.daisy.org/z3986/2005/ncx/" version="2005-1">
  <head>
    <meta name="dtb:uid" content="" />
    <meta name="dtb:depth" content="" />
  </head>
  <docTitle>
    <text></text>
  </docTitle>
  <docAuthor>
    <text></text>
  </docAuthor>
  <navMap>
  </navMap>
</ncx>`

	xmlnsEpub = "http://www.idpf.org/2007/ops"
)

// ============================================ xhtml ==============================================

const (
	xhtmlDoctype = `<!DOCTYPE html>
`
	xhtmlLinkRel  = "stylesheet"
	xhtmlTemplate = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">
  <head>
    <title dir="auto"></title>
  </head>
  <body dir="auto"></body>
</html>
`
)

// ============================================ write ==============================================

const (
	containerFilename     = "container.xml"
	containerFileTemplate = `<?xml version="1.0" encoding="UTF-8"?>
<container version="1.0" xmlns="urn:oasis:names:tc:opendocument:xmlns:container">
  <rootfiles>
    <rootfile full-path="%s/%s" media-type="application/oebps-package+xml" />
  </rootfiles>
</container>
`
	// This seems to be the standard based on the latest EPUB spec:
	// http://www.idpf.org/epub/31/spec/epub-ocf.html
	contentFolderName    = "EPUB"
	coverImageProperties = "cover-image"
	// Permissions for any new directories we create
	dirPermissions = 0755
	// Permissions for any new files we create
	filePermissions   = 0644
	mediaTypeCSS      = "text/css"
	mediaTypeEpub     = "application/epub+zip"
	mediaTypeJpeg     = "image/jpeg"
	mediaTypeNcx      = "application/x-dtbncx+xml"
	mediaTypeXhtml    = "application/xhtml+xml"
	metaInfFolderName = "META-INF"
	mimetypeFilename  = "mimetype"
	pkgFilename       = "package.opf"
	tempDirPrefix     = "go-epub"
	xhtmlFolderName   = "xhtml"
)

// ============================================ pkg ==============================================

const (
	pkgAuthorID       = "role"
	pkgAuthorData     = "aut"
	pkgAuthorProperty = "role"
	pkgAuthorRefines  = "#creator"
	pkgAuthorScheme   = "marc:relators"
	pkgCreatorID      = "creator"
	pkgFileTemplate   = `<?xml version="1.0" encoding="UTF-8"?>
<package version="3.0" unique-identifier="pub-id" xmlns="http://www.idpf.org/2007/opf">
  <metadata xmlns:dc="http://purl.org/dc/elements/1.1/">
    <dc:identifier id="pub-id"></dc:identifier>
    <dc:title></dc:title>
    <dc:language></dc:language>
    <dc:description></dc:description>
  </metadata>
  <manifest>
  </manifest>
  <spine toc="ncx">
  </spine>
</package>
`
	pkgModifiedProperty = "dcterms:modified"
	pkgUniqueIdentifier = "pub-id"

	xmlnsDc = "http://purl.org/dc/elements/1.1/"
)
