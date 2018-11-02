load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    url = "https://github.com/bazelbuild/rules_go/releases/download/0.15.4/rules_go-0.15.4.tar.gz",
    sha256 = "7519e9e1c716ae3c05bd2d984a42c3b02e690c5df728dc0a84b23f90c355c5a1",
)

http_archive(
    name = "bazel_gazelle",
    urls = ["https://github.com/bazelbuild/bazel-gazelle/releases/download/0.14.0/bazel-gazelle-0.14.0.tar.gz"],
    sha256 = "c0a5739d12c6d05b6c1ad56f2200cb0b57c5a70e03ebd2f7b87ce88cabf09c7b",
)

http_archive(
    name = "com_github_atlassian_bazel_tools",
    strip_prefix = "bazel-tools-6fef37f33dfa0189be9df4d3d60e6291bfe71177",
    urls = ["https://github.com/atlassian/bazel-tools/archive/6fef37f33dfa0189be9df4d3d60e6291bfe71177.zip"],
)

git_repository(
    name = "io_bazel_rules_docker",
    commit = "7401cb256222615c497c0dee5a4de5724a4f4cc7",  # 2018-06-22
    remote = "https://github.com/bazelbuild/rules_docker.git",
)

load("@io_bazel_rules_docker//docker:docker.bzl", "docker_repositories")

docker_repositories()

git_repository(
    name = "build_bazel_rules_nodejs",
    remote = "https://github.com/bazelbuild/rules_nodejs.git",
    tag = "0.15.0",
)

load("@build_bazel_rules_nodejs//:package.bzl", "rules_nodejs_dependencies")

rules_nodejs_dependencies()

load("@build_bazel_rules_nodejs//:defs.bzl", "node_repositories", "yarn_install")

node_repositories()

yarn_install(
    name = "npm",
    package_json = "//:package.json",
    yarn_lock = "//:yarn.lock",
)

# This requires rules_docker to be fully instantiated before it is pulled in.
git_repository(
    name = "io_bazel_rules_k8s",
    commit = "2054f7bf4d51f9e439313c56d7a208960a8a179f",  # 2018-07-29
    remote = "https://github.com/bazelbuild/rules_k8s.git",
)

load("@io_bazel_rules_k8s//k8s:k8s.bzl", "k8s_repositories", "k8s_defaults")

k8s_repositories()

_CLUSTER = "minikube"

_NAMESPACE = "default"

[k8s_defaults(
    name = "k8s_" + kind,
    cluster = _CLUSTER,
    #context = _CONTEXT,
    kind = kind,
    namespace = _NAMESPACE,
) for kind in [
    "deploy",
    "job",
    "pod",
    "priority_class",
    "secret",
    "service",
]]

load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains")

go_rules_dependencies()

go_register_toolchains()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

load("@com_github_atlassian_bazel_tools//gometalinter:deps.bzl", "gometalinter_dependencies")

gometalinter_dependencies()

load(
    "@io_bazel_rules_docker//go:image.bzl",
    _go_image_repos = "repositories",
)

_go_image_repos()

git_repository(
    name = "io_kubernetes_build",
    commit = "4ce715fbe67d8fbed05ec2bb47a148e754100a4b",
    remote = "https://github.com/kubernetes/repo-infra.git",
)

git_repository(
    name = "com_github_jmhodges_bazel_gomock",
    commit = "5b73edb74e569ff404b3beffc809d6d9f205e0e4",
    remote = "https://github.com/jmhodges/bazel_gomock.git",
)

go_repository(
    name = "com_github_golang_mock",
    commit = "503ea999614e3b9d10f0060b8182b6fc7ea664f8",
    importpath = "github.com/golang/mock",
)

# External dependencies

