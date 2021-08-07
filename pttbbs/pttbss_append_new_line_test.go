package pttbbs

import (
	"bytes"
	"os"
	"testing"
)

func TestNewComment(t *testing.T) {
	c := Connector{"./testcase"}

	fileName := "M.1606672292.A.B23"

	expectedNewLine := "\u001B[1;31m→ \u001B[33mSYSOP\u001B[m\u001B[33m:test                                                   \u001B[m推 01/17 11:36"
	boardPath := "./testcase/boards/t/test/"
	filePath := boardPath + "/" + fileName
	stat, _ := os.Stat(filePath)
	oriFileSize := stat.Size()

	err := c.AppendBoardArticleFile(filePath, []byte(expectedNewLine))

	if err != nil {
		t.Errorf("Unexpected Error happened: %s", err)
	}

	stat, _ = os.Stat(filePath)
	newFileSize := stat.Size()
	newLineSize := newFileSize - oriFileSize
	buf := make([]byte, newLineSize)
	fileHandle, _ := os.Open(filePath)
	_, err = fileHandle.ReadAt(buf, oriFileSize)
	actualNewLine := bytes.NewBuffer(buf)

	if expectedNewLine != actualNewLine.String() {
		t.Errorf("newline not matched, expected: %s, \ngot: %s", expectedNewLine, actualNewLine)
	}
}
