package alt

// int64
func int64ToBytes(i int64) []byte {
	return uint64ToBytes(uint64(i))
}

func bytesToInt64(b []byte) (uint64, error) {
	return bytesToUint64(b)

}
