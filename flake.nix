{
  description = "Go environment for the Go Exercism track";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
  };

  outputs = { self, nixpkgs }:
    let
      supportedSystems = [ "x86_64-linux" "aarch64-linux" "x86_64-darwin" "aarch64-darwin" ];

      forAllSystems = function:
        nixpkgs.lib.genAttrs supportedSystems (system: function nixpkgs.legacyPackages.${system});
    in
    {
      devShells = forAllSystems (pkgs: {
        default = pkgs.mkShell {
          packages = with pkgs; [
            go
            exercism
            gopls           # The official Go language server
            gotools         # Essential utilities (like goimports)
            golangci-lint   # The paradigmatic linter for Go codebases
            delve           # Go debugger
          ];

          shellHook = ''
            echo "🚀 Exercism Go Workspace Synchronized"
            echo "Using Go: $(go version)"
            echo ""
            echo "If this is your inaugural run, remember to authenticate the CLI:"
            echo "  exercism configure --token=YOUR_API_TOKEN"
            echo ""
          '';
        };
      });
    };
}
