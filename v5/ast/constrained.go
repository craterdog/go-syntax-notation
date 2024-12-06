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

func ConstrainedClass() ConstrainedClassLike {
	return constrainedClassReference()
}

// Constructor Methods

func (c *constrainedClass_) Make(
	any_ any,
) ConstrainedLike {
	if uti.IsUndefined(any_) {
		panic("The \"any\" attribute is required by this class.")
	}
	var instance = &constrained_{
		// Initialize the instance attributes.
		any_: any_,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *constrained_) GetClass() ConstrainedClassLike {
	return constrainedClassReference()
}

// Attribute Methods

func (v *constrained_) GetAny() any {
	return v.any_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type constrained_ struct {
	// Declare the instance attributes.
	any_ any
}

// Class Structure

type constrainedClass_ struct {
	// Declare the class constants.
}

// Class Reference

func constrainedClassReference() *constrainedClass_ {
	return constrainedClassReference_
}

var constrainedClassReference_ = &constrainedClass_{
	// Initialize the class constants.
}
