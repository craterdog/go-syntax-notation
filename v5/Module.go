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
the packages contained in this module.  It also provides constructors for each
commonly used class that is exported by the module.  Each constructor delegates
the actual construction process to its corresponding concrete class declared in
the corresponding package contained within this module.

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

// CLASS CONSTRUCTORS

// Ast/Alternative

func Alternative(
	option OptionLike,
) ast.AlternativeLike {
	return ast.AlternativeClass().Alternative(
		option,
	)
}

// Ast/Cardinality

func Cardinality(
	any_ any,
) ast.CardinalityLike {
	return ast.CardinalityClass().Cardinality(
		any_,
	)
}

// Ast/Character

func Character(
	any_ any,
) ast.CharacterLike {
	return ast.CharacterClass().Character(
		any_,
	)
}

// Ast/Constrained

func Constrained(
	any_ any,
) ast.ConstrainedLike {
	return ast.ConstrainedClass().Constrained(
		any_,
	)
}

// Ast/Definition

func Definition(
	any_ any,
) ast.DefinitionLike {
	return ast.DefinitionClass().Definition(
		any_,
	)
}

// Ast/Element

func Element(
	any_ any,
) ast.ElementLike {
	return ast.ElementClass().Element(
		any_,
	)
}

// Ast/Explicit

func Explicit(
	glyph string,
	optionalExtent ExtentLike,
) ast.ExplicitLike {
	return ast.ExplicitClass().Explicit(
		glyph,
		optionalExtent,
	)
}

// Ast/Expression

func Expression(
	lowercase string,
	pattern PatternLike,
	optionalNote string,
) ast.ExpressionLike {
	return ast.ExpressionClass().Expression(
		lowercase,
		pattern,
		optionalNote,
	)
}

// Ast/Extent

func Extent(
	glyph string,
) ast.ExtentLike {
	return ast.ExtentClass().Extent(
		glyph,
	)
}

// Ast/Filter

func Filter(
	optionalExcluded string,
	characters abs.Sequential[CharacterLike],
) ast.FilterLike {
	return ast.FilterClass().Filter(
		optionalExcluded,
		characters,
	)
}

// Ast/Group

func Group(
	pattern PatternLike,
) ast.GroupLike {
	return ast.GroupClass().Group(
		pattern,
	)
}

// Ast/Identifier

func Identifier(
	any_ any,
) ast.IdentifierLike {
	return ast.IdentifierClass().Identifier(
		any_,
	)
}

// Ast/Inline

func Inline(
	terms abs.Sequential[TermLike],
	optionalNote string,
) ast.InlineLike {
	return ast.InlineClass().Inline(
		terms,
		optionalNote,
	)
}

// Ast/Limit

func Limit(
	optionalNumber string,
) ast.LimitLike {
	return ast.LimitClass().Limit(
		optionalNumber,
	)
}

// Ast/Line

func Line(
	identifier IdentifierLike,
	optionalNote string,
) ast.LineLike {
	return ast.LineClass().Line(
		identifier,
		optionalNote,
	)
}

// Ast/Multiline

func Multiline(
	lines abs.Sequential[LineLike],
) ast.MultilineLike {
	return ast.MultilineClass().Multiline(
		lines,
	)
}

// Ast/Notice

func Notice(
	comment string,
) ast.NoticeLike {
	return ast.NoticeClass().Notice(
		comment,
	)
}

// Ast/Option

func Option(
	repetitions abs.Sequential[RepetitionLike],
) ast.OptionLike {
	return ast.OptionClass().Option(
		repetitions,
	)
}

// Ast/Pattern

func Pattern(
	option OptionLike,
	alternatives abs.Sequential[AlternativeLike],
) ast.PatternLike {
	return ast.PatternClass().Pattern(
		option,
		alternatives,
	)
}

// Ast/Quantified

func Quantified(
	number string,
	optionalLimit LimitLike,
) ast.QuantifiedLike {
	return ast.QuantifiedClass().Quantified(
		number,
		optionalLimit,
	)
}

// Ast/Reference

func Reference(
	identifier IdentifierLike,
	optionalCardinality CardinalityLike,
) ast.ReferenceLike {
	return ast.ReferenceClass().Reference(
		identifier,
		optionalCardinality,
	)
}

// Ast/Repetition

func Repetition(
	element ElementLike,
	optionalCardinality CardinalityLike,
) ast.RepetitionLike {
	return ast.RepetitionClass().Repetition(
		element,
		optionalCardinality,
	)
}

// Ast/Rule

func Rule(
	uppercase string,
	definition DefinitionLike,
) ast.RuleLike {
	return ast.RuleClass().Rule(
		uppercase,
		definition,
	)
}

// Ast/Syntax

func Syntax(
	notice NoticeLike,
	comment1 string,
	rules abs.Sequential[RuleLike],
	comment2 string,
	expressions abs.Sequential[ExpressionLike],
) ast.SyntaxLike {
	return ast.SyntaxClass().Syntax(
		notice,
		comment1,
		rules,
		comment2,
		expressions,
	)
}

// Ast/Term

func Term(
	any_ any,
) ast.TermLike {
	return ast.TermClass().Term(
		any_,
	)
}

// Ast/Text

func Text(
	any_ any,
) ast.TextLike {
	return ast.TextClass().Text(
		any_,
	)
}

// Grammar/Formatter

func Formatter() gra.FormatterLike {
	return gra.FormatterClass().Formatter()
}

// Grammar/Parser

func Parser() gra.ParserLike {
	return gra.ParserClass().Parser()
}

// Grammar/Processor

func Processor() gra.ProcessorLike {
	return gra.ProcessorClass().Processor()
}

// Grammar/Scanner

func Scanner(
	source string,
	tokens abs.QueueLike[TokenLike],
) gra.ScannerLike {
	return gra.ScannerClass().Scanner(
		source,
		tokens,
	)
}

// Grammar/Token

func Token(
	line uint,
	position uint,
	type_ TokenType,
	value string,
) gra.TokenLike {
	return gra.TokenClass().Token(
		line,
		position,
		type_,
		value,
	)
}

// Grammar/Validator

func Validator() gra.ValidatorLike {
	return gra.ValidatorClass().Validator()
}

// Grammar/Visitor

func Visitor(
	processor Methodical,
) gra.VisitorLike {
	return gra.VisitorClass().Visitor(
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
