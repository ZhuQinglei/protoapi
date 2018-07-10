// Code generated by "esc -o generator/data/tpl/tpl.go -pkg=tpl generator/template"; DO NOT EDIT.

package tpl

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type _escLocalFS struct{}

var _escLocal _escLocalFS

type _escStaticFS struct{}

var _escStatic _escStaticFS

type _escDirectory struct {
	fs   http.FileSystem
	name string
}

type _escFile struct {
	compressed string
	size       int64
	modtime    int64
	local      string
	isDir      bool

	once sync.Once
	data []byte
	name string
}

func (_escLocalFS) Open(name string) (http.File, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	return os.Open(f.local)
}

func (_escStaticFS) prepare(name string) (*_escFile, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	var err error
	f.once.Do(func() {
		f.name = path.Base(name)
		if f.size == 0 {
			return
		}
		var gr *gzip.Reader
		b64 := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(f.compressed))
		gr, err = gzip.NewReader(b64)
		if err != nil {
			return
		}
		f.data, err = ioutil.ReadAll(gr)
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fs _escStaticFS) Open(name string) (http.File, error) {
	f, err := fs.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.File()
}

func (dir _escDirectory) Open(name string) (http.File, error) {
	return dir.fs.Open(dir.name + name)
}

func (f *_escFile) File() (http.File, error) {
	type httpFile struct {
		*bytes.Reader
		*_escFile
	}
	return &httpFile{
		Reader:   bytes.NewReader(f.data),
		_escFile: f,
	}, nil
}

func (f *_escFile) Close() error {
	return nil
}

func (f *_escFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}

func (f *_escFile) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *_escFile) Name() string {
	return f.name
}

func (f *_escFile) Size() int64 {
	return f.size
}

func (f *_escFile) Mode() os.FileMode {
	return 0
}

func (f *_escFile) ModTime() time.Time {
	return time.Unix(f.modtime, 0)
}

func (f *_escFile) IsDir() bool {
	return f.isDir
}

func (f *_escFile) Sys() interface{} {
	return f
}

// FS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func FS(useLocal bool) http.FileSystem {
	if useLocal {
		return _escLocal
	}
	return _escStatic
}

// Dir returns a http.Filesystem for the embedded assets on a given prefix dir.
// If useLocal is true, the filesystem's contents are instead used.
func Dir(useLocal bool, name string) http.FileSystem {
	if useLocal {
		return _escDirectory{fs: _escLocal, name: name}
	}
	return _escDirectory{fs: _escStatic, name: name}
}

// FSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func FSByte(useLocal bool, name string) ([]byte, error) {
	if useLocal {
		f, err := _escLocal.Open(name)
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(f)
		_ = f.Close()
		return b, err
	}
	f, err := _escStatic.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.data, nil
}

// FSMustByte is the same as FSByte, but panics if name is not present.
func FSMustByte(useLocal bool, name string) []byte {
	b, err := FSByte(useLocal, name)
	if err != nil {
		panic(err)
	}
	return b
}

// FSString is the string version of FSByte.
func FSString(useLocal bool, name string) (string, error) {
	b, err := FSByte(useLocal, name)
	return string(b), err
}

// FSMustString is the string version of FSMustByte.
func FSMustString(useLocal bool, name string) string {
	return string(FSMustByte(useLocal, name))
}

