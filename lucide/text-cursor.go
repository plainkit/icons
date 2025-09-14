package lucide

import x "github.com/bloxui/blox"

// TextCursor creates a Text Cursor Lucide icon.
func TextCursor(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-text-cursor", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M17 22h-1a4 4 0 0 1-4-4V6a4 4 0 0 1 4-4h1"))),
		x.Child(x.Path(x.D("M7 22h1a4 4 0 0 0 4-4v-1"))),
		x.Child(x.Path(x.D("M7 2h1a4 4 0 0 1 4 4v1"))),
	)
	return x.Svg(svgArgs...)
}
