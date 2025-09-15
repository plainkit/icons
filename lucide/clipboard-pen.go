package lucide

import x "github.com/plainkit/blox"

// ClipboardPen creates a Clipboard Pen Lucide icon.
func ClipboardPen(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-clipboard-pen", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("8"), x.RectHeight("4"), x.X("8"), x.Y("2"), x.Rx("1"))),
		x.Child(x.Path(x.D("M16 4h2a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2h-5.5"))),
		x.Child(x.Path(x.D("M4 13.5V6a2 2 0 0 1 2-2h2"))),
		x.Child(x.Path(x.D("M13.378 15.626a1 1 0 1 0-3.004-3.004l-5.01 5.012a2 2 0 0 0-.506.854l-.837 2.87a.5.5 0 0 0 .62.62l2.87-.837a2 2 0 0 0 .854-.506z"))),
	)
	return x.Svg(svgArgs...)
}
