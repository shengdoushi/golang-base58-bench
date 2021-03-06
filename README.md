# golang base58 lib bench

# golang base58 lib

m0t0k1ch1/base58

https://github.com/m0t0k1ch1/base58

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
➜  base58-bench git:(master) ✗ go test -bench=BenchmarkEncode
goos: darwin
goarch: amd64
pkg: shengdoushi/golang-base58-bench
BenchmarkEncodeItchyny-8       	      20	  92408674 ns/op
BenchmarkEncodeJbenet-8        	      10	 122789581 ns/op
BenchmarkEncodeM0t0k1ch1-8     	      20	  92152212 ns/op
BenchmarkEncodeMrtron-8        	      50	  23120254 ns/op
BenchmarkEncodeShengdoushi-8   	      50	  23946466 ns/op
BenchmarkEncodeTv42-8          	      10	 113029698 ns/op
PASS
ok  	shengdoushi/golang-base58-bench	10.347s
```

Decode performance

```
➜  base58-bench git:(master) ✗ go test -bench=BenchmarkDecode
goos: darwin
goarch: amd64
pkg: shengdoushi/golang-base58-bench
BenchmarkDecodeItchyny-8       	      10	 130006794 ns/op
BenchmarkDecodeJbenet-8        	       5	 304503864 ns/op
BenchmarkDecodeM0t0k1ch1-8     	      10	 178443403 ns/op
BenchmarkDecodeMrtron-8        	     100	  17393332 ns/op
BenchmarkDecodeShengdoushi-8   	     100	  14995003 ns/op
BenchmarkDecodeTv42-8          	      10	 139954525 ns/op
PASS
ok  	shengdoushi/golang-base58-bench	12.410s
```

You can change the file bench_test.go to test your self:

```golang
	// testcase count
	testcaseCount = 20000
	// every testcase's decoded bytes length
	testcaseDecodedBytesLength = 32
```

# Algorithms

tv42/base58 and itchyny/base58-go are based by math/big, so they are slower, and their api are not easy to use, espesially tv42/base58. And tv42/base58 does not support the big number with a prefix of 0. tv42/base58 only support flickr alphabet.

m0t0k1ch1/base58 is also based by math/big, performance is just like tv42/base58 and itchyny/base58-go. but the api is easy to use.

shengdoushi/base58 and mr-tron/base58 are both fastest. In terms of encoding methods, they are roughly the same logic. But in decoding, they are not the same. shengdoushi/base58 can use custom alphabet just like '一二三四五六七八九十壹贰叁肆伍陆柒捌玖零拾佰仟万亿圆甲乙丙丁戊己庚辛壬癸子丑寅卯辰巳午未申酉戌亥金木水火土雷电风雨福', mr-tron/base58 only support non-unicode custom alphabet.


