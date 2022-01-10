package portego

type Pkgbuild struct {
	pkgname			[]string
	pkgver			string
	pkgrel			uint
	epoch			uint
	pkgdesc			string
	url				string
	license			[]string
	install			string
	changelog		string
	source			[]string
	validgpgkeys	[]string
	noextract		[]string
	cksums			[]string
	md5sums			[]string
	sha1sums		[]string
	sha224sums		[]string
	sha256sums		[]string
	sha384sums		[]string
	sha512sums		[]string
	b2sums			[]string
	groups			[]string
	arch			[]string
	backup			[]string
	depends			[]string
	makedepends		[]string
	checkdepends	[]string
	optdepends		[]string
	conflicts		[]string
	provides		[]string
	replaces		[]string
	
	options struct {
		strip		string
		docs		string
		libtool		string
		staticlibs	string
		emptydirs	string
		zipman		string
		ccache		string
		distcc		string
		buildflags	string
		makeflags	string
		debug		string
		lto			string
	}
	
	packageSteps	[]string
	prepareSteps	[]string
	buildSteps		[]string
	checkSteps		[]string
	
	srcdir			string
	pkgdir			string
	startdir		string
	
	pkgbase			string
	
	preInstall		string
	postInstall		string
	preUpgrade		string
	postUpgrade 	string
	preRemove		string
	postRemove		string
	
	vcsdirectory	string
	vcsurl			string
	vcsfragment		string
	vcsquery		string
}

// Set Pkgname
func (p *Pkgbuild) AddPkgname(val string) {
    p.pkgname = append(p.pkgname, val)
}

// Get Pkgname
func (p Pkgbuild) Pkgname() []string {
    return p.pkgname
}

// Set Pkgver
func (p *Pkgbuild) SetPkgver(val string) {
	p.pkgver = val
}

// Get Pkgver
func (p Pkgbuild) Pkgver() string {
    return p.pkgver
}

// Set Pkgrel
func (p *Pkgbuild) SetPkgrel(val uint) {
	p.pkgrel = val
}

// Get Pkgrel
func (p Pkgbuild) Pkgrel() uint {
    return p.pkgrel
}

// Set Epoch
func (p *Pkgbuild) SetEpoch(val uint) {
	p.epoch = val
}

// Get Epoch
func (p Pkgbuild) Epoch() uint {
    return p.epoch
}

// Set Pkgdesc
func (p *Pkgbuild) SetPkgdesc(val string) {
	p.pkgdesc = val
}

// Get Pkgdesc
func (p Pkgbuild) Pkgdesc() string {
    return p.pkgdesc
}

// Set Url
func (p *Pkgbuild) SetUrl(val string) {
	p.url = val
}

// Get Pkgver
func (p Pkgbuild) Url() string {
    return p.url
}
