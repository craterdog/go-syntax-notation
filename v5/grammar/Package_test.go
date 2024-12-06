/*
................................................................................
.    Copyright (c) 2009-2024 Crater Dog Technologies.  All Rights Reserved.    .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

package grammar_test

import (
	gra "github.com/craterdog/go-syntax-notation/v5/grammar"
	ass "github.com/stretchr/testify/assert"
	osx "os"
	tes "testing"
)

func TestRoundTrips(t *tes.T) {
	var syntaxFile = "../Syntax.cdsn"
	var bytes, err = osx.ReadFile(syntaxFile)
	if err != nil {
		panic(err)
	}
	var source = string(bytes)
	var parser = gra.ParserClass().Make()
	var syntax = parser.ParseSource(source)
	var validator = gra.ValidatorClass().Make()
	validator.ValidateSyntax(syntax)
	var formatter = gra.FormatterClass().Make()
	var actual = formatter.FormatSyntax(syntax)
	ass.Equal(t, source, actual)
}
