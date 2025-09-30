package lucide

import (
	html "github.com/plainkit/html"
)

// UserStar creates a User Star Lucide icon.
func UserStar(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-user-star", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M16.051 12.616a1 1 0 0 1 1.909.024l.737 1.452a1 1 0 0 0 .737.535l1.634.256a1 1 0 0 1 .588 1.806l-1.172 1.168a1 1 0 0 0-.282.866l.259 1.613a1 1 0 0 1-1.541 1.134l-1.465-.75a1 1 0 0 0-.912 0l-1.465.75a1 1 0 0 1-1.539-1.133l.258-1.613a1 1 0 0 0-.282-.866l-1.156-1.153a1 1 0 0 1 .572-1.822l1.633-.256a1 1 0 0 0 .737-.535z")),
		html.SvgPath(html.AD("M8 15H7a4 4 0 0 0-4 4v2")),
		html.SvgCircle(html.ACx("10"), html.ACy("7"), html.AR("4")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
