package lucide

import x "github.com/bloxui/blox"

// UserStar creates a User Star Lucide icon.
func UserStar(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-user-star", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M16.051 12.616a1 1 0 0 1 1.909.024l.737 1.452a1 1 0 0 0 .737.535l1.634.256a1 1 0 0 1 .588 1.806l-1.172 1.168a1 1 0 0 0-.282.866l.259 1.613a1 1 0 0 1-1.541 1.134l-1.465-.75a1 1 0 0 0-.912 0l-1.465.75a1 1 0 0 1-1.539-1.133l.258-1.613a1 1 0 0 0-.282-.866l-1.156-1.153a1 1 0 0 1 .572-1.822l1.633-.256a1 1 0 0 0 .737-.535z"))),
		x.Child(x.Path(x.D("M8 15H7a4 4 0 0 0-4 4v2"))),
		x.Child(x.Circle(x.Cx("10"), x.Cy("7"), x.R("4"))),
	)
	return x.Svg(svgArgs...)
}
