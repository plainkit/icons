package lucide

import x "github.com/bloxui/blox"

// Bookmark creates a Bookmark Lucide icon.
func Bookmark(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-bookmark", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m19 21-7-4-7 4V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v16z"))),
	)
	return x.Svg(svgArgs...)
}
