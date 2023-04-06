package ee

func (f *fundamental) MarshalText() ([]byte, error) {
	return []byte(f.Error()), nil
}

func (w *withStack) MarshalText() ([]byte, error) {
	return []byte(w.Error()), nil
}

func (w *withMessage) MarshalText() ([]byte, error) {
	return []byte(w.Error()), nil
}