go_repository(
    name = "com_github_ethereum_go_ethereum",
    # Last updated September 09, 2018
    # Note: go-ethereum is not bazel-friendly with regards to cgo. We have a
    # a fork that has resolved these issues by disabling HID/USB support and
    # some manual fixes for c imports in the crypto package. This is forked
    # branch should be updated from time to time with the latest go-ethereum
    # code.
    remote = "https://github.com/prysmaticlabs/bazel-go-ethereum",
    vcs = "git",
    # Last updated September 09, 2018
    commit = "f4b3f83362a4cf2928e57914af040aea76c8a7d6",
    importpath = "github.com/ethereum/go-ethereum",
)

go_repository(
    name = "com_github_urfave_cli",
    importpath = "github.com/urfave/cli",
    commit = "8e01ec4cd3e2d84ab2fe90d8210528ffbb06d8ff",
)

go_repository(
    name = "com_github_x_cray_logrus_prefixed_formatter",
    importpath = "github.com/x-cray/logrus-prefixed-formatter",
    commit = "bb2702d423886830dee131692131d35648c382e2",
)

go_repository(
    name = "com_github_mgutz_ansi",
    importpath = "github.com/mgutz/ansi",
    commit = "9520e82c474b0a04dd04f8a40959027271bab992",
)

go_repository(
    name = "com_github_fjl_memsize",
    importpath = "github.com/fjl/memsize",
    commit = "ca190fb6ffbc076ff49197b7168a760f30182d2e",
)

go_repository(
    name = "com_github_libp2p_go_floodsub",
    commit = "900095341be1eeba1a0d6c1e08ab61d1242c6cb4",
    importpath = "github.com/libp2p/go-floodsub",
    build_file_proto_mode = "disable_global",
)

go_repository(
    name = "com_github_libp2p_go_libp2p",
    commit = "57ed88e170d64b6fb98253654fd66e8634af0b11",
    importpath = "github.com/libp2p/go-libp2p",
    build_file_proto_mode = "disable_global",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_peer",
    commit = "993d742bc29dcf4894b7730ba610fd78900be76c",
    importpath = "github.com/libp2p/go-libp2p-peer",
    build_file_proto_mode = "disable_global",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_crypto",
    build_file_proto_mode = "disable_global",
    commit = "7240b40a3ddc47c4d17c15baabcbe45e5219171b",
    importpath = "github.com/libp2p/go-libp2p-crypto",
    build_file_proto_mode = "disable_global",
)

go_repository(
    name = "com_github_multiformats_go_multiaddr",
    commit = "fe1c46f8be5af4aff4db286e08839295bd922efb",
    importpath = "github.com/multiformats/go-multiaddr",
    build_file_proto_mode = "disable_global",
)

go_repository(
    name = "com_github_ipfs_go_log",
    commit = "4040da784a6be3829112e12ed1ee4b8e2fc69341",
    importpath = "github.com/ipfs/go-log",
)

go_repository(
    name = "com_github_multiformats_go_multihash",
    commit = "66cd79b386e51ddf63270f2b71c57436a60035f3",
    importpath = "github.com/multiformats/go-multihash",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_swarm",
    commit = "67f7e37245d17c8308dcf6ac5f5f0f66797694ef",
    importpath = "github.com/libp2p/go-libp2p-swarm",
    build_file_proto_mode = "disable_global",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_host",
    commit = "28ec0a42060315368874a2e09d7cd4dc6102814c",
    importpath = "github.com/libp2p/go-libp2p-host",
    build_file_proto_mode = "disable_global",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_peerstore",
    commit = "0df1488a3f842adeb536c8cb3ac395fbae1d69e1",
    importpath = "github.com/libp2p/go-libp2p-peerstore",
    build_file_proto_mode = "disable_global",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_circuit",
    commit = "eca2b86a1bcf00adba4f71b3654c3898b19ef421",
    importpath = "github.com/libp2p/go-libp2p-circuit",
    build_file_proto_mode = "disable_global",
)

