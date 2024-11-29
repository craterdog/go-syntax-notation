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

/*
┌────────────────────────────────── WARNING ───────────────────────────────────┐
│              This "Package.go" file was automatically generated.             │
│                     Any updates to it may be overwritten.                    │
└──────────────────────────────────────────────────────────────────────────────┘

Package "ast" provides the abstract syntax tree (AST) classes for this module
based on the "Syntax.cdsn" grammar for the module.  Each AST class manages the
attributes associated with its corresponding rule definition found in the
grammar.

For detailed documentation on this package refer to the wiki:
  - https://github.com/craterdog/go-syntax-notation/wiki

This package follows the Crater Dog Technologies™ Go Coding Conventions located
here:
  - https://github.com/craterdog/go-class-model/wiki

Additional concrete implementations of the classes declared by this package can
be developed and used seamlessly since the interface declarations only depend on
other interfaces and intrinsic types—and the class implementations only depend
on interfaces, not on each other.
*/
package ast

import (
	abs "github.com/craterdog/go-collection-framework/v4/collection"
)

// TYPE DECLARATIONS

// FUNCTIONAL DECLARATIONS

// CLASS DECLARATIONS

/*
AlternativeClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete alternative-like class.
*/
type AlternativeClassLike interface {
	// Constructor Methods
	Make(
		option OptionLike,
	) AlternativeLike
}

/*
CardinalityClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete cardinality-like class.
*/
type CardinalityClassLike interface {
	// Constructor Methods
	Make(
		any_ any,
	) CardinalityLike
}

/*
CharacterClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete character-like class.
*/
type CharacterClassLike interface {
	// Constructor Methods
	Make(
		any_ any,
	) CharacterLike
}

/*
ConstrainedClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete constrained-like class.
*/
type ConstrainedClassLike interface {
	// Constructor Methods
	Make(
		any_ any,
	) ConstrainedLike
}

/*
DefinitionClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete definition-like class.
*/
type DefinitionClassLike interface {
	// Constructor Methods
	Make(
		any_ any,
	) DefinitionLike
}

/*
ElementClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete element-like class.
*/
type ElementClassLike interface {
	// Constructor Methods
	Make(
		any_ any,
	) ElementLike
}

/*
ExplicitClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete explicit-like class.
*/
type ExplicitClassLike interface {
	// Constructor Methods
	Make(
		glyph string,
		optionalExtent ExtentLike,
	) ExplicitLike
}

/*
ExpressionClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete expression-like class.
*/
type ExpressionClassLike interface {
	// Constructor Methods
	Make(
		lowercase string,
		pattern PatternLike,
		optionalNote string,
	) ExpressionLike
}

/*
ExtentClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete extent-like class.
*/
type ExtentClassLike interface {
	// Constructor Methods
	Make(
		glyph string,
	) ExtentLike
}

/*
FilterClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete filter-like class.
*/
type FilterClassLike interface {
	// Constructor Methods
	Make(
		optionalExcluded string,
		characters abs.Sequential[CharacterLike],
	) FilterLike
}

/*
GroupClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete group-like class.
*/
type GroupClassLike interface {
	// Constructor Methods
	Make(
		pattern PatternLike,
	) GroupLike
}

/*
IdentifierClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete identifier-like class.
*/
type IdentifierClassLike interface {
	// Constructor Methods
	Make(
		any_ any,
	) IdentifierLike
}

/*
InlineClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete inline-like class.
*/
type InlineClassLike interface {
	// Constructor Methods
	Make(
		terms abs.Sequential[TermLike],
		optionalNote string,
	) InlineLike
}

/*
LimitClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete limit-like class.
*/
type LimitClassLike interface {
	// Constructor Methods
	Make(
		optionalNumber string,
	) LimitLike
}

/*
LineClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete line-like class.
*/
type LineClassLike interface {
	// Constructor Methods
	Make(
		identifier IdentifierLike,
		optionalNote string,
	) LineLike
}

