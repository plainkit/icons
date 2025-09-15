package lucide

import x "github.com/plainkit/html"

// LibraryBig creates a Library Big Lucide icon.
func LibraryBig(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-library-big", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("8"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("1"))),
		x.Child(x.Path(x.D("M7 3v18"))),
		x.Child(x.Path(x.D("M20.4 18.9c.2.5-.1 1.1-.6 1.3l-1.9.7c-.5.2-1.1-.1-1.3-.6L11.1 5.1c-.2-.5.1-1.1.6-1.3l1.9-.7c.5-.2 1.1.1 1.3.6Z"))),
	)
	return x.Svg(svgArgs...)
}
