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
│              This "Module.go" file was automatically generated.              │
│      Updates to any part of this file—other than the Module Description      │
│             and the Global Functions sections may be overwritten.            │
└──────────────────────────────────────────────────────────────────────────────┘

Package "module" declares type aliases for the commonly used types declared in
the packages contained in this module.  It also provides a default constructor
for each commonly used class that is exported by the module.  Each constructor
delegates the actual construction process to its corresponding concrete class
declared in the corresponding package contained within this module.

For detailed documentation on this entire module refer to the wiki:
  - https://github.com/craterdog/go-syntax-notation/wiki
*/
package module

import (
	abs "github.com/craterdog/go-collection-framework/v5/collection"
	ast "github.com/craterdog/go-syntax-notation/v5/ast"
	gra "github.com/craterdog/go-syntax-notation/v5/grammar"
)

// TYPE ALIASES

// Ast

type (
	AlternativeLike = ast.AlternativeLike
	CardinalityLike = ast.CardinalityLike
	CharacterLike   = ast.CharacterLike
	ConstrainedLike = ast.ConstrainedLike
	DefinitionLike  = ast.DefinitionLike
	ElementLike     = ast.ElementLike
	ExplicitLike    = ast.ExplicitLike
	ExpressionLike  = ast.ExpressionLike
	ExtentLike      = ast.ExtentLike
	FilterLike      = ast.FilterLike
	GroupLike       = ast.GroupLike
	IdentifierLike  = ast.IdentifierLike
	InlineLike      = ast.InlineLike
	LimitLike       = ast.LimitLike
	LineLike        = ast.LineLike
	MultilineLike   = ast.MultilineLike
	NoticeLike      = ast.NoticeLike
	OptionLike      = ast.OptionLike
	PatternLike     = ast.PatternLike
	QuantifiedLike  = ast.QuantifiedLike
	ReferenceLike   = ast.ReferenceLike
	RepetitionLike  = ast.RepetitionLike
	RuleLike        = ast.RuleLike
	SyntaxLike      = ast.SyntaxLike
	TermLike        = ast.TermLike
	TextLike        = ast.TextLike
)

// Grammar

type (
	TokenType = gra.TokenType
)

const (
	ErrorToken     = gra.ErrorToken
	CommentToken   = gra.CommentToken
	DelimiterToken = gra.DelimiterToken
	ExcludedToken  = gra.ExcludedToken
	GlyphToken     = gra.GlyphToken
	IntrinsicToken = gra.IntrinsicToken
	LiteralToken   = gra.LiteralToken
	LowercaseToken = gra.LowercaseToken
	NewlineToken   = gra.NewlineToken
	NoteToken      = gra.NoteToken
	NumberToken    = gra.NumberToken
	OptionalToken  = gra.OptionalToken
	RepeatedToken  = gra.RepeatedToken
	SpaceToken     = gra.SpaceToken
	UppercaseToken = gra.UppercaseToken
)

type (
	FormatterLike = gra.FormatterLike
	ParserLike    = gra.ParserLike
	ProcessorLike = gra.ProcessorLike
	ScannerLike   = gra.ScannerLike
	TokenLike     = gra.TokenLike
	ValidatorLike = gra.ValidatorLike
	VisitorLike   = gra.VisitorLike
)

type (
	Methodical = gra.Methodical
)

// DEFAULT CONSTRUCTORS

// Ast

func Alternative(
	option OptionLike,
) AlternativeLike {
	return ast.AlternativeClass().Make(
		option,
	)
}

func Cardinality(
	any_ any,
) CardinalityLike {
	return ast.CardinalityClass().Make(
		any_,
	)
}

func Character(
	any_ any,
) CharacterLike {
	return ast.CharacterClass().Make(
		any_,
	)
}

func Constrained(
	any_ any,
) ConstrainedLike {
	return ast.ConstrainedClass().Make(
		any_,
	)
}

func Definition(
	any_ any,
) DefinitionLike {
	return ast.DefinitionClass().Make(
		any_,
	)
}

func Element(
	any_ any,
) ElementLike {
	return ast.ElementClass().Make(
		any_,
	)
}

func Explicit(
	glyph string,
	optionalExtent ExtentLike,
) ExplicitLike {
	return ast.ExplicitClass().Make(
		glyph,
		optionalExtent,
	)
}

func Expression(
	lowercase string,
	pattern PatternLike,
	optionalNote string,
) ExpressionLike {
	return ast.ExpressionClass().Make(
		lowercase,
		pattern,
		optionalNote,
	)
}

