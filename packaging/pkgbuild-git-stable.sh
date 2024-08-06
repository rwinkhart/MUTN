#!/bin/sh

create_pkgbuild_git_stable() {
    printf '\ngenerating PKGBUILD...\n'
    local source="git+https://github.com/rwinkhart/MUTN.git#tag=v\${pkgver}"
    printf "# Maintainer: Randall Winkhart <idgr at tutanota dot com>
pkgname=mutn
pkgver="$version"
pkgrel="$revision"
pkgdesc='A simple, self-hosted, SSH-synchronized password/note manager for the CLI (based on libmutton)'
arch=('x86_64' 'i686' 'i486' 'pentium4' 'aarch64' 'armv7h' 'riscv64')
url='https://github.com/rwinkhart/MUTN'
license=('MIT')
makedepends=(go util-linux gzip)
optdepends=(
    'wl-clipboard: Wayland clipboard support'
    'xclip: X11 clipboard support'
    'bash-completion: Bash completion support'
)
source=(\""$source"\")
sha512sums=(SKIP)

build() {
    cd \${srcdir}/MUTN

    # compress man page
    gzip -kf ./docs/man

    # determine microarchitecture feature level
    case \$CARCH in
        'x86_64')
            lscpuOutput=\$(lscpu | grep Flags)
            if [ ! -z \"\$(echo \$lscpuOutput | grep avx512f)\" ]; then
                export GOAMD64=v4
            elif [ ! -z \"\$(echo \$lscpuOutput | grep avx2)\" ]; then
                export GOAMD64=v3
            elif [ ! -z \"\$(echo \$lscpuOutput | grep sse4_2)\" ]; then
                export GOAMD64=v2
            else
                export GOAMD64=v1
            fi
            ;;
        # TODO check aarch64 feature level
    esac

    # compile binary
    CGO_ENABLED=1 go build -ldflags=\"-s -w\" -trimpath ./mutn.go
}

package() {
    cd \${srcdir}/MUTN
    install -Dm755 ./mutn \${pkgdir}/usr/bin/mutn
    install -Dm644 ./LICENSE \${pkgdir}/usr/share/licenses/mutn/LICENSE
    install -Dm644 ./completions/zsh/_mutn \${pkgdir}/usr/share/zsh/site-functions/_mutn
    install -Dm644 ./completions/bash/mutn \${pkgdir}/usr/share/bash-completion/completions/mutn
    install -Dm644 ./docs/man.gz \${pkgdir}/usr/share/man/man1/mutn.1.gz
}
" > output/PKGBUILD
    printf '\nPKGBUILD generated\n\n'
}
