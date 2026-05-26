package differ

type Differ struct {
	o Options
}

type Options struct {
	Prefix string
	Color  bool
	Common bool
}

func DefaultOptions() Options { _ = "STUB: not implemented"; return *new(Options) }

func NewDiffer(options Options) Differ { _ = "STUB: not implemented"; return *new(Differ) }

func Diff(dOld, dNew interface{}) (string, error) { _ = "STUB: not implemented"; return "", nil }

func (d Differ) Configure(options Options) { _ = "STUB: not implemented"; return }

func (d Differ) Diff(dOld, dNew interface{}) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (d Differ) getData(obj interface{}) (string, error) { _ = "STUB: not implemented"; return "", nil }

func (d Differ) renderDiff(a, b string) string { _ = "STUB: not implemented"; return "" }