go_repository(
    name = "com_github_coreos_go_semver",
    commit = "e214231b295a8ea9479f11b70b35d5acf3556d9b",
    importpath = "github.com/coreos/go-semver",
    build_file_proto_mode = "disable_global",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_interface_connmgr",
    commit = "74fba35f582dc5026b6dc1c43a7616f01c4598a9",
    importpath = "github.com/libp2p/go-libp2p-interface-connmgr",
    build_file_proto_mode = "disable_global",
)

go_repository(
    name = "com_github_libp2p_go_conn_security_multistream",
    commit = "9b6cfc8078f3d7030037e5686acb2989ba2a3683",
    importpath = "github.com/libp2p/go-conn-security-multistream",
    build_file_proto_mode = "disable_global",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_metrics",
    commit = "a10ff6e75dae3c868023867e8caa534a04bdc624",
    importpath = "github.com/libp2p/go-libp2p-metrics",
    build_file_proto_mode = "disable_global",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_net",
    commit = "c070e8fb6eade68612d96e0fbace731f89897d33",
    importpath = "github.com/libp2p/go-libp2p-net",
    build_file_proto_mode = "disable_global",
)

go_repository(
    name = "com_github_whyrusleeping_mafmt",
    commit = "1dc32401ee9fdd3f6cdb3405ec984d5dae877b2a",
    importpath = "github.com/whyrusleeping/mafmt",
    build_file_proto_mode = "disable_global",
)

go_repository(
    name = "com_github_multiformats_go_multiaddr_net",
    commit = "1cb9a0e8a6de3c8a10f6cee60d01d793603c4f7e",
    importpath = "github.com/multiformats/go-multiaddr-net",
    build_file_proto_mode = "disable_global",
)

go_repository(
    name = "com_github_agl_ed25519",
    commit = "5312a61534124124185d41f09206b9fef1d88403",
    importpath = "github.com/agl/ed25519",
)

go_repository(
    name = "com_github_minio_blake2b_simd",
    commit = "3f5f724cb5b182a5c278d6d3d55b40e7f8c2efb4",
    importpath = "github.com/minio/blake2b-simd",
)

go_repository(
    name = "com_github_gxed_hashland",
    commit = "d9f6b97f8db22dd1e090fd0bbbe98f09cc7dd0a8",
    importpath = "github.com/gxed/hashland",
)

go_repository(
    name = "com_github_mattn_go_colorable",
    commit = "efa589957cd060542a26d2dd7832fd6a6c6c3ade",
    importpath = "github.com/mattn/go-colorable",
)

go_repository(
    name = "com_github_whyrusleeping_mdns",
    commit = "ef14215e6b30606f4ce84174ed7a644a05cb1af3",
    importpath = "github.com/whyrusleeping/mdns",
)

go_repository(
    name = "com_github_btcsuite_btcd",
    commit = "2a560b2036bee5e3679ec2133eb6520b2f195213",
    importpath = "github.com/btcsuite/btcd",
)

go_repository(
    name = "com_github_minio_sha256_simd",
    commit = "ad98a36ba0da87206e3378c556abbfeaeaa98668",
    importpath = "github.com/minio/sha256-simd",
)

go_repository(
    name = "com_github_mr_tron_base58",
    commit = "9ad991d48a423a0e52c77eac173b1b589665492a",
    importpath = "github.com/mr-tron/base58",
)

go_repository(
    name = "com_github_whyrusleeping_go_smux_yamux",
    commit = "f60148d42be1861638cbb16dfd32bee46251f86b",
    importpath = "github.com/whyrusleeping/go-smux-yamux",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_secio",
    build_file_proto_mode = "disable_global",
    commit = "7b61c65c05f627ae0d02c4e9f5c5815ef1486d68",
    importpath = "github.com/libp2p/go-libp2p-secio",
    build_file_proto_mode = "disable_global",
)

