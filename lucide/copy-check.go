package lucide

import x "github.com/bloxui/blox"

// CopyCheck creates a Copy Check Lucide icon.
func CopyCheck(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-copy-check", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m12 15 2 2 4-4"))),
		x.Child(x.Rect(x.RectWidth("14"), x.RectHeight("14"), x.X("8"), x.Y("8"), x.Rx("2"), x.Ry("2"))),
		x.Child(x.Path(x.D("M4 16c-1.1 0-2-.9-2-2V4c0-1.1.9-2 2-2h10c1.1 0 2 .9 2 2"))),
	)
	return x.Svg(svgArgs...)
}
