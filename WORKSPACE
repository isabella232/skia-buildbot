workspace(
    name = "skia_infra",

    # Must be kept in sync with the npm_install rules defined below invoked below.
    managed_directories = {
        "@npm": ["node_modules"],
    },
)

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

# Read the instructions in //bazel/rbe/README.md before updating this repository.
#
# We load bazel-toolchains here, rather than closer where it's first used (RBE container toolchain),
# because the grpc_deps() macro (invoked below) will pull an old version of bazel-toolchains if it's
# not already defined.
http_archive(
    name = "bazel_toolchains",
    sha256 = "179ec02f809e86abf56356d8898c8bd74069f1bd7c56044050c2cd3d79d0e024",
    strip_prefix = "bazel-toolchains-4.1.0",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-toolchains/releases/download/4.1.0/bazel-toolchains-4.1.0.tar.gz",
        "https://github.com/bazelbuild/bazel-toolchains/releases/download/4.1.0/bazel-toolchains-4.1.0.tar.gz",
    ],
)

###############
# Buildifier. #
###############

http_archive(
    name = "com_github_bazelbuild_buildtools",
    sha256 = "2adaafee16c53b80adff742b88bc90b2a5e99bf6889a5d82f22ef66655dc467b",
    strip_prefix = "buildtools-4.0.0",
    url = "https://github.com/bazelbuild/buildtools/archive/4.0.0.zip",
)

##############################
# Go rules and dependencies. #
##############################

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "69de5c704a05ff37862f7e0f5534d4f479418afc21806c887db544a316f3cb6b",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.27.0/rules_go-v0.27.0.tar.gz",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.27.0/rules_go-v0.27.0.tar.gz",
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "62ca106be173579c0a167deb23358fdfe71ffa1e4cfdddf5582af26520f1c66f",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.23.0/bazel-gazelle-v0.23.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.23.0/bazel-gazelle-v0.23.0.tar.gz",
    ],
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("//:go_repositories.bzl", "go_repositories")

# gazelle:repository_macro go_repositories.bzl%go_repositories
go_repositories()

go_rules_dependencies()

# Gazelle fails for toolchain versions < 1.14 with an error like the following:
#
#     gazelle: [...]: go: updates to go.mod needed, disabled by -mod=readonly
go_register_toolchains(version = "1.16.3")

gazelle_dependencies()

##########################
# Other Go dependencies. #
##########################

# Needed by @com_github_bazelbuild_remote_apis.
load("@com_github_bazelbuild_remote_apis//:repository_rules.bzl", "switched_rules_by_language")

switched_rules_by_language(
    name = "bazel_remote_apis_imports",
    go = True,
)

# Needed by @com_github_bazelbuild_remote_apis.
http_archive(
    name = "com_google_protobuf",
    sha256 = "9748c0d90e54ea09e5e75fb7fac16edce15d2028d4356f32211cfa3c0e956564",
    strip_prefix = "protobuf-3.11.4",
    urls = ["https://github.com/protocolbuffers/protobuf/archive/v3.11.4.zip"],
)

load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")

protobuf_deps()

# Needed by @com_github_bazelbuild_remote_apis for the googleapis protos.
http_archive(
    name = "googleapis",
    build_file = "BUILD.googleapis",
    sha256 = "7b6ea252f0b8fb5cd722f45feb83e115b689909bbb6a393a873b6cbad4ceae1d",
    strip_prefix = "googleapis-143084a2624b6591ee1f9d23e7f5241856642f4d",
    urls = ["https://github.com/googleapis/googleapis/archive/143084a2624b6591ee1f9d23e7f5241856642f4d.zip"],
)

# Needed by @com_github_bazelbuild_remote_apis for gRPC.
http_archive(
    name = "com_github_grpc_grpc",
    sha256 = "b391a327429279f6f29b9ae7e5317cd80d5e9d49cc100e6d682221af73d984a6",
    strip_prefix = "grpc-93e8830070e9afcbaa992c75817009ee3f4b63a0",  # v1.24.3 with fixes
    urls = ["https://github.com/grpc/grpc/archive/93e8830070e9afcbaa992c75817009ee3f4b63a0.zip"],
)

