package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

const (
	// ModuleName defines the module name
	ModuleName = "blog"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_blog"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	//...

	// Keep track of the index of posts
	PostKey      = "Post-value-"
	PostCountKey = "Post-count-"
)

const (
	CommentKey      = "Comment-value-"
	CommentCountKey = "Comment-count-"
)

var (
	ErrCommentOld = sdkerrors.Register(ModuleName, 1300, "")
	ErrID         = sdkerrors.Register(ModuleName, 1400, "")
)
