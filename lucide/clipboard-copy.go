package lucide

import x "github.com/plainkit/html"

// ClipboardCopy creates a Clipboard Copy Lucide icon.
func ClipboardCopy(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-clipboard-copy", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("8"), x.RectHeight("4"), x.X("8"), x.Y("2"), x.Rx("1"), x.Ry("1"))),
		x.Child(x.Path(x.D("M8 4H6a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-2"))),
		x.Child(x.Path(x.D("M16 4h2a2 2 0 0 1 2 2v4"))),
		x.Child(x.Path(x.D("M21 14H11"))),
		x.Child(x.Path(x.D("m15 10-4 4 4 4"))),
	)
	return x.Svg(svgArgs...)
}
