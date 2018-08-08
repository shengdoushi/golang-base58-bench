package main

import (
	"crypto/rand"
	"math/big"
	"testing"

	itchynyAlgo "github.com/itchyny/base58-go"
	jbenetAlgo "github.com/jbenet/go-base58"
	mrtronAlgo "github.com/mr-tron/base58/base58"
	shengdoushiAlgo "github.com/shengdoushi/base58"
	tv42Algo "github.com/tv42/base58"
)

const (
	// testcase count
	testcaseCount = 20000
	// every testcase's decoded bytes length
	testcaseDecodedBytesLength = 32
)

type testCase struct {
	decoded []byte
	encoded string
}

var testcases = make([]testCase, 0, testcaseCount)

// init test cases
func init() {
	for i := 0; i < testcaseCount; i++ {
		data := make([]byte, testcaseDecodedBytesLength)
		rand.Read(data)
		testcases = append(testcases, testCase{decoded: data, encoded: mrtronAlgo.FastBase58Encoding(data)})
	}
}

// bench encode
func BenchmarkEncodeItchyny(b *testing.B) {
	testItems := make([][]byte, 0, len(testcases))
	for _, testcase := range testcases {
		_, shouldDecoded := toItchynyCase(testcase)
		testItems = append(testItems, shouldDecoded)
	}
	b.ResetTimer()
	alphabet := itchynyAlgo.BitcoinEncoding
	for i := 0; i < b.N; i++ {
		for _, testcase := range testItems {
			alphabet.Encode(testcase)
		}
	}
}

func BenchmarkEncodeJbenet(b *testing.B) {
	alphabet := jbenetAlgo.BTCAlphabet
	for i := 0; i < b.N; i++ {
		for _, testcase := range testcases {
			jbenetAlgo.EncodeAlphabet(testcase.decoded, alphabet)
		}
	}
}

func BenchmarkEncodeMrtron(b *testing.B) {
	alphabet := mrtronAlgo.BTCAlphabet
	for i := 0; i < b.N; i++ {
		for _, testcase := range testcases {
			mrtronAlgo.EncodeAlphabet(testcase.decoded, alphabet)
		}
	}
}

func BenchmarkEncodeShengdoushi(b *testing.B) {
	alphabet := shengdoushiAlgo.BitcoinAlphabet
	for i := 0; i < b.N; i++ {
		for _, testcase := range testcases {
			shengdoushiAlgo.Encode(testcase.decoded, alphabet)
		}
	}
}

func BenchmarkEncodeTv42(b *testing.B) {
	testItems := make([]*big.Int, 0, len(testcases))
	for _, testcase := range testcases {
		_, shouldDecoded := toTv42Case(testcase)
		if shouldDecoded == nil {
			continue
		}
		testItems = append(testItems, shouldDecoded)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, testcase := range testItems {
			tv42Algo.EncodeBig(nil, testcase)
		}
	}
}

// decode bench
func BenchmarkDecodeItchyny(b *testing.B) {
	testItems := make([]string, 0, len(testcases))
	for _, testcase := range testcases {
		shouldEncoded, _ := toItchynyCase(testcase)
		testItems = append(testItems, shouldEncoded)
	}

	b.ResetTimer()
	alphabet := itchynyAlgo.BitcoinEncoding
	for i := 0; i < b.N; i++ {
		for _, testcase := range testItems {
			alphabet.Decode([]byte(testcase))
		}
	}
}

func BenchmarkDecodeJbenet(b *testing.B) {
	alphabet := jbenetAlgo.BTCAlphabet
	for i := 0; i < b.N; i++ {
		for _, testcase := range testcases {
			jbenetAlgo.DecodeAlphabet(testcase.encoded, alphabet)
		}
	}
}

func BenchmarkDecodeMrtron(b *testing.B) {
	alphabet := mrtronAlgo.BTCAlphabet
	for i := 0; i < b.N; i++ {
		for _, testcase := range testcases {
			mrtronAlgo.DecodeAlphabet(testcase.encoded, alphabet)
		}
	}
}

func BenchmarkDecodeShengdoushi(b *testing.B) {
	alphabet := shengdoushiAlgo.BitcoinAlphabet
	for i := 0; i < b.N; i++ {
		for _, testcase := range testcases {
			shengdoushiAlgo.Decode(testcase.encoded, alphabet)
		}
	}
}

func BenchmarkDecodeTv42(b *testing.B) {
	testItems := make([]string, 0, len(testcases))
	for _, testcase := range testcases {
		shouldEncoded, shouldDecoded := toTv42Case(testcase)
		if shouldDecoded == nil {
			continue
		}
		testItems = append(testItems, shouldEncoded)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, testcase := range testcases {
			tv42Algo.DecodeToBig([]byte(testcase.encoded))
		}
	}
}
