package lucide

import x "github.com/plainkit/blox"

// ClipboardPlus creates a Clipboard Plus Lucide icon.
func ClipboardPlus(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-clipboard-plus", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("8"), x.RectHeight("4"), x.X("8"), x.Y("2"), x.Rx("1"), x.Ry("1"))),
		x.Child(x.Path(x.D("M16 4h2a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h2"))),
		x.Child(x.Path(x.D("M9 14h6"))),
		x.Child(x.Path(x.D("M12 17v-6"))),
	)
	return x.Svg(svgArgs...)
}
