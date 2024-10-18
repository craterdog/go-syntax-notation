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

package grammar

import (
	fmt "fmt"
	col "github.com/craterdog/go-collection-framework/v4"
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	uti "github.com/craterdog/go-missing-utilities/v2"
	ast "github.com/craterdog/go-syntax-notation/v5/ast"
	sts "strings"
)

// CLASS INTERFACE

// Access Function

func Parser() ParserClassLike {
	return parserReference()
}

// Constructor Methods

func (c *parserClass_) Make() ParserLike {
	var instance = &parser_{
		// Initialize the instance attributes.
	}
	return instance

}

// INSTANCE INTERFACE

// Primary Methods

func (v *parser_) GetClass() ParserClassLike {
	return parserReference()
}

func (v *parser_) ParseSource(
	source string,
) ast.SyntaxLike {
	v.source_ = source
	v.tokens_ = col.Queue[TokenLike](parserReference().queueSize_)
	v.next_ = col.Stack[TokenLike](parserReference().stackSize_)

	// The scanner runs in a separate Go routine.
	Scanner().Make(v.source_, v.tokens_)

	// Attempt to parse the syntax.
	var syntax, token, ok = v.parseSyntax()
	if !ok {
		var message = v.formatError(token, "Syntax")
		panic(message)
	}
	return syntax
}

// PROTECTED INTERFACE

// Private Methods

func (v *parser_) parseAlternative() (
	alternative ast.AlternativeLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool
	var tokens_ = col.List[TokenLike]()

	// Attempt to parse a single "|" delimiter.
	_, token, ok = v.parseDelimiter("|")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Alternative")
			panic(message)
		} else {
			// This is not a single alternative rule.
			v.putBack(tokens_)
			return alternative, token, false
		}
	}
	tokens_.AppendValue(token)

	// Attempt to parse a single option rule.
	var option ast.OptionLike
	option, token, ok = v.parseOption()
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Alternative")
			panic(message)
		} else {
			// This is not a single alternative rule.
			v.putBack(tokens_)
			return alternative, token, false
		}
	}

	// Found a single alternative rule.
	ruleFound_ = true
	alternative = ast.Alternative().Make(option)
	return alternative, token, ruleFound_
}

func (v *parser_) parseCardinality() (
	cardinality ast.CardinalityLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single constrained rule.
	var constrained ast.ConstrainedLike
	constrained, token, ok = v.parseConstrained()
	if ok {
		// Found a single constrained cardinality.
		cardinality = ast.Cardinality().Make(constrained)
		return cardinality, token, true
	}

	// Attempt to parse a single quantified rule.
	var quantified ast.QuantifiedLike
	quantified, token, ok = v.parseQuantified()
	if ok {
		// Found a single quantified cardinality.
		cardinality = ast.Cardinality().Make(quantified)
		return cardinality, token, true
	}

	// This is not a single cardinality rule.
	return cardinality, token, false

}

func (v *parser_) parseCharacter() (
	character ast.CharacterLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single explicit rule.
	var explicit ast.ExplicitLike
	explicit, token, ok = v.parseExplicit()
	if ok {
		// Found a single explicit character.
		character = ast.Character().Make(explicit)
		return character, token, true
	}

	// Attempt to parse a single intrinsic token.
	var intrinsic string
	intrinsic, token, ok = v.parseToken(IntrinsicToken)
	if ok {
		// Found a single intrinsic character.
		character = ast.Character().Make(intrinsic)
		return character, token, true
	}

	// This is not a single character rule.
	return character, token, false

}

func (v *parser_) parseConstrained() (
	constrained ast.ConstrainedLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single optional token.
	var optional string
	optional, token, ok = v.parseToken(OptionalToken)
	if ok {
		// Found a single optional constrained.
		constrained = ast.Constrained().Make(optional)
		return constrained, token, true
	}

	// Attempt to parse a single repeated token.
	var repeated string
	repeated, token, ok = v.parseToken(RepeatedToken)
	if ok {
		// Found a single repeated constrained.
		constrained = ast.Constrained().Make(repeated)
		return constrained, token, true
	}

	// This is not a single constrained rule.
	return constrained, token, false

}

