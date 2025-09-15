package lucide

import x "github.com/plainkit/blox"

// Mails creates a Mails Lucide icon.
func Mails(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-mails", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M17 19a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2v-8a2 2 0 0 1 1-1.732"))),
		x.Child(x.Path(x.D("m22 5.5-6.419 4.179a2 2 0 0 1-2.162 0L7 5.5"))),
		x.Child(x.Rect(x.RectWidth("15"), x.RectHeight("12"), x.X("7"), x.Y("3"), x.Rx("2"))),
	)
	return x.Svg(svgArgs...)
}
