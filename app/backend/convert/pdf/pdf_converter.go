// #############################################################################
// # File: pdf_converter.go                                                    #
// # Project: pdf                                                              #
// # Created Date: 2023/09/11 16:21:03                                         #
// # Author: realjf                                                            #
// # -----                                                                     #
// # Last Modified: 2023/09/11 16:50:33                                        #
// # Modified By: realjf                                                       #
// # -----                                                                     #
// # Copyright (c) 2023 realjf                                                 #
// #############################################################################
package pdf

type PDFConverter interface {
}

type pdfConverter struct {
}

func NewPDFConverter() PDFConverter {
	return &pdfConverter{}
}