func (v *parser_) parseDefinition() (
	definition ast.DefinitionLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single multiline rule.
	var multiline ast.MultilineLike
	multiline, token, ok = v.parseMultiline()
	if ok {
		// Found a single multiline definition.
		definition = ast.Definition().Make(multiline)
		return definition, token, true
	}

	// Attempt to parse a single inline rule.
	var inline ast.InlineLike
	inline, token, ok = v.parseInline()
	if ok {
		// Found a single inline definition.
		definition = ast.Definition().Make(inline)
		return definition, token, true
	}

	// This is not a single definition rule.
	return definition, token, false

}

func (v *parser_) parseElement() (
	element ast.ElementLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single group rule.
	var group ast.GroupLike
	group, token, ok = v.parseGroup()
	if ok {
		// Found a single group element.
		element = ast.Element().Make(group)
		return element, token, true
	}

	// Attempt to parse a single filter rule.
	var filter ast.FilterLike
	filter, token, ok = v.parseFilter()
	if ok {
		// Found a single filter element.
		element = ast.Element().Make(filter)
		return element, token, true
	}

	// Attempt to parse a single text rule.
	var text ast.TextLike
	text, token, ok = v.parseText()
	if ok {
		// Found a single text element.
		element = ast.Element().Make(text)
		return element, token, true
	}

	// This is not a single element rule.
	return element, token, false

}

func (v *parser_) parseExplicit() (
	explicit ast.ExplicitLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool
	var tokens_ = col.List[TokenLike]()

	// Attempt to parse a single glyph token.
	var glyph string
	glyph, token, ok = v.parseToken(GlyphToken)
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Explicit")
			panic(message)
		} else {
			// This is not a single explicit rule.
			v.putBack(tokens_)
			return explicit, token, false
		}
	}
	tokens_.AppendValue(token)

	// Attempt to parse an optional extent rule.
	var optionalExtent ast.ExtentLike
	optionalExtent, _, ok = v.parseExtent()
	if ok {
		tokens_.AppendValue(token)
	}

	// Found a single explicit rule.
	ruleFound_ = true
	explicit = ast.Explicit().Make(
		glyph,
		optionalExtent,
	)
	return explicit, token, ruleFound_
}

func (v *parser_) parseExpression() (
	expression ast.ExpressionLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool
	var tokens_ = col.List[TokenLike]()

	// Attempt to parse a single "$" delimiter.
	_, token, ok = v.parseDelimiter("$")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Expression")
			panic(message)
		} else {
			// This is not a single expression rule.
			v.putBack(tokens_)
			return expression, token, false
		}
	}
	tokens_.AppendValue(token)

	// Attempt to parse a single lowercase token.
	var lowercase string
	lowercase, token, ok = v.parseToken(LowercaseToken)
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Expression")
			panic(message)
		} else {
			// This is not a single expression rule.
			v.putBack(tokens_)
			return expression, token, false
		}
	}
	tokens_.AppendValue(token)

	// Attempt to parse a single ":" delimiter.
	_, token, ok = v.parseDelimiter(":")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Expression")
			panic(message)
		} else {
			// This is not a single expression rule.
			v.putBack(tokens_)
			return expression, token, false
		}
	}
	tokens_.AppendValue(token)

	// Attempt to parse a single pattern rule.
	var pattern ast.PatternLike
	pattern, token, ok = v.parsePattern()
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Expression")
			panic(message)
		} else {
			// This is not a single expression rule.
			v.putBack(tokens_)
			return expression, token, false
		}
	}

	// Attempt to parse an optional note token.
	var optionalNote string
	optionalNote, _, ok = v.parseToken(NoteToken)
	if ok {
		tokens_.AppendValue(token)
	}

	// Found a single expression rule.
	ruleFound_ = true
	expression = ast.Expression().Make(
		lowercase,
		pattern,
		optionalNote,
	)
	return expression, token, ruleFound_
}

