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

/*
┌────────────────────────────────── WARNING ───────────────────────────────────┐
│                 This class file was automatically generated.                 │
│                     Any updates to it may be overwritten.                    │
└──────────────────────────────────────────────────────────────────────────────┘
*/

package ast

import (
	col "github.com/craterdog/go-collection-framework/v7"
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func MultiLiteralClass() MultiLiteralClassLike {
	return multiLiteralClass()
}

// Constructor Methods

func (c *multiLiteralClass_) MultiLiteral(
	literalOptions col.Sequential[LiteralOptionLike],
) MultiLiteralLike {
	if uti.IsUndefined(literalOptions) {
		panic("The \"literalOptions\" attribute is required by this class.")
	}
	var instance = &multiLiteral_{
		// Initialize the instance attributes.
		literalOptions_: literalOptions,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *multiLiteral_) GetClass() MultiLiteralClassLike {
	return multiLiteralClass()
}

// Attribute Methods

func (v *multiLiteral_) GetLiteralOptions() col.Sequential[LiteralOptionLike] {
	return v.literalOptions_
}

// PROTECTED INTERFACE

// Instance Structure

type multiLiteral_ struct {
	// Declare the instance attributes.
	literalOptions_ col.Sequential[LiteralOptionLike]
}

// Class Structure

type multiLiteralClass_ struct {
	// Declare the class constants.
}

// Class Reference

func multiLiteralClass() *multiLiteralClass_ {
	return multiLiteralClassReference_
}

var multiLiteralClassReference_ = &multiLiteralClass_{
	// Initialize the class constants.
}
