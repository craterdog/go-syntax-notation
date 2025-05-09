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

func MultiRuleClass() MultiRuleClassLike {
	return multiRuleClass()
}

// Constructor Methods

func (c *multiRuleClass_) MultiRule(
	ruleOptions col.Sequential[RuleOptionLike],
) MultiRuleLike {
	if uti.IsUndefined(ruleOptions) {
		panic("The \"ruleOptions\" attribute is required by this class.")
	}
	var instance = &multiRule_{
		// Initialize the instance attributes.
		ruleOptions_: ruleOptions,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *multiRule_) GetClass() MultiRuleClassLike {
	return multiRuleClass()
}

// Attribute Methods

func (v *multiRule_) GetRuleOptions() col.Sequential[RuleOptionLike] {
	return v.ruleOptions_
}

// PROTECTED INTERFACE

// Instance Structure

type multiRule_ struct {
	// Declare the instance attributes.
	ruleOptions_ col.Sequential[RuleOptionLike]
}

// Class Structure

type multiRuleClass_ struct {
	// Declare the class constants.
}

// Class Reference

func multiRuleClass() *multiRuleClass_ {
	return multiRuleClassReference_
}

var multiRuleClassReference_ = &multiRuleClass_{
	// Initialize the class constants.
}
