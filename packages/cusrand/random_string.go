package cusrand

func UniqueRandomString(count int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, count)
	for i := range b {
		b[i] = charset[RandomInt(0, len(charset))]
	}
	return string(b)
}
