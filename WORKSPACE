# Copyright 2017 The Bazel Authors. All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "com_github_bazelbuild_bazel_integration_testing",
    sha256 = "c147892a10fd48e6555832cd98b89974c147b051d395046cc6bfaafb3a95c3f5",
    strip_prefix = "bazel-integration-testing-89c26c5ef02a91f20a451eed76ce98126ddd4028",
    url = "https://github.com/bazelbuild/bazel-integration-testing/archive/89c26c5ef02a91f20a451eed76ce98126ddd4028.tar.gz",
)

load("@com_github_bazelbuild_bazel_integration_testing//tools:repositories.bzl", "bazel_binaries")

bazel_binaries()

http_archive(
    name = "bazel_skylib",
    sha256 = "2ea8a5ed2b448baf4a6855d3ce049c4c452a6470b1efd1504fdb7c1c134d220a",
    strip_prefix = "bazel-skylib-0.8.0",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-skylib/archive/0.8.0.tar.gz",
        "https://github.com/bazelbuild/bazel-skylib/archive/0.8.0.tar.gz",
    ],
)

# NOTE: URLs are mirrored by an asynchronous review process. They must
#       be greppable for that to happen. It's OK to submit broken mirror
#       URLs, so long as they're correctly formatted. Bazel's downloader
#       has fast failover.

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "77dfd303492f2634de7a660445ee2d3de2960cbd52f97d8c0dffa9362d3ddef9",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/0.18.1/rules_go-0.18.1.tar.gz",
        "https://github.com/bazelbuild/rules_go/releases/download/0.18.1/rules_go-0.18.1.tar.gz",
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "3c681998538231a2d24d0c07ed5a7658cb72bfb5fd4bf9911157c0e9ac6a2687",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/0.17.0/bazel-gazelle-0.17.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/0.17.0/bazel-gazelle-0.17.0.tar.gz",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

go_repository(
    name = "com_github_fsnotify_fsnotify",
    importpath = "github.com/fsnotify/fsnotify",
    tag = "v1.4.7",
)

go_repository(
    name = "com_github_jaschaephraim_lrserver",
    importpath = "github.com/jaschaephraim/lrserver",
    tag = "3.0.1",
)

go_repository(
    name = "com_github_gorilla_websocket",
    importpath = "github.com/gorilla/websocket",
    tag = "v1.4.0",
)

go_repository(
    name = "org_golang_x_sys",
    commit = "cc5685c2db1239775905f3911f0067c0fa74762f",
    importpath = "golang.org/x/sys",
)

# NOTE: this must match rules_go version located at
# https://github.com/bazelbuild/rules_go/blob/master/go/private/repositories.bzl
go_repository(
    name = "com_github_golang_protobuf",
    importpath = "github.com/golang/protobuf",
    tag = "v1.3.1",
)

go_repository(
    name = "com_github_gorilla_websocket",
    importpath = "github.com/gorilla/websocket",
    tag = "v1.4.0",
)

# NOTE: this must match rules_go version above
go_repository(
    name = "com_github_bazelbuild_rules_go",
    importpath = "github.com/bazelbuild/rules_go",
    tag = "0.16.2",
)
