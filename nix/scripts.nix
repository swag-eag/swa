{ pkgs
, config
, swa ? (import ../. { inherit pkgs; })
}: rec {
  start-swa = pkgs.writeShellScriptBin "start-swa" ''
    # rely on environment to provide swad
    export PATH=${pkgs.test-env}/bin:$PATH
    ${../scripts/start-swa} ${config.swa-config} ${config.dotenv} $@
  '';
  start-geth = pkgs.writeShellScriptBin "start-geth" ''
    export PATH=${pkgs.test-env}/bin:${pkgs.go-ethereum}/bin:$PATH
    source ${config.dotenv}
    ${../scripts/start-geth} ${config.geth-genesis} $@
  '';
  start-scripts = pkgs.symlinkJoin {
    name = "start-scripts";
    paths = [ start-swa start-geth ];
  };
}
