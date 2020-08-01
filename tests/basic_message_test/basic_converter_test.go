package basic_message_test

import (
	"goprotoconv/goprotoconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadProtobufGoFileReadsProtobufGoFile(t *testing.T) {
	protobufFile, err := goprotoconv.LoadProtobufGoFile("./tests/basic_message_test/a.pb.go")

	assert.NotNil(t, protobufFile)
	assert.NoError(t, err)
}

func TestLoadProtobufGoFileReadsPackageName(t *testing.T) {
	protobufFile, _ := goprotoconv.LoadProtobufGoFile("./tests/basic_message_test/a.pb.go")
	assert.NotNil(t, protobufFile)
	assert.Equal(t, "basicmessage_packagea", protobufFile.PackageName)
	assert.False(t, true)
}