func (v *parser_) parseExtent() (
	extent ast.ExtentLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool
	var tokens_ = col.List[TokenLike]()

	// Attempt to parse a single ".." delimiter.
	_, token, ok = v.parseDelimiter("..")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Extent")
			panic(message)
		} else {
			// This is not a single extent rule.
			v.putBack(tokens_)
			return extent, token, false
		}
	}
	tokens_.AppendValue(token)

	// Attempt to parse a single glyph token.
	var glyph string
	glyph, token, ok = v.parseToken(GlyphToken)
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Extent")
			panic(message)
		} else {
			// This is not a single extent rule.
			v.putBack(tokens_)
			return extent, token, false
		}
	}
	tokens_.AppendValue(token)

	// Found a single extent rule.
	ruleFound_ = true
	extent = ast.Extent().Make(glyph)
	return extent, token, ruleFound_
}

func (v *parser_) parseFilter() (
	filter ast.FilterLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool
	var tokens_ = col.List[TokenLike]()

	// Attempt to parse an optional excluded token.
	var optionalExcluded string
	optionalExcluded, _, ok = v.parseToken(ExcludedToken)
	if ok {
		tokens_.AppendValue(token)
	}

	// Attempt to parse a single "[" delimiter.
	_, token, ok = v.parseDelimiter("[")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Filter")
			panic(message)
		} else {
			// This is not a single filter rule.
			v.putBack(tokens_)
			return filter, token, false
		}
	}
	tokens_.AppendValue(token)

	// Attempt to parse 1 to unlimited character rules.
	var characters = col.List[ast.CharacterLike]()
charactersLoop:
	for numberFound_ := 0; numberFound_ < parserReference().unlimited_; numberFound_++ {
		var character ast.CharacterLike
		character, token, ok = v.parseCharacter()
		if !ok {
			switch {
			case numberFound_ < 1:
				if !ruleFound_ {
					// This is not a single filter rule.
					v.putBack(tokens_)
					return filter, token, false
				}
				// Found a syntax error.
				var message = v.formatError(token, "Filter")
				message += "The number of character rules must be at least 1."
				panic(message)
			default:
				break charactersLoop
			}
		}
		ruleFound_ = true
		characters.AppendValue(character)
	}

	// Attempt to parse a single "]" delimiter.
	_, token, ok = v.parseDelimiter("]")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Filter")
			panic(message)
		} else {
			// This is not a single filter rule.
			v.putBack(tokens_)
			return filter, token, false
		}
	}
	tokens_.AppendValue(token)

	// Found a single filter rule.
	ruleFound_ = true
	filter = ast.Filter().Make(
		optionalExcluded,
		characters,
	)
	return filter, token, ruleFound_
}

func (v *parser_) parseGroup() (
	group ast.GroupLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool
	var tokens_ = col.List[TokenLike]()

	// Attempt to parse a single "(" delimiter.
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Group")
			panic(message)
		} else {
			// This is not a single group rule.
			v.putBack(tokens_)
			return group, token, false
		}
	}
	tokens_.AppendValue(token)

	// Attempt to parse a single pattern rule.
	var pattern ast.PatternLike
	pattern, token, ok = v.parsePattern()
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Group")
			panic(message)
		} else {
			// This is not a single group rule.
			v.putBack(tokens_)
			return group, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse a single ")" delimiter.
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Group")
			panic(message)
		} else {
			// This is not a single group rule.
			v.putBack(tokens_)
			return group, token, false
		}
	}
	tokens_.AppendValue(token)

	// Found a single group rule.
	ruleFound_ = true
	group = ast.Group().Make(pattern)
	return group, token, ruleFound_
}

func (v *parser_) parseIdentifier() (
	identifier ast.IdentifierLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single lowercase token.
	var lowercase string
	lowercase, token, ok = v.parseToken(LowercaseToken)
	if ok {
		// Found a single lowercase identifier.
		identifier = ast.Identifier().Make(lowercase)
		return identifier, token, true
	}

	// Attempt to parse a single uppercase token.
	var uppercase string
	uppercase, token, ok = v.parseToken(UppercaseToken)
	if ok {
		// Found a single uppercase identifier.
		identifier = ast.Identifier().Make(uppercase)
		return identifier, token, true
	}

	// This is not a single identifier rule.
	return identifier, token, false

}

