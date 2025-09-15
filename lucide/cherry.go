package lucide

import x "github.com/plainkit/html"

// Cherry creates a Cherry Lucide icon.
func Cherry(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-cherry", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M2 17a5 5 0 0 0 10 0c0-2.76-2.5-5-5-3-2.5-2-5 .24-5 3Z"))),
		x.Child(x.Path(x.D("M12 17a5 5 0 0 0 10 0c0-2.76-2.5-5-5-3-2.5-2-5 .24-5 3Z"))),
		x.Child(x.Path(x.D("M7 14c3.22-2.91 4.29-8.75 5-12 1.66 2.38 4.94 9 5 12"))),
		x.Child(x.Path(x.D("M22 9c-4.29 0-7.14-2.33-10-7 5.71 0 10 4.67 10 7Z"))),
	)
	return x.Svg(svgArgs...)
}
