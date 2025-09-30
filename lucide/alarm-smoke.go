package lucide

import (
	html "github.com/plainkit/html"
)

// AlarmSmoke creates a Alarm Smoke Lucide icon.
func AlarmSmoke(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-alarm-smoke", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M11 21c0-2.5 2-2.5 2-5")),
		html.SvgPath(html.AD("M16 21c0-2.5 2-2.5 2-5")),
		html.SvgPath(html.AD("m19 8-.8 3a1.25 1.25 0 0 1-1.2 1H7a1.25 1.25 0 0 1-1.2-1L5 8")),
		html.SvgPath(html.AD("M21 3a1 1 0 0 1 1 1v2a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V4a1 1 0 0 1 1-1z")),
		html.SvgPath(html.AD("M6 21c0-2.5 2-2.5 2-5")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
