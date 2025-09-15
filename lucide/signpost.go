package lucide

import x "github.com/plainkit/html"

// Signpost creates a Signpost Lucide icon.
func Signpost(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-signpost", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 13v8"))),
		x.Child(x.Path(x.D("M12 3v3"))),
		x.Child(x.Path(x.D("M18 6a2 2 0 0 1 1.387.56l2.307 2.22a1 1 0 0 1 0 1.44l-2.307 2.22A2 2 0 0 1 18 13H6a2 2 0 0 1-1.387-.56l-2.306-2.22a1 1 0 0 1 0-1.44l2.306-2.22A2 2 0 0 1 6 6z"))),
	)
	return x.Svg(svgArgs...)
}
