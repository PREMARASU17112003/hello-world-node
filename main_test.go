package main

import (
    "testing"
    "bytes"
    "os"
)

func TestMain(t *testing.T) {
    var buf bytes.Buffer
    out := os.Stdout
    os.Stdout = &buf
    defer func() { os.Stdout = out }()

    main()

    if buf.String() != "Hello, World!\n" {
        t.Errorf("expected 'Hello, World!', got '%s'", buf.String())
    }
}