go_repository(
    name = "com_github_libp2p_go_tcp_transport",
    commit = "5e52db593970b614b25507ad331d098cdc25998b",
    importpath = "github.com/libp2p/go-tcp-transport",
    build_file_proto_mode = "disable_global",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_protocol",
    commit = "b29f3d97e3a2fb8b29c5d04290e6cb5c5018004b",
    importpath = "github.com/libp2p/go-libp2p-protocol",
    build_file_proto_mode = "disable_global",
)

go_repository(
    name = "com_github_jbenet_goprocess",
    commit = "b497e2f366b8624394fb2e89c10ab607bebdde0b",
    importpath = "github.com/jbenet/goprocess",
)

go_repository(
    name = "com_github_multiformats_go_multistream",
    commit = "0e509f6b999d3be262c0702d25fe868c77de1565",
    importpath = "github.com/multiformats/go-multistream",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_loggables",
    commit = "fc4c4dc5d90bcd55e0af5b5f744005519e070dc3",
    importpath = "github.com/libp2p/go-libp2p-loggables",
    build_file_proto_mode = "disable_global",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_nat",
    commit = "fcc8db1a9963a3d489abbc632f2c0f404e4ff607",
    importpath = "github.com/libp2p/go-libp2p-nat",
    build_file_proto_mode = "disable_global",
)

go_repository(
    name = "com_github_multiformats_go_multiaddr_dns",
    commit = "78f39e8892d4f8c5c9f18679cc06d0b40ecab8cf",
    importpath = "github.com/multiformats/go-multiaddr-dns",
)

go_repository(
    name = "com_github_fd_go_nat",
    commit = "bad65a492f32121a87197f4a085905c35e2a367e",
    importpath = "github.com/fd/go-nat",
)

go_repository(
    name = "com_github_whyrusleeping_go_logging",
    commit = "0457bb6b88fc1973573aaf6b5145d8d3ae972390",
    importpath = "github.com/whyrusleeping/go-logging",
)

go_repository(
    name = "com_github_mattn_go_isatty",
    commit = "3fb116b820352b7f0c281308a4d6250c22d94e27",
    importpath = "github.com/mattn/go-isatty",
)

go_repository(
    name = "com_github_libp2p_go_stream_muxer",
    commit = "9c6bd93eecbbab56630bb2688bb435d9fd1dfeb1",
    importpath = "github.com/libp2p/go-stream-muxer",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_transport_upgrader",
    commit = "49139764f899e43870cd4c0485a8bbbb89f467db",
    importpath = "github.com/libp2p/go-libp2p-transport-upgrader",
    build_file_proto_mode = "disable_global",
)

go_repository(
    name = "com_github_libp2p_go_testutil",
    commit = "f967bbd5fcb7fb6337504e5d78c53c865e80733c",
    importpath = "github.com/libp2p/go-testutil",
    build_file_proto_mode = "disable_global",
)

go_repository(
    name = "com_github_whyrusleeping_go_smux_multistream",
    commit = "e799b10bc6eea2cc5ce18f7b7b4d8e1a0439476d",
    importpath = "github.com/whyrusleeping/go-smux-multistream",
)

go_repository(
    name = "com_github_libp2p_go_maddr_filter",
    commit = "57fd7e2ed649ba28b4f2c7bcab3a606e7cc4b12c",
    importpath = "github.com/libp2p/go-maddr-filter",
    build_file_proto_mode = "disable_global",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_transport",
    commit = "bba5e63c9faea32bcb9b5416d70180bf2bc8033b",
    importpath = "github.com/libp2p/go-libp2p-transport",
    build_file_proto_mode = "disable_global",
)

go_repository(
    name = "com_github_libp2p_go_addr_util",
    commit = "2dc53609d9aa745f7bf5755bef6f54fb9ff77cd6",
    importpath = "github.com/libp2p/go-addr-util",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_interface_pnet",
    commit = "d240acf619f63dfb776821a1d4d28a918f77edd5",
    importpath = "github.com/libp2p/go-libp2p-interface-pnet",
    build_file_proto_mode = "disable_global",
)