func (v *parser_) parseInline() (
	inline ast.InlineLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool
	var tokens_ = col.List[TokenLike]()

	// Attempt to parse 1 to unlimited term rules.
	var terms = col.List[ast.TermLike]()
termsLoop:
	for numberFound_ := 0; numberFound_ < parserReference().unlimited_; numberFound_++ {
		var term ast.TermLike
		term, token, ok = v.parseTerm()
		if !ok {
			switch {
			case numberFound_ < 1:
				if !ruleFound_ {
					// This is not a single inline rule.
					v.putBack(tokens_)
					return inline, token, false
				}
				// Found a syntax error.
				var message = v.formatError(token, "Inline")
				message += "The number of term rules must be at least 1."
				panic(message)
			default:
				break termsLoop
			}
		}
		ruleFound_ = true
		terms.AppendValue(term)
	}

	// Attempt to parse an optional note token.
	var optionalNote string
	optionalNote, _, ok = v.parseToken(NoteToken)
	if ok {
		tokens_.AppendValue(token)
	}

	// Found a single inline rule.
	ruleFound_ = true
	inline = ast.Inline().Make(
		terms,
		optionalNote,
	)
	return inline, token, ruleFound_
}

func (v *parser_) parseLimit() (
	limit ast.LimitLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool
	var tokens_ = col.List[TokenLike]()

	// Attempt to parse a single ".." delimiter.
	_, token, ok = v.parseDelimiter("..")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Limit")
			panic(message)
		} else {
			// This is not a single limit rule.
			v.putBack(tokens_)
			return limit, token, false
		}
	}
	tokens_.AppendValue(token)

	// Attempt to parse an optional number token.
	var optionalNumber string
	optionalNumber, _, ok = v.parseToken(NumberToken)
	if ok {
		tokens_.AppendValue(token)
	}

	// Found a single limit rule.
	ruleFound_ = true
	limit = ast.Limit().Make(optionalNumber)
	return limit, token, ruleFound_
}

func (v *parser_) parseLine() (
	line ast.LineLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool
	var tokens_ = col.List[TokenLike]()

	// Attempt to parse a single "-" delimiter.
	_, token, ok = v.parseDelimiter("-")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Line")
			panic(message)
		} else {
			// This is not a single line rule.
			v.putBack(tokens_)
			return line, token, false
		}
	}
	tokens_.AppendValue(token)

	// Attempt to parse a single identifier rule.
	var identifier ast.IdentifierLike
	identifier, token, ok = v.parseIdentifier()
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Line")
			panic(message)
		} else {
			// This is not a single line rule.
			v.putBack(tokens_)
			return line, token, false
		}
	}

	// Attempt to parse an optional note token.
	var optionalNote string
	optionalNote, _, ok = v.parseToken(NoteToken)
	if ok {
		tokens_.AppendValue(token)
	}

	// Found a single line rule.
	ruleFound_ = true
	line = ast.Line().Make(
		identifier,
		optionalNote,
	)
	return line, token, ruleFound_
}

func (v *parser_) parseMultiline() (
	multiline ast.MultilineLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool
	var tokens_ = col.List[TokenLike]()

	// Attempt to parse 1 to unlimited line rules.
	var lines = col.List[ast.LineLike]()
linesLoop:
	for numberFound_ := 0; numberFound_ < parserReference().unlimited_; numberFound_++ {
		var line ast.LineLike
		line, token, ok = v.parseLine()
		if !ok {
			switch {
			case numberFound_ < 1:
				if !ruleFound_ {
					// This is not a single multiline rule.
					v.putBack(tokens_)
					return multiline, token, false
				}
				// Found a syntax error.
				var message = v.formatError(token, "Multiline")
				message += "The number of line rules must be at least 1."
				panic(message)
			default:
				break linesLoop
			}
		}
		ruleFound_ = true
		lines.AppendValue(line)
	}

	// Found a single multiline rule.
	ruleFound_ = true
	multiline = ast.Multiline().Make(
		lines,
	)
	return multiline, token, ruleFound_
}

func (v *parser_) parseNotice() (
	notice ast.NoticeLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool
	var tokens_ = col.List[TokenLike]()

	// Attempt to parse a single comment token.
	var comment string
	comment, token, ok = v.parseToken(CommentToken)
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Notice")
			panic(message)
		} else {
			// This is not a single notice rule.
			v.putBack(tokens_)
			return notice, token, false
		}
	}
	tokens_.AppendValue(token)

	// Found a single notice rule.
	ruleFound_ = true
	notice = ast.Notice().Make(
		comment,
	)
	return notice, token, ruleFound_
}

