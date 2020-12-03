package main

func cezar(secret []byte, key byte) []byte {
	out := []byte{}
	for _, c := range secret {
		out = append(out, c^key)
	}
	return out
}