load("@com_github_grpc_grpc//bazel:grpc_deps.bzl", "grpc_deps")

grpc_deps()

###################################################
# JavaScript / TypeScript rules and dependencies. #
###################################################

http_archive(
    name = "build_bazel_rules_nodejs",
    sha256 = "8f5f192ba02319254aaf2cdcca00ec12eaafeb979a80a1e946773c520ae0a2c9",
    urls = ["https://github.com/bazelbuild/rules_nodejs/releases/download/3.7.0/rules_nodejs-3.7.0.tar.gz"],
)

# The npm_install rule runs anytime the package.json or package-lock.json file changes. It also
# extracts any Bazel rules distributed in an npm package.
#
# There must be one npm_install rule for each package.json file in this repository. Any node_modules
# directories managed by npm_install rules must be mentioned in the workspace() rule at the top of
# this file.
load("@build_bazel_rules_nodejs//:index.bzl", "npm_install")

# Manages the node_modules directory.
npm_install(
    name = "npm",
    package_json = "//:package.json",
    package_lock_json = "//:package-lock.json",
)

################################
# Sass rules and dependencies. #
################################

http_archive(
    name = "io_bazel_rules_sass",
    sha256 = "9ad74e6e75a86939f4349b31d43bb1db4279e4f2a139c5ebaf56cf99feea1faa",
    strip_prefix = "rules_sass-1.32.8",
    url = "https://github.com/bazelbuild/rules_sass/archive/1.32.8.zip",
)

load("@io_bazel_rules_sass//:package.bzl", "rules_sass_dependencies")

rules_sass_dependencies()

load("@io_bazel_rules_sass//:defs.bzl", "sass_repositories")

sass_repositories()

##################################
# Docker rules and dependencies. #
##################################

http_archive(
    name = "io_bazel_rules_docker",
    sha256 = "59d5b42ac315e7eadffa944e86e90c2990110a1c8075f1cd145f487e999d22b3",
    strip_prefix = "rules_docker-0.17.0",
    urls = ["https://github.com/bazelbuild/rules_docker/releases/download/v0.17.0/rules_docker-v0.17.0.tar.gz"],
)

load(
    "@io_bazel_rules_docker//repositories:repositories.bzl",
    container_repositories = "repositories",
)

container_repositories()

# This is required by the toolchain_container rule.
load(
    "@io_bazel_rules_docker//repositories:go_repositories.bzl",
    container_go_deps = "go_deps",
)

container_go_deps()

load(
    "@io_bazel_rules_docker//container:container.bzl",
    "container_pull",
)

# Provides the pkg_tar rule, needed by the skia_app_container macro.
#
# See https://github.com/bazelbuild/rules_pkg/tree/main/pkg.
http_archive(
    name = "rules_pkg",
    sha256 = "038f1caa773a7e35b3663865ffb003169c6a71dc995e39bf4815792f385d837d",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_pkg/releases/download/0.4.0/rules_pkg-0.4.0.tar.gz",
        "https://github.com/bazelbuild/rules_pkg/releases/download/0.4.0/rules_pkg-0.4.0.tar.gz",
    ],
)

load("@rules_pkg//:deps.bzl", "rules_pkg_dependencies")

rules_pkg_dependencies()

##################
# Miscellaneous. #
##################

load("@bazel_toolchains//rules/exec_properties:exec_properties.bzl", "rbe_exec_properties")

# Defines a local repository named "exec_properties" which defines constants such as NETWORK_ON.
# These constants are used by the //:rbe_custom_platform build rule.
#
# See https://github.com/bazelbuild/bazel-toolchains/tree/master/rules/exec_properties.
rbe_exec_properties(
    name = "exec_properties",
)

######################
# Docker containers. #
######################

