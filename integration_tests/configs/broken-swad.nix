{ pkgs ? import ../../nix { } }:
let swad = (pkgs.callPackage ../../. { });
in
swad.overrideAttrs (oldAttrs: {
  patches = oldAttrs.patches or [ ] ++ [
    ./broken-swad.patch
  ];
})
