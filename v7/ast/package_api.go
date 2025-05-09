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
│            This "package_api.go" file was automatically generated.           │
│                     Any updates to it may be overwritten.                    │
└──────────────────────────────────────────────────────────────────────────────┘

Package "ast" provides the abstract syntax tree (AST) classes for this module
based on the "syntax.cdsn" grammar for the module.  Each AST class manages the
attributes associated with its corresponding rule definition found in the
grammar.

For detailed documentation on this package refer to the wiki:
  - https://github.com/craterdog/go-syntax-notation/wiki

This package follows the Crater Dog Technologies™ Go Coding Conventions located
here:
  - https://github.com/craterdog/go-development-tools/wiki/Coding-Conventions

Additional concrete implementations of the classes declared by this package can
be developed and used seamlessly since the interface declarations only depend on
other interfaces and intrinsic types—and the class implementations only depend
on interfaces, not on each other.
*/
package ast

import (
	col "github.com/craterdog/go-collection-framework/v7"
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
	Alternative(
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
	Cardinality(
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
	Character(
		any_ any,
	) CharacterLike
}

/*
ComponentClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete component-like class.
*/
type ComponentClassLike interface {
	// Constructor Methods
	Component(
		any_ any,
	) ComponentLike
}

/*
ConstrainedClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete constrained-like class.
*/
type ConstrainedClassLike interface {
	// Constructor Methods
	Constrained(
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
	Definition(
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
	Element(
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
	Explicit(
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
	Expression(
		lowercase string,
		pattern PatternLike,
		optionalNote string,
	) ExpressionLike
}

/*
ExpressionOptionClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete expression-option-like class.
*/
type ExpressionOptionClassLike interface {
	// Constructor Methods
	ExpressionOption(
		newline string,
		lowercase string,
		optionalNote string,
	) ExpressionOptionLike
}

/*
ExtentClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete extent-like class.
*/
type ExtentClassLike interface {
	// Constructor Methods
	Extent(
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
	Filter(
		optionalExcluded string,
		characters col.Sequential[CharacterLike],
	) FilterLike
}

/*
GroupClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete group-like class.
*/
type GroupClassLike interface {
	// Constructor Methods
	Group(
		pattern PatternLike,
	) GroupLike
}

/*
ImplicitClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete implicit-like class.
*/
type ImplicitClassLike interface {
	// Constructor Methods
	Implicit(
		intrinsic string,
	) ImplicitLike
}

/*
InlineRuleClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete inline-rule-like class.
*/
type InlineRuleClassLike interface {
	// Constructor Methods
	InlineRule(
		terms col.Sequential[TermLike],
		optionalNote string,
	) InlineRuleLike
}

/*
LimitClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete limit-like class.
*/
type LimitClassLike interface {
	// Constructor Methods
	Limit(
		optionalNumber string,
	) LimitLike
}

/*
LiteralOptionClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete literal-option-like class.
*/
type LiteralOptionClassLike interface {
	// Constructor Methods
	LiteralOption(
		newline string,
		literal string,
		optionalNote string,
	) LiteralOptionLike
}

/*
MultiExpressionClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete multi-expression-like class.
*/
type MultiExpressionClassLike interface {
	// Constructor Methods
	MultiExpression(
		expressionOptions col.Sequential[ExpressionOptionLike],
	) MultiExpressionLike
}

/*
MultiLiteralClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete multi-literal-like class.
*/
type MultiLiteralClassLike interface {
	// Constructor Methods
	MultiLiteral(
		literalOptions col.Sequential[LiteralOptionLike],
	) MultiLiteralLike
}

/*
MultiRuleClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete multi-rule-like class.
*/
type MultiRuleClassLike interface {
	// Constructor Methods
	MultiRule(
		ruleOptions col.Sequential[RuleOptionLike],
	) MultiRuleLike
}

/*
NoticeClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete notice-like class.
*/
type NoticeClassLike interface {
	// Constructor Methods
	Notice(
		comment string,
		newline string,
	) NoticeLike
}

/*
OptionClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete option-like class.
*/
type OptionClassLike interface {
	// Constructor Methods
	Option(
		repetitions col.Sequential[RepetitionLike],
	) OptionLike
}

/*
PatternClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete pattern-like class.
*/
type PatternClassLike interface {
	// Constructor Methods
	Pattern(
		option OptionLike,
		alternatives col.Sequential[AlternativeLike],
	) PatternLike
}

/*
QuantifiedClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete quantified-like class.
*/
type QuantifiedClassLike interface {
	// Constructor Methods
	Quantified(
		number string,
		optionalLimit LimitLike,
	) QuantifiedLike
}

/*
RepetitionClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete repetition-like class.
*/
type RepetitionClassLike interface {
	// Constructor Methods
	Repetition(
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
	Rule(
		uppercase string,
		definition DefinitionLike,
	) RuleLike
}

/*
RuleOptionClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete rule-option-like class.
*/
type RuleOptionClassLike interface {
	// Constructor Methods
	RuleOption(
		newline string,
		uppercase string,
		optionalNote string,
	) RuleOptionLike
}

/*
SyntaxClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete syntax-like class.
*/
type SyntaxClassLike interface {
	// Constructor Methods
	Syntax(
		notice NoticeLike,
		comment1 string,
		rules col.Sequential[RuleLike],
		comment2 string,
		expressions col.Sequential[ExpressionLike],
	) SyntaxLike
}

/*
TermClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete term-like class.
*/
type TermClassLike interface {
	// Constructor Methods
	Term(
		component ComponentLike,
		optionalCardinality CardinalityLike,
	) TermLike
}

/*
TextClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete text-like class.
*/
type TextClassLike interface {
	// Constructor Methods
	Text(
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
ComponentLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete component-like class.
*/
type ComponentLike interface {
	// Principal Methods
	GetClass() ComponentClassLike

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
ExpressionOptionLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete expression-option-like class.
*/
type ExpressionOptionLike interface {
	// Principal Methods
	GetClass() ExpressionOptionClassLike

	// Attribute Methods
	GetNewline() string
	GetLowercase() string
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
	GetCharacters() col.Sequential[CharacterLike]
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
ImplicitLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete implicit-like class.
*/
type ImplicitLike interface {
	// Principal Methods
	GetClass() ImplicitClassLike

	// Attribute Methods
	GetIntrinsic() string
}

/*
InlineRuleLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete inline-rule-like class.
*/
type InlineRuleLike interface {
	// Principal Methods
	GetClass() InlineRuleClassLike

	// Attribute Methods
	GetTerms() col.Sequential[TermLike]
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
LiteralOptionLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete literal-option-like class.
*/
type LiteralOptionLike interface {
	// Principal Methods
	GetClass() LiteralOptionClassLike

	// Attribute Methods
	GetNewline() string
	GetLiteral() string
	GetOptionalNote() string
}

/*
MultiExpressionLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete multi-expression-like class.
*/
type MultiExpressionLike interface {
	// Principal Methods
	GetClass() MultiExpressionClassLike

	// Attribute Methods
	GetExpressionOptions() col.Sequential[ExpressionOptionLike]
}

/*
MultiLiteralLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete multi-literal-like class.
*/
type MultiLiteralLike interface {
	// Principal Methods
	GetClass() MultiLiteralClassLike

	// Attribute Methods
	GetLiteralOptions() col.Sequential[LiteralOptionLike]
}

/*
MultiRuleLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete multi-rule-like class.
*/
type MultiRuleLike interface {
	// Principal Methods
	GetClass() MultiRuleClassLike

	// Attribute Methods
	GetRuleOptions() col.Sequential[RuleOptionLike]
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
	GetNewline() string
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
	GetRepetitions() col.Sequential[RepetitionLike]
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
	GetAlternatives() col.Sequential[AlternativeLike]
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
RuleOptionLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete rule-option-like class.
*/
type RuleOptionLike interface {
	// Principal Methods
	GetClass() RuleOptionClassLike

	// Attribute Methods
	GetNewline() string
	GetUppercase() string
	GetOptionalNote() string
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
	GetRules() col.Sequential[RuleLike]
	GetComment2() string
	GetExpressions() col.Sequential[ExpressionLike]
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
	GetComponent() ComponentLike
	GetOptionalCardinality() CardinalityLike
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