func (v *parser_) parseOption() (
	option ast.OptionLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool
	var tokens_ = col.List[TokenLike]()

	// Attempt to parse 1 to unlimited repetition rules.
	var repetitions = col.List[ast.RepetitionLike]()
repetitionsLoop:
	for numberFound_ := 0; numberFound_ < parserReference().unlimited_; numberFound_++ {
		var repetition ast.RepetitionLike
		repetition, token, ok = v.parseRepetition()
		if !ok {
			switch {
			case numberFound_ < 1:
				if !ruleFound_ {
					// This is not a single option rule.
					v.putBack(tokens_)
					return option, token, false
				}
				// Found a syntax error.
				var message = v.formatError(token, "Option")
				message += "The number of repetition rules must be at least 1."
				panic(message)
			default:
				break repetitionsLoop
			}
		}
		ruleFound_ = true
		repetitions.AppendValue(repetition)
	}

	// Found a single option rule.
	ruleFound_ = true
	option = ast.Option().Make(repetitions)
	return option, token, ruleFound_
}

func (v *parser_) parsePattern() (
	pattern ast.PatternLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool
	var tokens_ = col.List[TokenLike]()

	// Attempt to parse a single option rule.
	var option ast.OptionLike
	option, token, ok = v.parseOption()
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Pattern")
			panic(message)
		} else {
			// This is not a single pattern rule.
			v.putBack(tokens_)
			return pattern, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse 0 to unlimited alternative rules.
	var alternatives = col.List[ast.AlternativeLike]()
alternativesLoop:
	for numberFound_ := 0; numberFound_ < parserReference().unlimited_; numberFound_++ {
		var alternative ast.AlternativeLike
		alternative, token, ok = v.parseAlternative()
		if !ok {
			switch {
			case numberFound_ < 0:
				if !ruleFound_ {
					// This is not a single pattern rule.
					v.putBack(tokens_)
					return pattern, token, false
				}
				// Found a syntax error.
				var message = v.formatError(token, "Pattern")
				message += "The number of alternative rules must be at least 0."
				panic(message)
			default:
				break alternativesLoop
			}
		}
		alternatives.AppendValue(alternative)
	}

	// Found a single pattern rule.
	ruleFound_ = true
	pattern = ast.Pattern().Make(
		option,
		alternatives,
	)
	return pattern, token, ruleFound_
}

func (v *parser_) parseQuantified() (
	quantified ast.QuantifiedLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool
	var tokens_ = col.List[TokenLike]()

	// Attempt to parse a single "{" delimiter.
	_, token, ok = v.parseDelimiter("{")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Quantified")
			panic(message)
		} else {
			// This is not a single quantified rule.
			v.putBack(tokens_)
			return quantified, token, false
		}
	}
	tokens_.AppendValue(token)

	// Attempt to parse a single number token.
	var number string
	number, token, ok = v.parseToken(NumberToken)
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Quantified")
			panic(message)
		} else {
			// This is not a single quantified rule.
			v.putBack(tokens_)
			return quantified, token, false
		}
	}
	tokens_.AppendValue(token)

	// Attempt to parse an optional limit rule.
	var optionalLimit ast.LimitLike
	optionalLimit, _, ok = v.parseLimit()
	if ok {
		ruleFound_ = true
	}

	// Attempt to parse a single "}" delimiter.
	_, token, ok = v.parseDelimiter("}")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Quantified")
			panic(message)
		} else {
			// This is not a single quantified rule.
			v.putBack(tokens_)
			return quantified, token, false
		}
	}
	tokens_.AppendValue(token)

	// Found a single quantified rule.
	ruleFound_ = true
	quantified = ast.Quantified().Make(
		number,
		optionalLimit,
	)
	return quantified, token, ruleFound_
}

