
go_test(
    name="basic_converter_test",
    srcs=glob(["*_test.go"]),
    deps=[
        "//goprotoconv",
        "//third_party:goassert",
    ],
    data=[
        "a.pb.go",
        "b.pb.go",
    ],
    visibility=["//tests/..."],
)
