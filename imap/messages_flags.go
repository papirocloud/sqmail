package imap

import (
	"github.com/emersion/go-imap/v2"
	"github.com/emersion/go-imap/v2/imapclient"
)

// AddFlags adds the specified flag(s) to the specified message(s).
func AddFlags(c *imapclient.Client, seqSet imap.SeqSet, flags []imap.Flag, silent bool) ([]*imapclient.FetchMessageBuffer, error) {
	var flagStrs []string
	for _, flag := range flags {
		flagStrs = append(flagStrs, string(flag))
	}

	logger.Info().Strs("flags", flagStrs).Msg("adding flags to message(s)")

	return Store(c, seqSet, &imap.StoreFlags{
		Op:     imap.StoreFlagsAdd,
		Flags:  flags,
		Silent: silent,
	})
}

// DeleteFlags removes the specified flag(s) from the specified message(s).
func DeleteFlags(c *imapclient.Client, seqSet imap.SeqSet, flags []imap.Flag, silent bool) ([]*imapclient.FetchMessageBuffer, error) {
	var flagStrs []string
	for _, flag := range flags {
		flagStrs = append(flagStrs, string(flag))
	}

	logger.Info().Strs("flags", flagStrs).Msg("deleting flags from message(s)")

	return Store(c, seqSet, &imap.StoreFlags{
		Op:     imap.StoreFlagsDel,
		Flags:  flags,
		Silent: silent,
	})
}

// SetFlags replaces the flags of the specified message(s) with the specified flag(s).
func SetFlags(c *imapclient.Client, seqSet imap.SeqSet, flags []imap.Flag, silent bool) ([]*imapclient.FetchMessageBuffer, error) {
	var flagStrs []string
	for _, flag := range flags {
		flagStrs = append(flagStrs, string(flag))
	}

	logger.Info().Strs("flags", flagStrs).Msg("setting flags of message(s)")

	return Store(c, seqSet, &imap.StoreFlags{
		Op:     imap.StoreFlagsSet,
		Flags:  flags,
		Silent: silent,
	})
}

// UIDAddFlags adds the specified flag(s) to the specified message(s).
func UIDAddFlags(c *imapclient.Client, uidSet imap.UIDSet, flags []imap.Flag, silent bool) ([]*imapclient.FetchMessageBuffer, error) {
	var flagStrs []string
	for _, flag := range flags {
		flagStrs = append(flagStrs, string(flag))
	}

	logger.Info().Strs("flags", flagStrs).Msg("adding flags to message(s)")

	return UIDStore(c, uidSet, &imap.StoreFlags{
		Op:     imap.StoreFlagsAdd,
		Flags:  flags,
		Silent: silent,
	})
}

// UIDDeleteFlags removes the specified flag(s) from the specified message(s).
func UIDDeleteFlags(c *imapclient.Client, uidSet imap.UIDSet, flags []imap.Flag, silent bool) ([]*imapclient.FetchMessageBuffer, error) {
	var flagStrs []string
	for _, flag := range flags {
		flagStrs = append(flagStrs, string(flag))
	}

	logger.Info().Strs("flags", flagStrs).Msg("deleting flags from message(s)")

	return UIDStore(c, uidSet, &imap.StoreFlags{
		Op:     imap.StoreFlagsDel,
		Flags:  flags,
		Silent: silent,
	})
}

// UIDSetFlags replaces the flags of the specified message(s) with the specified flag(s).
func UIDSetFlags(c *imapclient.Client, uidSet imap.UIDSet, flags []imap.Flag, silent bool) ([]*imapclient.FetchMessageBuffer, error) {
	var flagStrs []string
	for _, flag := range flags {
		flagStrs = append(flagStrs, string(flag))
	}

	logger.Info().Strs("flags", flagStrs).Msg("setting flags of message(s)")

	return UIDStore(c, uidSet, &imap.StoreFlags{
		Op:     imap.StoreFlagsSet,
		Flags:  flags,
		Silent: silent,
	})
}
