# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class Kiln < Formula
  desc ""
  homepage ""
  version "0.70.0"

  on_macos do
    if Hardware::CPU.intel?
      url "https://github.com/pivotal-cf/kiln/releases/download/v0.70.0/kiln-darwin-amd64-0.70.0.tar.gz"
      sha256 "7fae97a124411f5df59e76adb88383c620155b02836b79c8be7187cbaa51f7f4"

      def install
        bin.install "kiln"
      end
    end
    if Hardware::CPU.arm?
      url "https://github.com/pivotal-cf/kiln/releases/download/v0.70.0/kiln-darwin-arm64-0.70.0.tar.gz"
      sha256 "2dd1dce58ce5402990b568e3fc1b5da9491b41b3edc82ccc74ddd980505ca57a"

      def install
        bin.install "kiln"
      end
    end
  end

  on_linux do
    if Hardware::CPU.intel?
      url "https://github.com/pivotal-cf/kiln/releases/download/v0.70.0/kiln-linux-amd64-0.70.0.tar.gz"
      sha256 "d304bf98a5f208dfad141a41875317e46fffcf6bcc9482f511937a14b84e32ef"

      def install
        bin.install "kiln"
      end
    end
  end

  test do
    system "#{bin}/kiln --version"
  end
end
