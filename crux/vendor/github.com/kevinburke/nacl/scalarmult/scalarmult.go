// Package scalarmult provides an implementation of scalar multiplication.
//
// scalarmult is designed to be strong as a component of various well-known
// "hashed Diffie–Hellman" applications. In particular, it is designed to make
// the "computational Diffie–Hellman" problem (CDH) difficult with respect to
// the standard base.
// crypto_scalarmult is also designed to make CDH difficult with respect to
// other nontrivial bases. In particular, if a represented group element has
// small order, then it is annihilated by all represented scalars. This feature
// allows protocols to avoid validating membership in the subgroup generated by
// the standard base.
//
// scalarmult does not make any promises regarding the "decisional
// Diffie–Hellman" problem (DDH), the "static Diffie–Hellman" problem (SDH), etc.
// Users are responsible for hashing group elements.
//
// The current primitive is the function crypto_scalarmult_curve25519 specified
// in "Cryptography in NaCl", Sections 2, 3, and 4. This function is conjectured
// to be strong. For background see Bernstein, "Curve25519: new Diffie-Hellman
// speed records," Lecture Notes in Computer Science 3958 (2006), 207–228,
// https://cr.yp.to/papers.html#curve25519.
//
// scalarmult is compatible with NaCL: https://nacl.cr.yp.to/scalarmult.html
package scalarmult

import "golang.org/x/crypto/curve25519"

// Size is the size, in bytes, of a value for use in scalar multiplication
const Size = 32

// Base returns the product in*base where dst and base are the x coordinates of
// group points, base is the standard generator and all values are in
// little-endian form.
func Base(secretkey *[Size]byte) *[Size]byte {
	dst := new([Size]byte)
	curve25519.ScalarBaseMult(dst, secretkey)
	return dst
}

// Mult returns the product in*base where dst and base are the x coordinates of
// group points and all values are in little-endian form.
func Mult(in, base *[Size]byte) *[Size]byte {
	key := new([Size]byte)
	curve25519.ScalarMult(key, in, base)
	return key
}
