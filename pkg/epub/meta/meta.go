// #############################################################################
// # File: meta.go                                                             #
// # Project: meta                                                             #
// # Created Date: 2024/12/09 23:00:41                                         #
// # Author: realjf                                                            #
// # -----                                                                     #
// # Last Modified: 2024/12/10 07:26:50                                        #
// # Modified By: realjf                                                       #
// # -----                                                                     #
// #                                                                           #
// #############################################################################
package meta

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
)

type MetaCommand struct {
	cacheDir string
	filePath string

	container Container

	contentOPFPath string
	pkg            Package
}

func New(cacheDir, fileName string) *MetaCommand {
	if fileName == "" {
		fileName = "META-INF/container.xml"
	}
	filePath := filepath.Join(cacheDir, fileName)
	return &MetaCommand{
		filePath: filePath,
		cacheDir: cacheDir,
	}
}

func (m *MetaCommand) Close() {

}

func (m *MetaCommand) Execute() (err error) {
	// parse container.xml
	err = m.parseContainer()
	if err != nil {
		fmt.Printf("[MetaCommand] parse container error: %v", err)
		return
	}
	// parse content.opf
	err = m.parseContentOPF()
	if err != nil {
		fmt.Printf("[MetaCommand] parse content opf error: %v", err)
		return
	}
	return
}

func (m *MetaCommand) parseContainer() (err error) {
	content, err := os.ReadFile(m.filePath)
	if err != nil {
		fmt.Printf("[MetaCommand] reader from %s error: %v", m.filePath, err)
		return
	}

	if err := xml.Unmarshal(content, &m.container); err != nil {
		fmt.Printf("[MetaCommand] unmarshal xml from %s error: %v", m.filePath, err)
		return err
	}

	m.contentOPFPath = filepath.Join(m.cacheDir, m.container.Rootfiles.Rootfile[0].FullPath)
	fmt.Printf("Rootfile path: %s\n", m.contentOPFPath)
	return
}

func (m *MetaCommand) parseContentOPF() (err error) {
	content, err := os.ReadFile(m.contentOPFPath)
	if err != nil {
		fmt.Printf("[MetaCommand] reader from %s error: %v", m.contentOPFPath, err)
		return
	}

	if err := xml.Unmarshal(content, &m.pkg); err != nil {
		fmt.Printf("[MetaCommand] unmarshal xml from %s error: %v", m.contentOPFPath, err)
		return err
	}

	fmt.Println("Metadata:")
	fmt.Printf("  Title: %s\n", m.pkg.Metadata.Title)
	fmt.Printf("  Description: %s\n", m.pkg.Metadata.Description)
	fmt.Printf("  Creator: %s\n", m.pkg.Metadata.Creator)
	fmt.Printf("  Language: %s\n", m.pkg.Metadata.Language)
	fmt.Printf("  Identifier: %s\n", m.pkg.Metadata.Identifier)
	fmt.Printf("  Publisher: %s\n", m.pkg.Metadata.Publisher)
	fmt.Printf("  PublishedAt: %s\n", m.pkg.Metadata.PublishedAt)
	return
}
