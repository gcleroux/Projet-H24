{
  description = "Projet H24";
  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixos-23.11";
  inputs.flake-utils.url = "github:numtide/flake-utils";

  outputs = inputs@{ self, ... }:
    inputs.flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import inputs.nixpkgs { inherit system; };
        buildDeps = with pkgs; [ git go kubectl k9s kind ];
        xorgLibs = with pkgs; [
          xorg.libX11.dev
          xorg.libXrandr
          xorg.libXcursor
          xorg.libXinerama
          xorg.libXi
          xorg.libXxf86vm
          libGL
        ];
      in {
        devShell = pkgs.mkShell {
          buildInputs = buildDeps ++ xorgLibs;
          hardeningDisable = [ "fortify" ];
        };
      });
}
