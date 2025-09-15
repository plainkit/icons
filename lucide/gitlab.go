package lucide

import x "github.com/plainkit/html"

// Gitlab creates a Gitlab Lucide icon.
func Gitlab(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-gitlab", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m22 13.29-3.33-10a.42.42 0 0 0-.14-.18.38.38 0 0 0-.22-.11.39.39 0 0 0-.23.07.42.42 0 0 0-.14.18l-2.26 6.67H8.32L6.1 3.26a.42.42 0 0 0-.1-.18.38.38 0 0 0-.26-.08.39.39 0 0 0-.23.07.42.42 0 0 0-.14.18L2 13.29a.74.74 0 0 0 .27.83L12 21l9.69-6.88a.71.71 0 0 0 .31-.83Z"))),
	)
	return x.Svg(svgArgs...)
}
