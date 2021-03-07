package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"io/ioutil"
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/niubaoshu/gotiny"
	"google.golang.org/protobuf/proto"

	"gotest/pb"
)

var testData = Person{
	ID:    "ntrrg",
	Name:  "Miguel Angel Rivera Notararigo",
	Email: "ntrrg@example.com",
	Alive: true,

	Numbers: []int{11, 2, 0},

	Car: Car{
		ID:    1,
		Brand: "Toyota",
		Model: "Corolla Araya",
	},

	Family: []*Person{
		{ID: "bongo", Name: "Bongo Notararigo"},
		{ID: "assdro", Name: "Alessandro Notararigo"},
	},

	Data: []byte(`{
		"username": "test",
		"anime": [
			"One Piece",
			"Fullmetal Alchemist",
			"Fate",
			"Hellsing",
			"Naruto",
			"Dragon Ball"
		]
	}`),
}

var testDataPb = pb.Person{
	Id:    "ntrrg",
	Name:  "Miguel Angel Rivera Notararigo",
	Email: "ntrrg@example.com",
	Alive: true,

	Numbers: []int64{11, 2, 0},

	Car: &pb.Car{
		Id:    1,
		Brand: "Toyota",
		Model: "Corolla Araya",
	},

	Family: []*pb.Person{
		{Id: "bongo", Name: "Bongo Notararigo"},
		{Id: "assdro", Name: "Alessandro Notararigo"},
	},

	Data: []byte(`{
		"username": "test",
		"anime": [
			"One Piece",
			"Fullmetal Alchemist",
			"Fate",
			"Hellsing",
			"Naruto",
			"Dragon Ball"
		]
	}`),
}

/* Doesn't support []Person inside Person
func BenchmarkBinary(b *testing.B) {
	b.Run("marshal", func(b *testing.B) {
		for i := 0; i <= b.N; i++ {
			data, err := binary.Marshal(&testData)
			if err != nil {
				b.Fatal(err)
			}

			_ = data
		}
	})

	data, err := binary.Marshal(&testData)
	if err != nil {
		b.Fatal(err)
	}

	b.Run("unmarshal", func(b *testing.B) {
		for i := 0; i <= b.N; i++ {
			var result Person

			if err := binary.Unmarshal(data, &result); err != nil {
				b.Fatal(err)
			}

			if n := result.Numbers; n[0] != 11 || n[1] != 2 || n[2] != 0 {
				b.Fatalf("bad decoding. got: %v", result)
			}

			if result.Data["username"] != "test" {
				b.Fatalf("bad decoding. got: %v", result)
			}
		}
	})
}
*/

func BenchmarkGOB(b *testing.B) {
	b.Run("encode", func(b *testing.B) {
		for i := 0; i <= b.N; i++ {
			if err := gob.NewEncoder(ioutil.Discard).Encode(testData); err != nil {
				b.Fatal(err)
			}
		}
	})

	buf := bytes.NewBuffer(nil)

	if err := gob.NewEncoder(buf).Encode(testData); err != nil {
		b.Fatal(err)
	}

	reader := bytes.NewReader(buf.Bytes())

	b.Run("decode", func(b *testing.B) {
		for i := 0; i <= b.N; i++ {
			reader.Seek(0, 0)
			var result Person

			if err := gob.NewDecoder(reader).Decode(&result); err != nil {
				b.Fatal(err)
			}

			if n := result.Numbers; n[0] != 11 || n[1] != 2 || n[2] != 0 {
				b.Fatalf("bad decoding. got: %v", result)
			}
		}
	})
}

