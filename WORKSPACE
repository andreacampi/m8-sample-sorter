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
# XXX we're not using the "02. Kits" directory because some filenames are duplicate.
new_local_repository(
    name = "a808_from_mars",
    path = "/Users/andreacampi/Music/Samples/808 From Mars/WAV",
    build_file_content = """
package(default_visibility = ["//visibility:public"])
filegroup(
    name = "BD",
    srcs = glob([
        "01. Individual Hits/01. Bass Drum/**/*.wav",
    ]),
)
filegroup(
    name = "SD",
    srcs = glob([
        "01. Individual Hits/02. Snare Drum/**/*.wav",
    ]),
)
filegroup(
    name = "Tom",
    srcs = glob([
        "01. Individual Hits/03. Low Tom/**/*.wav",
        "01. Individual Hits/04. Mid Tom/**/*.wav",
        "01. Individual Hits/05. Hi Tom/**/*.wav",
    ]),
)
filegroup(
    name = "Perc",
    srcs = glob([
        "01. Individual Hits/06. Low Conga/**/*.wav",
        "01. Individual Hits/07. Mid Conga/**/*.wav",
        "01. Individual Hits/08. Hi Conga/**/*.wav",
        "01. Individual Hits/08. Hi Conga/**/*.wav",
        "01. Individual Hits/11. Claves/**/*.wav",
        "01. Individual Hits/12. Maracas/**/*.wav",
    ]),
)
filegroup(
    name = "Rim",
    srcs = glob([
        "01. Individual Hits/09. Rim Shot/**/*.wav",
    ]),
)
filegroup(
    name = "Clap",
    srcs = glob([
        "01. Individual Hits/10. Hand Clap/**/*.wav",
    ]),
)
filegroup(
    name = "Cowbell",
    srcs = glob([
        "01. Individual Hits/12. Cowbell/**/*.wav",
    ]),
)
filegroup(
    name = "Cymbal",
    srcs = glob([
        "01. Individual Hits/14. Cymbal/**/*.wav",
    ]),
)
filegroup(
    name = "OH",
    srcs = glob([
        "01. Individual Hits/15. Open HH/**/*.wav",
    ]),
)
filegroup(
    name = "CH",
    srcs = glob([
        "01. Individual Hits/16. Closed HH/**/*.wav",
    ]),
)
"""
)
