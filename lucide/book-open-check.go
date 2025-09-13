package lucide

import x "github.com/bloxui/blox"

// BookOpenCheck creates a Book Open Check Lucide icon.
func BookOpenCheck(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-book-open-check", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 21V7"))),
		x.Child(x.Path(x.D("m16 12 2 2 4-4"))),
		x.Child(x.Path(x.D("M22 6V4a1 1 0 0 0-1-1h-5a4 4 0 0 0-4 4 4 4 0 0 0-4-4H3a1 1 0 0 0-1 1v13a1 1 0 0 0 1 1h6a3 3 0 0 1 3 3 3 3 0 0 1 3-3h6a1 1 0 0 0 1-1v-1.3"))),
	)
	return x.Svg(svgArgs...)
}
