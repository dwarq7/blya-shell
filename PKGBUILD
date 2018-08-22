pkgname=blya-shell
pkgver=1.0
pkgrel=0
pkgdesc="Bilyk Ilya Shell written in Go"
arch=('x86_64' 'armv7h' 'aarch64')
url="https://github.com/dwarq7/$pkgname"
license=('MIT')
makedepends=('git' 'go')
backup=()
options=("!strip")
source=("git://github.com/dwarq7/$pkgname.git")
sha256sums=('SKIP')

build() {
  cd "$pkgname"

  go build -o blya main.go
}

package() {
  cd "$pkgname"

  install -Dm755 "blya" "$pkgdir/usr/bin/$pkgname"
}