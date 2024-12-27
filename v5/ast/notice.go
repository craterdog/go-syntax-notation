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
│                 This class file was automatically generated.                 │
│                     Any updates to it may be overwritten.                    │
└──────────────────────────────────────────────────────────────────────────────┘
*/

package ast

import (
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func NoticeClass() NoticeClassLike {
	return noticeClass()
}

// Constructor Methods

func (c *noticeClass_) Notice(
	comment string,
) NoticeLike {
	if uti.IsUndefined(comment) {
		panic("The \"comment\" attribute is required by this class.")
	}
	var instance = &notice_{
		// Initialize the instance attributes.
		comment_: comment,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *notice_) GetClass() NoticeClassLike {
	return noticeClass()
}

// Attribute Methods

func (v *notice_) GetComment() string {
	return v.comment_
}

// PROTECTED INTERFACE

// Instance Structure

type notice_ struct {
	// Declare the instance attributes.
	comment_ string
}

// Class Structure

type noticeClass_ struct {
	// Declare the class constants.
}

// Class Reference

func noticeClass() *noticeClass_ {
	return noticeClassReference_
}

var noticeClassReference_ = &noticeClass_{
	// Initialize the class constants.
}