func Extent(
	glyph string,
) ExtentLike {
	return ast.ExtentClass().Make(
		glyph,
	)
}

func Filter(
	optionalExcluded string,
	characters abs.Sequential[CharacterLike],
) FilterLike {
	return ast.FilterClass().Make(
		optionalExcluded,
		characters,
	)
}

func Group(
	pattern PatternLike,
) GroupLike {
	return ast.GroupClass().Make(
		pattern,
	)
}

func Identifier(
	any_ any,
) IdentifierLike {
	return ast.IdentifierClass().Make(
		any_,
	)
}

func Inline(
	terms abs.Sequential[TermLike],
	optionalNote string,
) InlineLike {
	return ast.InlineClass().Make(
		terms,
		optionalNote,
	)
}

func Limit(
	optionalNumber string,
) LimitLike {
	return ast.LimitClass().Make(
		optionalNumber,
	)
}

func Line(
	identifier IdentifierLike,
	optionalNote string,
) LineLike {
	return ast.LineClass().Make(
		identifier,
		optionalNote,
	)
}

func Multiline(
	lines abs.Sequential[LineLike],
) MultilineLike {
	return ast.MultilineClass().Make(
		lines,
	)
}

func Notice(
	comment string,
) NoticeLike {
	return ast.NoticeClass().Make(
		comment,
	)
}

func Option(
	repetitions abs.Sequential[RepetitionLike],
) OptionLike {
	return ast.OptionClass().Make(
		repetitions,
	)
}

func Pattern(
	option OptionLike,
	alternatives abs.Sequential[AlternativeLike],
) PatternLike {
	return ast.PatternClass().Make(
		option,
		alternatives,
	)
}

func Quantified(
	number string,
	optionalLimit LimitLike,
) QuantifiedLike {
	return ast.QuantifiedClass().Make(
		number,
		optionalLimit,
	)
}

func Reference(
	identifier IdentifierLike,
	optionalCardinality CardinalityLike,
) ReferenceLike {
	return ast.ReferenceClass().Make(
		identifier,
		optionalCardinality,
	)
}

func Repetition(
	element ElementLike,
	optionalCardinality CardinalityLike,
) RepetitionLike {
	return ast.RepetitionClass().Make(
		element,
		optionalCardinality,
	)
}

func Rule(
	uppercase string,
	definition DefinitionLike,
) RuleLike {
	return ast.RuleClass().Make(
		uppercase,
		definition,
	)
}

func Syntax(
	notice NoticeLike,
	comment1 string,
	rules abs.Sequential[RuleLike],
	comment2 string,
	expressions abs.Sequential[ExpressionLike],
) SyntaxLike {
	return ast.SyntaxClass().Make(
		notice,
		comment1,
		rules,
		comment2,
		expressions,
	)
}

func Term(
	any_ any,
) TermLike {
	return ast.TermClass().Make(
		any_,
	)
}

func Text(
	any_ any,
) TextLike {
	return ast.TextClass().Make(
		any_,
	)
}

// Grammar

func Formatter() FormatterLike {
	return gra.FormatterClass().Make()
}

func Parser() ParserLike {
	return gra.ParserClass().Make()
}

func Processor() ProcessorLike {
	return gra.ProcessorClass().Make()
}

func Scanner(
	source string,
	tokens abs.QueueLike[TokenLike],
) ScannerLike {
	return gra.ScannerClass().Make(
		source,
		tokens,
	)
}

func Token(
	line uint,
	position uint,
	type_ TokenType,
	value string,
) TokenLike {
	return gra.TokenClass().Make(
		line,
		position,
		type_,
		value,
	)
}

func Validator() ValidatorLike {
	return gra.ValidatorClass().Make()
}

func Visitor(
	processor Methodical,
) VisitorLike {
	return gra.VisitorClass().Make(
		processor,
	)
}

// GLOBAL FUNCTIONS

func FormatSyntax(
	syntax SyntaxLike,
) string {
	var formatter = Formatter()
	return formatter.FormatSyntax(syntax)
}

func MatchesType(
	tokenValue string,
	tokenType TokenType,
) bool {
	var scannerClass = gra.ScannerClass()
	return scannerClass.MatchesType(tokenValue, tokenType)
}

func ParseSource(
	source string,
) SyntaxLike {
	var parser = Parser()
	return parser.ParseSource(source)
}

func ValidateSyntax(
	syntax SyntaxLike,
) {
	var validator = Validator()
	validator.ValidateSyntax(syntax)
}