go_repository(
    name = "com_github_libp2p_go_conn_security",
    commit = "49eda33ab1bbdd3ac1115e573a4f8dee777bd901",
    importpath = "github.com/libp2p/go-conn-security",
    build_file_proto_mode = "disable_global",
)

go_repository(
    name = "com_github_whyrusleeping_timecache",
    commit = "cfcb2f1abfee846c430233aef0b630a946e0a5a6",
    importpath = "github.com/whyrusleeping/timecache",
)

go_repository(
    name = "com_github_miekg_dns",
    commit = "ba6747e8a94115e9dc7738afb87850687611df1b",
    importpath = "github.com/miekg/dns",
)

go_repository(
    name = "com_github_opentracing_opentracing_go",
    commit = "6aa6febac7b98f836100ecaea478c04f30b6dbd0",
    importpath = "github.com/opentracing/opentracing-go",
)

go_repository(
    name = "com_github_libp2p_go_reuseport",
    commit = "dd0c37d7767bc38280bd9813145b65f8bd560629",
    importpath = "github.com/libp2p/go-reuseport",
    build_file_proto_mode = "disable_global",
)

go_repository(
    name = "com_github_huin_goupnp",
    commit = "1395d1447324cbea88d249fbfcfd70ea878fdfca",
    importpath = "github.com/huin/goupnp",
)

go_repository(
    name = "com_github_spaolacci_murmur3",
    commit = "f09979ecbc725b9e6d41a297405f65e7e8804acc",
    importpath = "github.com/spaolacci/murmur3",
)

go_repository(
    name = "com_github_jbenet_go_temp_err_catcher",
    commit = "aac704a3f4f27190b4ccc05f303a4931fd1241ff",
    importpath = "github.com/jbenet/go-temp-err-catcher",
)

go_repository(
    name = "com_github_satori_go_uuid",
    commit = "36e9d2ebbde5e3f13ab2e25625fd453271d6522e",
    importpath = "github.com/satori/go.uuid",
)

go_repository(
    name = "com_github_sirupsen_logrus",
    commit = "1ed61965b9e594bf37539680d7f63eccd060314f",
    importpath = "github.com/sirupsen/logrus",
)

go_repository(
    name = "org_golang_x_sys",
    commit = "e4b3c5e9061176387e7cea65e4dc5853801f3fb7",
    importpath = "golang.org/x/sys",
)

go_repository(
    name = "com_github_whyrusleeping_yamux",
    commit = "3f61f992de5b927691ca81b814e04061f5ca625a",
    importpath = "github.com/whyrusleeping/yamux",
)

go_repository(
    name = "com_github_libp2p_go_flow_metrics",
    commit = "7e5a55af485341567f98d6847a373eb5ddcdcd43",
    importpath = "github.com/libp2p/go-flow-metrics",
    build_file_proto_mode = "disable_global",
)

go_repository(
    name = "com_github_libp2p_go_msgio",
    commit = "5b7d4ee2db24d582f5c255f0e3411d7144835311",
    importpath = "github.com/libp2p/go-msgio",
    build_file_proto_mode = "disable_global",
)

go_repository(
    name = "com_github_jackpal_gateway",
    commit = "cbcf4e3f3baee7952fc386c8b2534af4d267c875",
    importpath = "github.com/jackpal/gateway",
)

go_repository(
    name = "com_github_whyrusleeping_multiaddr_filter",
    commit = "e903e4adabd70b78bc9293b6ee4f359afb3f9f59",
    importpath = "github.com/whyrusleeping/multiaddr-filter",
)

go_repository(
    name = "com_github_libp2p_go_ws_transport",
    commit = "ac3f59a8924f35d6403210da5287096dfde75bd2",
    importpath = "github.com/libp2p/go-ws-transport",
    build_file_proto_mode = "disable_global",
)

go_repository(
    name = "org_golang_x_crypto",
    commit = "e3636079e1a4c1f337f212cc5cd2aca108f6c900",
    importpath = "golang.org/x/crypto",
)

