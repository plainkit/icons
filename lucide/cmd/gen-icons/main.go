package main

import (
	"bytes"
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"
)

// Supported SVG element names
var supportedElems = map[string]bool{
	"path":     true,
	"line":     true,
	"rect":     true,
	"circle":   true,
	"ellipse":  true,
	"polyline": true,
	"polygon":  true,
}

// Attribute mapping for painting/transform/opacity and extras
var paintAndMore = map[string]string{
	// painting
	"fill":              "Fill",
	"fill-opacity":      "FillOpacity",
	"fill-rule":         "FillRule",
	"stroke":            "Stroke",
	"stroke-width":      "StrokeWidth",
	"stroke-dasharray":  "StrokeDasharray",
	"stroke-dashoffset": "StrokeDashoffset",
	"stroke-linecap":    "StrokeLinecap",
	"stroke-linejoin":   "StrokeLinejoin",
	"stroke-opacity":    "StrokeOpacity",
	"stroke-miterlimit": "StrokeMiterlimit",
	// transform/opacity
	"transform": "Transform",
	"opacity":   "Opacity",
	// path length
	"pathLength": "PathLength",
}

type childCall struct {
	// element call, e.g. Path, Line, Rect
	elem string
	// args inside x.<Elem>(...), already formatted fragments like x.D("M...")
	attrCalls []string
}

func main() {
	iconsDir := flag.String("icons", "icons", "Path to the icons directory containing .svg files")
	outDir := flag.String("out", ".", "Output directory for generated .go files (defaults to current dir)")
	overwrite := flag.Bool("overwrite", false, "Overwrite existing Go files if present")
	only := flag.String("only", "", "Comma-separated list of base icon names (without .svg) to generate; empty for all")
	flag.Parse()

	iconsAbs, err := filepath.Abs(*iconsDir)
	if err != nil {
		fatalf("resolve icons dir: %v", err)
	}
	outAbs, err := filepath.Abs(*outDir)
	if err != nil {
		fatalf("resolve out dir: %v", err)
	}

	// Filter set
	var filter map[string]struct{}
	if strings.TrimSpace(*only) != "" {
		filter = map[string]struct{}{}
		for _, p := range strings.Split(*only, ",") {
			p = strings.TrimSpace(p)
			if p == "" {
				continue
			}
			// normalize without extension
			p = strings.TrimSuffix(p, ".svg")
			filter[p] = struct{}{}
		}
	}

	entries, err := os.ReadDir(iconsAbs)
	if err != nil {
		fatalf("read icons dir %s: %v", iconsAbs, err)
	}

	// Deterministic order
	sort.Slice(entries, func(i, j int) bool { return entries[i].Name() < entries[j].Name() })

	var generated, skipped int
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		name := e.Name()
		if !strings.HasSuffix(name, ".svg") {
			continue
		}
		base := strings.TrimSuffix(name, ".svg")
		if filter != nil {
			if _, ok := filter[base]; !ok {
				continue
			}
		}

		outFile := filepath.Join(outAbs, base+".go")
		if fileExists(outFile) && !*overwrite {
			skipped++
			continue
		}

		// Parse and generate
		absSVG := filepath.Join(iconsAbs, name)
		children, err := parseSVGChildren(absSVG)
		if err != nil {
			fatalf("parse %s: %v", relPath(absSVG), err)
		}

		funcName := toFuncName(base)
		docTitle := toDocTitle(base)
		className := fmt.Sprintf("lucide lucide-%s", base)
		contents := renderIconFile(funcName, docTitle, className, children)

		if err := os.WriteFile(outFile, []byte(contents), fs.FileMode(0644)); err != nil {
			fatalf("write %s: %v", relPath(outFile), err)
		}
		fmt.Printf("generated %s\n", relPath(outFile))
		generated++
	}

	fmt.Printf("done: generated=%d skipped=%d\n", generated, skipped)
}

