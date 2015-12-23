package catfruits

// Infomation is a repository infomation.
type Infomation struct {
	Flameworks []string `json:"flameworks"`
}

// Scanner is scanner for repository.
type Scanner struct {
	dir    string
	fwFunc map[string]FlameworkFunc
}

// FlameworkFunc is a function that detects framework
type FlameworkFunc func(*Scanner) (ok bool, err error)

// NewScanner returns a Scanner
func NewScanner(dir string) *Scanner {
	return &Scanner{
		dir:    dir,
		fwFunc: DefaultFlameworkFuncs,
	}
}

// Scan scans a repository.
func (sc *Scanner) Scan() (*Infomation, error) {
	info := &Infomation{
		Flameworks: make([]string, 0),
	}

	for name, f := range sc.fwFunc {
		ok, err := f(sc)
		if err != nil {
			return nil, err
		}
		if !ok {
			continue
		}

		info.Flameworks = append(info.Flameworks, name)
	}

	return info, nil
}
