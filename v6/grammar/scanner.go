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

package grammar

import (
	fmt "fmt"
	fra "github.com/craterdog/go-collection-framework/v5"
	col "github.com/craterdog/go-collection-framework/v5/collection"
	uti "github.com/craterdog/go-missing-utilities/v2"
	reg "regexp"
	sts "strings"
	uni "unicode"
)

// CLASS INTERFACE

// Access Function

func ScannerClass() ScannerClassLike {
	return scannerClass()
}

// Constructor Methods

func (c *scannerClass_) Scanner(
	source string,
	tokens col.QueueLike[TokenLike],
) ScannerLike {
	if uti.IsUndefined(source) {
		panic("The \"source\" attribute is required by this class.")
	}
	if uti.IsUndefined(tokens) {
		panic("The \"tokens\" attribute is required by this class.")
	}
	var instance = &scanner_{
		// Initialize the instance attributes.
		line_:     1,
		position_: 1,
		runes_:    []rune(source),
		tokens_:   tokens,
	}
	go instance.scanTokens() // Start scanning tokens in the background.
	return instance
}

// Function Methods

func (c *scannerClass_) FormatToken(
	token TokenLike,
) string {
	var value = token.GetValue()
	value = fmt.Sprintf("%q", value)
	if len(value) > 40 {
		value = fmt.Sprintf("%.40q...", value)
	}
	return fmt.Sprintf(
		"Token [type: %s, line: %d, position: %d]: %s",
		c.tokens_.GetValue(token.GetType()),
		token.GetLine(),
		token.GetPosition(),
		value,
	)
}

func (c *scannerClass_) FormatType(
	tokenType TokenType,
) string {
	return c.tokens_.GetValue(tokenType)
}

func (c *scannerClass_) MatchesType(
	tokenValue string,
	tokenType TokenType,
) bool {
	var matcher = c.matchers_.GetValue(tokenType)
	var match = matcher.FindString(tokenValue)
	return uti.IsDefined(match)
}

// INSTANCE INTERFACE

// Principal Methods

func (v *scanner_) GetClass() ScannerClassLike {
	return scannerClass()
}

// PROTECTED INTERFACE

// Private Methods

func (v *scanner_) emitToken(
	tokenType TokenType,
) {
	var value = string(v.runes_[v.first_:v.next_])
	switch value {
	case "\x00":
		value = "<NULL>"
	case "\a":
		value = "<BELL>"
	case "\b":
		value = "<BKSP>"
	case "\t":
		value = "<HTAB>"
	case "\f":
		value = "<FMFD>"
	case "\r":
		value = "<CRTN>"
	case "\v":
		value = "<VTAB>"
	}
	var token = TokenClass().Token(v.line_, v.position_, tokenType, value)
	//fmt.Println(ScannerClass().FormatToken(token)) // Uncomment when debugging.
	v.tokens_.AddValue(token) // This will block if the queue is full.
}

func (v *scanner_) foundError() {
	v.next_++
	v.emitToken(ErrorToken)
}

func (v *scanner_) foundToken(
	tokenType TokenType,
) bool {
	// Attempt to match the specified token type.
	var class = scannerClass()
	var matcher = class.matchers_.GetValue(tokenType)
	var text = string(v.runes_[v.next_:])
	var match = matcher.FindString(text)
	if uti.IsUndefined(match) {
		return false
	}

	// Check for false delimiter matches.
	var token = []rune(match)
	var length = uint(len(token))
	var previous = token[length-1]
	if tokenType == DelimiterToken && uint(len(v.runes_)) > v.next_+length {
		var next = v.runes_[v.next_+length]
		if (uni.IsLetter(previous) || uni.IsNumber(previous)) &&
			(uni.IsLetter(next) || uni.IsNumber(next) || next == '_') {
			return false
		}
	}

	// Found the requested token type.
	v.next_ += length
	v.emitToken(tokenType)
	var count = uint(sts.Count(match, "\n"))
	if count > 0 {
		v.line_ += count
		v.position_ = v.indexOfLastEol(token)
	} else {
		v.position_ += v.next_ - v.first_
	}
	v.first_ = v.next_
	return true
}

func (v *scanner_) indexOfLastEol(
	runes []rune,
) uint {
	var length = uint(len(runes))
	for index := length; index > 0; index-- {
		if runes[index-1] == '\n' {
			return length - index + 1
		}
	}
	return 0
}

func (v *scanner_) scanTokens() {
loop:
	for v.next_ < uint(len(v.runes_)) {
		switch {
		// Find the next token type.
		case v.foundToken(DelimiterToken):
		case v.foundToken(NewlineToken):
		case v.foundToken(SpaceToken):
		case v.foundToken(CommentToken):
		case v.foundToken(ExcludedToken):
		case v.foundToken(GlyphToken):
		case v.foundToken(IntrinsicToken):
		case v.foundToken(LowercaseToken):
		case v.foundToken(NoteToken):
		case v.foundToken(NumberToken):
		case v.foundToken(OptionalToken):
		case v.foundToken(QuoteToken):
		case v.foundToken(RepeatedToken):
		case v.foundToken(UppercaseToken):
		default:
			v.foundError()
			break loop
		}
	}
	v.tokens_.CloseChannel()
}

