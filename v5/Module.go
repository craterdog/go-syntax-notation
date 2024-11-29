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
│  Updates to any section other than the GLOBAL FUNCTIONS may be overwritten.  │
└──────────────────────────────────────────────────────────────────────────────┘

Package "module" declares type aliases for the commonly used types declared in
the packages contained in this module.  It also provides a universal constructor
for each commonly used class that is exported by the module.  Each constructor
delegates the actual construction process to its corresponding concrete class
declared in the corresponding package contained within this module.

For detailed documentation on this entire module refer to the wiki:
  - https://github.com/craterdog/go-syntax-notation/wiki
*/
package module

import (
	fmt "fmt"
	abs "github.com/craterdog/go-collection-framework/v4/collection"
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

// UNIVERSAL CONSTRUCTORS

// Ast

func Alternative(arguments ...any) AlternativeLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case OptionLike:
			argumentTypes += "OptionLike, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Alternative constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma and space.
		argumentTypes = argumentTypes[:length-2]
	}

	// Call the corresponding constructor.
	var instance_ AlternativeLike
	switch argumentTypes {
	case "OptionLike":
		var option = arguments[0].(OptionLike)
		instance_ = ast.Alternative().Make(
			option,
		)
	default:
		var message = fmt.Sprintf(
			"No Alternative constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func Cardinality(arguments ...any) CardinalityLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case any:
			argumentTypes += "any, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Cardinality constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma and space.
		argumentTypes = argumentTypes[:length-2]
	}

	// Call the corresponding constructor.
	var instance_ CardinalityLike
	switch argumentTypes {
	case "any":
		var any_ = arguments[0]
		instance_ = ast.Cardinality().Make(
			any_,
		)
	default:
		var message = fmt.Sprintf(
			"No Cardinality constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func Character(arguments ...any) CharacterLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case any:
			argumentTypes += "any, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Character constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma and space.
		argumentTypes = argumentTypes[:length-2]
	}

	// Call the corresponding constructor.
	var instance_ CharacterLike
	switch argumentTypes {
	case "any":
		var any_ = arguments[0]
		instance_ = ast.Character().Make(
			any_,
		)
	default:
		var message = fmt.Sprintf(
			"No Character constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func Constrained(arguments ...any) ConstrainedLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case any:
			argumentTypes += "any, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Constrained constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma and space.
		argumentTypes = argumentTypes[:length-2]
	}

	// Call the corresponding constructor.
	var instance_ ConstrainedLike
	switch argumentTypes {
	case "any":
		var any_ = arguments[0]
		instance_ = ast.Constrained().Make(
			any_,
		)
	default:
		var message = fmt.Sprintf(
			"No Constrained constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func Definition(arguments ...any) DefinitionLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case any:
			argumentTypes += "any, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Definition constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma and space.
		argumentTypes = argumentTypes[:length-2]
	}

	// Call the corresponding constructor.
	var instance_ DefinitionLike
	switch argumentTypes {
	case "any":
		var any_ = arguments[0]
		instance_ = ast.Definition().Make(
			any_,
		)
	default:
		var message = fmt.Sprintf(
			"No Definition constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func Element(arguments ...any) ElementLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case any:
			argumentTypes += "any, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Element constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma and space.
		argumentTypes = argumentTypes[:length-2]
	}

	// Call the corresponding constructor.
	var instance_ ElementLike
	switch argumentTypes {
	case "any":
		var any_ = arguments[0]
		instance_ = ast.Element().Make(
			any_,
		)
	default:
		var message = fmt.Sprintf(
			"No Element constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func Explicit(arguments ...any) ExplicitLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			argumentTypes += "string, "
		case ExtentLike:
			argumentTypes += "ExtentLike, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Explicit constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma and space.
		argumentTypes = argumentTypes[:length-2]
	}

	// Call the corresponding constructor.
	var instance_ ExplicitLike
	switch argumentTypes {
	case "string, ExtentLike":
		var glyph = arguments[0].(string)
		var optionalExtent = arguments[1].(ExtentLike)
		instance_ = ast.Explicit().Make(
			glyph,
			optionalExtent,
		)
	default:
		var message = fmt.Sprintf(
			"No Explicit constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func Expression(arguments ...any) ExpressionLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			argumentTypes += "string, "
		case PatternLike:
			argumentTypes += "PatternLike, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Expression constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma and space.
		argumentTypes = argumentTypes[:length-2]
	}

	// Call the corresponding constructor.
	var instance_ ExpressionLike
	switch argumentTypes {
	case "string, PatternLike, string":
		var lowercase = arguments[0].(string)
		var pattern = arguments[1].(PatternLike)
		var optionalNote = arguments[2].(string)
		instance_ = ast.Expression().Make(
			lowercase,
			pattern,
			optionalNote,
		)
	default:
		var message = fmt.Sprintf(
			"No Expression constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func Extent(arguments ...any) ExtentLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			argumentTypes += "string, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Extent constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma and space.
		argumentTypes = argumentTypes[:length-2]
	}

	// Call the corresponding constructor.
	var instance_ ExtentLike
	switch argumentTypes {
	case "string":
		var glyph = arguments[0].(string)
		instance_ = ast.Extent().Make(
			glyph,
		)
	default:
		var message = fmt.Sprintf(
			"No Extent constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func Filter(arguments ...any) FilterLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			argumentTypes += "string, "
		case abs.Sequential[CharacterLike]:
			argumentTypes += "abs.Sequential[CharacterLike], "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Filter constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma and space.
		argumentTypes = argumentTypes[:length-2]
	}

	// Call the corresponding constructor.
	var instance_ FilterLike
	switch argumentTypes {
	case "string, abs.Sequential[CharacterLike]":
		var optionalExcluded = arguments[0].(string)
		var characters = arguments[1].(abs.Sequential[CharacterLike])
		instance_ = ast.Filter().Make(
			optionalExcluded,
			characters,
		)
	default:
		var message = fmt.Sprintf(
			"No Filter constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func Group(arguments ...any) GroupLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case PatternLike:
			argumentTypes += "PatternLike, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Group constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma and space.
		argumentTypes = argumentTypes[:length-2]
	}

	// Call the corresponding constructor.
	var instance_ GroupLike
	switch argumentTypes {
	case "PatternLike":
		var pattern = arguments[0].(PatternLike)
		instance_ = ast.Group().Make(
			pattern,
		)
	default:
		var message = fmt.Sprintf(
			"No Group constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func Identifier(arguments ...any) IdentifierLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case any:
			argumentTypes += "any, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Identifier constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma and space.
		argumentTypes = argumentTypes[:length-2]
	}

	// Call the corresponding constructor.
	var instance_ IdentifierLike
	switch argumentTypes {
	case "any":
		var any_ = arguments[0]
		instance_ = ast.Identifier().Make(
			any_,
		)
	default:
		var message = fmt.Sprintf(
			"No Identifier constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func Inline(arguments ...any) InlineLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case abs.Sequential[TermLike]:
			argumentTypes += "abs.Sequential[TermLike], "
		case string:
			argumentTypes += "string, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Inline constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma and space.
		argumentTypes = argumentTypes[:length-2]
	}

	// Call the corresponding constructor.
	var instance_ InlineLike
	switch argumentTypes {
	case "abs.Sequential[TermLike], string":
		var terms = arguments[0].(abs.Sequential[TermLike])
		var optionalNote = arguments[1].(string)
		instance_ = ast.Inline().Make(
			terms,
			optionalNote,
		)
	default:
		var message = fmt.Sprintf(
			"No Inline constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func Limit(arguments ...any) LimitLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			argumentTypes += "string, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Limit constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma and space.
		argumentTypes = argumentTypes[:length-2]
	}

	// Call the corresponding constructor.
	var instance_ LimitLike
	switch argumentTypes {
	case "string":
		var optionalNumber = arguments[0].(string)
		instance_ = ast.Limit().Make(
			optionalNumber,
		)
	default:
		var message = fmt.Sprintf(
			"No Limit constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func Line(arguments ...any) LineLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case IdentifierLike:
			argumentTypes += "IdentifierLike, "
		case string:
			argumentTypes += "string, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Line constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma and space.
		argumentTypes = argumentTypes[:length-2]
	}

	// Call the corresponding constructor.
	var instance_ LineLike
	switch argumentTypes {
	case "IdentifierLike, string":
		var identifier = arguments[0].(IdentifierLike)
		var optionalNote = arguments[1].(string)
		instance_ = ast.Line().Make(
			identifier,
			optionalNote,
		)
	default:
		var message = fmt.Sprintf(
			"No Line constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func Multiline(arguments ...any) MultilineLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case abs.Sequential[LineLike]:
			argumentTypes += "abs.Sequential[LineLike], "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Multiline constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma and space.
		argumentTypes = argumentTypes[:length-2]
	}

	// Call the corresponding constructor.
	var instance_ MultilineLike
	switch argumentTypes {
	case "abs.Sequential[LineLike]":
		var lines = arguments[0].(abs.Sequential[LineLike])
		instance_ = ast.Multiline().Make(
			lines,
		)
	default:
		var message = fmt.Sprintf(
			"No Multiline constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func Notice(arguments ...any) NoticeLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			argumentTypes += "string, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Notice constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma and space.
		argumentTypes = argumentTypes[:length-2]
	}

	// Call the corresponding constructor.
	var instance_ NoticeLike
	switch argumentTypes {
	case "string":
		var comment = arguments[0].(string)
		instance_ = ast.Notice().Make(
			comment,
		)
	default:
		var message = fmt.Sprintf(
			"No Notice constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func Option(arguments ...any) OptionLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case abs.Sequential[RepetitionLike]:
			argumentTypes += "abs.Sequential[RepetitionLike], "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Option constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma and space.
		argumentTypes = argumentTypes[:length-2]
	}

	// Call the corresponding constructor.
	var instance_ OptionLike
	switch argumentTypes {
	case "abs.Sequential[RepetitionLike]":
		var repetitions = arguments[0].(abs.Sequential[RepetitionLike])
		instance_ = ast.Option().Make(
			repetitions,
		)
	default:
		var message = fmt.Sprintf(
			"No Option constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func Pattern(arguments ...any) PatternLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case OptionLike:
			argumentTypes += "OptionLike, "
		case abs.Sequential[AlternativeLike]:
			argumentTypes += "abs.Sequential[AlternativeLike], "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Pattern constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma and space.
		argumentTypes = argumentTypes[:length-2]
	}

	// Call the corresponding constructor.
	var instance_ PatternLike
	switch argumentTypes {
	case "OptionLike, abs.Sequential[AlternativeLike]":
		var option = arguments[0].(OptionLike)
		var alternatives = arguments[1].(abs.Sequential[AlternativeLike])
		instance_ = ast.Pattern().Make(
			option,
			alternatives,
		)
	default:
		var message = fmt.Sprintf(
			"No Pattern constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func Quantified(arguments ...any) QuantifiedLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			argumentTypes += "string, "
		case LimitLike:
			argumentTypes += "LimitLike, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Quantified constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma and space.
		argumentTypes = argumentTypes[:length-2]
	}

	// Call the corresponding constructor.
	var instance_ QuantifiedLike
	switch argumentTypes {
	case "string, LimitLike":
		var number = arguments[0].(string)
		var optionalLimit = arguments[1].(LimitLike)
		instance_ = ast.Quantified().Make(
			number,
			optionalLimit,
		)
	default:
		var message = fmt.Sprintf(
			"No Quantified constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func Reference(arguments ...any) ReferenceLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case IdentifierLike:
			argumentTypes += "IdentifierLike, "
		case CardinalityLike:
			argumentTypes += "CardinalityLike, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Reference constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma and space.
		argumentTypes = argumentTypes[:length-2]
	}

	// Call the corresponding constructor.
	var instance_ ReferenceLike
	switch argumentTypes {
	case "IdentifierLike, CardinalityLike":
		var identifier = arguments[0].(IdentifierLike)
		var optionalCardinality = arguments[1].(CardinalityLike)
		instance_ = ast.Reference().Make(
			identifier,
			optionalCardinality,
		)
	default:
		var message = fmt.Sprintf(
			"No Reference constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func Repetition(arguments ...any) RepetitionLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case ElementLike:
			argumentTypes += "ElementLike, "
		case CardinalityLike:
			argumentTypes += "CardinalityLike, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Repetition constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma and space.
		argumentTypes = argumentTypes[:length-2]
	}

	// Call the corresponding constructor.
	var instance_ RepetitionLike
	switch argumentTypes {
	case "ElementLike, CardinalityLike":
		var element = arguments[0].(ElementLike)
		var optionalCardinality = arguments[1].(CardinalityLike)
		instance_ = ast.Repetition().Make(
			element,
			optionalCardinality,
		)
	default:
		var message = fmt.Sprintf(
			"No Repetition constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func Rule(arguments ...any) RuleLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			argumentTypes += "string, "
		case DefinitionLike:
			argumentTypes += "DefinitionLike, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Rule constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma and space.
		argumentTypes = argumentTypes[:length-2]
	}

	// Call the corresponding constructor.
	var instance_ RuleLike
	switch argumentTypes {
	case "string, DefinitionLike":
		var uppercase = arguments[0].(string)
		var definition = arguments[1].(DefinitionLike)
		instance_ = ast.Rule().Make(
			uppercase,
			definition,
		)
	default:
		var message = fmt.Sprintf(
			"No Rule constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func Syntax(arguments ...any) SyntaxLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case NoticeLike:
			argumentTypes += "NoticeLike, "
		case string:
			argumentTypes += "string, "
		case abs.Sequential[RuleLike]:
			argumentTypes += "abs.Sequential[RuleLike], "
		case abs.Sequential[ExpressionLike]:
			argumentTypes += "abs.Sequential[ExpressionLike], "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Syntax constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma and space.
		argumentTypes = argumentTypes[:length-2]
	}

	// Call the corresponding constructor.
	var instance_ SyntaxLike
	switch argumentTypes {
	case "NoticeLike, string, abs.Sequential[RuleLike], string, abs.Sequential[ExpressionLike]":
		var notice = arguments[0].(NoticeLike)
		var comment1 = arguments[1].(string)
		var rules = arguments[2].(abs.Sequential[RuleLike])
		var comment2 = arguments[3].(string)
		var expressions = arguments[4].(abs.Sequential[ExpressionLike])
		instance_ = ast.Syntax().Make(
			notice,
			comment1,
			rules,
			comment2,
			expressions,
		)
	default:
		var message = fmt.Sprintf(
			"No Syntax constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func Term(arguments ...any) TermLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case any:
			argumentTypes += "any, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Term constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma and space.
		argumentTypes = argumentTypes[:length-2]
	}

	// Call the corresponding constructor.
	var instance_ TermLike
	switch argumentTypes {
	case "any":
		var any_ = arguments[0]
		instance_ = ast.Term().Make(
			any_,
		)
	default:
		var message = fmt.Sprintf(
			"No Term constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func Text(arguments ...any) TextLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case any:
			argumentTypes += "any, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Text constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma and space.
		argumentTypes = argumentTypes[:length-2]
	}

	// Call the corresponding constructor.
	var instance_ TextLike
	switch argumentTypes {
	case "any":
		var any_ = arguments[0]
		instance_ = ast.Text().Make(
			any_,
		)
	default:
		var message = fmt.Sprintf(
			"No Text constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

// Grammar

func Formatter(arguments ...any) FormatterLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Formatter constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma and space.
		argumentTypes = argumentTypes[:length-2]
	}

	// Call the corresponding constructor.
	var instance_ FormatterLike
	switch argumentTypes {
	case "":
		instance_ = gra.Formatter().Make()
	default:
		var message = fmt.Sprintf(
			"No Formatter constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func Parser(arguments ...any) ParserLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Parser constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma and space.
		argumentTypes = argumentTypes[:length-2]
	}

	// Call the corresponding constructor.
	var instance_ ParserLike
	switch argumentTypes {
	case "":
		instance_ = gra.Parser().Make()
	default:
		var message = fmt.Sprintf(
			"No Parser constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func Processor(arguments ...any) ProcessorLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Processor constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma and space.
		argumentTypes = argumentTypes[:length-2]
	}

	// Call the corresponding constructor.
	var instance_ ProcessorLike
	switch argumentTypes {
	case "":
		instance_ = gra.Processor().Make()
	default:
		var message = fmt.Sprintf(
			"No Processor constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func Scanner(arguments ...any) ScannerLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			argumentTypes += "string, "
		case abs.QueueLike[TokenLike]:
			argumentTypes += "abs.QueueLike[TokenLike], "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Scanner constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma and space.
		argumentTypes = argumentTypes[:length-2]
	}

	// Call the corresponding constructor.
	var instance_ ScannerLike
	switch argumentTypes {
	case "string, abs.QueueLike[TokenLike]":
		var source = arguments[0].(string)
		var tokens = arguments[1].(abs.QueueLike[TokenLike])
		instance_ = gra.Scanner().Make(
			source,
			tokens,
		)
	default:
		var message = fmt.Sprintf(
			"No Scanner constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func Token(arguments ...any) TokenLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case uint:
			argumentTypes += "uint, "
		case TokenType:
			argumentTypes += "TokenType, "
		case string:
			argumentTypes += "string, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Token constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma and space.
		argumentTypes = argumentTypes[:length-2]
	}

	// Call the corresponding constructor.
	var instance_ TokenLike
	switch argumentTypes {
	case "uint, uint, TokenType, string":
		var line = arguments[0].(uint)
		var position = arguments[1].(uint)
		var type_ = arguments[2].(TokenType)
		var value = arguments[3].(string)
		instance_ = gra.Token().Make(
			line,
			position,
			type_,
			value,
		)
	default:
		var message = fmt.Sprintf(
			"No Token constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func Validator(arguments ...any) ValidatorLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Validator constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma and space.
		argumentTypes = argumentTypes[:length-2]
	}

	// Call the corresponding constructor.
	var instance_ ValidatorLike
	switch argumentTypes {
	case "":
		instance_ = gra.Validator().Make()
	default:
		var message = fmt.Sprintf(
			"No Validator constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
}

func Visitor(arguments ...any) VisitorLike {
	// Analyze the arguments.
	var argumentTypes string
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case Methodical:
			argumentTypes += "Methodical, "
		default:
			var message = fmt.Sprintf(
				"An unexpected argument type was passed into the Visitor constructor: %v of type %T",
				argument,
				actual,
			)
			panic(message)
		}
	}
	var length = len(argumentTypes)
	if length > 0 {
		// Remove the trailing comma and space.
		argumentTypes = argumentTypes[:length-2]
	}

	// Call the corresponding constructor.
	var instance_ VisitorLike
	switch argumentTypes {
	case "Methodical":
		var processor = arguments[0].(Methodical)
		instance_ = gra.Visitor().Make(
			processor,
		)
	default:
		var message = fmt.Sprintf(
			"No Visitor constructor matching the arguments was found: %v\n",
			arguments,
		)
		panic(message)
	}
	return instance_
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
	var scannerClass = gra.Scanner()
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
