package fdlimit

func RaiseMax() (uint64, error) {
	return Raise(MaxUint64)
}
