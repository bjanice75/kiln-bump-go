# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class Kiln < Formula
  desc ""
  homepage ""
  version "0.94.0"

  on_macos do
    on_intel do
      url "https://github.com/pivotal-cf/kiln/releases/download/v0.94.0/kiln-darwin-amd64-0.94.0.tar.gz"
      sha256 "ff8773f4710529fc490da05a89eb239ede91ffa462f4a410bc3a397996405d49"

      def install
        bin.install "kiln"
      end
    end
    on_arm do
      url "https://github.com/pivotal-cf/kiln/releases/download/v0.94.0/kiln-darwin-arm64-0.94.0.tar.gz"
      sha256 "96dd21ef3c1fde148ef78cec36cc99100e043f616d81f45240217842f43e0f52"

      def install
        bin.install "kiln"
      end
    end
  end

  on_linux do
    on_intel do
      if Hardware::CPU.is_64_bit?
        url "https://github.com/pivotal-cf/kiln/releases/download/v0.94.0/kiln-linux-amd64-0.94.0.tar.gz"
        sha256 "fbe2e5c08f90cd060d50ebe7d426d954d34360a1a6c77f1112db25524cee1ce0"

        def install
          bin.install "kiln"
        end
      end
    end
  end

  test do
    system "#{bin}/kiln --version"
  end
end
