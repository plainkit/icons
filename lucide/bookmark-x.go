package lucide

import x "github.com/plainkit/html"

// BookmarkX creates a Bookmark X Lucide icon.
func BookmarkX(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-bookmark-x", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m19 21-7-4-7 4V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2Z"))),
		x.Child(x.Path(x.D("m14.5 7.5-5 5"))),
		x.Child(x.Path(x.D("m9.5 7.5 5 5"))),
	)
	return x.Svg(svgArgs...)
}
