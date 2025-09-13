package lucide

import x "github.com/bloxui/blox"

// BookmarkCheck creates a Bookmark Check Lucide icon.
func BookmarkCheck(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-bookmark-check", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m19 21-7-4-7 4V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2Z"))),
		x.Child(x.Path(x.D("m9 10 2 2 4-4"))),
	)
	return x.Svg(svgArgs...)
}
