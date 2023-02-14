package scanner

type (
	scanner struct {
		source string
	}

	Scanner interface {
		ScanTokens() []string
	}
)

func NewScanner(source string) Scanner {
	return &scanner{
		source: source,
	}
}

func (s *scanner) ScanTokens() []string {
	return []string{s.source}
}
