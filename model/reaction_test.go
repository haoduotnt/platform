// Copyright (c) 2016 Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package model

import (
	"strings"
	"testing"
)

func TestReactionIsValid(t *testing.T) {
	reaction := Reaction{
		UserId:    NewId(),
		PostId:    NewId(),
		EmojiName: "emoji",
	}

	if err := reaction.IsValid(); err != nil {
		t.Fatal(err)
	}

	reaction.UserId = ""
	if err := reaction.IsValid(); err == nil {
		t.Fatal("user id should be invalid")
	}

	reaction.UserId = "1234garbage"
	if err := reaction.IsValid(); err == nil {
		t.Fatal("user id should be invalid")
	}

	reaction.UserId = NewId()
	reaction.PostId = ""
	if err := reaction.IsValid(); err == nil {
		t.Fatal("post id should be invalid")
	}

	reaction.PostId = "1234garbage"
	if err := reaction.IsValid(); err == nil {
		t.Fatal("post id should be invalid")
	}

	reaction.PostId = NewId()
	reaction.EmojiName = ""
	if err := reaction.IsValid(); err == nil {
		t.Fatal("emoji name should be invalid")
	}

	reaction.EmojiName = strings.Repeat("a", 65)
	if err := reaction.IsValid(); err == nil {
		t.Fatal("emoji name should be invalid")
	}

	reaction.EmojiName = strings.Repeat("a", 64)
	if err := reaction.IsValid(); err != nil {
		t.Fatal(err)
	}
}
