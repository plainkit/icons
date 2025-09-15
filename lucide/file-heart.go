package lucide

import x "github.com/plainkit/blox"

// FileHeart creates a File Heart Lucide icon.
func FileHeart(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-file-heart", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M14 2v4a2 2 0 0 0 2 2h4"))),
		x.Child(x.Path(x.D("M2.62 13.8A2.25 2.25 0 1 1 6 10.836a2.25 2.25 0 1 1 3.38 2.966l-2.626 2.856a.998.998 0 0 1-1.507 0z"))),
		x.Child(x.Path(x.D("M4 6.005V4a2 2 0 0 1 2-2h9l5 5v13a2 2 0 0 1-2 2H6a2 2 0 0 1-1.9-1.376"))),
	)
	return x.Svg(svgArgs...)
}
