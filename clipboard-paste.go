package lucide

import x "github.com/bloxui/blox"

// ClipboardPaste creates a Clipboard Paste Lucide icon.
func ClipboardPaste(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-clipboard-paste", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M11 14h10"))),
		x.Child(x.Path(x.D("M16 4h2a2 2 0 0 1 2 2v1.344"))),
		x.Child(x.Path(x.D("m17 18 4-4-4-4"))),
		x.Child(x.Path(x.D("M8 4H6a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h12a2 2 0 0 0 1.793-1.113"))),
		x.Child(x.Rect(x.X("8"), x.Y("2"), x.RectWidth("8"), x.RectHeight("4"), x.Rx("1"))),
	)
	return x.Svg(svgArgs...)
}
