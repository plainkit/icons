package lucide

import x "github.com/bloxui/blox"

// TextCursorInput creates a Text Cursor Input Lucide icon.
func TextCursorInput(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-text-cursor-input", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 20h-1a2 2 0 0 1-2-2 2 2 0 0 1-2 2H6"))),
		x.Child(x.Path(x.D("M13 8h7a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2h-7"))),
		x.Child(x.Path(x.D("M5 16H4a2 2 0 0 1-2-2v-4a2 2 0 0 1 2-2h1"))),
		x.Child(x.Path(x.D("M6 4h1a2 2 0 0 1 2 2 2 2 0 0 1 2-2h1"))),
		x.Child(x.Path(x.D("M9 6v12"))),
	)
	return x.Svg(svgArgs...)
}
