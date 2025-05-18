# To learn more about how to use Nix to configure your environment
# see: https://firebase.google.com/docs/studio/customize-workspace
{ pkgs, ... }: {
  # Which nixpkgs channel to use.
  channel = "stable-24.05"; # or "unstable"

  # Use https://search.nixos.org/packages to find packages
  packages = [
    pkgs.go_1_22  # Cài đặt Go phiên bản 1.22 (thay bằng phiên bản bạn cần)
    pkgs.gopls    # Công cụ hỗ trợ lập trình Go (tương tự VS Code extension)
    pkgs.gnumake
    pkgs.openssh
    pkgs.rsync
    pkgs.cloudflared
    pkgs.zip
  ];

  # Sets environment variables in the workspace
  env = {
    # Ví dụ: đặt GOROOT hoặc GOPATH nếu cần
    # GOROOT = "${pkgs.go_1_22}/share/go";
  };

  idx = {
    # Search for the extensions you want on https://open-vsx.org/ and use "publisher.id"
    extensions = [
      "golang.go"  # Tiện ích mở rộng Go cho VS Code
    ];

    # Enable previews
    previews = {
      enable = true;
      previews = {
        # Nếu dự án Go là server web, cấu hình preview
        web = {
          command = ["go" "run" "main.go"];
          manager = "web";
          env = {
            PORT = "$PORT"; # PORT do Firebase Studio cung cấp
          };
        };
      };
    };

    # Workspace lifecycle hooks
    workspace = {
      # Runs when a workspace is first created
      onCreate = {
        # Cài đặt dependencies khi workspace được tạo
        go-mod-tidy = "go mod tidy";
      };
      # Runs when the workspace is (re)started
      onStart = {
        # Ví dụ: chạy server Go ở chế độ background
        # run-server = "go run main.go &";
      };
    };
  };
}