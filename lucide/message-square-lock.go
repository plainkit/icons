package lucide

import x "github.com/plainkit/html"

// MessageSquareLock creates a Message Square Lock Lucide icon.
func MessageSquareLock(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-message-square-lock", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M22 8.5V5a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v16.286a.71.71 0 0 0 1.212.502l2.202-2.202A2 2 0 0 1 6.828 19H10"))),
		x.Child(x.Path(x.D("M20 15v-2a2 2 0 0 0-4 0v2"))),
		x.Child(x.Rect(x.RectWidth("8"), x.RectHeight("5"), x.X("14"), x.Y("15"), x.Rx("1"))),
	)
	return x.Svg(svgArgs...)
}