// Instance Structure

type scanner_ struct {
	// Declare the instance attributes.
	first_    uint // A zero based index of the first possible rune in the next token.
	next_     uint // A zero based index of the next possible rune in the next token.
	line_     uint // The line number in the source string of the next rune.
	position_ uint // The position in the current line of the next rune.
	runes_    []rune
	tokens_   col.QueueLike[TokenLike]
}

// Class Structure

type scannerClass_ struct {
	// Declare the class constants.
	tokens_   col.CatalogLike[TokenType, string]
	matchers_ col.CatalogLike[TokenType, *reg.Regexp]
}

// Class Reference

func scannerClass() *scannerClass_ {
	return scannerClassReference_
}

var scannerClassReference_ = &scannerClass_{
	// Initialize the class constants.
	tokens_: fra.CatalogFromMap[TokenType, string](
		map[TokenType]string{
			// Define identifiers for each type of token.
			ErrorToken:     "error",
			CommentToken:   "comment",
			DelimiterToken: "delimiter",
			ExcludedToken:  "excluded",
			GlyphToken:     "glyph",
			IntrinsicToken: "intrinsic",
			LowercaseToken: "lowercase",
			NewlineToken:   "newline",
			NoteToken:      "note",
			NumberToken:    "number",
			OptionalToken:  "optional",
			QuoteToken:     "quote",
			RepeatedToken:  "repeated",
			SpaceToken:     "space",
			UppercaseToken: "uppercase",
		},
	),
	matchers_: fra.CatalogFromMap[TokenType, *reg.Regexp](
		map[TokenType]*reg.Regexp{
			// Define pattern matchers for each type of token.
			CommentToken:   reg.MustCompile("^" + comment_),
			DelimiterToken: reg.MustCompile("^" + delimiter_),
			ExcludedToken:  reg.MustCompile("^" + excluded_),
			GlyphToken:     reg.MustCompile("^" + glyph_),
			IntrinsicToken: reg.MustCompile("^" + intrinsic_),
			LowercaseToken: reg.MustCompile("^" + lowercase_),
			NewlineToken:   reg.MustCompile("^" + newline_),
			NoteToken:      reg.MustCompile("^" + note_),
			NumberToken:    reg.MustCompile("^" + number_),
			OptionalToken:  reg.MustCompile("^" + optional_),
			QuoteToken:     reg.MustCompile("^" + quote_),
			RepeatedToken:  reg.MustCompile("^" + repeated_),
			SpaceToken:     reg.MustCompile("^" + space_),
			UppercaseToken: reg.MustCompile("^" + uppercase_),
		},
	),
}

// Private Constants

/*
NOTE:
These private constants define the regular expression sub-patterns that make up
the intrinsic types and token types.  Unfortunately there is no way to make them
private to the scanner class since they must be TRUE Go constants to be used in
this way.  We append an underscore to each name to lessen the chance of a name
collision with other private Go class constants in this package.
*/
const (
	// Define the regular expression patterns for each intrinsic type.
	any_     = "." // This does NOT include newline characters.
	control_ = "\\p{Cc}"
	digit_   = "\\p{Nd}"
	eol_     = "\\r?\\n"
	lower_   = "\\p{Ll}"
	upper_   = "\\p{Lu}"

	// Define the regular expression patterns for each token type.
	delimiter_ = "(?:\\}|\\||\\{|\\]|\\[|\\.\\.|\\)|\\(|\\$|:)"
	newline_   = "(?:" + eol_ + ")"
	space_     = "(?:[ \\t]+)"
	base16_    = "(?:[0-9a-f])"
	comment_   = "(?:!>" + eol_ + "(" + any_ + "|" + eol_ + ")*?" + eol_ + "<!" + eol_ + ")"
	escape_    = "(?:\\\\((?:" + unicode_ + ")|[abfnrtv\"\\\\]))"
	excluded_  = "(?:~)"
	glyph_     = "(?:'[^" + control_ + "]')"
	intrinsic_ = "(?:ANY|CONTROL|DIGIT|EOL|LOWER|UPPER)"
	lowercase_ = "(?:" + lower_ + "(" + digit_ + "|" + lower_ + "|" + upper_ + ")*)"
	note_      = "(?:! [^" + control_ + "]*)"
	number_    = "(?:" + digit_ + "+)"
	optional_  = "(?:\\?)"
	quote_     = "(?:\"((?:" + escape_ + ")|[^\"" + control_ + "])+\")"
	repeated_  = "(?:\\*|\\+)"
	unicode_   = "(?:(x(?:" + base16_ + "){2})|(u(?:" + base16_ + "){4})|(U(?:" + base16_ + "){8}))"
	uppercase_ = "(?:" + upper_ + "(" + digit_ + "|" + lower_ + "|" + upper_ + ")*)"
)
