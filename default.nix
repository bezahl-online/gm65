# build using 
# nix-build -E 'let pkgs = import <nixpkgs> { }; in pkgs.callPackage ./default.nix {}'
#
{ lib
, buildGoModule
, fetchFromGitHub
, nixosTests
, testers
, installShellFiles
}:
let
  version = "1.0";
  owner = "bezahl-online";
  repo = "gm65";
  rev = "v${version}";
  sha256 = "";
in
buildGoModule {
  pname = "gm65server";
  inherit version;

  src = fetchFromGitHub {
    owner = "bezahl-online";
    repo = repo;
    rev = "53209f6781ac9c0e113ce7fafb58366f07002b57";
    sha256 = "sha256-gHHLmOEy05rBL/10ck67ZgIwRO4x6YvBI5e1aQN4L/E=";
  };
  # src = ../gm65/.;
 
  vendorSha256 = "sha256-GoJ2XMiml03UCi7Ow09pXPbc960n+w1nmhdOAAthoR8=";

  buildPhase = ''
    runHook preBuild
    CGO_ENABLED=0 go build -o gm65server .
    runHook postBuild
  '';

  installPhase = ''
    mkdir -p $out/bin
    mv gm65server $out/bin
    cp localhost.crt localhost.key $out/bin
  '';

  meta = with lib; {
    homepage = "https://github.com/bezahl-online/gm65";
    description = "gm65 server code";
    license = licenses.mit;
    maintainers = with maintainers; [ /* list of maintainers here */ ];
  };
}

