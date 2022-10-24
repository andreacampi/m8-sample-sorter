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

# TODO 3AM entire phrases

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

# TODO 909 from Mars

# Acoustic Funk Vol 1 from Producer Loops (via HumbleBundle)
# We only use one-shots.
# TODO maybe synths and risers
new_local_repository(
    name = "acoustic_funk_vol_1",
    path = "/Users/andreacampi/Music/Samples/Acoustic Funk Vol 1",
    build_file_content = """
package(default_visibility = ["//visibility:public"])
filegroup(
    name = "BD",
    srcs = glob([
        "One Shots/Drums/AFV1_Drums_Kick*.wav",
    ]),
)
filegroup(
    name = "SD",
    srcs = glob([
        "One Shots/Drums/AFV1_Drums_Snare*.wav",
    ]),
)
filegroup(
    name = "Tom",
    srcs = glob([
        "One Shots/Drums/AFV1_Drums_FTom*.wav",
        "One Shots/Drums/AFV1_Drums_RTom*.wav",
    ]),
)
filegroup(
    name = "Cymbal",
    srcs = glob([
        "One Shots/Drums/AFV1_Drums_Crash*.wav",
        "One Shots/Drums/AFV1_Drums_Ride*.wav",
    ]),
)
filegroup(
    name = "OH",
    srcs = glob([
        "One Shots/Drums/AFV1_Drums_HH6.wav",
        "One Shots/Drums/AFV1_Drums_HH7.wav",
        "One Shots/Drums/AFV1_Drums_HH8.wav",
        "One Shots/Drums/AFV1_Drums_HH9.wav",
    ]),
)
filegroup(
    name = "CH",
    srcs = glob([
        "One Shots/Drums/AFV1_Drums_HH1.wav",
        "One Shots/Drums/AFV1_Drums_HH2.wav",
        "One Shots/Drums/AFV1_Drums_HH3.wav",
        "One Shots/Drums/AFV1_Drums_HH4.wav",
        "One Shots/Drums/AFV1_Drums_HH5.wav",
    ]),
)
"""
)

# TODO Bliss

# TODO Calculate

# TODO Chill Trap & Melodic RnB – good vocals

# TODO CIMMERIUS

# TODO Cinematic

# TODO Dark Realm Horror FX

# TODO Darksynth & Cyberpunk 

# TODO Disco Pimp 

# TODO Epic Cinematic Anthems Kit 

# TODO Essence of Analogue Vol 1 - House

# TODO Essence of Analogue Vol 2 - Cinematic Synthwave

# TODO Essence of Analogue Vol 3 - Ambient Odyssey

# TODO Essential WAV From Mars - 16bit

# TODO FEM sample pack

# Tom Ferry - UK House from Loopmasters
# House, Deep House, UK House, Dance, Pop, EDM
# It also has drum loops which we don't use
# TODO loads of synths and some risers/FX
new_local_repository(
    name = "tom_ferry_uk_house",
    path = "/Users/andreacampi/Music/Samples/Tom Ferry - UK House/Samples",
    build_file_content = """
package(default_visibility = ["//visibility:public"])
filegroup(
    name = "claps",
    srcs = glob([
        "Loops/Clap Loops/**/*.wav",
    ]),
)
filegroup(
    name = "pad_loops",
    srcs = glob([
        "Loops/Music Loops/**/*.wav",
    ]),
)
filegroup(
    name = "BD",
    srcs = glob([
        "One Shots/Kicks/*.wav",
    ]),
)
"""
)

# Vocal stems
# Various vocal stems, bought from misc sites
new_local_repository(
    name = "vocal_stems",
    path = "/Users/andreacampi/Music/Samples/Vocal stems",
    build_file_content = """
package(default_visibility = ["//visibility:public"])
filegroup(
    name = "vocal_stems",
    srcs = glob([
        "**/*.wav",
    ]),
)
"""
)
