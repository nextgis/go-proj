/*
 * File: alg.go
 * Project: go-proj
 * File Created: Saturday, 27th February 2021 7:34:03 pm
 * Author: Dmitry Baryshnikov <dmitry.baryshnikov@nextgis.com>
 * -----
 * Last Modified: Saturday, 27th February 2021 7:34:16 pm
 * Modified By: Dmitry Baryshnikov, <dmitry.baryshnikov@nextgis.com>
 * -----
 * Copyright 2019 - 2021 NextGIS, <info@nextgis.com>
 *
 *   Permission is hereby granted, free of charge, to any person obtaining a copy
 *   of this software and associated documentation files (the "Software"), to deal
 *   in the Software without restriction, including without limitation the rights
 *   to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 *   copies of the Software, and to permit persons to whom the Software is
 *   furnished to do so, subject to the following conditions:
 *
 *   The above copyright notice and this permission notice shall be included in
 *   all copies or substantial portions of the Software.
 *
 *   THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 *   IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 *   FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 *   AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 *   LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 *   OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 *   THE SOFTWARE.
 */

package proj

// #cgo LDFLAGS: -lproj
// #include <geodesic.h>
// struct PJconsts {
//     void *ctx;
//     const char *descr;
//     void *params;
//     char *def_full;
//     struct PJconsts *parent;
//     char *def_size;
//     char *def_shape;
//     char *def_spherification;
//     char *def_ellps;
//     struct geod_geodesic *geod;
// };
import "C"

import (
	"errors"
	"unsafe"
)

// GeodesicArea Calculate geodesic area
func (p *Proj) GeodesicArea(xs, ys []float64) (float64, float64, error) {
	if p == nil {
		return 0.0, 0.0, errors.New("Missing or invalid projection")
	}

	coordCountX := len(xs)
	coordCountY := len(ys)

	if coordCountX != coordCountY {
		return 0.0, 0.0, errors.New("Number of x and y coordinates differs")
	}

	if xs == nil || ys == nil {
		return 0.0, 0.0, nil
	}

	coordCount := C.int(coordCountX)
	var area C.double = 0.0
	var dist C.double = 0.0

	// NOTE: Reinterpret cast here
	PJ := (*C.struct_PJconsts)(p.p)

// void GEOD_DLL geod_polygonarea(const struct geod_geodesic* g,
// 	double lats[], double lons[], int n,
// 	double* pA, double* pP);

	C.geod_polygonarea(
		PJ.geod, 
		(*C.double)(unsafe.Pointer(&ys[0])),
		(*C.double)(unsafe.Pointer(&xs[0])), 
		coordCount, &area, &dist)

	return float64(area), float64(dist), nil
}

// GeodesicDistance  Calculate geodesic distance
func (p *Proj) GeodesicDistance(xs, ys []float64) (float64, error) {
	if p == nil {
		return 0.0, errors.New("Missing or invalid projection")
	}

	coordCountX := len(xs)
	coordCountY := len(ys)

	if coordCountX != coordCountY {
		return 0.0, errors.New("Number of x and y coordinates differs")
	}

	if xs == nil || ys == nil {
		return 0.0, nil
	}
	
	fullDist := 0.0

	// NOTE: Reinterpret cast here
	PJ := (*C.struct_PJconsts)(p.p)


// void GEOD_DLL geod_inverse(const struct geod_geodesic* g,
// 	double lat1, double lon1,
// 	double lat2, double lon2,
// 	double* ps12, double* pazi1, double* pazi2);
	for i := 0; i < coordCountX - 1; i++ {
		var dist C.double = 0.0
		C.geod_inverse(PJ.geod,
			C.double(ys[i]), C.double(xs[i]),
			C.double(ys[i+1]), C.double(xs[i+1]), 
			&dist, nil, nil)
		fullDist += float64(dist)
	}

	return fullDist, nil
}
