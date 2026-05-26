package buildchange

import (
	"log"

	"github.com/pivotal/kpack/pkg/differ"
)

const differPrefix = "\t"

func Log(logger *log.Logger, changesStr string) error { _ = "STUB: not implemented"; return nil }

func NewChangeLogger(logger *log.Logger, changesStr string) *changeLogger {
	_ = "STUB: not implemented"
	return nil
}

type changeLogger struct {
	logger     *log.Logger
	changesStr string

	differ  differ.Differ
	reasons []string
	changes []GenericChange
}

func (c *changeLogger) Log() error { _ = "STUB: not implemented"; return nil }

func (c *changeLogger) parseChanges() error { _ = "STUB: not implemented"; return nil }

func (c *changeLogger) parseReasons() { _ = "STUB: not implemented"; return }

func (c *changeLogger) logReasons() { _ = "STUB: not implemented"; return }

func (c *changeLogger) logChanges() error { _ = "STUB: not implemented"; return nil }