/*
MultilineClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete multiline-like class.
*/
type MultilineClassLike interface {
	// Constructor Methods
	Make(
		lines abs.Sequential[LineLike],
	) MultilineLike
}

/*
NoticeClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete notice-like class.
*/
type NoticeClassLike interface {
	// Constructor Methods
	Make(
		comment string,
	) NoticeLike
}

/*
OptionClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete option-like class.
*/
type OptionClassLike interface {
	// Constructor Methods
	Make(
		repetitions abs.Sequential[RepetitionLike],
	) OptionLike
}

/*
PatternClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete pattern-like class.
*/
type PatternClassLike interface {
	// Constructor Methods
	Make(
		option OptionLike,
		alternatives abs.Sequential[AlternativeLike],
	) PatternLike
}

/*
QuantifiedClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete quantified-like class.
*/
type QuantifiedClassLike interface {
	// Constructor Methods
	Make(
		number string,
		optionalLimit LimitLike,
	) QuantifiedLike
}

/*
ReferenceClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete reference-like class.
*/
type ReferenceClassLike interface {
	// Constructor Methods
	Make(
		identifier IdentifierLike,
		optionalCardinality CardinalityLike,
	) ReferenceLike
}

/*
RepetitionClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete repetition-like class.
*/
type RepetitionClassLike interface {
	// Constructor Methods
	Make(
		element ElementLike,
		optionalCardinality CardinalityLike,
	) RepetitionLike
}

/*
RuleClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete rule-like class.
*/
type RuleClassLike interface {
	// Constructor Methods
	Make(
		uppercase string,
		definition DefinitionLike,
	) RuleLike
}

/*
SyntaxClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete syntax-like class.
*/
type SyntaxClassLike interface {
	// Constructor Methods
	Make(
		notice NoticeLike,
		comment1 string,
		rules abs.Sequential[RuleLike],
		comment2 string,
		expressions abs.Sequential[ExpressionLike],
	) SyntaxLike
}

/*
TermClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete term-like class.
*/
type TermClassLike interface {
	// Constructor Methods
	Make(
		any_ any,
	) TermLike
}

/*
TextClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete text-like class.
*/
type TextClassLike interface {
	// Constructor Methods
	Make(
		any_ any,
	) TextLike
}

// INSTANCE DECLARATIONS

/*
AlternativeLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete alternative-like class.
*/
type AlternativeLike interface {
	// Principal Methods
	GetClass() AlternativeClassLike

	// Attribute Methods
	GetOption() OptionLike
}

/*
CardinalityLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete cardinality-like class.
*/
type CardinalityLike interface {
	// Principal Methods
	GetClass() CardinalityClassLike

	// Attribute Methods
	GetAny() any
}

/*
CharacterLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete character-like class.
*/
type CharacterLike interface {
	// Principal Methods
	GetClass() CharacterClassLike

	// Attribute Methods
	GetAny() any
}

/*
ConstrainedLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete constrained-like class.
*/
type ConstrainedLike interface {
	// Principal Methods
	GetClass() ConstrainedClassLike

	// Attribute Methods
	GetAny() any
}

/*
DefinitionLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete definition-like class.
*/
type DefinitionLike interface {
	// Principal Methods
	GetClass() DefinitionClassLike

	// Attribute Methods
	GetAny() any
}

/*
ElementLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete element-like class.
*/
type ElementLike interface {
	// Principal Methods
	GetClass() ElementClassLike

	// Attribute Methods
	GetAny() any
}

/*
ExplicitLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete explicit-like class.
*/
type ExplicitLike interface {
	// Principal Methods
	GetClass() ExplicitClassLike

	// Attribute Methods
	GetGlyph() string
	GetOptionalExtent() ExtentLike
}

/*
ExpressionLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete expression-like class.
*/
type ExpressionLike interface {
	// Principal Methods
	GetClass() ExpressionClassLike

	// Attribute Methods
	GetLowercase() string
	GetPattern() PatternLike
	GetOptionalNote() string
}

/*
ExtentLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete extent-like class.
*/
type ExtentLike interface {
	// Principal Methods
	GetClass() ExtentClassLike

	// Attribute Methods
	GetGlyph() string
}

