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

package module_test

import (
	fmt "fmt"
	syn "github.com/craterdog/go-syntax-notation/v5"
	ass "github.com/stretchr/testify/assert"
	osx "os"
	tes "testing"
)

var testDirectories = []string{
	"../../go-test-framework/v5/",
}

func TestRoundTrips(t *tes.T) {
	fmt.Println("Round Trip Tests:")
	for _, directory := range testDirectories {
		fmt.Printf("   %v\n", directory)
		var bytes, err = osx.ReadFile(directory + "Syntax.cdsn")
		if err != nil {
			panic(err)
		}
		var source = string(bytes)
		var syntax = syn.ParseSource(source)
		syn.ValidateSyntax(syntax)
		var actual = syn.FormatSyntax(syntax)
		ass.Equal(t, source, actual)
	}
	fmt.Println("Done.")
}
