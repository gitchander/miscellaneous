package learning

import (
	"testing"
)

func TestPolygon(t *testing.T) {

	type testSample struct {
		a, b []Point
		ok   bool
	}

	var samples = map[string]testSample{
		"both-nil": {
			a:  nil,
			b:  nil,
			ok: false,
		},
		"both-empty": {
			a:  []Point{},
			b:  []Point{},
			ok: false,
		},
		"both-(0,0)": {
			a:  []Point{Pt(0, 0)},
			b:  []Point{Pt(0, 0)},
			ok: false,
		},
		"(0,0)-(1,0)": {
			a:  []Point{Pt(0, 0)},
			b:  []Point{Pt(1, 0)},
			ok: true,
		},
		"empty-some": {
			a:  []Point{},
			b:  []Point{Pt(1, 1), Pt(1, 2), Pt(2, 1)},
			ok: false,
		},
		"some-empty": {
			a:  []Point{Pt(1, 3), Pt(2, 2), Pt(3, 1)},
			b:  []Point{},
			ok: false,
		},
		"some-some-1": {
			a:  []Point{Pt(1, 3), Pt(2, 2), Pt(3, 1)},
			b:  []Point{Pt(1, 1), Pt(1, 2), Pt(2, 1)},
			ok: true,
		},
		"some-some-2": {
			a:  []Point{Pt(1, 3), Pt(2, 2), Pt(1, 1)},
			b:  []Point{Pt(1, 2), Pt(3, 1), Pt(2, 1)},
			ok: false,
		},
		"some-some-3": {
			a:  []Point{Pt(1, 0), Pt(2, 0), Pt(3, 0)},
			b:  []Point{Pt(4, 0), Pt(5, 0), Pt(6, 0)},
			ok: true,
		},
		"lines-both-equal": {
			a: LinesToPoints(
				TextSample{
					Target: 'A',
					Lines: []string{
						"A-",
						"--",
					},
				}),
			b: LinesToPoints(
				TextSample{
					Target: 'B',
					Lines: []string{
						"B-",
						"--",
					},
				}),
			ok: false,
		},
		"lines-B inside A": {
			a: LinesToPoints(
				TextSample{
					Target: 'A',
					Lines: []string{
						"--A",
						"---",
						"A--",
					},
				}),
			b: LinesToPoints(
				TextSample{
					Target: 'B',
					Lines: []string{
						"---",
						"-B-",
						"---",
					},
				}),
			ok: false,
		},
		"point on line": {
			a: LinesToPoints(
				TextSample{
					Target: 'A',
					Lines: []string{
						"---",
						"-A-",
						"---",
					},
				}),
			b: LinesToPoints(
				TextSample{
					Target: 'B',
					Lines: []string{
						"--B",
						"-B-",
						"B--",
					},
				}),
			ok: false,
		},
		"cross lines": {
			a: LinesToPoints(
				TextSample{
					Target: 'A',
					Lines: []string{
						"A--",
						"-A-",
						"--A",
					},
				}),
			b: LinesToPoints(
				TextSample{
					Target: 'B',
					Lines: []string{
						"--B",
						"-B-",
						"B--",
					},
				}),
			ok: false,
		},
		"cross lines 2": {
			a: LinesToPoints(
				TextSample{
					Target: 'A',
					Lines: []string{
						"A---",
						"-A--",
						"--A-",
					},
				}),
			b: LinesToPoints(
				TextSample{
					Target: 'B',
					Lines: []string{
						"-B--",
						"--B-",
						"---B",
					},
				}),
			ok: true,
		},
		"cross lines 3": {
			a: LinesToPoints(
				TextSample{
					Target: 'A',
					Lines: []string{
						"-A--",
						"--A-",
						"---A",
					},
				}),
			b: LinesToPoints(
				TextSample{
					Target: 'B',
					Lines: []string{
						"----",
						"-B--",
						"B---",
					},
				}),
			ok: true,
		},
		"cross lines 4": {
			a: LinesToPoints(
				TextSample{
					Target: 'A',
					Lines: []string{
						"-A--",
						"--A-",
						"---A",
					},
				}),
			b: LinesToPoints(
				TextSample{
					Target: 'B',
					Lines: []string{
						"--B-",
						"-B--",
						"B---",
					},
				}),
			ok: false,
		},
		"cross lines flat 1": {
			a: LinesToPoints(
				TextSample{
					Target: 'A',
					Lines: []string{
						"AAA",
					},
				}),
			b: LinesToPoints(
				TextSample{
					Target: 'B',
					Lines: []string{
						"BBB",
					},
				}),
			ok: false,
		},
		"cross lines flat 2": {
			a: LinesToPoints(
				TextSample{
					Target: 'A',
					Lines: []string{
						"-AAA-",
					},
				}),
			b: LinesToPoints(
				TextSample{
					Target: 'B',
					Lines: []string{
						"BBBBB",
					},
				}),
			ok: false,
		},
		"cross lines flat 3": {
			a: LinesToPoints(
				TextSample{
					Target: 'A',
					Lines: []string{
						"AAA--",
					},
				}),
			b: LinesToPoints(
				TextSample{
					Target: 'B',
					Lines: []string{
						"--BBB",
					},
				}),
			ok: false,
		},
		"cross lines flat 4": {
			a: LinesToPoints(
				TextSample{
					Target: 'A',
					Lines: []string{
						"AAA---",
					},
				}),
			b: LinesToPoints(
				TextSample{
					Target: 'B',
					Lines: []string{
						"---BBB",
					},
				}),
			ok: true,
		},
		"cross lines flat 5": {
			a: LinesToPoints(
				TextSample{
					Target: 'A',
					Lines: []string{
						"AAA------",
					},
				}),
			b: LinesToPoints(
				TextSample{
					Target: 'B',
					Lines: []string{
						"------BBB",
					},
				}),
			ok: true,
		},
		"cross lines flat 6": {
			a: LinesToPoints(
				TextSample{
					Target: 'A',
					Lines: []string{
						"A",
						"A",
						"A",
					},
				}),
			b: LinesToPoints(
				TextSample{
					Target: 'B',
					Lines: []string{
						"-",
						"-",
						"-",
						"B",
						"B",
						"B",
					},
				}),
			ok: true,
		},
		"cross lines flat 7": {
			a: LinesToPoints(
				TextSample{
					Target: 'A',
					Lines: []string{
						"A",
						"A",
						"A",
					},
				}),
			b: LinesToPoints(
				TextSample{
					Target: 'B',
					Lines: []string{
						"-",
						"-",
						"B",
						"B",
						"B",
					},
				}),
			ok: false,
		},
		"lines-triangles": {
			a: LinesToPoints(
				TextSample{
					Target: 'A',
					Lines: []string{
						"AAAA-",
						"AAA--",
						"AA---",
						"A----",
					},
				}),
			b: LinesToPoints(
				TextSample{
					Target: 'B',
					Lines: []string{
						"----B",
						"---BB",
						"--BBB",
						"-BBBB",
					},
				}),
			ok: true,
		},
		"in rhombus": {
			a: LinesToPoints(
				TextSample{
					Target: 'A',
					Lines: []string{
						"--A--",
						"-A-A-",
						"A---A",
						"-A-A-",
						"--A--",
					},
				}),
			b: LinesToPoints(
				TextSample{
					Target: 'B',
					Lines: []string{
						"-----",
						"-----",
						"--B--",
						"-----",
						"-----",
					},
				}),
			ok: false,
		},
		"rhombus in square": {
			a: LinesToPoints(
				TextSample{
					Target: 'A',
					Lines: []string{
						"--A--",
						"-A-A-",
						"A---A",
						"-A-A-",
						"--A--",
					},
				}),
			b: LinesToPoints(
				TextSample{
					Target: 'B',
					Lines: []string{
						"B---B",
						"-----",
						"-----",
						"-----",
						"B---B",
					},
				}),
			ok: false,
		},
		"rhombus in rhombus": {
			a: LinesToPoints(
				TextSample{
					Target: 'A',
					Lines: []string{
						"-------",
						"---A---",
						"--A-A--",
						"-A---A-",
						"--A-A--",
						"---A---",
						"-------",
					},
				}),
			b: LinesToPoints(
				TextSample{
					Target: 'B',
					Lines: []string{
						"---B---",
						"--B-B--",
						"-B---B-",
						"B-----B",
						"-B---B-",
						"--B-B--",
						"---B---",
					},
				}),
			ok: false,
		},
		"square in square": {
			a: LinesToPoints(
				TextSample{
					Target: 'A',
					Lines: []string{
						"-------",
						"-------",
						"--AAA--",
						"--A-A--",
						"--AAA--",
						"-------",
						"-------",
					},
				}),
			b: LinesToPoints(
				TextSample{
					Target: 'B',
					Lines: []string{
						"BBBBBBB",
						"B-----B",
						"B-----B",
						"B-----B",
						"B-----B",
						"B-----B",
						"BBBBBBB",
					},
				}),
			ok: false,
		},
		"square in square 2": {
			a: LinesToPoints(
				TextSample{
					Target: 'A',
					Lines: []string{
						"------------",
						"------------",
						"-------AAA--",
						"-------A-A--",
						"-------AAA--",
						"------------",
						"------------",
					},
				}),
			b: LinesToPoints(
				TextSample{
					Target: 'B',
					Lines: []string{
						"BBBBBBB",
						"B-----B",
						"B-----B",
						"B-----B",
						"B-----B",
						"B-----B",
						"BBBBBBB",
					},
				}),
			ok: true,
		},
	}

	for key, sample := range samples {
		ok := SeparableStraight(sample.a, sample.b)
		if ok != sample.ok {
			t.Fatalf("sample %q invalid result: have %t, want %t", key, ok, sample.ok)
		}
	}
}

type TextSample struct {
	Target byte
	Lines  []string
}

func LinesToPoints(ts TextSample) []Point {
	var ps []Point
	for y, line := range ts.Lines {
		bs := []byte(line)
		for x, b := range bs {
			if b == ts.Target {
				ps = append(ps, Pt(x, y))
			}
		}
	}
	return ps
}