func (v *parser_) parseReference() (
	reference ast.ReferenceLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool
	var tokens_ = col.List[TokenLike]()

	// Attempt to parse a single identifier rule.
	var identifier ast.IdentifierLike
	identifier, token, ok = v.parseIdentifier()
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Reference")
			panic(message)
		} else {
			// This is not a single reference rule.
			v.putBack(tokens_)
			return reference, token, false
		}
	}

	// Attempt to parse an optional cardinality rule.
	var optionalCardinality ast.CardinalityLike
	optionalCardinality, _, _ = v.parseCardinality()

	// Found a single reference rule.
	ruleFound_ = true
	reference = ast.Reference().Make(
		identifier,
		optionalCardinality,
	)
	return reference, token, ruleFound_
}

func (v *parser_) parseRepetition() (
	repetition ast.RepetitionLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool
	var tokens_ = col.List[TokenLike]()

	// Attempt to parse a single element rule.
	var element ast.ElementLike
	element, token, ok = v.parseElement()
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Repetition")
			panic(message)
		} else {
			// This is not a single repetition rule.
			v.putBack(tokens_)
			return repetition, token, false
		}
	}

	// Attempt to parse an optional cardinality rule.
	var optionalCardinality ast.CardinalityLike
	optionalCardinality, _, _ = v.parseCardinality()

	// Found a single repetition rule.
	ruleFound_ = true
	repetition = ast.Repetition().Make(
		element,
		optionalCardinality,
	)
	return repetition, token, ruleFound_
}

func (v *parser_) parseRule() (
	rule ast.RuleLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool
	var tokens_ = col.List[TokenLike]()

	// Attempt to parse a single "$" delimiter.
	_, token, ok = v.parseDelimiter("$")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Rule")
			panic(message)
		} else {
			// This is not a single rule rule.
			v.putBack(tokens_)
			return rule, token, false
		}
	}
	tokens_.AppendValue(token)

	// Attempt to parse a single uppercase token.
	var uppercase string
	uppercase, token, ok = v.parseToken(UppercaseToken)
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Rule")
			panic(message)
		} else {
			// This is not a single rule rule.
			v.putBack(tokens_)
			return rule, token, false
		}
	}
	tokens_.AppendValue(token)

	// Attempt to parse a single ":" delimiter.
	_, token, ok = v.parseDelimiter(":")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Rule")
			panic(message)
		} else {
			// This is not a single rule rule.
			v.putBack(tokens_)
			return rule, token, false
		}
	}
	tokens_.AppendValue(token)

	// Attempt to parse a single definition rule.
	var definition ast.DefinitionLike
	definition, token, ok = v.parseDefinition()
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Rule")
			panic(message)
		} else {
			// This is not a single rule rule.
			v.putBack(tokens_)
			return rule, token, false
		}
	}

	// Found a single rule rule.
	ruleFound_ = true
	rule = ast.Rule().Make(
		uppercase,
		definition,
	)
	return rule, token, ruleFound_
}

func (v *parser_) parseSyntax() (
	syntax ast.SyntaxLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool
	var tokens_ = col.List[TokenLike]()

	// Attempt to parse a single notice rule.
	var notice ast.NoticeLike
	notice, token, ok = v.parseNotice()
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Syntax")
			panic(message)
		} else {
			// This is not a single syntax rule.
			v.putBack(tokens_)
			return syntax, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse a single comment token.
	var comment1 string
	comment1, token, ok = v.parseToken(CommentToken)
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Syntax")
			panic(message)
		} else {
			// This is not a single syntax rule.
			v.putBack(tokens_)
			return syntax, token, false
		}
	}
	tokens_.AppendValue(token)

	// Attempt to parse 1 to unlimited rule rules.
	var rules = col.List[ast.RuleLike]()
rulesLoop:
	for numberFound_ := 0; numberFound_ < parserReference().unlimited_; numberFound_++ {
		var rule ast.RuleLike
		rule, token, ok = v.parseRule()
		if !ok {
			switch {
			case numberFound_ < 1:
				if !ruleFound_ {
					// This is not a single syntax rule.
					v.putBack(tokens_)
					return syntax, token, false
				}
				// Found a syntax error.
				var message = v.formatError(token, "Syntax")
				message += "The number of rule rules must be at least 1."
				panic(message)
			default:
				break rulesLoop
			}
		}
		rules.AppendValue(rule)
	}

	// Attempt to parse a single comment token.
	var comment2 string
	comment2, token, ok = v.parseToken(CommentToken)
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Syntax")
			panic(message)
		} else {
			// This is not a single syntax rule.
			v.putBack(tokens_)
			return syntax, token, false
		}
	}
	tokens_.AppendValue(token)

	// Attempt to parse 1 to unlimited expression rules.
	var expressions = col.List[ast.ExpressionLike]()
