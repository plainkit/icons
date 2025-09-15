package lucide

import x "github.com/plainkit/html"

// ToyBrick creates a Toy Brick Lucide icon.
func ToyBrick(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-toy-brick", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("12"), x.X("3"), x.Y("8"), x.Rx("1"))),
		x.Child(x.Path(x.D("M10 8V5c0-.6-.4-1-1-1H6a1 1 0 0 0-1 1v3"))),
		x.Child(x.Path(x.D("M19 8V5c0-.6-.4-1-1-1h-3a1 1 0 0 0-1 1v3"))),
	)
	return x.Svg(svgArgs...)
}