func parseSVGChildren(path string) ([]childCall, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	dec := xml.NewDecoder(f)
	dec.Strict = false
	dec.AutoClose = xml.HTMLAutoClose
	dec.Entity = xml.HTMLEntity

	var inSVG bool
	var children []childCall
	var tok xml.Token
	for {
		t, err := dec.Token()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		tok = t
		switch se := tok.(type) {
		case xml.StartElement:
			name := local(se.Name)
			if name == "svg" {
				inSVG = true
				// ignore root attributes; buildLucideArgs handles defaults
				continue
			}
			if !inSVG {
				continue
			}
			if !supportedElems[name] {
				// Unsupported element: skip its subtree
				if err := dec.Skip(); err != nil {
					return nil, err
				}
				continue
			}
			call := makeChildCall(name, se.Attr)
			children = append(children, call)
			// No need to process end; we're only mapping empty elements used by icons
			if err := dec.Skip(); err != nil {
				return nil, err
			}
		}
	}
	return children, nil
}

func makeChildCall(elem string, attrs []xml.Attr) childCall {
	// Map attributes for each element
	// We gather element-specific first, then generic painting/transform/opacity, in a stable order.

	// Collect attributes into a map for easy lookup while preserving string values as-is.
	m := map[string]string{}
	for _, a := range attrs {
		m[local(a.Name)] = a.Value
	}

	var args []string
	switch elem {
	case "path":
		if v, ok := m["d"]; ok {
			args = append(args, fmt.Sprintf("x.D(\"%s\")", escapeString(collapseWS(v))))
		}
	case "line":
		// Follow common ordering seen in repo: X1, X2, Y1, Y2
		if v, ok := m["x1"]; ok {
			args = append(args, fmt.Sprintf("x.X1(\"%s\")", escapeString(v)))
		}
		if v, ok := m["x2"]; ok {
			args = append(args, fmt.Sprintf("x.X2(\"%s\")", escapeString(v)))
		}
		if v, ok := m["y1"]; ok {
			args = append(args, fmt.Sprintf("x.Y1(\"%s\")", escapeString(v)))
		}
		if v, ok := m["y2"]; ok {
			args = append(args, fmt.Sprintf("x.Y2(\"%s\")", escapeString(v)))
		}
	case "rect":
		// Use RectWidth/RectHeight then X/Y then Rx/Ry if present
		if v, ok := m["width"]; ok {
			args = append(args, fmt.Sprintf("x.RectWidth(\"%s\")", escapeString(v)))
		}
		if v, ok := m["height"]; ok {
			args = append(args, fmt.Sprintf("x.RectHeight(\"%s\")", escapeString(v)))
		}
		if v, ok := m["x"]; ok {
			args = append(args, fmt.Sprintf("x.X(\"%s\")", escapeString(v)))
		}
		if v, ok := m["y"]; ok {
			args = append(args, fmt.Sprintf("x.Y(\"%s\")", escapeString(v)))
		}
		if v, ok := m["rx"]; ok {
			args = append(args, fmt.Sprintf("x.Rx(\"%s\")", escapeString(v)))
		}
		if v, ok := m["ry"]; ok {
			args = append(args, fmt.Sprintf("x.Ry(\"%s\")", escapeString(v)))
		}
	case "circle":
		if v, ok := m["cx"]; ok {
			args = append(args, fmt.Sprintf("x.Cx(\"%s\")", escapeString(v)))
		}
		if v, ok := m["cy"]; ok {
			args = append(args, fmt.Sprintf("x.Cy(\"%s\")", escapeString(v)))
		}
		if v, ok := m["r"]; ok {
			args = append(args, fmt.Sprintf("x.R(\"%s\")", escapeString(v)))
		}
	case "ellipse":
		if v, ok := m["cx"]; ok {
			args = append(args, fmt.Sprintf("x.EllipseCx(\"%s\")", escapeString(v)))
		}
		if v, ok := m["cy"]; ok {
			args = append(args, fmt.Sprintf("x.EllipseCy(\"%s\")", escapeString(v)))
		}
		if v, ok := m["rx"]; ok {
			args = append(args, fmt.Sprintf("x.EllipseRx(\"%s\")", escapeString(v)))
		}
		if v, ok := m["ry"]; ok {
			args = append(args, fmt.Sprintf("x.EllipseRy(\"%s\")", escapeString(v)))
		}
	case "polyline", "polygon":
		if v, ok := m["points"]; ok {
			args = append(args, fmt.Sprintf("x.Points(\"%s\")", escapeString(collapseWS(v))))
		}
	}

	// Append generic known attributes, if present, in a stable order based on paintAndMore map keys.
	// We'll iterate keys in sorted order for deterministic output.
	keys := make([]string, 0, len(paintAndMore))
	for k := range paintAndMore {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		if v, ok := m[k]; ok {
			helper := paintAndMore[k]
			args = append(args, fmt.Sprintf("x.%s(\"%s\")", helper, escapeString(collapseWS(v))))
		}
	}

	// Compose final childCall
	return childCall{
		elem:      toElemCall(elem),
		attrCalls: args,
	}
}

