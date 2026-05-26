package testhelpers

import (
	"strings"
	"testing"
)

const defaultPrefix = "\t"

type DiffOutBuilder struct {
	t  *testing.T
	sb strings.Builder
	o  DiffOptions
}

type DiffOptions struct {
	Prefix string
	Color  bool
}

func DefaultDiffOptions() DiffOptions { _ = "STUB: not implemented"; return *new(DiffOptions) }

func NewDiffOutBuilder(t *testing.T) *DiffOutBuilder { _ = "STUB: not implemented"; return nil }

func (d *DiffOutBuilder) Configure(options DiffOptions) *DiffOutBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (d *DiffOutBuilder) Reset() *DiffOutBuilder { _ = "STUB: not implemented"; return nil }

func (d *DiffOutBuilder) Txt(str string) *DiffOutBuilder { _ = "STUB: not implemented"; return nil }

func (d *DiffOutBuilder) NoD(str string) *DiffOutBuilder { _ = "STUB: not implemented"; return nil }

func (d *DiffOutBuilder) Old(str string) *DiffOutBuilder { _ = "STUB: not implemented"; return nil }

func (d *DiffOutBuilder) New(str string) *DiffOutBuilder { _ = "STUB: not implemented"; return nil }

func (d *DiffOutBuilder) Out() string { _ = "STUB: not implemented"; return "" }