/*
FilterLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete filter-like class.
*/
type FilterLike interface {
	// Principal Methods
	GetClass() FilterClassLike

	// Attribute Methods
	GetOptionalExcluded() string
	GetCharacters() abs.Sequential[CharacterLike]
}

/*
GroupLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete group-like class.
*/
type GroupLike interface {
	// Principal Methods
	GetClass() GroupClassLike

	// Attribute Methods
	GetPattern() PatternLike
}

/*
IdentifierLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete identifier-like class.
*/
type IdentifierLike interface {
	// Principal Methods
	GetClass() IdentifierClassLike

	// Attribute Methods
	GetAny() any
}

/*
InlineLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete inline-like class.
*/
type InlineLike interface {
	// Principal Methods
	GetClass() InlineClassLike

	// Attribute Methods
	GetTerms() abs.Sequential[TermLike]
	GetOptionalNote() string
}

/*
LimitLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete limit-like class.
*/
type LimitLike interface {
	// Principal Methods
	GetClass() LimitClassLike

	// Attribute Methods
	GetOptionalNumber() string
}

/*
LineLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete line-like class.
*/
type LineLike interface {
	// Principal Methods
	GetClass() LineClassLike

	// Attribute Methods
	GetIdentifier() IdentifierLike
	GetOptionalNote() string
}

/*
MultilineLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete multiline-like class.
*/
type MultilineLike interface {
	// Principal Methods
	GetClass() MultilineClassLike

	// Attribute Methods
	GetLines() abs.Sequential[LineLike]
}

/*
NoticeLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete notice-like class.
*/
type NoticeLike interface {
	// Principal Methods
	GetClass() NoticeClassLike

	// Attribute Methods
	GetComment() string
}

/*
OptionLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete option-like class.
*/
type OptionLike interface {
	// Principal Methods
	GetClass() OptionClassLike

	// Attribute Methods
	GetRepetitions() abs.Sequential[RepetitionLike]
}

/*
PatternLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete pattern-like class.
*/
type PatternLike interface {
	// Principal Methods
	GetClass() PatternClassLike

	// Attribute Methods
	GetOption() OptionLike
	GetAlternatives() abs.Sequential[AlternativeLike]
}

/*
QuantifiedLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete quantified-like class.
*/
type QuantifiedLike interface {
	// Principal Methods
	GetClass() QuantifiedClassLike

	// Attribute Methods
	GetNumber() string
	GetOptionalLimit() LimitLike
}

/*
ReferenceLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete reference-like class.
*/
type ReferenceLike interface {
	// Principal Methods
	GetClass() ReferenceClassLike

	// Attribute Methods
	GetIdentifier() IdentifierLike
	GetOptionalCardinality() CardinalityLike
}

/*
RepetitionLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete repetition-like class.
*/
type RepetitionLike interface {
	// Principal Methods
	GetClass() RepetitionClassLike

	// Attribute Methods
	GetElement() ElementLike
	GetOptionalCardinality() CardinalityLike
}

/*
RuleLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete rule-like class.
*/
type RuleLike interface {
	// Principal Methods
	GetClass() RuleClassLike

	// Attribute Methods
	GetUppercase() string
	GetDefinition() DefinitionLike
}

/*
SyntaxLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete syntax-like class.
*/
type SyntaxLike interface {
	// Principal Methods
	GetClass() SyntaxClassLike

	// Attribute Methods
	GetNotice() NoticeLike
	GetComment1() string
	GetRules() abs.Sequential[RuleLike]
	GetComment2() string
	GetExpressions() abs.Sequential[ExpressionLike]
}

/*
TermLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete term-like class.
*/
type TermLike interface {
	// Principal Methods
	GetClass() TermClassLike

	// Attribute Methods
	GetAny() any
}

/*
TextLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete text-like class.
*/
type TextLike interface {
	// Principal Methods
	GetClass() TextClassLike

	// Attribute Methods
	GetAny() any
}

// ASPECT DECLARATIONS
