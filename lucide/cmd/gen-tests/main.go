package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	dir := flag.String("dir", "lucide", "Path to the lucide directory")
	overwrite := flag.Bool("overwrite", true, "Overwrite existing test files")
	flag.Parse()

	// Resolve directory to absolute path for clarity
	lucideDir, err := filepath.Abs(*dir)
	if err != nil {
		fatalf("resolve dir: %v", err)
	}

	entries, err := os.ReadDir(lucideDir)
	if err != nil {
		fatalf("read dir %s: %v", lucideDir, err)
	}

	var generated, skipped int
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		name := e.Name()
		if !strings.HasSuffix(name, ".go") {
			continue
		}
		if strings.HasSuffix(name, "_test.go") {
			continue
		}
		if name == "lucide.go" { // support file, not an icon
			continue
		}

		base := strings.TrimSuffix(name, ".go")
		testFile := filepath.Join(lucideDir, base+"_test.go")

		// Skip if test file already exists unless overwrite is set
		if fileExists(testFile) && !*overwrite {
			skipped++
			continue
		}

		funcName := toFuncName(base)
		contents := renderTestFile(funcName, base)

		if err := os.WriteFile(testFile, []byte(contents), fs.FileMode(0644)); err != nil {
			fatalf("write %s: %v", testFile, err)
		}
		fmt.Printf("generated %s\n", relPath(testFile))
		generated++
	}

	fmt.Printf("done: generated=%d skipped=%d\n", generated, skipped)
}

func renderTestFile(funcName, fileBase string) string {
	// Keep the format consistent with existing tests
	return fmt.Sprintf(`package lucide

import (
    "os"
    "strings"
    "testing"

    x "github.com/plainkit/html"
    "github.com/stretchr/testify/require"
)

func Test%[1]s_SVGMatchesReference(t *testing.T) {
    // Render the component
    rendered := x.Render(%[1]s())

    // Parse rendered SVG
    gotNode, err := parseXMLToNode(strings.NewReader(rendered))
    require.NoError(t, err, "parse rendered SVG")
    gotNode = normalizeNode(gotNode)

    // Load reference SVG
    f, err := os.Open("icons/%[2]s.svg")
    require.NoError(t, err, "open reference SVG")
    defer f.Close()

    wantNode, err := parseXMLToNode(f)
    require.NoError(t, err, "parse reference SVG")
    wantNode = normalizeNode(wantNode)

    // Compare root element
    require.Equal(t, wantNode.Name.Local, gotNode.Name.Local, "root element name")
    require.Equal(t, wantNode.Attrs, gotNode.Attrs, "root attributes must match")

    // Compare recursively
    compareNodes(t, wantNode, gotNode)
}
`, funcName, fileBase)
}

func toFuncName(base string) string {
	parts := strings.Split(base, "-")
	var b strings.Builder
	for _, p := range parts {
		if p == "" {
			continue
		}
		r, size := utf8.DecodeRuneInString(p)
		if unicode.IsLetter(r) {
			// Uppercase the first letter, keep the rest intact
			b.WriteString(strings.ToUpper(string(r)))
			b.WriteString(p[size:])
		} else {
			// For tokens starting with digits or other runes, keep as is
			b.WriteString(p)
		}
	}
	return b.String()
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func relPath(path string) string {
	if wd, err := os.Getwd(); err == nil {
		if rp, err := filepath.Rel(wd, path); err == nil {
			return rp
		}
	}
	return path
}

func fatalf(format string, args ...any) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}
