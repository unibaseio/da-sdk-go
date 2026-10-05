package archive

// ShadowTar is the name a directory upload's tarball is given; directory
// walks skip it. (The Tar/UnTar helpers that lived here were unused, and
// UnTar wrote entries without containing their paths; see targz.go.)
const ShadowTar = ".config.tar.gz"
