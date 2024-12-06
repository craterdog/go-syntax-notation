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

func QuantifiedClass() QuantifiedClassLike {
	return quantifiedClassReference()
}

// Constructor Methods

func (c *quantifiedClass_) Make(
	number string,
	optionalLimit LimitLike,
) QuantifiedLike {
	if uti.IsUndefined(number) {
		panic("The \"number\" attribute is required by this class.")
	}
	var instance = &quantified_{
		// Initialize the instance attributes.
		number_:        number,
		optionalLimit_: optionalLimit,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *quantified_) GetClass() QuantifiedClassLike {
	return quantifiedClassReference()
}

// Attribute Methods

func (v *quantified_) GetNumber() string {
	return v.number_
}

func (v *quantified_) GetOptionalLimit() LimitLike {
	return v.optionalLimit_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type quantified_ struct {
	// Declare the instance attributes.
	number_        string
	optionalLimit_ LimitLike
}

// Class Structure

type quantifiedClass_ struct {
	// Declare the class constants.
}

// Class Reference

func quantifiedClassReference() *quantifiedClass_ {
	return quantifiedClassReference_
}

var quantifiedClassReference_ = &quantifiedClass_{
	// Initialize the class constants.
}
