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
│             This "module_api.go" file was automatically generated.           │
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
	col "github.com/craterdog/go-collection-framework/v7"
	ast "github.com/craterdog/go-syntax-notation/v7/ast"
	gra "github.com/craterdog/go-syntax-notation/v7/grammar"
)

// TYPE ALIASES

// Ast

type (
	AlternativeClassLike      = ast.AlternativeClassLike
	CardinalityClassLike      = ast.CardinalityClassLike
	CharacterClassLike        = ast.CharacterClassLike
	ComponentClassLike        = ast.ComponentClassLike
	ConstrainedClassLike      = ast.ConstrainedClassLike
	DefinitionClassLike       = ast.DefinitionClassLike
	ElementClassLike          = ast.ElementClassLike
	ExplicitClassLike         = ast.ExplicitClassLike
	ExpressionClassLike       = ast.ExpressionClassLike
	ExpressionOptionClassLike = ast.ExpressionOptionClassLike
	ExtentClassLike           = ast.ExtentClassLike
	FilterClassLike           = ast.FilterClassLike
	GroupClassLike            = ast.GroupClassLike
	ImplicitClassLike         = ast.ImplicitClassLike
	InlineRuleClassLike       = ast.InlineRuleClassLike
	LimitClassLike            = ast.LimitClassLike
	LiteralOptionClassLike    = ast.LiteralOptionClassLike
	MultiExpressionClassLike  = ast.MultiExpressionClassLike
	MultiLiteralClassLike     = ast.MultiLiteralClassLike
	MultiRuleClassLike        = ast.MultiRuleClassLike
	NoticeClassLike           = ast.NoticeClassLike
	OptionClassLike           = ast.OptionClassLike
	PatternClassLike          = ast.PatternClassLike
	QuantifiedClassLike       = ast.QuantifiedClassLike
	RepetitionClassLike       = ast.RepetitionClassLike
	RuleClassLike             = ast.RuleClassLike
	RuleOptionClassLike       = ast.RuleOptionClassLike
	SyntaxClassLike           = ast.SyntaxClassLike
	TermClassLike             = ast.TermClassLike
	TextClassLike             = ast.TextClassLike
)

