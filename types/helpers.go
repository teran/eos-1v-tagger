package types

import "time"

// PtrBool ...
func PtrBool(t bool) *bool { return &t }

// PtrInt64 ...
func PtrInt64(t int64) *int64 { return &t }

// PtrFloat64 ...
func PtrFloat64(t float64) *float64 { return &t }

// PtrString ...
func PtrString(t string) *string { return &t }

// PtrTime ...
func PtrTime(t time.Time) *time.Time { return &t }

// PtrFilmAdvanceMode ...
func PtrFilmAdvanceMode(fam FilmAdvanceMode) *FilmAdvanceMode {
	if fam == "" {
		return nil
	}
	return &fam
}

// PtrAFMode ...
func PtrAFMode(t AFMode) *AFMode {
	if t == "" {
		return nil
	}
	return &t
}

// PtrFlashMode ...
func PtrFlashMode(fm FlashMode) *FlashMode {
	if fm == "" {
		return nil
	}
	return &fm
}

// PtrMeteringMode ...
func PtrMeteringMode(mm MeteringMode) *MeteringMode {
	if mm == "" {
		return nil
	}
	return &mm
}

// PtrMultipleExposure ...
func PtrMultipleExposure(me MultipleExposure) *MultipleExposure {
	if me == "" {
		return nil
	}
	return &me
}

// PtrShootingMode ...
func PtrShootingMode(sm ShootingMode) *ShootingMode {
	if sm == "" {
		return nil
	}
	return &sm
}

// PtrAperture ...
func PtrAperture(av float64) *Aperture {
	f := Aperture(av)
	return &f
}
