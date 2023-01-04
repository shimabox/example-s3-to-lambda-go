package testhelper

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

// ExtractStdout はStdoutに書き込まれた文字列を抽出する関数です
// (Stderrも同じ要領で出力先を変更できます)
func ExtractStdout(t *testing.T, f func()) string {
	t.Helper()

	// 既存のStdoutを退避する
	orgStdout := os.Stdout
	defer func() {
		// 出力先を元に戻す
		os.Stdout = orgStdout
	}()

	// パイプの作成(r: Reader, w: Writer)
	r, w, _ := os.Pipe()

	// Stdoutの出力先をパイプのwriterに変更する
	os.Stdout = w

	// テスト対象の関数を実行する
	f()

	// Writerをクローズする
	// Writerオブジェクトはクローズするまで処理をブロックするので注意
	w.Close()

	// Bufferに書き込こまれた内容を読み出す
	var buf bytes.Buffer
	if _, err := buf.ReadFrom(r); err != nil {
		t.Fatalf("failed to read buf: %v", err)
	}

	// 文字列を取得して返却
	return strings.TrimRight(buf.String(), "\n")
}
