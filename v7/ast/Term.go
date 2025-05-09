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
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func TermClass() TermClassLike {
	return termClass()
}

// Constructor Methods

func (c *termClass_) Term(
	component ComponentLike,
	optionalCardinality CardinalityLike,
) TermLike {
	if uti.IsUndefined(component) {
		panic("The \"component\" attribute is required by this class.")
	}
	var instance = &term_{
		// Initialize the instance attributes.
		component_:           component,
		optionalCardinality_: optionalCardinality,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *term_) GetClass() TermClassLike {
	return termClass()
}

// Attribute Methods

func (v *term_) GetComponent() ComponentLike {
	return v.component_
}

func (v *term_) GetOptionalCardinality() CardinalityLike {
	return v.optionalCardinality_
}

// PROTECTED INTERFACE

// Instance Structure

type term_ struct {
	// Declare the instance attributes.
	component_           ComponentLike
	optionalCardinality_ CardinalityLike
}

// Class Structure

type termClass_ struct {
	// Declare the class constants.
}

// Class Reference

func termClass() *termClass_ {
	return termClassReference_
}

var termClassReference_ = &termClass_{
	// Initialize the class constants.
}
