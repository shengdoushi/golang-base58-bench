# golang base58 lib bench

# golang base58 lib

mr-tron/base58

https://github.com/mr-tron/base58

itchyny/base58-go

https://github.com/itchyny/base58-go

tv42/base58

https://github.com/tv42/base58

jbenet/go-base58

https://github.com/jbenet/go-base58

shengdoushi/base58

https://github.com/shengdoushi/base58

# Performance

Encode performance

```
➜  base58-bench go test -bench=BenchmarkEncode
goos: darwin
goarch: amd64
BenchmarkEncodeItchyny-8       	      20	  93046533 ns/op
BenchmarkEncodeJbenet-8        	      10	 129783558 ns/op
BenchmarkEncodeMrtron-8        	      50	  23767042 ns/op
BenchmarkEncodeShengdoushi-8   	      50	  23943148 ns/op
BenchmarkEncodeTv42-8          	      10	 109971479 ns/op
PASS
ok  	_/Users/macintoshhd/Documents/git/base58-bench	8.320s
```

Decode performance

```
➜  base58-bench go test -bench=BenchmarkDecode
goos: darwin
goarch: amd64
BenchmarkDecodeItchyny-8       	      10	 126616860 ns/op
BenchmarkDecodeJbenet-8        	       5	 297163192 ns/op
BenchmarkDecodeMrtron-8        	     100	  16515815 ns/op
BenchmarkDecodeShengdoushi-8   	     100	  14638295 ns/op
BenchmarkDecodeTv42-8          	      10	 143919340 ns/op
PASS
ok  	_/Users/macintoshhd/Documents/git/base58-bench	10.063s
```

You can change the file bench_test.go to test your self:

```golang
	// testcase count
	testcaseCount = 20000
	// every testcase's decoded bytes length
	testcaseDecodedBytesLength = 32
```

# Algorithms

tv42/base58  and itchyny/base58-go are based by math/big, so they are slower, and their api are not easy to use, espesially tv42/base58. And tv42/base58 does not support the big number with a prefix of 0. tv42/base58 only support flickr alphabet.


shengdoushi/base58 and mr-tron/base58 are both fastest. In terms of encoding methods, they are roughly the same logic. But in decoding, they are not the same.