go_repository(
    name = "com_github_jackpal_go_nat_pmp",
    commit = "28a68d0c24adce1da43f8df6a57340909ecd7fdd",
    importpath = "github.com/jackpal/go-nat-pmp",
)

go_repository(
    name = "com_github_libp2p_go_reuseport_transport",
    commit = "5cdb097c8035e75fc59d12f22509aeb700a272d0",
    importpath = "github.com/libp2p/go-reuseport-transport",
    build_file_proto_mode = "disable_global",
)

go_repository(
    name = "com_github_libp2p_go_sockaddr",
    commit = "a7494d4eefeb607c8bc491cf8850a6e8dbd41cab",
    importpath = "github.com/libp2p/go-sockaddr",
    build_file_proto_mode = "disable_global",
)

go_repository(
    name = "com_github_whyrusleeping_go_notifier",
    commit = "097c5d47330ff6a823f67e3515faa13566a62c6f",
    importpath = "github.com/whyrusleeping/go-notifier",
)

go_repository(
    name = "com_github_gorilla_websocket",
    commit = "3130e8d3f1986de7dc3fd86255f4ce9618677508",
    importpath = "github.com/gorilla/websocket",
)

go_repository(
    name = "com_github_whyrusleeping_go_smux_multiplex",
    commit = "b905dd063c9764201a381824e7191a7ccaf3f4bc",
    importpath = "github.com/whyrusleeping/go-smux-multiplex",
)

go_repository(
    name = "com_github_gxed_eventfd",
    commit = "80a92cca79a8041496ccc9dd773fcb52a57ec6f9",
    importpath = "github.com/gxed/eventfd",
)

go_repository(
    name = "com_github_whyrusleeping_go_multiplex",
    commit = "015295179194cbcc2eb7e13504222749af868544",
    importpath = "github.com/whyrusleeping/go-multiplex",
)

go_repository(
    name = "com_github_gxed_goendian",
    commit = "0f5c6873267e5abf306ffcdfcfa4bf77517ef4a7",
    importpath = "github.com/gxed/GoEndian",
)

go_repository(
    name = "com_github_syndtr_goleveldb",
    commit = "c4c61651e9e37fa117f53c5a906d3b63090d8445",
    importpath = "github.com/syndtr/goleveldb",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_blankhost",
    commit = "073f507db72de824e981aa0f15f158175a8d6be1",
    importpath = "github.com/libp2p/go-libp2p-blankhost",
    build_file_proto_mode = "disable_global",
)

go_repository(
    name = "com_github_steakknife_hamming",
    commit = "c99c65617cd3d686aea8365fe563d6542f01d940",
    importpath = "github.com/steakknife/hamming",
)

go_repository(
    name = "io_opencensus_go",
    commit = "f21fe3feadc5461b952191052818685a410428d4",
    importpath = "go.opencensus.io",
)

go_repository(
    name = "org_golang_google_api",
    commit = "7ca32eb868bf53ea2fc406698eb98583a8073d19",
    importpath = "google.golang.org/api",
)

go_repository(
    name = "org_golang_x_sync",
    commit = "1d60e4601c6fd243af51cc01ddf169918a5407ca",
    importpath = "golang.org/x/sync",
)

go_repository(
<<<<<<< HEAD
    name = "io_k8s_client_go",
    build_file_proto_mode = "disable_global",
    commit = "a5d40177db184c504978c4b08350ac9e367042b0",
    importpath = "k8s.io/client-go",
)

go_repository(
    name = "io_k8s_apimachinery",
    build_file_proto_mode = "disable_global",
    commit = "5ae511e0ed34ac0fbf6e17340df5605372591d00",
    importpath = "k8s.io/apimachinery",
)

go_repository(
    name = "com_github_google_gofuzz",
    commit = "24818f796faf91cd76ec7bddd72458fbced7a6c1",
    importpath = "github.com/google/gofuzz",
)

