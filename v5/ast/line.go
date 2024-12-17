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

/*
┌────────────────────────────────── WARNING ───────────────────────────────────┐
│                 This class file was automatically generated.                 │
│                     Any updates to it may be overwritten.                    │
└──────────────────────────────────────────────────────────────────────────────┘
*/

package ast

import (
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func LineClass() LineClassLike {
	return lineClassReference()
}

// Constructor Methods

func (c *lineClass_) Line(
	identifier IdentifierLike,
	optionalNote string,
) LineLike {
	if uti.IsUndefined(identifier) {
		panic("The \"identifier\" attribute is required by this class.")
	}
	var instance = &line_{
		// Initialize the instance attributes.
		identifier_:   identifier,
		optionalNote_: optionalNote,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *line_) GetClass() LineClassLike {
	return lineClassReference()
}

// Attribute Methods

func (v *line_) GetIdentifier() IdentifierLike {
	return v.identifier_
}

func (v *line_) GetOptionalNote() string {
	return v.optionalNote_
}

// PROTECTED INTERFACE

// Instance Structure

type line_ struct {
	// Declare the instance attributes.
	identifier_   IdentifierLike
	optionalNote_ string
}

// Class Structure

type lineClass_ struct {
	// Declare the class constants.
}

// Class Reference

func lineClassReference() *lineClass_ {
	return lineClassReference_
}

var lineClassReference_ = &lineClass_{
	// Initialize the class constants.
}
