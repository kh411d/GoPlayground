package main

import (
	"encoding/json"
	"fmt"
	"math"
)

// [{"x":25.119156,"y":55.364294},{"x":25.108791,"y":55.344119},{"x":25.076423,"y":55.303435},{"x":25.065724,"y":55.293763}]
/*
PHP
function GetCenterFromDegrees($data)
{
    if (!is_array($data)) return FALSE;

    $num_coords = count($data);

    $X = 0.0;
    $Y = 0.0;
    $Z = 0.0;

    foreach ($data as $coord)
    {
        $lat = $coord[0] * pi() / 180;
        $lon = $coord[1] * pi() / 180;

        $a = cos($lat) * cos($lon);
        $b = cos($lat) * sin($lon);
        $c = sin($lat);

        $X += $a;
        $Y += $b;
        $Z += $c;
    }

    $X /= $num_coords;
    $Y /= $num_coords;
    $Z /= $num_coords;

    $lon = atan2($Y, $X);
    $hyp = sqrt($X * $X + $Y * $Y);
    $lat = atan2($Z, $hyp);

    return array($lat * 180 / pi(), $lon * 180 / pi());
}

JS
function rad2degr(rad) { return rad * 180 / Math.PI; }
function degr2rad(degr) { return degr * Math.PI / 180; }

/**
 * @param latLngInDeg array of arrays with latitude and longtitude
 *   pairs in degrees. e.g. [[latitude1, longtitude1], [latitude2
 *   [longtitude2] ...]
 *
 * @return array with the center latitude longtitude pairs in
 *   degrees.
 *
 function getLatLngCenter(latLngInDegr) {
    var LATIDX = 0;
    var LNGIDX = 1;
    var sumX = 0;
    var sumY = 0;
    var sumZ = 0;

    for (var i=0; i<latLngInDegr.length; i++) {
        var lat = degr2rad(latLngInDegr[i][LATIDX]);
        var lng = degr2rad(latLngInDegr[i][LNGIDX]);
        // sum of cartesian coordinates
        sumX += Math.cos(lat) * Math.cos(lng);
        sumY += Math.cos(lat) * Math.sin(lng);
        sumZ += Math.sin(lat);
    }

    var avgX = sumX / latLngInDegr.length;
    var avgY = sumY / latLngInDegr.length;
    var avgZ = sumZ / latLngInDegr.length;

    // convert average x, y, z coordinate to latitude and longtitude
    var lng = Math.atan2(avgY, avgX);
    var hyp = Math.sqrt(avgX * avgX + avgY * avgY);
    var lat = Math.atan2(avgZ, hyp);

    return ([rad2degr(lat), rad2degr(lng)]);
}
*/
type latlng struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func GetCenterFromLatLngs(data []latlng) latlng {
	totalData := len(data)
	if totalData == 0 {
		return latlng{}
	}
	var x, y, z float64

	for _, c := range data {
		lat := c.X * math.Pi / 180
		lon := c.Y * math.Pi / 180

		x += math.Cos(lat) * math.Cos(lon)
		y += math.Cos(lat) * math.Sin(lon)
		z += math.Sin(lat)
	}

	x = x / float64(totalData)
	y = y / float64(totalData)
	z = z / float64(totalData)

	aLon := math.Atan2(y, x)
	hyp := math.Sqrt(x*x + y*y)
	aLat := math.Atan2(z, hyp)

	return latlng{
		X: aLat * 180 / math.Pi,
		Y: aLon * 180 / math.Pi,
	}
}

func main() {
	//data := `[{"x":24.796523,"y":46.722549},{"x":24.796523,"y":46.722549},{"x":24.8283,"y":46.724966},{"x":24.854495,"y":46.741987},{"x":24.878212,"y":46.731167},{"x":24.902339,"y":46.70544},{"x":24.979547,"y":46.660438},{"x":24.984833,"y":46.683124},{"x":25.044388,"y":46.695274},{"x":25.079395,"y":46.714447},{"x":25.037146,"y":46.755871},{"x":24.964778,"y":46.788427},{"x":24.970833,"y":46.805808},{"x":24.904847,"y":46.867544},{"x":24.885955,"y":46.88341},{"x":24.853458,"y":46.857902},{"x":24.796523,"y":46.722549}]`
	data := `[{"x":24.879628,"y":46.721753},{"x":24.879628,"y":46.721753},{"x":24.824244,"y":46.585304},{"x":24.801135,"y":46.55964},{"x":24.862107,"y":46.546369},{"x":24.878611,"y":46.586725},{"x":24.996205,"y":46.536064},{"x":25.026394,"y":46.617978},{"x":24.978273,"y":46.65686},{"x":24.904786,"y":46.697275},{"x":24.879628,"y":46.721753}]`
	var rdata []latlng
	json.Unmarshal([]byte(data), &rdata)
	l := GetCenterFromLatLngs(rdata)
	fmt.Printf("%#v\n", l)
}
