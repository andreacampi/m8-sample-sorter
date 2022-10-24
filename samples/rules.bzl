def _copy_wav_files_impl(ctx):
    all_input_files = [
        f for t in ctx.attr.data
            for f in t.files.to_list()
    ]

    all_outputs = []
    for f in all_input_files:
        out = ctx.actions.declare_file(f.basename)
        all_outputs += [out]
        ctx.actions.run(
            outputs=[out],
            executable="cp",
            arguments=["../../../../../%s" % (f.path), out.path],
        )

    # Small sanity check
    if len(all_input_files) != len(all_outputs):
        fail("Output count should be 1-to-1 with input count.")

    return [
        DefaultInfo(
            files=depset(all_outputs),
            runfiles=ctx.runfiles(files=all_outputs))
    ]

samples = rule(
    implementation=_copy_wav_files_impl,
    attrs={
        "data": attr.label_list(
            allow_files = True,
        ),
    },
)