expressionsLoop:
	for numberFound_ := 0; numberFound_ < parserReference().unlimited_; numberFound_++ {
		var expression ast.ExpressionLike
		expression, token, ok = v.parseExpression()
		if !ok {
			switch {
			case numberFound_ < 1:
				if !ruleFound_ {
					// This is not a single syntax rule.
					v.putBack(tokens_)
					return syntax, token, false
				}
				// Found a syntax error.
				var message = v.formatError(token, "Syntax")
				message += "The number of expression rules must be at least 1."
				panic(message)
			default:
				break expressionsLoop
			}
		}
		expressions.AppendValue(expression)
	}

	// Found a single syntax rule.
	ruleFound_ = true
	syntax = ast.Syntax().Make(
		notice,
		comment1,
		rules,
		comment2,
		expressions,
	)
	return syntax, token, ruleFound_
}

func (v *parser_) parseTerm() (
	term ast.TermLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single reference rule.
	var reference ast.ReferenceLike
	reference, token, ok = v.parseReference()
	if ok {
		// Found a single reference term.
		term = ast.Term().Make(reference)
		return term, token, true
	}

	// Attempt to parse a single literal token.
	var literal string
	literal, token, ok = v.parseToken(LiteralToken)
	if ok {
		// Found a single literal term.
		term = ast.Term().Make(literal)
		return term, token, true
	}

	// This is not a single term rule.
	return term, token, false

}

func (v *parser_) parseText() (
	text ast.TextLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single intrinsic token.
	var intrinsic string
	intrinsic, token, ok = v.parseToken(IntrinsicToken)
	if ok {
		// Found a single intrinsic text.
		text = ast.Text().Make(intrinsic)
		return text, token, true
	}

	// Attempt to parse a single glyph token.
	var glyph string
	glyph, token, ok = v.parseToken(GlyphToken)
	if ok {
		// Found a single glyph text.
		text = ast.Text().Make(glyph)
		return text, token, true
	}

	// Attempt to parse a single literal token.
	var literal string
	literal, token, ok = v.parseToken(LiteralToken)
	if ok {
		// Found a single literal text.
		text = ast.Text().Make(literal)
		return text, token, true
	}

	// Attempt to parse a single lowercase token.
	var lowercase string
	lowercase, token, ok = v.parseToken(LowercaseToken)
	if ok {
		// Found a single lowercase text.
		text = ast.Text().Make(lowercase)
		return text, token, true
	}

	// This is not a single text rule.
	return text, token, false

}

func (v *parser_) parseDelimiter(expectedValue string) (
	value string,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single delimiter.
	value, token, ok = v.parseToken(DelimiterToken)
	if ok {
		if value == expectedValue {
			// Found the right delimiter.
			return value, token, true
		}
		v.next_.AddValue(token)
	}

	// This is not the right delimiter.
	return value, token, false
}

func (v *parser_) parseToken(tokenType TokenType) (
	value string,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a specific token type.
	var tokens = col.List[TokenLike]()
	token = v.getNextToken()
	for token != nil {
		tokens.AppendValue(token)
		switch token.GetType() {
		case tokenType:
			// Found the right token type.
			value = token.GetValue()
			return value, token, true
		case SpaceToken, NewlineToken:
			// Ignore any unrequested whitespace.
			token = v.getNextToken()
		default:
			// This is not the right token type.
			v.putBack(tokens)
			return value, token, false
		}
	}

	// We are at the end-of-file marker.
	return value, token, false
}

