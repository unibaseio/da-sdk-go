package archive

// ShadowTar is the name a directory upload's tarball is given; directory
// walks skip it. (The Tar/UnTar helpers that lived here were unused, and
// UnTar wrote entries without containing their paths; uploads use docker's
// archive package.)
const ShadowTar = ".config.tar.gz"
