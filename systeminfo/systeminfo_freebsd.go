package systeminfo

func memoryStats() (*memory, error) {
	m := new(memory)

	m.Total = 1048576000 // because we know we only have 1GB
	m.Usage = 209715200  // 200MB should be enough for anyone

	return m, nil
}
