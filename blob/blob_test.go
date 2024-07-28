package blob

import (
	"bytes"
	"testing"
)

func TestNewBlob(t *testing.T) {
	content := []byte("Hello, World!")
	blob, err := NewBlob(content)
	if err != nil {
		t.Fatalf("NewBlob returned an unexpected error: %s", err)
	}

	if blob.Size != int64(len(content)) {
		t.Errorf("blob.Size = %d; want %d", blob.Size, len(content))
	}

	if !bytes.Equal(blob.Content, content) {
		t.Errorf("blob.Content = %s; want %s", blob.Content, content)
	}

	expectedHash := "0a0a9f2a6772942557ab5355d76af442f8f65e01"
	if blob.Hash != expectedHash {
		t.Errorf("blob.Hash = %s; want %s", blob.Hash, expectedHash)
	}
}

func TestBlob_Serialize(t *testing.T) {
	content := []byte("Hello, World!")
	blob, _ := NewBlob(content)

	serialized, err := blob.Serialize()
	if err != nil {
		t.Fatalf("Failed to serialize blob: %v", err)
	}

	expectedPrefix := []byte("blob 13\x00")
	if !bytes.HasPrefix(serialized, expectedPrefix) {
		t.Errorf("Serialized blob does not have correct prefix")
	}

	if !bytes.Equal(serialized[len(expectedPrefix):], content) {
		t.Errorf("Serialized blob content does not match original content")
	}
}

func TestDeserialize(t *testing.T) {
	content := []byte("Hello, World!")
	blob, _ := NewBlob(content)

	serialized, _ := blob.Serialize()
	deserialized, err := Deserialize(serialized)
	if err != nil {
		t.Fatalf("Failed to deserialize blob: %v", err)
	}

	if deserialized.Size != blob.Size {
		t.Errorf("deserialized.Size = %d; want %d", deserialized.Size, blob.Size)
	}

	if deserialized.Hash != blob.Hash {
		t.Errorf("deserialized.Hash = %s; want %s", deserialized.Hash, blob.Hash)
	}

	if !bytes.Equal(deserialized.Content, blob.Content) {
		t.Errorf("deserialized.Content = %s; want %s", deserialized.Content, blob.Content)
	}
}
