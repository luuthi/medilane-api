package utils

import "time"

func TimeUnixMilli() int64 {
	t := time.Now()
	return int64(time.Nanosecond) * t.UnixNano() / int64(time.Millisecond)
}

func TimeUnix() int64 {
	t := time.Now()
	return t.Unix()
}

func TimeUnixNano() int64 {
	t := time.Now()
	return t.UnixNano()
}

func TimeUnixMicro() int64 {
	t := time.Now()
	return int64(time.Nanosecond) * t.UnixNano() / int64(time.Microsecond)
}