# Pulls the gcr.io/google/rbe-ubuntu16-04 container, used as the base container for our custom RBE
# toolchain container.
container_pull(
    name = "google_debian10",
    digest = "sha256:96a0145e8bb84d6886abfb9f6a955d9ab3f8b1876b8f7572273598c86e902983",
    registry = "gcr.io",
    repository = "cloud-marketplace/google/debian10",
)

# Pulls the gcr.io/skia-public/skia-wasm-release container with the Skia WASM build.
container_pull(
    name = "container_pull_skia_wasm",
    registry = "gcr.io",
    repository = "skia-public/skia-wasm-release",
    # The container_pull documentation[1] recommends specifying a digest (via the "digest" argument)
    # for reproducible builds.
    #
    # We specify the "prod" tag here instead of a digest for simplicity, but this might cause Bazel
    # to fetch the "prod" image once, cache it, and use it for all subsequent builds, completely
    # ignoring any new images uploaded to GCR with the "prod" tag.
    #
    # This caching problem could be solved by replacing the "tag" argument with a "digest" argument
    # with the latest digest. But this requires setting up an autoroller to update the digest every
    # time a new container image is uploaded to GCR.
    tag = "prod",
)

# Pulls the gcr.io/skia-public/basealpine container, needed by the skia_app_container macro.
container_pull(
    name = "basealpine",
    digest = "sha256:d58f6c8ced7d5436ce87ed1904340f1d8c6775eb8d891ac768159241a20c1eb4",
    registry = "gcr.io",
    repository = "skia-public/basealpine",
)

# Pulls the gcr.io/skia-public/base-cipd container, needed by some apps that use the
# skia_app_container macro.
container_pull(
    name = "base-cipd",
    digest = "sha256:0ae30b768fb1bdcbea5b6721075b758806c4076a74a8a99a67ff3632df87cf5a",
    registry = "gcr.io",
    repository = "skia-public/base-cipd",
)

##############################
# Packages for RBE container #
##############################
# The following http_archives are used to download and verify files that will be installed on
# our RBE container.
http_archive(
    name = "go_sdk_external",
    # We are downloading a tar file of pre-compiled executables, libraries and such that does
    # NOT have a BUILD file for Bazel to read. As such, we can specify one here using
    # build_file_content that makes all the contents of the the Golang SDK tar file available
    # as a target called @go_sdk_external//:extracted_files
    #
    # Debugging tip: make an intentional typo in build_file_content and then try to `bazel build`
    # something that depends on this. The error message will show where the archive is being
    # downloaded/extracted in your bazel cache, which can help manual inspection.
    build_file_content = """
filegroup(
    name = "extracted_files",
    srcs = glob(["go/*"], exclude_directories=0),
    visibility = ["//visibility:public"]
)""",
    # From https://golang.org/dl/
    sha256 = "dab7d9c34361dc21ec237d584590d72500652e7c909bf082758fb63064fca0ef",
    urls = ["https://golang.org/dl/go1.17.1.linux-amd64.tar.gz"],
)

http_archive(
    name = "cockroachdb_external",
    # We are downloading a zip file of pre-compiled executables, libraries and such that does
    # NOT have a BUILD file for Bazel to read. As such, we can specify one here using
    # build_file_content that makes all the contents of the the Android NDK zip file available
    # as a target called @android_ndk_external//:extracted_files
    build_file_content = """
filegroup(
    name = "extracted_exe",
    srcs = ["cockroach-v21.1.9.linux-amd64/cockroach"],
    visibility = ["//visibility:public"]
)""",
    # https://www.cockroachlabs.com/docs/v21.1/install-cockroachdb-linux does not currently
    # provide SHA256 signatures. kjlubick@ downloaded this file and computed this sha256 signature.
    sha256 = "05293e76dfb6443790117b6c6c05b1152038b49c83bd4345589e15ced8717be3",
    url = "https://binaries.cockroachdb.com/cockroach-v21.1.9.linux-amd64.tgz",
)
