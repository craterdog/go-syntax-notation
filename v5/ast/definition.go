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

func Definition() DefinitionClassLike {
	return definitionReference()
}

// Constructor Methods

func (c *definitionClass_) Make(
	any_ any,
) DefinitionLike {
	if uti.IsUndefined(any_) {
		panic("The \"any\" attribute is required by this class.")
	}
	var instance = &definition_{
		// Initialize the instance attributes.
		any_: any_,
	}
	return instance

}

// INSTANCE INTERFACE

// Primary Methods

func (v *definition_) GetClass() DefinitionClassLike {
	return definitionReference()
}

// Attribute Methods

func (v *definition_) GetAny() any {
	return v.any_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type definition_ struct {
	// Declare the instance attributes.
	any_ any
}

// Class Structure

type definitionClass_ struct {
	// Declare the class constants.
}

// Class Reference

func definitionReference() *definitionClass_ {
	return definitionReference_
}

var definitionReference_ = &definitionClass_{
	// Initialize the class constants.
}
