package gpkg

import (
	"testing"
)

func uiptr(ui uint) *uint {
	return &ui
}

func TestReplaceTokens(t *testing.T) {
	type tcase struct {
		qtext    string
		zoom     *uint
		extent   *[4]float64
		expected string
	}

	fn := func(t *testing.T, tc tcase) {
		output := replaceTokens(tc.qtext, tc.zoom, tc.extent)

		if tc.expected != output {
			t.Errorf("expected %v\n got\n %v", tc.expected, output)
			return
		}
	}

	tests := map[string]tcase{
		"zoom": {
			qtext: `
				SELECT
					fid, geom, featurecla, min_zoom, 22 as max_zoom, minx, miny, maxx, maxy
				FROM
					ne_110m_land t JOIN rtree_ne_110m_land_geom si ON t.fid = si.id
				WHERE
					min_zoom <= !ZOOM! AND max_zoom >= !ZOOM!`,
			zoom: uiptr(9),
			expected: `
				SELECT
					fid, geom, featurecla, min_zoom, 22 as max_zoom, minx, miny, maxx, maxy
				FROM
					ne_110m_land t JOIN rtree_ne_110m_land_geom si ON t.fid = si.id
				WHERE
					min_zoom <= 9 AND max_zoom >= 9`,
		},
		"bbox": {
			qtext: `
				SELECT
					fid, geom, featurecla, min_zoom, 22 as max_zoom, minx, miny, maxx, maxy
				FROM
					ne_110m_land t JOIN rtree_ne_110m_land_geom si ON t.fid = si.id
				WHERE
					!BBOX!`,
			extent: &[4]float64{
				-180, -85.0511,
				180, 85.0511,
			},
			expected: `
				SELECT
					fid, geom, featurecla, min_zoom, 22 as max_zoom, minx, miny, maxx, maxy
				FROM
					ne_110m_land t JOIN rtree_ne_110m_land_geom si ON t.fid = si.id
				WHERE
					minx <= 180 AND maxx >= -180 AND miny <= 85.0511 AND maxy >= -85.0511`,
		},
		"bbox zoom": {
			qtext: `
				SELECT
					fid, geom, featurecla, min_zoom, 22 as max_zoom, minx, miny, maxx, maxy
				FROM
					ne_110m_land t JOIN rtree_ne_110m_land_geom si ON t.fid = si.id
				WHERE
					!BBOX! AND min_zoom = !ZOOM!`,
			extent: &[4]float64{
				-180, -85.0511,
				180, 85.0511,
			},
			zoom: uiptr(3),
			expected: `
				SELECT
					fid, geom, featurecla, min_zoom, 22 as max_zoom, minx, miny, maxx, maxy
				FROM
					ne_110m_land t JOIN rtree_ne_110m_land_geom si ON t.fid = si.id
				WHERE
					minx <= 180 AND maxx >= -180 AND miny <= 85.0511 AND maxy >= -85.0511 AND min_zoom = 3`,
		},
	}

	for name, tc := range tests {
		tc := tc
		t.Run(name, func(t *testing.T) {
			fn(t, tc)
		})
	}
}
