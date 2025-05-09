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

func InlineRuleClass() InlineRuleClassLike {
	return inlineRuleClass()
}

// Constructor Methods

func (c *inlineRuleClass_) InlineRule(
	terms col.Sequential[TermLike],
	optionalNote string,
) InlineRuleLike {
	if uti.IsUndefined(terms) {
		panic("The \"terms\" attribute is required by this class.")
	}
	var instance = &inlineRule_{
		// Initialize the instance attributes.
		terms_:        terms,
		optionalNote_: optionalNote,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *inlineRule_) GetClass() InlineRuleClassLike {
	return inlineRuleClass()
}

// Attribute Methods

func (v *inlineRule_) GetTerms() col.Sequential[TermLike] {
	return v.terms_
}

func (v *inlineRule_) GetOptionalNote() string {
	return v.optionalNote_
}

// PROTECTED INTERFACE

// Instance Structure

type inlineRule_ struct {
	// Declare the instance attributes.
	terms_        col.Sequential[TermLike]
	optionalNote_ string
}

// Class Structure

type inlineRuleClass_ struct {
	// Declare the class constants.
}

// Class Reference

func inlineRuleClass() *inlineRuleClass_ {
	return inlineRuleClassReference_
}

var inlineRuleClassReference_ = &inlineRuleClass_{
	// Initialize the class constants.
}
