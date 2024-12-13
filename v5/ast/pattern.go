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

package ast

import (
	abs "github.com/craterdog/go-collection-framework/v5/collection"
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func PatternClass() PatternClassLike {
	return patternClassReference()
}

// Constructor Methods

func (c *patternClass_) Pattern(
	option OptionLike,
	alternatives abs.Sequential[AlternativeLike],
) PatternLike {
	if uti.IsUndefined(option) {
		panic("The \"option\" attribute is required by this class.")
	}
	if uti.IsUndefined(alternatives) {
		panic("The \"alternatives\" attribute is required by this class.")
	}
	var instance = &pattern_{
		// Initialize the instance attributes.
		option_:       option,
		alternatives_: alternatives,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *pattern_) GetClass() PatternClassLike {
	return patternClassReference()
}

// Attribute Methods

func (v *pattern_) GetOption() OptionLike {
	return v.option_
}

func (v *pattern_) GetAlternatives() abs.Sequential[AlternativeLike] {
	return v.alternatives_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type pattern_ struct {
	// Declare the instance attributes.
	option_       OptionLike
	alternatives_ abs.Sequential[AlternativeLike]
}

// Class Structure

type patternClass_ struct {
	// Declare the class constants.
}

// Class Reference

func patternClassReference() *patternClass_ {
	return patternClassReference_
}

var patternClassReference_ = &patternClass_{
	// Initialize the class constants.
}
