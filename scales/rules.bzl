def _copy_m8n_file_impl(ctx):
    args = ctx.actions.args()
    args.add("%s/%s.m8n" % (ctx.attr.input_dir, ctx.attr.name))
    args.add(ctx.outputs.out)
    ctx.actions.run(
        outputs = [ctx.outputs.out],
        executable = "cp",
        arguments = [args],
    )

copy_m8n_file = rule(
    implementation=_copy_m8n_file_impl,
    attrs={
        "input_dir": attr.string(),
    },
    outputs = {"out": "%{name}.m8n"},
)
