package main

import (
	"strings"
	"bytes"
	"math/big"
	"testing"

	itchynyAlgo "github.com/itchyny/base58-go"
	jbenetAlgo "github.com/jbenet/go-base58"
	m0t0k1ch1Algo "github.com/m0t0k1ch1/base58"
	mrtronAlgo "github.com/mr-tron/base58/base58"
	shengdoushiAlgo "github.com/shengdoushi/base58"
	tv42Algo "github.com/tv42/base58"
)

const (
	// tv42's alphabet is not btc, but flicker
	tv42Alphabet = "123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ"
	btcAlphabet  = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
)

func toItchynyCase(testcase testCase)(encoded string, decoded []byte){
	zeros := 0
	for ; zeros < len(testcase.decoded) && testcase.decoded[zeros] == 0; zeros++ {
	}
	if (zeros==len(testcase.decoded)){
		decoded = []byte("0")
	}else{
		n := new(big.Int).SetBytes(testcase.decoded)
		decoded = []byte(strings.Repeat("0", zeros) + n.String())
	}
	encoded = testcase.encoded
	return
}

func TestItchyny(t *testing.T) {
	alphabet := itchynyAlgo.BitcoinEncoding
	for _, testcase := range testcases {
		// prefix zero should be ignored
		shouldEncoded, shouldDecoded := toItchynyCase(testcase)
		encoded, err := alphabet.Encode(shouldDecoded)
		if err != nil {
			t.Errorf("Itchyny encode(%v) should be %s, but error: %s", shouldDecoded, shouldEncoded, err.Error())
		} else if string(encoded) != shouldEncoded {
			t.Errorf("Itchyny encode(%v) should be %v, but %v", shouldDecoded, shouldEncoded, string(encoded))
		}

		decoded, err := alphabet.Decode([]byte(shouldEncoded))
		if err != nil {
			t.Errorf("Itchyny decode(%s) should be %v, but error: %s", shouldEncoded, shouldDecoded, err.Error())
		} else if !bytes.Equal(shouldDecoded, decoded){
			t.Errorf("Itchyny decode(%s) should be %v, but %v", shouldEncoded, shouldDecoded, decoded)
		}
	}
}

func TestJbenet(t *testing.T) {
	alphabet := jbenetAlgo.BTCAlphabet
	for _, testcase := range testcases {
		encoded := jbenetAlgo.EncodeAlphabet(testcase.decoded, alphabet)
		if encoded != testcase.encoded {
			t.Errorf("Jbenet encode(%v) should be %s, but %s", testcase.decoded, testcase.encoded, encoded)
		}
		decoded := jbenetAlgo.DecodeAlphabet(testcase.encoded, alphabet)
		if !bytes.Equal(decoded, testcase.decoded) {
			t.Errorf("Jbenet decode(%s) should be %v, but %v", testcase.encoded, testcase.decoded, decoded)
		}
	}
}

func TestM0t0k1ch1(t *testing.T) {
	alphabet := m0t0k1ch1Algo.NewBitcoinBase58()
	for _, testcase := range testcases {
		encoded, _ := alphabet.EncodeToString(testcase.decoded)
		if encoded != testcase.encoded {
			t.Errorf("Mrtron encode(%v) should be %s, but %s", testcase.decoded, testcase.encoded, encoded)
		}
		decoded, _ := alphabet.DecodeString(testcase.encoded)
		if !bytes.Equal(decoded, testcase.decoded) {
			t.Errorf("Mrtron decode(%s) should be %v, but %v", testcase.encoded, testcase.decoded, decoded)
		}
	}
}

func TestMrtron(t *testing.T) {
	alphabet := mrtronAlgo.BTCAlphabet
	for _, testcase := range testcases {
		encoded := mrtronAlgo.EncodeAlphabet(testcase.decoded, alphabet)
		if encoded != testcase.encoded {
			t.Errorf("Mrtron encode(%v) should be %s, but %s", testcase.decoded, testcase.encoded, encoded)
		}
		decoded, _ := mrtronAlgo.DecodeAlphabet(testcase.encoded, alphabet)
		if !bytes.Equal(decoded, testcase.decoded) {
			t.Errorf("Mrtron decode(%s) should be %v, but %v", testcase.encoded, testcase.decoded, decoded)
		}
	}
}

func TestShengdoushi(t *testing.T) {
	alphabet := shengdoushiAlgo.BitcoinAlphabet
	for _, testcase := range testcases {
		encoded := shengdoushiAlgo.Encode([]byte(testcase.decoded), alphabet)
		if encoded != testcase.encoded {
			t.Errorf("Shengdoushi encode(%v) should be %s, but %s", testcase.decoded, testcase.encoded, encoded)
		}
		decoded, _ := shengdoushiAlgo.Decode(testcase.encoded, alphabet)
		if !bytes.Equal(decoded, testcase.decoded) {
			t.Errorf("Shengdoushi decode(%s) should be %v, but %v", testcase.encoded, testcase.decoded, decoded)
		}
	}
}

func toTv42Case(testcase testCase)(encoded string, decoded *big.Int){
	zeros := 0
	for ; zeros < len(testcase.decoded) && testcase.decoded[zeros] == 0; zeros++ {
	}
	if (zeros==len(testcase.decoded)){
		decoded = nil
	}else{
		decoded = new(big.Int).SetBytes(testcase.decoded)
	}
	encodedBytes := make([]byte, len(testcase.encoded)-zeros)
	for idx, ch := range testcase.encoded[zeros:]{
		index := strings.IndexRune(btcAlphabet, ch)
		encodedBytes[idx] = tv42Alphabet[index]
	}
	encoded = string(encodedBytes)
	return
}

func TestTv42(t *testing.T) {
	for _, testcase := range testcases {
		shouldEncoded, shouldDecoded := toTv42Case(testcase)
		if shouldDecoded == nil {
			continue
		}
		// prefix zero should be ignored
		encoded := tv42Algo.EncodeBig(nil, shouldDecoded)
		if string(encoded) != shouldEncoded {
			t.Errorf("Tv42 encode(%v) should be %v, but %v", shouldDecoded, shouldEncoded, string(encoded))
		}

		decoded, err := tv42Algo.DecodeToBig([]byte(shouldEncoded))
		if err != nil {
		 	t.Errorf("Tv42 decode(%s) should be %v, but error: %s", shouldEncoded, shouldDecoded, err.Error())
		} else if !bytes.Equal(shouldDecoded.Bytes(), decoded.Bytes()){
		 	t.Errorf("Tv42 decode(%s) should be %v, but %v", shouldEncoded, shouldDecoded.Bytes(), decoded.Bytes())
		}
	}
}

