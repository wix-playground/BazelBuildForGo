workspace(name = "com_github_wix_private_bazelbuildforgo")

http_archive(
    name = "io_bazel_rules_go",
    url = "https://github.com/bazelbuild/rules_go/releases/download/0.10.1/rules_go-0.10.1.tar.gz",
    sha256 = "4b14d8dd31c6dbaf3ff871adcd03f28c3274e42abc855cb8fb4d01233c0154dc",
)

http_archive(
    name = "bazel_gazelle",
    url = "https://github.com/bazelbuild/bazel-gazelle/releases/download/0.10.1/bazel-gazelle-0.10.1.tar.gz",
    sha256 = "d03625db67e9fb0905bbd206fa97e32ae9da894fe234a493e7517fd25faec914",
)

load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains","go_repository")
go_rules_dependencies()
go_register_toolchains()


go_repository(
    name = "com_github_coreos_etcd",
    commit = "f5c56401d74401f040258afeb5fbdd875fafaf69",
    build_file_generation = "on",
    build_file_name = "BUILD.bazel",
    importpath = "github.com/coreos/etcd",
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
gazelle_dependencies()


git_repository(
    name = "io_bazel_rules_docker",
    remote = "https://github.com/bazelbuild/rules_docker.git",
    tag = "v0.4.0",
)

load(
    "@io_bazel_rules_docker//container:container.bzl",
    "container_image","container_pull",
    container_repositories = "repositories",
)

container_repositories()

container_pull(
   name = "alpine",
   registry = "index.docker.io",
   repository = "library/alpine",
   tag = "3.4",
)