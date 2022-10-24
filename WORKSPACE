load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
http_archive(
    name = "bazel_skylib",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-skylib/releases/download/1.3.0/bazel-skylib-1.3.0.tar.gz",
        "https://github.com/bazelbuild/bazel-skylib/releases/download/1.3.0/bazel-skylib-1.3.0.tar.gz",
    ],
    sha256 = "74d544d96f4a5bb630d465ca8bbcfe231e3594e5aae57e1edbf17a6eb3ca2506",
)
load("@bazel_skylib//:workspace.bzl", "bazel_skylib_workspace")
bazel_skylib_workspace()

# 808 From Mars
# High quality, multi-format pack. We are only using the WAV directory.
new_local_repository(
    name = "a808_from_mars",
    path = "/Users/andreacampi/Music/Samples/808 From Mars/WAV",
    build_file_content = """
package(default_visibility = ["//visibility:public"])
filegroup(
    name = "BD",
    srcs = glob([
        "01. Individual Hits/01. Bass Drum/Clean/Digital/A/*.wav"
    ]),
)
"""
)
