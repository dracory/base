package img_test

import (
	"net/url"
	"sort"
	"strings"
	"testing"

	"github.com/dracory/base/img" // Adjust import path if necessary
)

func TestPicsumURL(t *testing.T) {
	tests := []struct {
		name   string
		width  int
		height int
		opts   img.PicsumURLOptions
		want   string // Expected URL string
	}{
		{
			name:   "Basic URL",
			width:  200,
			height: 300,
			opts:   img.PicsumURLOptions{},
			want:   "https://picsum.photos/200/300",
		},
		{
			name:   "With Seed",
			width:  400,
			height: 200,
			opts:   img.PicsumURLOptions{Seed: "testseed"},
			want:   "https://picsum.photos/seed/testseed/400/200",
		},
		{
			name:   "With ID",
			width:  150,
			height: 150,
			opts:   img.PicsumURLOptions{ID: 123},
			want:   "https://picsum.photos/id/123/150/150",
		},
		{
			name:   "With Seed and ID",
			width:  300,
			height: 250,
			opts:   img.PicsumURLOptions{Seed: "anotha", ID: 456},
			want:   "https://picsum.photos/seed/anotha/id/456/300/250", // Seed comes first
		},
		{
			name:   "With Grayscale",
			width:  500,
			height: 500,
			opts:   img.PicsumURLOptions{Grayscale: true},
			want:   "https://picsum.photos/500/500?grayscale=", // Expect 'grayscale='
		},
		{
			name:   "With Blur",
			width:  100,
			height: 400,
			opts:   img.PicsumURLOptions{Blur: 5},
			want:   "https://picsum.photos/100/400?blur=5", // No leading '&'
		},
		{
			name:   "With Grayscale and Blur",
			width:  600,
			height: 300,
			opts:   img.PicsumURLOptions{Grayscale: true, Blur: 10},
			// Order doesn't matter for query params with net/url, but Encode sorts them.
			// blur comes before grayscale alphabetically.
			want: "https://picsum.photos/600/300?blur=10&grayscale=",
		},
		{
			name:   "With Seed, ID, Grayscale, and Blur",
			width:  800,
			height: 600,
			opts:   img.PicsumURLOptions{Seed: "full", ID: 789, Grayscale: true, Blur: 2},
			// Order doesn't matter, but Encode sorts them.
			want: "https://picsum.photos/seed/full/id/789/800/600?blur=2&grayscale=",
		},
		{
			name:   "Zero Width and Height",
			width:  0,
			height: 0,
			opts:   img.PicsumURLOptions{},
			want:   "https://picsum.photos/0/0",
		},
		{
			name:   "Blur value > 10 (gets clamped)", // Updated description
			width:  100,
			height: 100,
			opts:   img.PicsumURLOptions{Blur: 15},
			want:   "https://picsum.photos/100/100?blur=10", // Expect clamped value 10
		},
		{
			name:   "Blur value < 1 (treated as 0, so no blur param)",
			width:  100,
			height: 100,
			opts:   img.PicsumURLOptions{Blur: -5},
			want:   "https://picsum.photos/100/100", // Blur <= 0 is ignored
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := img.PicsumURL(tt.width, tt.height, tt.opts)

			// Separate base URL and query string for comparison
			baseWant, queryWantStr, hasQueryWant := strings.Cut(tt.want, "?")
			baseGot, queryGotStr, hasQueryGot := strings.Cut(got, "?")

			// Compare base URLs
			if baseGot != baseWant {
				t.Errorf("PicsumURL() base URL mismatch:\ngot = %q\nwant= %q", baseGot, baseWant)
				return // No point checking query if base is wrong
			}

			// Compare query string presence
			if hasQueryGot != hasQueryWant {
				t.Errorf("PicsumURL() query string presence mismatch:\ngot query = %v (%q)\nwant query= %v (%q)", hasQueryGot, queryGotStr, hasQueryWant, queryWantStr)
				return
			}

			// Compare query strings if they exist (order-independent)
			if hasQueryWant {
				wantParams, errWant := url.ParseQuery(queryWantStr)
				if errWant != nil {
					t.Fatalf("Error parsing want query string %q: %v", queryWantStr, errWant)
				}
				gotParams, errGot := url.ParseQuery(queryGotStr)
				if errGot != nil {
					t.Fatalf("Error parsing got query string %q: %v", queryGotStr, errGot)
				}

				if !compareURLValues(wantParams, gotParams) {
					t.Errorf("PicsumURL() query params mismatch:\ngot = %q\nwant= %q", queryGotStr, queryWantStr)
				}
			}
		})
	}
}

// compareURLValues checks if two url.Values are equivalent, ignoring order.
func compareURLValues(v1, v2 url.Values) bool {
	if len(v1) != len(v2) {
		return false
	}
	for k, val1 := range v1 {
		val2, ok := v2[k]
		if !ok {
			return false // Key missing in v2
		}
		if !equalStringSlices(val1, val2) {
			return false // Values for key differ
		}
	}
	// We don't need to check the other way around because lengths are equal
	return true
}

// equalStringSlices checks if two string slices are equal (order matters within slice).
// For query params, usually there's only one value per key, but this handles multivalue cases.
func equalStringSlices(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	// If order doesn't matter within the slice for a given key, sort them first:
	sort.Strings(s1)
	sort.Strings(s2)
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

// TestPicsumURL_QueryParamOrderIndependence is less critical now since net/url.Values.Encode() sorts keys,
// but we can keep it to verify our comparison logic works.
func TestPicsumURL_QueryParamOrderIndependence(t *testing.T) {
	opts := img.PicsumURLOptions{Grayscale: true, Blur: 5}
	width, height := 300, 300

	// Expected result (net/url sorts keys: blur, grayscale)
	want := "https://picsum.photos/300/300?blur=5&grayscale="

	got := img.PicsumURL(width, height, opts)

	// Direct comparison should work now
	if got != want {
		t.Errorf("PicsumURL() mismatch:\ngot = %q\nwant= %q", got, want)
	}

	// Optional: Verify using the comparison helper as well
	baseWant, queryWantStr, _ := strings.Cut(want, "?")
	baseGot, queryGotStr, _ := strings.Cut(got, "?")

	if baseGot != baseWant {
		t.Errorf("PicsumURL() base URL mismatch:\ngot = %q\nwant= %q", baseGot, baseWant)
	}

	wantParams, _ := url.ParseQuery(queryWantStr)
	gotParams, _ := url.ParseQuery(queryGotStr)

	if !compareURLValues(wantParams, gotParams) {
		t.Errorf("PicsumURL() query params mismatch (using helper):\ngot = %q\nwant= %q", queryGotStr, queryWantStr)
	}
}