var _escData = map[string]*_escFile{

	"/generator/template/helper.gots": {
		local:   "generator/template/helper.gots",
		size:    2363,
		modtime: 1530694642,
		compressed: `
H4sIAAAAAAAC/4xWbW/bNhD+7l9xE7ZKil0p2zqgUKamL+u6DkMdxMknwx8Y62SzpkmNopxoTf77cKQk
y29bggC27p57Ht4LT47PzgZA//C2YJqtYcMEMCiN5nJhzRpNpWUJTALKucowa7xglszAnEm4Q2BFgTID
o4AVHCotBnAWD/ChUNpAXsm54aolCDZMJA1J2H6BbwMAAKfWAG+vP39Q60JJlIaCwshiHK4QbI5B/MOr
83jBR+C/9Y+6f37n3Mlx90+v4sUI/O9PeD+44NEJ97kNHh73/vLeBU9PuH9z7pkfXgyeBgPXCXhfcZEB
g9vrv+CubupK9bHtKanCZomAMus1rdICEmu/YyXax1xpayhRb/gce+CGyOE1/l1haUDdfcW5iQCjRWRj
/0Ah1HXjvVeVyKjNFOKgoHIwdYE7wP7AJNBNDuSVEHCrRdvrl1bi08cbSnGFdbxhokIoGNel5cAHti4E
JvRAKVFsCt7SmCKJY6HmTCxVaZLX56/PPQIxvYAUvkm2xgS8e5SLe+TeCCSfr5zBcCa9J8K250rfwFHC
S2JJO5IXxJF2FEcHe4ESNTP4Cc2tFr/evAkq3c34qCl5Ajf7485zCL5z3rAx9W5BpcWFtT0N7IdAQ1Sm
hBSms4uBs1IpA3KtsAYu4ZCOnBvWcLWyY9fxQiujqJHRkpXje3mlVYHa1MEK67BPQn+0HNJGYLrCeral
bI7YsltkmoKkzj8+2lFRObRmv5IZ5lxi5u9ruNyPEtskR0RSbv1xTIK8BKY1q/cPERk1sQUPQic8bcb3
HaFnB+orSG0dh+BPZ37vFICixOPg/RLZ9myYOF4dAkS50h/ZfBlsaAh3Sbt8MmZwx2NTOpnQ71zgYT5W
EVKguM+TcRt6cQDqZB3djr9JnvTbPjpthz0p+udk/CVy887zOtjsyT7tPNnBjoqqXAbNa2IVUhtSH4bd
iyMMtyUN27vRjkaJmjPB/8Hsyq24tCH9qrgM/Be0aLtbtw/u50D7c5gC3eGIywwfxnngX/qu4C9/hEvw
L31IgChheKDbP9f2Jm93/N7ifubabixf2Br7CKBl1YO1K6nDdTuqAzbL9dm7ta+cgmcX/sTZrH9HMgVv
wmqL8Z6zbOM+XdSG/ueSvVLl0S3bO+fW2D/c//zosG0HL/ZcT7uch+BF3nCHid7Y/wYAAP//Ubm8iDsJ
AAA=
`,
	},

	"/generator/template/interface.gots": {
		local:   "generator/template/interface.gots",
		size:    134,
		modtime: 1530694642,
		compressed: `
H4sIAAAAAAAC/6quVihKzEtPVdBzSSxJDKksSC1WqK3lSq0oyC8qUcjMK0ktSktMTlWortbzS8xNra1V
qOZSUFBQQOhzy0zNSQFrUoBIQBVagdTATYXJV1crpOalgHi1XAg2IAAA//9CAdkdhgAAAA==
`,
	},

	"/generator/template/spring_service.gojava": {
		local:   "generator/template/spring_service.gojava",
		size:    709,
		modtime: 1531219236,
		compressed: `
H4sIAAAAAAAC/6yRMW7DMAxFd52CyOQM1QW8BFmCDmmDohegbdYRElOqRCMIBN69sB3EHtqhaTUJ+vr6
T58B6xO2BDnbw7RVLY1xXfBRwMfWphAdtx8RO7r4eLIXqmzluLHI7AXFebY7kj2G4Lgtf2s9+PSw941S
8Jxo65vrA+bPnpJMXhP66uxqwCpJxFqgPmNKQysv2JHqFhNBNgAAOT9BRG4J7J7k6JsEqqOyWfylWI2N
ylF1tZ7UJe14csvM2b72Enp5vwZSnUOH54rNgnOQnnm+6Xh9gxpWJOkjz/bC8bocVTUTwTym/8DbkRR/
A7q3/WPGNwHlfQrEzVC9mq8AAAD//zeMKUrFAgAA
`,
	},

	"/generator/template/spring_struct.gojava": {
		local:   "generator/template/spring_struct.gojava",
		size:    415,
		modtime: 1531219021,
		compressed: `
H4sIAAAAAAAC/4SPwUoDMRCG73mK/1gPzQsUT4IHEemhLzBu4zqaTUMyuyBh3l2m20iLoDnNBP7v+yfT
8EljQGt+v46qO+d4yqci+KCF/Cwc/TNX2TmX59fIA4ZItVrmwYYXmoIqmgOA1rYolMYA/8ghHqvq+T8X
XkgC3jhRtOgTLXT4yudkgzcIzN0hIR2h6tbwqr0Vbmw9JSnzIHsqNKne/VPCnrxz9VdG3P/pt/XSorUb
Jrb9tJ92VzeNQQx7YInG3fRq9kqQuaTfWr1oulvddwAAAP//MpEfWZ8BAAA=
`,
	},

	"/generator/template/tsconfig.gojson": {
		local:   "generator/template/tsconfig.gojson",
		size:    663,
		modtime: 1531207706,
		compressed: `
H4sIAAAAAAAC/2SRT2/UQAzF7/sprLlwoH+gEhxy7qWHCqlIXBBSZycvWdPJeLA9XRDiu6Nkt4HdHqLI
evYb/56vr8ktSRl4vPpuUja/N0REIclUOUM/VWcpFjo6CETBpGnCfayhI9eGi6MwW+3YKGYei9GefUdf
Gt4YbVX2BiVrtYr6i5FHHeGhowD7EM5tUOI2w8hcOTmUuAxQlAQaRKmPHqmqVKgzjKTQ4zz3uK65zIWO
hphtXTIUuZtq5sT+AG+6oJ1R8EDNuIy0x7bG9EQ3b0mUVHJu9YJcKOMZGkeQK0C2i09cxu7liUn6lnHg
unn3/h/aUXmASW5zrnNPkR7/wyt+NFb0C+Vhj+eGy5Sj2eV8Fikoa4b4WaE8oXjMt0ii0UXPmQIm9lW9
h8c5vfOmmLPsP/8qvoNzusUQW/a7ab7YK8PM29DR12NJFHqZVgg6PehSfjwt51iuqsrEhnAUvi3/P5v5
+xsAAP//+QR6tJcCAAA=
`,
	},

	"/generator/template/vue.gots": {
		local:   "generator/template/vue.gots",
		size:    1341,
		modtime: 1531216370,
		compressed: `
H4sIAAAAAAAC/6xUT2/TThC976cYWT+pdpXaPVbOrxWipYCEaAVtL4jDYk8dC3t3mZ3tH5n97mhtJ7WT
CDiQU+btm5k3byapW6OJoYNz3RqtUPEC7hyCh3vSLRw8ODwypA0SPx+VWGiSrOlgKcbEO4ef0GpHBU4y
aIReeJ0AAOg6kqpCSC+dKrjWyno/PqTvlXF882zQe1gE4MrxBBlpqErvxVpdmnVdeiFZBtZl3aD3k5ZQ
oUKSjNfa8i01iw3wFkMML2XeYWOQDoS4c5g6i/FkrmQpxKuNPQKf+uol3kvXMBSNtDbIPQ9fPso2qMUn
RlXa3slh8iyDQivL5ArW1EOTOE5GWvhYZ5DipI+9WGe/lhbhlpo+NlQ/SEb4Ji3eUpODZapVBacQrZhN
nmWNLmSz0pbzk+OT40j80f5BeWwkyTbf2keSwzXptrb4of6O/2/t5stX+AlviDSdQSc2Y4y6b64urnIo
kZHaWiE4i1AhZ0ZbBlkUmsqgnDUQ/nBoeadCfb9Omr309gV0bsB8xTGvapuONi2gny7sc3+LoGn21HXZ
4djIDDe0p9N4XfNW0fwiogVEG5OjZHmYeS/+gVfTAoTsSEEv479wBmmFHA8GJSmvUMUz/pBjjVYW4fRs
coHbNusG00ZX8ZqdLPdSRwFrVlpKliDt9m95N9kvdiAM9/R3qnrq7yUpfBwuNI4+6xZ5FWx8JK2qaE+i
nyETgp//CYlfAQAA///yBF45PQUAAA==
`,
	},

	"/": {
		isDir: true,
		local: "",
	},

	"/generator": {
		isDir: true,
		local: "generator",
	},

	"/generator/template": {
		isDir: true,
		local: "generator/template",
	},
}
