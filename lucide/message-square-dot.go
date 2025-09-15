package lucide

import x "github.com/plainkit/blox"

// MessageSquareDot creates a Message Square Dot Lucide icon.
func MessageSquareDot(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-message-square-dot", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12.7 3H4a2 2 0 0 0-2 2v16.286a.71.71 0 0 0 1.212.502l2.202-2.202A2 2 0 0 1 6.828 19H20a2 2 0 0 0 2-2v-4.7"))),
		x.Child(x.Circle(x.Cx("19"), x.Cy("6"), x.R("3"))),
	)
	return x.Svg(svgArgs...)
}
