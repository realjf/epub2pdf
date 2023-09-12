// #############################################################################
// # File: error.go                                                            #
// # Project: epub                                                             #
// # Created Date: 2023/09/12 08:08:33                                         #
// # Author: realjf                                                            #
// # -----                                                                     #
// # Last Modified: 2023/09/12 09:05:42                                        #
// # Modified By: realjf                                                       #
// # -----                                                                     #
// # Copyright (c) 2023 realjf                                                 #
// #############################################################################
package epub

import "fmt"

// FilenameAlreadyUsedError is thrown by AddCSS, AddFont, AddImage, or AddSection
// if the same filename is used more than once.
type FilenameAlreadyUsedError struct {
	Filename string // Filename that caused the error
}

func (e *FilenameAlreadyUsedError) Error() string {
	return fmt.Sprintf("Filename already used: %s", e.Filename)
}

// FileRetrievalError is thrown by AddCSS, AddFont, AddImage, or Write if there was a
// problem retrieving the source file that was provided.
type FileRetrievalError struct {
	Source string // The source of the file whose retrieval failed
	Err    error  // The underlying error that was thrown
}

func (e *FileRetrievalError) Error() string {
	return fmt.Sprintf("Error retrieving %q from source: %+v", e.Source, e.Err)
}

// ParentDoesNotExistError is thrown by AddSubSection if the parent with the
// previously defined internal filename does not exist.
type ParentDoesNotExistError struct {
	Filename string // Filename that caused the error
}

func (e *ParentDoesNotExistError) Error() string {
	return fmt.Sprintf("Parent with the internal filename %s does not exist", e.Filename)
}

// UnableToCreateEpubError is thrown by Write if it cannot create the destination EPUB file
type UnableToCreateEpubError struct {
	Path string // The path that was given to Write to create the EPUB
	Err  error  // The underlying error that was thrown
}

func (e *UnableToCreateEpubError) Error() string {
	return fmt.Sprintf("Error creating EPUB at %q: %+v", e.Path, e.Err)
}
