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

func MultiExpressionClass() MultiExpressionClassLike {
	return multiExpressionClass()
}

// Constructor Methods

func (c *multiExpressionClass_) MultiExpression(
	expressionOptions col.Sequential[ExpressionOptionLike],
) MultiExpressionLike {
	if uti.IsUndefined(expressionOptions) {
		panic("The \"expressionOptions\" attribute is required by this class.")
	}
	var instance = &multiExpression_{
		// Initialize the instance attributes.
		expressionOptions_: expressionOptions,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *multiExpression_) GetClass() MultiExpressionClassLike {
	return multiExpressionClass()
}

// Attribute Methods

func (v *multiExpression_) GetExpressionOptions() col.Sequential[ExpressionOptionLike] {
	return v.expressionOptions_
}

// PROTECTED INTERFACE

// Instance Structure

type multiExpression_ struct {
	// Declare the instance attributes.
	expressionOptions_ col.Sequential[ExpressionOptionLike]
}

// Class Structure

type multiExpressionClass_ struct {
	// Declare the class constants.
}

// Class Reference

func multiExpressionClass() *multiExpressionClass_ {
	return multiExpressionClassReference_
}

var multiExpressionClassReference_ = &multiExpressionClass_{
	// Initialize the class constants.
}