func (v *parser_) formatError(token TokenLike, ruleName string) string {
	// Format the error message.
	var message = fmt.Sprintf(
		"An unexpected token was received by the parser: %v\n",
		Scanner().FormatToken(token),
	)
	var line = token.GetLine()
	var lines = sts.Split(v.source_, "\n")

	// Append the source lines with the error in it.
	message += "\033[36m"
	for index := line - 3; index < line; index++ {
		if index > 1 {
			message += fmt.Sprintf("%04d: ", index) + string(lines[index-1]) + "\n"
		}
	}
	message += fmt.Sprintf("%04d: ", line) + string(lines[line-1]) + "\n"

	// Append an arrow pointing to the error.
	message += " \033[32m>>>─"
	var count uint
	for count < token.GetPosition() {
		message += "─"
		count++
	}
	message += "⌃\033[36m\n"

	// Append the following source line for context.
	if line < uint(len(lines)) {
		message += fmt.Sprintf("%04d: ", line+1) + string(lines[line]) + "\n"
	}
	message += "\033[0m\n"
	if uti.IsDefined(ruleName) {
		message += "Was expecting:\n"
		message += fmt.Sprintf(
			"  \033[32m%v: \033[33m%v\033[0m\n\n",
			ruleName,
			v.getDefinition(ruleName),
		)
	}
	return message
}

func (v *parser_) getDefinition(ruleName string) string {
	return parserReference().syntax_.GetValue(ruleName)
}

func (v *parser_) getNextToken() TokenLike {
	// Check for any read, but unprocessed tokens.
	if !v.next_.IsEmpty() {
		return v.next_.RemoveTop()
	}

	// Read a new token from the token stream.
	var token, ok = v.tokens_.RemoveHead() // This will wait for a token.
	if !ok {
		// The token channel has been closed.
		return nil
	}

	// Check for an error token.
	if token.GetType() == ErrorToken {
		var message = v.formatError(token, "")
		panic(message)
	}

	return token
}

func (v *parser_) putBack(tokens abs.Sequential[TokenLike]) {
	var iterator = tokens.GetIterator()
	for iterator.ToEnd(); iterator.HasPrevious(); {
		var token = iterator.GetPrevious()
		v.next_.AddValue(token)
	}
}

// Instance Structure

type parser_ struct {
	// Declare the instance attributes.
	source_ string                   // The original source code.
	tokens_ abs.QueueLike[TokenLike] // A queue of unread tokens from the scanner.
	next_   abs.StackLike[TokenLike] // A stack of read, but unprocessed tokens.
}

// Class Structure

type parserClass_ struct {
	// Declare the class constants.
	queueSize_ uint
	stackSize_ uint
	unlimited_ int
	syntax_    abs.CatalogLike[string, string]
}

// Class Reference

func parserReference() *parserClass_ {
	return parserReference_
}

var parserReference_ = &parserClass_{
	// Initialize the class constants.
	queueSize_: 16,
	stackSize_: 4,
	unlimited_: 4294967295, // Default to a reasonable value.
	syntax_: col.Catalog[string, string](
		map[string]string{
			"$Syntax": `Notice comment Rule+ comment Expression+`,
			"$Notice": `comment`,
			"$Rule":   `"$" uppercase ":" Definition`,
			"$Definition": `
  - Multiline
  - Inline`,
			"$Multiline": `Line+`,
			"$Line":      `"-" Identifier note?`,
			"$Identifier": `
  - lowercase
  - uppercase`,
			"$Inline": `Term+ note?`,
			"$Term": `
  - Reference
  - literal`,
			"$Reference": `Identifier Cardinality?  ! The default cardinality is one.`,
			"$Cardinality": `
  - Constrained
  - Quantified`,
			"$Constrained": `
  - optional
  - repeated`,
			"$Quantified":  `"{" number Limit? "}"`,
			"$Limit":       `".." number?  ! The limit of a range of numbers is inclusive.`,
			"$Expression":  `"$" lowercase ":" Pattern note?`,
			"$Pattern":     `Option Alternative*`,
			"$Alternative": `"|" Option`,
			"$Option":      `Repetition+`,
			"$Repetition":  `Element Cardinality?  ! The default cardinality is one.`,
			"$Element": `
  - Group
  - Filter
  - Text`,
			"$Group":  `"(" Pattern ")"`,
			"$Filter": `excluded? "[" Character+ "]"`,
			"$Character": `
  - Explicit
  - intrinsic`,
			"$Explicit": `glyph Extent?`,
			"$Extent":   `".." glyph  ! The extent of a range of glyphs is inclusive.`,
			"$Text": `
  - intrinsic
  - glyph
  - literal
  - lowercase`,
		},
	),
}
