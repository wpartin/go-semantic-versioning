# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class GoSemanticVersioning < Formula
  desc "Tag your repository with Go!"
  homepage "https://github.com/wpartin/go-semantic-versioning"
  version "0.4.21"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/wpartin/go-semantic-versioning/releases/download/v0.4.21/go-semantic-versioning_Darwin_arm64.tar.gz"
      sha256 "1f0cdb0d50b6b14db914c7816b2aeffaf618c56b5ca5327d33e81a3d551c5c3c"

      def install
        bin.install "go-semantic-versioning"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/wpartin/go-semantic-versioning/releases/download/v0.4.21/go-semantic-versioning_Darwin_x86_64.tar.gz"
      sha256 "170c7e19ad865d4ca4dad85fbf1f95aec2a9e16b8b8ba6a502ef19c4bbd3873a"

      def install
        bin.install "go-semantic-versioning"
      end
    end
  end

  on_linux do
    if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
      url "https://github.com/wpartin/go-semantic-versioning/releases/download/v0.4.21/go-semantic-versioning_Linux_arm64.tar.gz"
      sha256 "deaf1889018da66ef1fb3b2bb7be5e24cebd2a299ce32d6346f720b3445332cb"

      def install
        bin.install "go-semantic-versioning"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/wpartin/go-semantic-versioning/releases/download/v0.4.21/go-semantic-versioning_Linux_x86_64.tar.gz"
      sha256 "e2091ff3ec74b251e7dff9aa6a26722ef918026f17f0418a4cd09cb50f85a1ea"

      def install
        bin.install "go-semantic-versioning"
      end
    end
  end
end