go_repository(
    name = "io_k8s_api",
    build_file_proto_mode = "disable_global",
    commit = "a191abe0b71e00ce4cde58af8002aa4c1a8bb068",
    importpath = "k8s.io/api",
)

go_repository(
    name = "com_github_googleapis_gnostic",
    commit = "48a0ecefe2e4190c7a2d63b477875854a2f993b3",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/googleapis/gnostic",
)

go_repository(
    name = "com_github_gregjones_httpcache",
    commit = "9cad4c3443a7200dd6400aef47183728de563a38",
    importpath = "github.com/gregjones/httpcache",
)

go_repository(
    name = "com_github_peterbourgon_diskv",
    commit = "0646ccaebea1ed1539efcab30cae44019090093f",
    importpath = "github.com/peterbourgon/diskv",
)

go_repository(
    name = "org_golang_x_time",
    commit = "fbb02b2291d28baffd63558aa44b4b56f178d650",
    importpath = "golang.org/x/time",
)

go_repository(
    name = "com_github_google_gofuzz",
    commit = "24818f796faf91cd76ec7bddd72458fbced7a6c1",
    importpath = "github.com/google/gofuzz",
)

go_repository(
    name = "com_github_grpc_ecosystem_go_grpc_prometheus",
    commit = "93bf4626fba73b751b0f3cdf2649be4ce0c420cd",
    importpath = "github.com/grpc-ecosystem/go-grpc-prometheus",
)

go_repository(
    name = "com_github_prometheus_client_golang",
    commit = "73edb9af667d9467d99b2001722770062a0086c3",
    importpath = "github.com/prometheus/client_golang",
)

go_repository(
    name = "com_github_beorn7_perks",
    commit = "3a771d992973f24aa725d07868b467d1ddfceafb",
    importpath = "github.com/beorn7/perks",
)

go_repository(
    name = "com_github_prometheus_procfs",
    commit = "418d78d0b9a7b7de3a6bbc8a23def624cc977bb2",
    importpath = "github.com/prometheus/procfs",
)

go_repository(
    name = "com_github_prometheus_client_model",
    commit = "5c3871d89910bfb32f5fcab2aa4b9ec68e65a99f",
    importpath = "github.com/prometheus/client_model",
)

go_repository(
    name = "com_github_prometheus_common",
    commit = "c7de2306084e37d54b8be01f3541a8464345e9a5",
    importpath = "github.com/prometheus/common",
)

go_repository(
    name = "com_github_matttproud_golang_protobuf_extensions",
    commit = "c12348ce28de40eed0136aa2b644d0ee0650e56c",
    importpath = "github.com/matttproud/golang_protobuf_extensions",
)

go_repository(
    name = "com_github_grpc_ecosystem_go_grpc_middleware",
    commit = "498ae206fc3cfe81cd82e48c1d4354026fa5f9ec",
    importpath = "github.com/grpc-ecosystem/go-grpc-middleware",
)

go_repository(
    name = "in_gopkg_yaml_v2",
    commit = "5420a8b6744d3b0345ab293f6fcba19c978f1183",
    importpath = "gopkg.in/yaml.v2",
)

go_repository(
    name = "com_github_google_btree",
    commit = "4030bb1f1f0c35b30ca7009e9ebd06849dd45306",
    importpath = "github.com/google/btree",
)

go_repository(
    name = "com_github_gogo_protobuf",
    commit = "61dbc136cf5d2f08d68a011382652244990a53a9",
    importpath = "github.com/gogo/protobuf",
)

go_repository(
    name = "com_github_google_uuid",
    commit = "9b3b1e0f5f99ae461456d768e7d301a7acdaa2d8",
    importpath = "github.com/google/uuid",
)

go_repository(
    name = "com_github_hashicorp_golang_lru",
    commit = "20f1fb78b0740ba8c3cb143a61e86ba5c8669768",
    importpath = "github.com/hashicorp/golang-lru",
)

