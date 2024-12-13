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
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func AlternativeClass() AlternativeClassLike {
	return alternativeClassReference()
}

// Constructor Methods

func (c *alternativeClass_) Alternative(
	option OptionLike,
) AlternativeLike {
	if uti.IsUndefined(option) {
		panic("The \"option\" attribute is required by this class.")
	}
	var instance = &alternative_{
		// Initialize the instance attributes.
		option_: option,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *alternative_) GetClass() AlternativeClassLike {
	return alternativeClassReference()
}

// Attribute Methods

func (v *alternative_) GetOption() OptionLike {
	return v.option_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type alternative_ struct {
	// Declare the instance attributes.
	option_ OptionLike
}

// Class Structure

type alternativeClass_ struct {
	// Declare the class constants.
}

// Class Reference

func alternativeClassReference() *alternativeClass_ {
	return alternativeClassReference_
}

var alternativeClassReference_ = &alternativeClass_{
	// Initialize the class constants.
}
