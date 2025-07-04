/*
................................................................................
.    Copyright (c) 2009-2025 Crater Dog Technologies.  All Rights Reserved.    .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

package module_test

import (
	fmt "fmt"
	uti "github.com/craterdog/go-missing-utilities/v7"
	not "github.com/craterdog/go-syntax-notation/v7"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

var syntaxFiles = []string{
	"./syntax.cdsn",
	"./test/basm.cdsn",
	"./test/bali.cdsn",
	"./test/gcmn.cdsn",
	"./test/full.cdsn",
}

func TestRoundTrips(t *tes.T) {
	fmt.Println("Round Trip Tests:")
	for _, syntaxFile := range syntaxFiles {
		fmt.Printf("   %v\n", syntaxFile)
		var source = uti.ReadFile(syntaxFile)
		var syntax = not.ParseSource(source)
		not.ValidateSyntax(syntax)
		var actual = not.FormatSyntax(syntax)
		ass.Equal(t, source, actual)
	}
	fmt.Println("Done.")
}
