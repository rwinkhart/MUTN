#!/bin/sh

create_apkbuild_source_stable_optimized() {
    printf '\nGenerating APKBUILD...\n'
    local source="https://github.com/rwinkhart/MUTN/archive/refs/tags/v"$version".tar.gz"
    local checksum=$(curl -sL "$source" | sha512sum | cut -d ' ' -f1)
    printf "# Maintainer: Randall Winkhart <idgr@tutanota.com>
pkgname=mutn
pkgver="$version"
pkgrel="$((revision-1))"
pkgdesc='A simple, self-hosted, SSH-synchronized password/note manager for the CLI (based on libmutton)'
arch='all'
url='https://github.com/rwinkhart/MUTN'
license='MIT'
depends=''
makedepends='busybox go'
options='!check net'
source=\"\$pkgname-\$pkgver.tar.gz::"$source"\"

build() {
    cd \${srcdir}/MUTN-\$pkgver

    # compress man page
    gzip -kf ./docs/man

    # determine microarchitecture feature level
    case \$CARCH in
        'x86_64')
            cpuFlags=\$(grep -E 'flags\s+:\s' /proc/cpuinfo)
            if [ ! -z \"\$(echo \"\$cpuFlags\" | grep 'avx512f')\" ]; then
                export GOAMD64=v4
            elif [ ! -z \"\$(echo \"\$cpuFlags\" | grep 'avx2')\" ]; then
                export GOAMD64=v3
            elif [ ! -z \"\$(echo \"\$cpuFlags\" | grep 'sse4_2')\" ]; then
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
    cd \${srcdir}/MUTN-\$pkgver
    install -Dm755 ./mutn \${pkgdir}/usr/bin/mutn
    install -Dm644 ./LICENSE \${pkgdir}/usr/share/licenses/mutn/LICENSE
    install -Dm644 ./completions/zsh/_mutn \${pkgdir}/usr/share/zsh/site-functions/_mutn
    install -Dm644 ./completions/bash/mutn \${pkgdir}/usr/share/bash-completion/completions/mutn
    install -Dm644 ./docs/man.gz \${pkgdir}/usr/share/man/man1/mutn.1.gz
}

sha512sums=\"
$checksum  \$pkgname-\$pkgver.tar.gz
\"
" > 1output/APKBUILD
    printf '\nAPKBUILD generated\n\n'
}
