package lucide

import x "github.com/plainkit/html"

// BugPlay creates a Bug Play Lucide icon.
func BugPlay(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-bug-play", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 19.655A6 6 0 0 1 6 14v-3a4 4 0 0 1 4-4h4a4 4 0 0 1 4 3.97"))),
		x.Child(x.Path(x.D("M14 15.003a1 1 0 0 1 1.517-.859l4.997 2.997a1 1 0 0 1 0 1.718l-4.997 2.997a1 1 0 0 1-1.517-.86z"))),
		x.Child(x.Path(x.D("M14.12 3.88 16 2"))),
		x.Child(x.Path(x.D("M21 5a4 4 0 0 1-3.55 3.97"))),
		x.Child(x.Path(x.D("M3 21a4 4 0 0 1 3.81-4"))),
		x.Child(x.Path(x.D("M3 5a4 4 0 0 0 3.55 3.97"))),
		x.Child(x.Path(x.D("M6 13H2"))),
		x.Child(x.Path(x.D("m8 2 1.88 1.88"))),
		x.Child(x.Path(x.D("M9 7.13V6a3 3 0 1 1 6 0v1.13"))),
	)
	return x.Svg(svgArgs...)
}
