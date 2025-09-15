package lucide

import x "github.com/plainkit/html"

// MessageCircleDashed creates a Message Circle Dashed Lucide icon.
func MessageCircleDashed(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-message-circle-dashed", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10.1 2.182a10 10 0 0 1 3.8 0"))),
		x.Child(x.Path(x.D("M13.9 21.818a10 10 0 0 1-3.8 0"))),
		x.Child(x.Path(x.D("M17.609 3.72a10 10 0 0 1 2.69 2.7"))),
		x.Child(x.Path(x.D("M2.182 13.9a10 10 0 0 1 0-3.8"))),
		x.Child(x.Path(x.D("M20.28 17.61a10 10 0 0 1-2.7 2.69"))),
		x.Child(x.Path(x.D("M21.818 10.1a10 10 0 0 1 0 3.8"))),
		x.Child(x.Path(x.D("M3.721 6.391a10 10 0 0 1 2.7-2.69"))),
		x.Child(x.Path(x.D("m6.163 21.117-2.906.85a1 1 0 0 1-1.236-1.169l.965-2.98"))),
	)
	return x.Svg(svgArgs...)
}
