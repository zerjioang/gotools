/*
It is a digital signature scheme known for its simplicity,
among the first whose security is based on the intractability
of certain discrete logarithm problems.It is efficient and
generates short signatures.

It was covered by U.S. Patent 4,995,082 which expired in February 2008.
*/
package generic

/*
A Schnorr group, proposed by Claus P. Schnorr, is a large prime-order
subgroup of Zpx, the multiplicative group of integers modulo p for some prime p.
To generate such a group, generate p,q,r such that

p = qr + 1

with p,q prime numbers

Then, choose an h value so that
1 < h < p
until you find one such that
h^r != 1 (mod p)

Finally compute
g = h^r (mod p)

being g, the generator of a subgroup Zpx of order q
*/
type SchnorrGroup struct {
}

/*
Schnorr

At its core, Schnorr signatures use a particular function, defined as

H'(m, s, e) = H(m || s * G + e * P)

* H is a hash function, for instance SHA256.
* s and e are 2 numbers forming the signature itself.
* m is the message Alice wants to sign.
* P is Alice’s public key.
*/
type Schnorr struct {
}