go_repository(
    name = "com_github_ipfs_go_cid",
    commit = "6e296c5c49ad84dc6a44af69fa1fe4e1245cd0cf",
    importpath = "github.com/ipfs/go-cid",
)

go_repository(
    name = "com_github_ipfs_go_datastore",
    commit = "19b8f34d345ede2d054ee7645b61e9c169992d7f",
    importpath = "github.com/ipfs/go-datastore",
)

go_repository(
    name = "com_github_ipfs_go_ipfs_util",
    commit = "10d786c5ed859afd22223df76a89bf57b24b2ee1",
    importpath = "github.com/ipfs/go-ipfs-util",
)

go_repository(
    name = "com_github_ipfs_go_todocounter",
    commit = "1e832b829506383050e6eebd12e05ea41a451532",
    importpath = "github.com/ipfs/go-todocounter",
)

go_repository(
    name = "com_github_jbenet_go_context",
    commit = "d14ea06fba99483203c19d92cfcd13ebe73135f4",
    importpath = "github.com/jbenet/go-context",
)

go_repository(
    name = "com_github_libp2p_go_buffer_pool",
    commit = "1d28ab5fdb424e5b3523d357553082e68df9b068",
    importpath = "github.com/libp2p/go-buffer-pool",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_kad_dht",
    commit = "fd8d798b512f65a2def035c1efac18081550e759",
    importpath = "github.com/libp2p/go-libp2p-kad-dht",
    build_file_proto_mode = "disable_global",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_kbucket",
    commit = "36076693e35f150f4e170db360cb020632f57a8d",
    importpath = "github.com/libp2p/go-libp2p-kbucket",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_record",
    commit = "7a182bb5ae667ca5c930d807be14d0655afd3d57",
    importpath = "github.com/libp2p/go-libp2p-record",
    build_file_proto_mode = "disable_global",
)

go_repository(
    name = "com_github_libp2p_go_libp2p_routing",
    commit = "65ad443586481b73cf4d391b07a139ae0388af7d",
    importpath = "github.com/libp2p/go-libp2p-routing",
)

go_repository(
    name = "com_github_libp2p_go_mplex",
    commit = "e717d3c8e1c30cadaab678b6909f7d1a35847c35",
    importpath = "github.com/libp2p/go-mplex",
)

go_repository(
    name = "com_github_multiformats_go_multibase",
    commit = "bb91b53e5695e699a86654d77d03db7bc7506d12",
    importpath = "github.com/multiformats/go-multibase",
)

go_repository(
    name = "com_github_whyrusleeping_base32",
    commit = "c30ac30633ccdabefe87eb12465113f06f1bab75",
    importpath = "github.com/whyrusleeping/base32",
)

go_repository(
    name = "com_github_whyrusleeping_go_keyspace",
    commit = "5b898ac5add1da7178a4a98e69cb7b9205c085ee",
    importpath = "github.com/whyrusleeping/go-keyspace",
)

go_repository(
    name = "org_golang_x_net",
    commit = "4dfa2610cdf3b287375bbba5b8f2a14d3b01d8de",
    importpath = "golang.org/x/net",
)

go_repository(
    name = "org_golang_x_text",
    commit = "905a57155faa8230500121607930ebb9dd8e139c",
    importpath = "golang.org/x/text",
)

go_repository(
    name = "org_golang_x_lint",
    commit = "06c8688daad7faa9da5a0c2f163a3d14aac986ca",
    importpath = "golang.org/x/lint",
)

go_repository(
    name = "com_github_aristanetworks_goarista",
    commit = "ff33da284e760fcdb03c33d37a719e5ed30ba844",
    importpath = "github.com/aristanetworks/goarista",
)

go_repository(
    name = "com_github_boltdb_bolt",
    commit = "fd01fc79c553a8e99d512a07e8e0c63d4a3ccfc5",
    importpath = "github.com/boltdb/bolt",
)
