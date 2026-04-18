func VerifySignature(publicKeyBytes []byte, data []byte, signature []byte) bool {

	r := new(big.Int).SetBytes(signature[:len(signature)/2])
	s := new(big.Int).SetBytes(signature[len(signature)/2:])

	curve := elliptic.P256()
	x := new(big.Int).SetBytes(publicKeyBytes[:len(publicKeyBytes)/2])
	y := new(big.Int).SetBytes(publicKeyBytes[len(publicKeyBytes)/2:])
	rawPubKey := ecdsa.PublicKey{Curve: curve, X: x, Y: y}

	hash := sha256.Sum256(data)

	return ecdsa.Verify(&rawPubKey, hash[:], r, s)
}