type (
	AlternativeLike      = ast.AlternativeLike
	CardinalityLike      = ast.CardinalityLike
	CharacterLike        = ast.CharacterLike
	ComponentLike        = ast.ComponentLike
	ConstrainedLike      = ast.ConstrainedLike
	DefinitionLike       = ast.DefinitionLike
	ElementLike          = ast.ElementLike
	ExplicitLike         = ast.ExplicitLike
	ExpressionLike       = ast.ExpressionLike
	ExpressionOptionLike = ast.ExpressionOptionLike
	ExtentLike           = ast.ExtentLike
	FilterLike           = ast.FilterLike
	GroupLike            = ast.GroupLike
	ImplicitLike         = ast.ImplicitLike
	InlineRuleLike       = ast.InlineRuleLike
	LimitLike            = ast.LimitLike
	LiteralOptionLike    = ast.LiteralOptionLike
	MultiExpressionLike  = ast.MultiExpressionLike
	MultiLiteralLike     = ast.MultiLiteralLike
	MultiRuleLike        = ast.MultiRuleLike
	NoticeLike           = ast.NoticeLike
	OptionLike           = ast.OptionLike
	PatternLike          = ast.PatternLike
	QuantifiedLike       = ast.QuantifiedLike
	RepetitionLike       = ast.RepetitionLike
	RuleLike             = ast.RuleLike
	RuleOptionLike       = ast.RuleOptionLike
	SyntaxLike           = ast.SyntaxLike
	TermLike             = ast.TermLike
	TextLike             = ast.TextLike
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
	FormatterClassLike = gra.FormatterClassLike
	ParserClassLike    = gra.ParserClassLike
	ProcessorClassLike = gra.ProcessorClassLike
	ScannerClassLike   = gra.ScannerClassLike
	TokenClassLike     = gra.TokenClassLike
	ValidatorClassLike = gra.ValidatorClassLike
	VisitorClassLike   = gra.VisitorClassLike
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

// CLASS ACCESSORS

// Ast

func AlternativeClass() AlternativeClassLike {
	return ast.AlternativeClass()
}

func Alternative(
	option ast.OptionLike,
) AlternativeLike {
	return AlternativeClass().Alternative(
		option,
	)
}

func CardinalityClass() CardinalityClassLike {
	return ast.CardinalityClass()
}

func Cardinality(
	any_ any,
) CardinalityLike {
	return CardinalityClass().Cardinality(
		any_,
	)
}

func CharacterClass() CharacterClassLike {
	return ast.CharacterClass()
}

func Character(
	any_ any,
) CharacterLike {
	return CharacterClass().Character(
		any_,
	)
}

func ComponentClass() ComponentClassLike {
	return ast.ComponentClass()
}

func Component(
	any_ any,
) ComponentLike {
	return ComponentClass().Component(
		any_,
	)
}

func ConstrainedClass() ConstrainedClassLike {
	return ast.ConstrainedClass()
}

func Constrained(
	any_ any,
) ConstrainedLike {
	return ConstrainedClass().Constrained(
		any_,
	)
}

func DefinitionClass() DefinitionClassLike {
	return ast.DefinitionClass()
}

func Definition(
	any_ any,
) DefinitionLike {
	return DefinitionClass().Definition(
		any_,
	)
}

func ElementClass() ElementClassLike {
	return ast.ElementClass()
}

func Element(
	any_ any,
) ElementLike {
	return ElementClass().Element(
		any_,
	)
}

func ExplicitClass() ExplicitClassLike {
	return ast.ExplicitClass()
}

func Explicit(
	glyph string,
	optionalExtent ast.ExtentLike,
) ExplicitLike {
	return ExplicitClass().Explicit(
		glyph,
		optionalExtent,
	)
}

func ExpressionClass() ExpressionClassLike {
	return ast.ExpressionClass()
}

func Expression(
	lowercase string,
	pattern ast.PatternLike,
	optionalNote string,
) ExpressionLike {
	return ExpressionClass().Expression(
		lowercase,
		pattern,
		optionalNote,
	)
}

func ExpressionOptionClass() ExpressionOptionClassLike {
	return ast.ExpressionOptionClass()
}

func ExpressionOption(
	newline string,
	lowercase string,
	optionalNote string,
) ExpressionOptionLike {
	return ExpressionOptionClass().ExpressionOption(
		newline,
		lowercase,
		optionalNote,
	)
}

func ExtentClass() ExtentClassLike {
	return ast.ExtentClass()
}

func Extent(
	glyph string,
) ExtentLike {
	return ExtentClass().Extent(
		glyph,
	)
}

func FilterClass() FilterClassLike {
	return ast.FilterClass()
}

func Filter(
	optionalExcluded string,
	characters col.Sequential[ast.CharacterLike],
) FilterLike {
	return FilterClass().Filter(
		optionalExcluded,
		characters,
	)
}

func GroupClass() GroupClassLike {
	return ast.GroupClass()
}

func Group(
	pattern ast.PatternLike,
) GroupLike {
	return GroupClass().Group(
		pattern,
	)
}

func ImplicitClass() ImplicitClassLike {
	return ast.ImplicitClass()
}

func Implicit(
	intrinsic string,
) ImplicitLike {
	return ImplicitClass().Implicit(
		intrinsic,
	)
}

func InlineRuleClass() InlineRuleClassLike {
	return ast.InlineRuleClass()
}

func InlineRule(
	terms col.Sequential[ast.TermLike],
	optionalNote string,
) InlineRuleLike {
	return InlineRuleClass().InlineRule(
		terms,
		optionalNote,
	)
}

func LimitClass() LimitClassLike {
	return ast.LimitClass()
}

func Limit(
	optionalNumber string,
) LimitLike {
	return LimitClass().Limit(
		optionalNumber,
	)
}

func LiteralOptionClass() LiteralOptionClassLike {
	return ast.LiteralOptionClass()
}

func LiteralOption(
	newline string,
	literal string,
	optionalNote string,
) LiteralOptionLike {
	return LiteralOptionClass().LiteralOption(
		newline,
		literal,
		optionalNote,
	)
}

func MultiExpressionClass() MultiExpressionClassLike {
	return ast.MultiExpressionClass()
}

func MultiExpression(
	expressionOptions col.Sequential[ast.ExpressionOptionLike],
) MultiExpressionLike {
	return MultiExpressionClass().MultiExpression(
		expressionOptions,
	)
}

func MultiLiteralClass() MultiLiteralClassLike {
	return ast.MultiLiteralClass()
}

func MultiLiteral(
	literalOptions col.Sequential[ast.LiteralOptionLike],
) MultiLiteralLike {
	return MultiLiteralClass().MultiLiteral(
		literalOptions,
	)
}

func MultiRuleClass() MultiRuleClassLike {
	return ast.MultiRuleClass()
}

func MultiRule(
	ruleOptions col.Sequential[ast.RuleOptionLike],
) MultiRuleLike {
	return MultiRuleClass().MultiRule(
		ruleOptions,
	)
}

func NoticeClass() NoticeClassLike {
	return ast.NoticeClass()
}

func Notice(
	comment string,
	newline string,
) NoticeLike {
	return NoticeClass().Notice(
		comment,
		newline,
	)
}

func OptionClass() OptionClassLike {
	return ast.OptionClass()
}

func Option(
	repetitions col.Sequential[ast.RepetitionLike],
) OptionLike {
	return OptionClass().Option(
		repetitions,
	)
}

func PatternClass() PatternClassLike {
	return ast.PatternClass()
}

func Pattern(
	option ast.OptionLike,
	alternatives col.Sequential[ast.AlternativeLike],
) PatternLike {
	return PatternClass().Pattern(
		option,
		alternatives,
	)
}

func QuantifiedClass() QuantifiedClassLike {
	return ast.QuantifiedClass()
}

func Quantified(
	number string,
	optionalLimit ast.LimitLike,
) QuantifiedLike {
	return QuantifiedClass().Quantified(
		number,
		optionalLimit,
	)
}

func RepetitionClass() RepetitionClassLike {
	return ast.RepetitionClass()
}

func Repetition(
	element ast.ElementLike,
	optionalCardinality ast.CardinalityLike,
) RepetitionLike {
	return RepetitionClass().Repetition(
		element,
		optionalCardinality,
	)
}

func RuleClass() RuleClassLike {
	return ast.RuleClass()
}

func Rule(
	uppercase string,
	definition ast.DefinitionLike,
) RuleLike {
	return RuleClass().Rule(
		uppercase,
		definition,
	)
}

func RuleOptionClass() RuleOptionClassLike {
	return ast.RuleOptionClass()
}

func RuleOption(
	newline string,
	uppercase string,
	optionalNote string,
) RuleOptionLike {
	return RuleOptionClass().RuleOption(
		newline,
		uppercase,
		optionalNote,
	)
}

func SyntaxClass() SyntaxClassLike {
	return ast.SyntaxClass()
}

func Syntax(
	notice ast.NoticeLike,
	comment1 string,
	rules col.Sequential[ast.RuleLike],
	comment2 string,
	expressions col.Sequential[ast.ExpressionLike],
) SyntaxLike {
	return SyntaxClass().Syntax(
		notice,
		comment1,
		rules,
		comment2,
		expressions,
	)
}

func TermClass() TermClassLike {
	return ast.TermClass()
}

func Term(
	component ast.ComponentLike,
	optionalCardinality ast.CardinalityLike,
) TermLike {
	return TermClass().Term(
		component,
		optionalCardinality,
	)
}

func TextClass() TextClassLike {
	return ast.TextClass()
}

func Text(
	any_ any,
) TextLike {
	return TextClass().Text(
		any_,
	)
}

// Grammar

func FormatterClass() FormatterClassLike {
	return gra.FormatterClass()
}

func Formatter() FormatterLike {
	return FormatterClass().Formatter()
}

func ParserClass() ParserClassLike {
	return gra.ParserClass()
}

func Parser() ParserLike {
	return ParserClass().Parser()
}

func ProcessorClass() ProcessorClassLike {
	return gra.ProcessorClass()
}

func Processor() ProcessorLike {
	return ProcessorClass().Processor()
}

func ScannerClass() ScannerClassLike {
	return gra.ScannerClass()
}

func Scanner(
	source string,
	tokens col.QueueLike[gra.TokenLike],
) ScannerLike {
	return ScannerClass().Scanner(
		source,
		tokens,
	)
}

func TokenClass() TokenClassLike {
	return gra.TokenClass()
}

func Token(
	line uint,
	position uint,
	type_ gra.TokenType,
	value string,
) TokenLike {
	return TokenClass().Token(
		line,
		position,
		type_,
		value,
	)
}

func ValidatorClass() ValidatorClassLike {
	return gra.ValidatorClass()
}

func Validator() ValidatorLike {
	return ValidatorClass().Validator()
}

func VisitorClass() VisitorClassLike {
	return gra.VisitorClass()
}

func Visitor(
	processor gra.Methodical,
) VisitorLike {
	return VisitorClass().Visitor(
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