func toElemCall(elem string) string {
	switch elem {
	case "path":
		return "Path"
	case "line":
		return "Line"
	case "rect":
		return "Rect"
	case "circle":
		return "Circle"
	case "ellipse":
		return "Ellipse"
	case "polyline":
		return "Polyline"
	case "polygon":
		return "Polygon"
	default:
		return strings.Title(elem)
	}
}

func renderIconFile(funcName, docTitle, className string, children []childCall) string {
	var b bytes.Buffer
	b.WriteString("package lucide\n\n")
	b.WriteString("import x \"github.com/bloxui/blox\"\n\n")
	fmt.Fprintf(&b, "// %s creates a %s Lucide icon.\n", funcName, docTitle)
	fmt.Fprintf(&b, "func %s(args ...x.SvgArg) x.Node {\n", funcName)
	fmt.Fprintf(&b, "\tsvgArgs := buildLucideArgs(\"%s\", args)\n", className)
	b.WriteString("\tsvgArgs = append(svgArgs,\n")
	for _, c := range children {
		// x.Child(x.<Elem>(<args>)),
		b.WriteString("\t\tx.Child(x.")
		b.WriteString(c.elem)
		b.WriteString("(")
		for i, a := range c.attrCalls {
			if i > 0 {
				b.WriteString(", ")
			}
			b.WriteString(a)
		}
		b.WriteString(")),\n")
	}
	b.WriteString("\t)\n")
	b.WriteString("\treturn x.Svg(svgArgs...)\n")
	b.WriteString("}\n")
	return b.String()
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

func toDocTitle(base string) string {
	parts := strings.Split(base, "-")
	var out []string
	for _, p := range parts {
		if p == "" {
			continue
		}
		r, size := utf8.DecodeRuneInString(p)
		if unicode.IsLetter(r) {
			out = append(out, strings.ToUpper(string(r))+p[size:])
		} else {
			out = append(out, p)
		}
	}
	return strings.Join(out, " ")
}

func local(n xml.Name) string {
	if n.Space != "" {
		return n.Local
	}
	return n.Local
}

func collapseWS(s string) string {
	// Trim leading/trailing and collapse internal sequences of spaces/newlines/tabs to single spaces
	s = strings.TrimSpace(s)
	var b strings.Builder
	prevSpace := false
	for _, r := range s {
		if r == ' ' || r == '\n' || r == '\t' || r == '\r' {
			if !prevSpace {
				b.WriteRune(' ')
				prevSpace = true
			}
			continue
		}
		prevSpace = false
		b.WriteRune(r)
	}
	return b.String()
}

func escapeString(s string) string {
	// Basic escaping for quotes and backslashes
	s = strings.ReplaceAll(s, "\\", "\\\\")
	s = strings.ReplaceAll(s, "\"", "\\\"")
	return s
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