func BenchmarkJSON(b *testing.B) {
	b.Run("marshal", func(b *testing.B) {
		for i := 0; i <= b.N; i++ {
			data, err := json.Marshal(testData)
			if err != nil {
				b.Fatal(err)
			}

			_ = data
		}
	})

	data, err := json.Marshal(testData)
	if err != nil {
		b.Fatal(err)
	}

	b.Run("unmarshal", func(b *testing.B) {
		for i := 0; i <= b.N; i++ {
			var result Person

			if err = json.Unmarshal(data, &result); err != nil {
				b.Fatal(err)
			}

			if n := result.Numbers; n[0] != 11 || n[1] != 2 || n[2] != 0 {
				b.Fatalf("bad decoding. got: %v", result)
			}
		}
	})

	b.Run("encode", func(b *testing.B) {
		for i := 0; i <= b.N; i++ {
			if err := json.NewEncoder(ioutil.Discard).Encode(testData); err != nil {
				b.Fatal(err)
			}
		}
	})

	buf := bytes.NewBuffer(nil)

	if err := json.NewEncoder(buf).Encode(testData); err != nil {
		b.Fatal(err)
	}

	reader := bytes.NewReader(buf.Bytes())

	b.Run("decode", func(b *testing.B) {
		for i := 0; i <= b.N; i++ {
			reader.Seek(0, 0)
			var result Person

			if err := json.NewDecoder(reader).Decode(&result); err != nil {
				b.Fatal(err)
			}

			if n := result.Numbers; n[0] != 11 || n[1] != 2 || n[2] != 0 {
				b.Fatalf("bad decoding. got: %v", result)
			}
		}
	})
}

func BenchmarkJSONIter(b *testing.B) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	b.Run("marshal", func(b *testing.B) {
		for i := 0; i <= b.N; i++ {
			data, err := json.Marshal(testData)
			if err != nil {
				b.Fatal(err)
			}

			_ = data
		}
	})

	data, err := json.Marshal(testData)
	if err != nil {
		b.Fatal(err)
	}

	b.Run("unmarshal", func(b *testing.B) {
		for i := 0; i <= b.N; i++ {
			var result Person

			if err = json.Unmarshal(data, &result); err != nil {
				b.Fatal(err)
			}

			if n := result.Numbers; n[0] != 11 || n[1] != 2 || n[2] != 0 {
				b.Fatalf("bad decoding. got: %v", result)
			}
		}
	})

	b.Run("encode", func(b *testing.B) {
		for i := 0; i <= b.N; i++ {
			if err := json.NewEncoder(ioutil.Discard).Encode(testData); err != nil {
				b.Fatal(err)
			}
		}
	})

	buf := bytes.NewBuffer(nil)

	if err := json.NewEncoder(buf).Encode(testData); err != nil {
		b.Fatal(err)
	}

	reader := bytes.NewReader(buf.Bytes())

	b.Run("decode", func(b *testing.B) {
		for i := 0; i <= b.N; i++ {
			reader.Seek(0, 0)
			var result Person

			if err := json.NewDecoder(reader).Decode(&result); err != nil {
				b.Fatal(err)
			}

			if n := result.Numbers; n[0] != 11 || n[1] != 2 || n[2] != 0 {
				b.Fatalf("bad decoding. got: %v", result)
			}
		}
	})
}

func BenchmarkTiny(b *testing.B) {
	b.Run("marshal", func(b *testing.B) {
		for i := 0; i <= b.N; i++ {
			data := gotiny.Marshal(&testData)
			_ = data
		}
	})

	data := gotiny.Marshal(&testData)

	b.Run("unmarshal", func(b *testing.B) {
		for i := 0; i <= b.N; i++ {
			var result Person

			gotiny.Unmarshal(data, &result)

			if n := result.Numbers; n[0] != 11 || n[1] != 2 || n[2] != 0 {
				b.Fatalf("bad decoding. got: %v", result)
			}
		}
	})
}

func BenchmarkProtobuf(b *testing.B) {
	b.Run("marshal", func(b *testing.B) {
		for i := 0; i <= b.N; i++ {
			data, err := proto.Marshal(&testDataPb)
			if err != nil {
				b.Fatal(err)
			}

			_ = data
		}
	})

	b.Run("unmarshal", func(b *testing.B) {
		data, err := proto.Marshal(&testDataPb)
		if err != nil {
			b.Fatal(err)
		}

		for i := 0; i <= b.N; i++ {
			var result pb.Person

			if err = proto.Unmarshal(data, &result); err != nil {
				b.Fatal(err)
			}

			if n := result.GetNumbers(); n[0] != 11 || n[1] != 2 || n[2] != 0 {
				b.Fatalf("bad decoding. got: %v", result)
			}
		}
	})
}
