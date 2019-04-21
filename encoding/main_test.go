package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/niubaoshu/gotiny"
)

func BenchmarkGOB(b *testing.B) {
	buf := bytes.NewBuffer(nil)

	b.Run("encode", func(b *testing.B) {
		for i := 0; i <= b.N; i++ {
			buf.Reset()

			if err := gob.NewEncoder(buf).Encode(testData); err != nil {
				panic(err)
			}
		}
	})

	reader := bytes.NewReader(buf.Bytes())

	b.Run("decode", func(b *testing.B) {
		for i := 0; i <= b.N; i++ {
			reader.Seek(0, 0)
			var result User

			if err := gob.NewDecoder(reader).Decode(&result); err != nil {
				panic(err)
			}

			if result.Data["username"] != "test" {
				b.Fatalf("bad decoding. got: %v", result)
			}
		}
	})
}

func BenchmarkJSON(b *testing.B) {
	b.Run("encode", func(b *testing.B) {
		for i := 0; i <= b.N; i++ {
			_, err := json.Marshal(testData)
			if err != nil {
				panic(err)
			}
		}
	})

	b.Run("decode", func(b *testing.B) {
		data, err := json.Marshal(testData)
		if err != nil {
			panic(err)
		}

		for i := 0; i <= b.N; i++ {
			var result User

			if err = json.Unmarshal(data, &result); err != nil {
				panic(err)
			}

			if result.Data["username"] != "test" {
				b.Fatalf("bad decoding. got: %v", result)
			}
		}
	})

	buf := bytes.NewBuffer(nil)

	b.Run("encode-stream", func(b *testing.B) {
		for i := 0; i <= b.N; i++ {
			buf.Reset()

			if err := json.NewEncoder(buf).Encode(testData); err != nil {
				panic(err)
			}
		}
	})

	reader := bytes.NewReader(buf.Bytes())

	b.Run("decode-stream", func(b *testing.B) {
		for i := 0; i <= b.N; i++ {
			reader.Seek(0, 0)
			var result User

			if err := json.NewDecoder(reader).Decode(&result); err != nil {
				panic(err)
			}

			if result.Data["username"] != "test" {
				b.Fatalf("bad decoding. got: %v", result)
			}
		}
	})
}

func BenchmarkTiny(b *testing.B) {
	b.Run("encode", func(b *testing.B) {
		for i := 0; i <= b.N; i++ {
			gotiny.Marshal(&testData)
		}
	})

	b.Run("decode", func(b *testing.B) {
		data := gotiny.Marshal(&testData)

		for i := 0; i <= b.N; i++ {
			var result User

			gotiny.Unmarshal(data, &result)

			if result.Data["username"] != "test" {
				b.Fatalf("bad decoding. got: %v", result)
			}
		}
	})
}

func BenchmarkProtobuf(b *testing.B) {
	b.Run("encode", func(b *testing.B) {
		for i := 0; i <= b.N; i++ {
			if _, err := proto.Marshal(testDataProtobuf); err != nil {
				panic(err)
			}
		}
	})

	b.Run("decode", func(b *testing.B) {
		data, err := proto.Marshal(testDataProtobuf)
		if err != nil {
			panic(err)
		}

		for i := 0; i <= b.N; i++ {
			var result UserProtobuf

			if err = proto.Unmarshal(data, &result); err != nil {
				panic(err)
			}

			if result.Data["username"] != "test" {
				b.Fatalf("bad decoding. got: %v", result)
			}
		}
	})
}
