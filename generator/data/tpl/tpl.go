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

	"/generator/template/echo_enum.gogo": {
		local:   "generator/template/echo_enum.gogo",
		size:    326,
		modtime: 1535537130,
		compressed: `
H4sIAAAAAAAC/3yQsWrEMBBEawv0D8ORwobcfcDBtSlDIJDmcCHkjTGx10aSi7Do34Ok2Imb60ajmd3H
LsZ+mZ4gcnkrMkattArfSzZfzUQxYuCQXDuzD6i1AgCRM5zhnnB5GWjsPFK1/Gw9kadN3pL9YcaV/mJn
EHfp2aTpnytb1Hbu/m1u8B7cwH3dwGcB0apiM5HH9YbJLPc925aElOmPEQ+YV5x2fXo+1n8Bq3yVylFY
HSPvvyfSVquo1U8AAAD//5wcsfVGAQAA
`,
	},

	"/generator/template/echo_service.gogo": {
		local:   "generator/template/echo_service.gogo",
		size:    1367,
		modtime: 1535429018,
		compressed: `
H4sIAAAAAAAC/5xUTU/bQBA925L/w9SqKhsZh16DuIBS+iEIClGvsKwnyRaz6+6Oocja/17t2DGhJapU
Tubtm5m3s++lEfJerBG6rrzqP71P4iRWD42xBFkSR6lGmmyImjQcROla0aa9K6V5mNTizpGQ9xOUG5Mm
cR4Yk0nodike0HtQDmiDoDShXQmJII0mobQDUdd8FABr6hqtS2J6bnC3eizrkjjqukOwQq8Rygukjakc
BLEAEEqWimr0PjvouvKLblpaPjfofQ6MzFsaoQIOZtYam29rDwF1xb347ixiga4x2iE4sq0kFjBio67O
w+0PZ/Q0tcNZepvEEbePhjHAfwMNAxI4PGnPjZJ41WoJN+Mibj4LXdVoM2cfoeveD3AOYe/lcPgp1ASZ
Fqm1GkKPTPaUM6MJf1EOGVoLLCJnbqQ0TE9A41P2x974sSO1Cmw4AVmeKl1lSufHjLw7Aa3qvkcULh/a
jAvq4X4PU7jEJ/7KPh4dHRWhPC+Y4Idq1ivLr9fzyyw4rbwmQa07FdUCf7boqIAwIg903wszLRX9RRbD
cGcfyx0fKM30PdK2/04hNGJoEBvBS9tiZ+BelfNvL+p6C+06ivOwwLVyhPZVLlqHFZCBO6UrsKYlTgC/
/F/8DOGAH3ImN6aA3gajC7oXI+/NxyGoFWiEct6QMrr8LuoWIb2aXy/TkYbl+WyZpfxjQBvv02KPCfM3
svOPSeezV4PC4P+d9GZqJ5PRZ3BmURC6/iFhG81huaMbpakQ+PPMVMhu2kZjSO5Omj4w0gXmFCTzL9A5
scZpKCr7lrk/Zjm/AwAA//9AECDDVwUAAA==
`,
	},

	"/generator/template/echo_struct.gogo": {
		local:   "generator/template/echo_struct.gogo",
		size:    198,
		modtime: 1535097442,
		compressed: `
H4sIAAAAAAAC/1zNvQ7CMAwE4D1S3uHUnXRnRWJEDDxAo9ZUhTYNSTpEVt4duRn42U6fdWdv+6cdCczm
WmMpWmnVtkKn2cZ4scuOKXv6Q8QUtj6BtQIA5gOCdSPBnCeahwjp1Yu5TWmWisTsJXWPuLpjw2zqWoP7
GpYfeG0U8rd0n0/khn2/aPUOAAD//5Sw0ZbGAAAA
`,
	},

	"/generator/template/spring_service.gojava": {
		local:   "generator/template/spring_service.gojava",
		size:    841,
		modtime: 1535097442,
		compressed: `
H4sIAAAAAAAC/6ySzWrrMBCF9wa/w+CVs7h6AW9C4BK6SBPa0P3YnjgiyUiVRoQg9O7F+bMXLaUhXhmd
OTPn08his8OOIEa1uvymVOVZnumDNU7AuE556zR3G4cHOhq3U0eqVa25VchsBEUbVnOSBVqruav+7F0Z
/7j5jbw17Glm2tMj7s9AXq7mPLOh3usGsPbisBFo9uh9fzmveKCUZugJYp4BAMT4DxxyR6AWJFvTekhp
kPQGmEAt7XnOB+4DQTH/vy7uVdMRd1mcFyDblIrJVR6TXY6u6WJUyyA2yPpkKaUhXt+wnI6YeumFh0rN
k1v8/nMkwfHgLzVPqos8IiFuz5l/YVst38dww4N4DtucpHw+zX3RP879Zmh163ZvlvLsKwAA///TJIC2
SQMAAA==
`,
	},

	"/generator/template/spring_struct.gojava": {
		local:   "generator/template/spring_struct.gojava",
		size:    565,
		modtime: 1535097442,
		compressed: `
H4sIAAAAAAAC/5SPQWrzMBCF9wbfYZb5F9EFwg+FQBehlCxygakySZXIkhiNTYPQ3cvYjXGgEKrVjIb3
3vcS2iueCUox+2msddM2beO6FFnAxs6cMAvxV+fNBe01x2AwhCgoLgazyzFsmVAib/6m2nNMxHJbxl1w
QNOL8+bNZRkvqf/wzoL1mLNibnV4x45qhdI2AAClrIExnAnMqyN/zLVOh8RuQCE4uYBexTsc8HBLo7aA
URsYC99tKBxB1dPPy6Ldj+NE88ix0jUG4d7KHhm7Wv89ZdMnny6bBQj8f4Kl+0xXyoMxrOfaM+Si75lE
vQ9OvJqvZkJ9TNJz+CW83qNmhNo23wEAAP//c5VntTUCAAA=
`,
	},

	"/generator/template/ts/data.gots": {
		local:   "generator/template/ts/data.gots",
		size:    604,
		modtime: 1535538479,
		compressed: `
H4sIAAAAAAAC/5yOQWsaQRiG7wv7Hz68CEvVe6EHi1soSJEivZRSxt1PHVhntzOzpTIMFGpJhBiEaA4e
klNAENRTIDE/x931Z4R12RgPuTinGd73feapWJZpWNDsUgFt6iFQAR1kyIlEF1p9KAbclz4JaPG45/hM
EsoEEM8D2UVwiSQgJA8dGXKEFlLWgVCgC5TtCwesFCCQ/6YOihRaOuWkw93qLr4+227uk8ltfD7ePl7m
smmaRdHF/2i8qDY+R8NRsljt1v+SyTwe/o2epslknswG8XQdj5bJ5iq+GUTL2fZhaBpWxTSUKgEnrINQ
tlnYE6C1aeCfwOcSkIU9UKr8hfRQa1CmAQCgVD74RNFzswVkSV79kN6/ES9Erd+9zJC5+7LOvs2frxRq
RJJmP8AjDcok8jZx8CSX92lJipR64GutFG0D/oJynbTQg0K9+tGu//xqN+xq064VtP7+Qylkbg58W/85
AAD//zpvH6BcAgAA
`,
	},

	"/generator/template/ts/helper.gots": {
		local:   "generator/template/ts/helper.gots",
		size:    2447,
		modtime: 1535097442,
		compressed: `
H4sIAAAAAAAC/5RWbW/bNhD+bsD/4SZslVS7UrZ1QKFMTV/WbR2GuouTT4Y/MNLJZk2TGkUp0dr894Gk
JEuxhaYBAoh3zz33yqPDp0+nEzD/8ConkuyhIgwIFEpSvrFyiaqUvADCAXkiUkwbNagtUZAQDjcIJM+R
p6AEkJxCKZk2DqcTvMuFVJCVPFFUtBReRVjU0PjtB3yeTgAArMMGeX35/q3Y54IjV9rKDyzIAnNGEvTC
H56fhRs6B/eVe1r/82urj0b0Pz0PN3Nwvx9Tv7Xm8zH9mTGfjah/eWPNV2P636x+7frn08n9dDKdNJ2B
NyVlKRC4vvwbbuqmyrpWpluFrrfaIiBP+00sJYPIKG5IgeaYCWkEBcqKJthHN1TWQOK/JRYKxM0nTFQA
GGwCY/wnMiYuG+2tKFmq265NLBREBqrOcQAcTFAE3ShBVjIG15K1rX9mfPzx7kpnucM6rAgrEXJCZWFJ
8I7sc4aROem0tHUMzlapPApDJhLCtqJQ0YuzF2eOQRG5gRg+c7LHCJxb5JtbpM4cOE12VqAo4c69Abex
xS/hJOWFpok7lieaJO44RqZ9gxwlUfhPibK+luzXq5deKbvRnze1j+Dq6BbQDLzvrNpvZb3bUUp2boVm
XvQHQ6X5VAExrNbnWmwVurSe1u6wBsrhBKvWVqSlbP0v7BDkUiihextsSbG45R+lyFGq2tth7Q9o9J/e
IHHjY7XDet0j7WJtPRhwHAPX8/Dli5kgkUErdkueYkY5pu6RH1uIMXKT7lwTFT1IGGqvtAAiJamPQgmU
WJomeL51v2pm+7WGr49j2EFsijoDd7V2+7EAsgJH4Ef1Mh2rCBstlcYEmZDvSLL1Kj2jD5i7xFKicKgy
uY1m9jtleCIx4xRi0Ibvl4vW9vwY1Xm2hENAUwYdQttZ696Cx/3+tVx8COyFoFntVQ893w+PZuyDvCy2
XvPA7HzdlNiFWffk+H6vuv7h9rTjUqCkhNH/MP1oN2Lc8H4SlHvuE7OdD5fzIXyQi165sxj0ZQ8oT/Fu
kXnuhWuL/+xHuAD3woUINCvMjlwPgztc+eHr8HDlP3bhN6IPZI99COgN18e1a6wDdnvtgGzW8jds5b73
GBzzXCytzAIGbmNwlqQ2IOdRezrsEwat7Vf286nV3IvzIOzH9rUfMGYGwAkd2+Au6Rk4gTMbUNln//8A
AAD//8c9h/ePCQAA
`,
	},

	"/generator/template/ts/ts_service.gots": {
		local:   "generator/template/ts/ts_service.gots",
		size:    1592,
		modtime: 1535537075,
		compressed: `
H4sIAAAAAAAC/4xUX28bxRd9t+TvcGX9pPyRs5v8JKrgKEgRApQXiETgfbJ7bU9Zzywzs02j1UoNwS2V
nCZSowpIoEFQFNEkBglK7Db0y+yunad+BTQ7u45d/oh52plz7r3n3Htte3a2XJqF9SaVUKceApXQQIaC
KHRhYwumfMEVJz6dMjw0NIczRSiTUBecKWQurKytgsNdBNUkCja5+AQ2qWqCaiJ4dEMQsQVT5Dblcqqq
HwXWucAqUDUlQeCnARXomuCcprVQJhXxPHSBsiyVL/hNdFQh5lprVptK2BRUKWSav77l44eOoP6InpF8
wW9RFyUQ2CCSOhBI0kCoc2FMEM8DwlxokS1giC4Q92YgVQuZAuI4XLiUNUBxkD46tE6dQlRhwzCZC5Kq
gCjKma4/99+Ppg+7T9JH9+LnzwYHj9Mv9uP+g2ISGjVQ0mkn+0+T+7uDp93hzzuDg5OVtdXB15/Hz78f
HG+/etFJes/iy5eDg5PB6Wl8cT991Ev6D7PuvnrRgeTwOD374erozvDH7fjlN8PudgYlZ18mRydx/8HV
d73B4Xl8cTZW8W7bJDdpJzs87D4xYnPu3k/p3n78x6ERuLK2ajQmj/vp0em1xvNvr75qpzvt5O5vyV53
uHNpFKXHvXT3PGn/Hl8+NErMt4Z++Sy+2E32O8M7nfjyKLnXTw9/TQ965dKsXS7Rls+FgsxLFUJY0R9r
greoRIj0xraKFVsascNyCQAgDAVhDQTr3YA5enIyigpkDqxV5gdKm4a5KKoWADJXs4rUlu0SRcZzj9b0
I+HBNa2Jno9CE8sle3aDyAxP9rrGjfbicCYVFNAyVJpK+TXbXnjz/9bCjUVrYeEN68Z8bXF+cb6yVC6F
4f8cj0j5Pmkh1JbBent00wq1h7/zh7cznfX8DcLQ0jFRNO0TQVqypl9G3qNopjbZ1Lx5HioIhFcDqYT+
iSyP+57OTVQr4yKjqFKFyqheZWbJpLJtUCiVuQhUgWBmoJb2QOvAEKwPfK3W+ph4AULlvXfWKxBFPpcq
DNGTGEUNVPl0pgPhVSEM/zUy81rEhmDuEOUpZowYfSzVRDYtUMLyW4X54tg2NAlzPQQUgotJkNZBh1l6
QawMn3k9vji56bzFlkD9D/N68GRo9E9KZOA4KCXoyEnOX6pI7t1CXWYsd6Snkq1Pvunl0p8BAAD//wIU
nek4BgAA
`,
	},

	"/generator/template/ts/ts_service.govue": {
		local:   "generator/template/ts/ts_service.govue",
		size:    2245,
		modtime: 1535097442,
		compressed: `
H4sIAAAAAAAC/4xVXW8bRRR9t+T/cGVVsh0567wgVS5FlFJQJEQimuZ9snvtnXY9s8zMJo3MSg0hLZWS
JlKjCkigQVAU0XyABCVJG/pndtfOU/8CGo/XXrtpyzzZd869c865M3erY2P53BjMuFRCnXoIVEIDGQqi
0IG5RSj6gitOfFo0ODQwmzNFKJNQF5wpZA5cmZ4EmzsIyiUKFri4BQtUuaBcBI/OCSIWoTgb4BcoeSBs
LFb0lsA6F1gBqooSBH4ZUIGOKTEE1rwok4p4HjpAWbesL/hNtFVKbMC7y4NKWBBUKWQaP7Po43VbUL8P
74J8weepgxIIzBFJbQgkaSDUuTCCiOcBYQ40ySIwRAeIczOQqolMAbFtLhzKGqA4SB9tWqd2SioVY5DM
AUlVQBTlTJ8//v+XhncOnySP7kXPn7U3HyffbkQnD9Ku6F2zFa+uxBtP4/tr7aeHnT+W25u7V6Yn2z98
Ez3/pb2z9OrFanz8LDp92d7cbe/tRUf3k0fH8cnDjMevXqxCvLWT7P96tn2n89tS9PLHzuFSBhDvfxdv
70YnD85+Pm5vHURH+5nT766Yg8wRw253Dp8Y4j3s+u/J+kb075Yhe2V60vCNH58k23sDvgc/nX2/kiyv
xHf/jtcPO8unhleyc5ysHcQr/0SnDw0T81tv/fl1dLQWb6x27qxGp9vxvZNk669k8zifG6vmc7Tpc6Gg
BVd50+cMmarAbIAQ6kvchOJ8gOO+4D4KtTjuoM0FUVwUL/UzM15kUkR6RwfAVj4HANBqCcIaCNYnAbN1
62UYpjvWJPMDpX0KQ6jowFSgMpEUh8zRSSlHq+oQRbJH9a/9DeHBAOai56Mo5nP53GyAViCxlKFfvqQ3
Puwbkc/h7W45B+sk8BTYHpFSs7qqf3xOmpoU3tYvXXZN60msVvUokEoEtuLCxDKBUjkF6iUDH0WpbAKh
ZtAr8RGRCDeEZwK+oPNEoX6QWlQNpBL6lV2GgquUX6tWPW4Tz+VS1S5OXJwopF5dsFOyULsMA+owsP0t
DTEiSz4RpClrIy0q12Ba8CaV+Bm9he+PtusruCYEFx9k1eo1/M9DBcGQoEzvSsql0upprhSyasKwUIFC
n2JBd+9NZ1SrMDP18VQNHFQompQhBFLPRlX1uRwZWnpGoVTDFQSqQOgBS6V1QRtutVrjQOvAEKwpX9tm
zRIvQCh8em2mAGGoC7da6EkMwwaq3qUtBcKraGffkdz1O01vmb9hr0bZUi6y0jBBQ1LC5dfsTpe+gdxD
y+ONkkA57NeQV6jbBi5hjkdZ43wYresiln54VhdffsOxGfcYLpgrMZp6fmZ4frhXrF9CoPQ5kwhEjk6M
cySGlddjRq92Tqvn+gtsYhJkYLu68HsTExVgHBgq/RV/t8NG2KW3Shj4UbjOm6hcfQMXBGeNwnmZI35k
IeHoZMzn/gsAAP//gM8Tf8UIAAA=
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

	"/generator/template/ts": {
		isDir: true,
		local: "generator/template/ts",
	},
}
