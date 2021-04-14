"""This module defines rules for building Skia Infrastructure web applications."""

load("@build_bazel_rules_nodejs//:index.bzl", _nodejs_test = "nodejs_test")
load("@npm//@bazel/rollup:index.bzl", "rollup_bundle")
load("@npm//@bazel/terser:index.bzl", "terser_minified")
load("@npm//html-insert-assets:index.bzl", "html_insert_assets")
load("@io_bazel_rules_docker//container:flatten.bzl", "container_flatten")
load("@io_bazel_rules_sass//:defs.bzl", "sass_binary", _sass_library = "sass_library")
load("//bazel/test_on_env:test_on_env.bzl", "test_on_env")
load("//infra-sk/html_insert_nonce_attribute:index.bzl", "html_insert_nonce_attribute")
load("//infra-sk/karma_test:index.bzl", _karma_test = "karma_test")
load("//infra-sk/sk_demo_page_server:index.bzl", _sk_demo_page_server = "sk_demo_page_server")
load(":ts_library.bzl", _ts_library = "ts_library")

# Re-export these common rules so we only have to load this .bzl file from our BUILD.bazel files.
karma_test = _karma_test
sass_library = _sass_library
sk_demo_page_server = _sk_demo_page_server
ts_library = _ts_library

def sk_element(
        name,
        ts_srcs,
        sass_srcs = [],
        ts_deps = [],
        sass_deps = [],
        sk_element_deps = [],
        visibility = None):
    """Defines a custom element for Skia Infrastructure web applications.

    This is just a convenience macro that generates the ts_library and sass_library targets required
    to build a custom element.

    By convention, if an sk_element includes a <name>.scss file in its sass_srcs, it will be
    considered the sk_element's entry-point Sass stylesheet. If sk_element A is depended on by
    sk_element B (via B's sk_element_deps argument), then A's entry-point Sass stylesheet will be
    automatically imported from B's styles.

    This macro generates a "ghost" Sass stylesheet with `@import` statements for the entry-point
    Sass stylesheet of this custom element, and those of each sk_element dependency in
    sk_element_deps. This ghost stylesheet will be included in the output CSS bundles of any sk_page
    that depends on this sk_element (directly or indirectly). It is therefore *not* necessary to
    explicitly import the Sass styles of the sk_element dependencies in sk_element_deps.

    Args:
      name: The name of the target.
      ts_srcs: TypeScript source files.
      sass_srcs: Sass source files.
      ts_deps: Any ts_library dependencies.
      sass_deps: Any sass_library dependencies.
      sk_element_deps: Any sk_element dependencies. Equivalent to adding the ts_library and
        sass_library of each sk_element to ts_deps and sass_deps, respectively.
      visibility: Visibility of the generated ts_library and sass_library targets.
    """

    # Find out whether a Sass entry-point stylesheet is provided.
    scss_entry_point = name + ".scss"
    if scss_entry_point not in sass_srcs:
        scss_entry_point = None

    # Generate a "ghost" Sass stylesheet with imports for the sk_element's entry-point stylesheet,
    # and the stylesheets with generated imports of each sk_element in sk_element_deps.
    #
    # This stylesheet recursively imports the Sass stylesheets of all the direct and indirect
    # sk_element dependencies (sk_element_deps argument).
    #
    # This stylesheet is not used by sk_elements directly. It is used in the sk_page macro to build
    # CSS bundles containing the Sass styles of all sk_elements in the sk_page's dependency graph.
    generate_sass_stylesheet_with_imports(
        name = name + "_sk_element_deps_scss",
        scss_files_to_import = ([scss_entry_point] if scss_entry_point else []) + [
            make_label_target_explicit(dep) + "_sk_element_deps_scss"
            for dep in sk_element_deps
        ],
        scss_output_file = name + "__generated_entrypoint_and_sk_element_deps_imports.scss",
        visibility = visibility,
    )

    # Extend ts_deps and sass_deps with the ts_library and sass_library targets produced by each
    # sk_element dependency in the sk_element_deps argument.
    all_ts_deps = [dep for dep in ts_deps]
    all_sass_deps = [dep for dep in sass_deps]
    for sk_element_dep in sk_element_deps:
        all_ts_deps.append(sk_element_dep)
        all_sass_deps.append(make_label_target_explicit(sk_element_dep) + "_styles")

    ts_library(
        name = name,
        srcs = ts_srcs,
        deps = all_ts_deps,
        visibility = visibility,
    )

    sass_library(
        name = name + "_styles",
        srcs = sass_srcs + [name + "_sk_element_deps_scss"],
        deps = all_sass_deps,
        visibility = visibility,
    )

def generate_sass_stylesheet_with_imports(name, scss_files_to_import, scss_output_file, visibility = None):
    """Generates a .scss file with one `@import` statement for each file in scss_files_to_import.

    Args:
      name: The name of the target.
      scss_files_to_import: A list of .scss files.
      scss_output_file: Name of the .scss file to generate.
      visibility: Visibility of the target.
    """

    # Build a list of shell commands to generate the output stylesheet.
    cmds = ["touch $@"]
    for scss_file in scss_files_to_import:
        import_statement = "@import '$(rootpath %s)';" % scss_file
        cmds.append("echo \"%s\" >> $@" % import_statement)

    native.genrule(
        name = name,
        srcs = scss_files_to_import,
        outs = [scss_output_file],
        cmd = " && ".join(cmds),
        visibility = visibility,
    )

def make_label_target_explicit(label):
    """Takes a label with a potentially implicit target name, and makes the target name explicit.

    For example, if the label is "//path/to/pkg", this macro will return "//path/to/pkg:pkg". If the
    label is already in the latter form, the label will be returned unchanged.

    Reference: https://docs.bazel.build/versions/master/build-ref.html#labels

    Args:
      label: A Bazel label.

    Returns:
      The given label, expanded to make the target name explicit.
     """
    if label.find(":") != -1:
        return label
    pkg_name = label.split("/").pop()
    return label + ":" + pkg_name

def nodejs_test(
        name,
        src,
        deps = [],
        tags = [],
        visibility = None,
        _internal_skip_naming_convention_enforcement = False):
    """Runs a Node.js unit test using the Mocha test runner.

    For tests that should run in the browser, please use karma_test instead.

    Args:
      name: Name of the target.
      src: A single TypeScript source file.
      deps: Any ts_library dependencies.
      tags: Tags for the generated nodejs_test rule.
      visibility: Visibility of the generated nodejs_test rule.
      _internal_skip_naming_convention_enforcement: Not part of the public API - do not use.
    """

    # This macro is called by sk_element_puppeteer_test, which uses a different naming convention.
    if not _internal_skip_naming_convention_enforcement and not src.endswith("_nodejs_test.ts"):
        fail("Node.js tests must end with \"_nodejs_test.ts\".")

    mocha_deps = [
        "@npm//mocha",
        "@npm//ts-node",
        "//:tsconfig.json",
    ]

    _nodejs_test(
        name = name,
        entry_point = "@npm//:node_modules/mocha/bin/mocha",
        data = [src] + deps + [dep for dep in mocha_deps if dep not in deps],
        templated_args = [
            "--require ts-node/register/transpile-only",
            "--timeout 60000",
            "--colors",
            "$(rootpath %s)" % src,
        ],
        tags = tags,
        visibility = visibility,
    )

def sk_element_puppeteer_test(name, src, sk_demo_page_server, deps = []):
    """Defines a Puppeteer test for the demo page served by an sk_demo_page_server.

    Puppeteer tests should save any screenshots inside the $TEST_UNDECLARED_OUTPUTS_DIR directory.
    To reduce the chances of name collisions, tests must save their screenshots under the
    $TEST_UNDECLARED_OUTPUTS_DIR/puppeteer-test-screenshots subdirectory. This convention will
    allow us to recover screenshots from multiple tests in a consistent way.

    Screenshots, and any other undeclared outputs of a test, can be found under //bazel-testlogs
    bundled as a single .zip file per test target. For example, if we run a Puppeteer test with e.g.
    "bazel test //path/to/my:puppeteer_test", any screenshots taken by this test will be found
    inside //bazel-testlogs/path/to/my/puppeteer_test/test.outputs/outputs.zip.

    To read more about undeclared test outputs, please see the following link:
    https://docs.bazel.build/versions/master/test-encyclopedia.html#test-interaction-with-the-filesystem.

    Args:
      name: Name of the rule.
      src: A single TypeScript source file.
      sk_demo_page_server: Label for the sk_demo_page_server target.
      deps: Any ts_library dependencies.
    """

    if not src.endswith("_puppeteer_test.ts"):
        fail("Puppeteer tests must end with \"_puppeteer_test.ts\".")

    nodejs_test(
        name = name + "_test_only",
        src = src,
        tags = ["manual"],  # Exclude it from wildcards, e.g. "bazel test all".
        deps = deps,
        _internal_skip_naming_convention_enforcement = True,
    )

    test_on_env(
        name = name,
        env = sk_demo_page_server,
        test = name + "_test_only",
    )

def copy_file(name, src, dst):
    """Copies a single file to a destination path, making parent directories as needed."""
    native.genrule(
        name = name,
        srcs = [src],
        outs = [dst],
        cmd = "mkdir -p $$(dirname $@) && cp $< $@",
    )

def sk_page(
        name,
        html_file,
        ts_entry_point,
        scss_entry_point = None,
        ts_deps = [],
        sass_deps = [],
        sk_element_deps = [],
        assets_serving_path = "/",
        nonce = None):
    """Builds a static HTML page, and its CSS and JavaScript development and production bundles.

    This macro generates the following files, where <name> is the given target name:

        development/<name>.html
        development/<name>.js
        development/<name>.css
        production/<name>.html
        production/<name>.js
        production/<name>.css

    The <name> target defined by this macro generates all of the above files.

    Tags <script> and <link> will be inserted into the output HTML pointing to the generated
    bundles. The serving path for said bundles defaults to "/" and can be overriden via the
    assets_serving_path argument.

    A timestamp will be appended to the URLs for any referenced assets for cache busting purposes,
    e.g. <script src="/index.js?v=27396986"></script>.

    If the nonce argument is provided, a nonce attribute will be inserted to all <link> and <script>
    tags. For example, if the nonce argument is set to "{% .Nonce %}", then the generated HTML will
    contain tags such as <script nonce="{% .Nonce %}" src="/index.js?v=27396986"></script>.

    This macro is designed to work side by side with the existing Webpack build without requiring
    any major changes to the pages in question.

    Args:
      name: The prefix used for the names of all the targets generated by this macro.
      html_file: The page's HTML file.
      ts_entry_point: TypeScript file used as the entry point for the JavaScript bundles.
      scss_entry_point: Sass file used as the entry point for the CSS bundles.
      ts_deps: Any ts_library dependencies.
      sass_deps: Any sass_library dependencies.
      sk_element_deps: Any sk_element dependencies. Equivalent to adding the ts_library and
        sass_library of each sk_element to deps and sass_deps, respectively.
      assets_serving_path: Path prefix for the inserted <script> and <link> tags.
      nonce: If set, its contents will be added as a "nonce" attributes to any inserted <script> and
        <link> tags.
    """

    # Extend ts_deps and sass_deps with the ts_library and sass_library targets produced by each
    # sk_element dependency in the sk_element_deps argument.
    all_ts_deps = [dep for dep in ts_deps]
    all_sass_deps = [dep for dep in sass_deps]
    for sk_element_dep in sk_element_deps:
        all_ts_deps.append(sk_element_dep)
        all_sass_deps.append(make_label_target_explicit(sk_element_dep) + "_styles")

    # Output directories.
    DEV_OUT_DIR = "development"
    PROD_OUT_DIR = "production"

    #######################
    # JavaScript bundles. #
    #######################

    ts_library(
        name = "%s_ts_lib" % name,
        srcs = [ts_entry_point],
        deps = all_ts_deps,
    )

    # Generates file <name>_js_bundle.js. Intermediate result; do not use.
    rollup_bundle(
        name = "%s_js_bundle" % name,
        deps = [
            ":%s_ts_lib" % name,
            "@npm//@rollup/plugin-node-resolve",
            "@npm//@rollup/plugin-commonjs",
            "@npm//rollup-plugin-sourcemaps",
        ],
        entry_point = ts_entry_point,
        format = "umd",
        config_file = "//infra-sk:rollup.config.js",
    )

    # Generates file <name>_js_bundle_minified.js. Intermediate result; do not use.
    terser_minified(
        name = "%s_js_bundle_minified" % name,
        src = "%s_js_bundle.js" % name,
        sourcemap = False,
    )

    # Generates file development/<name>.js.
    copy_file(
        name = "%s_js_dev" % name,
        src = "%s_js_bundle.js" % name,
        dst = "%s/%s.js" % (DEV_OUT_DIR, name),
    )

    # Generates file production/<name>.js.
    copy_file(
        name = "%s_js_prod" % name,
        # For some reason the output of the terser_minified rule above is not directly visible as a
        # source file, so we use the rule name instead (i.e. we drop the ".js" extension).
        src = "%s_js_bundle_minified" % name,
        dst = "%s/%s.js" % (PROD_OUT_DIR, name),
    )

    ################
    # CSS Bundles. #
    ################

    # Generate a blank Sass entry-point file to appease the sass_library rule, if one is not given.
    if not scss_entry_point:
        scss_entry_point = name + "__generated_empty_scss_entry_point"
        native.genrule(
            name = scss_entry_point,
            outs = [scss_entry_point + ".scss"],
            cmd = "touch $@",
        )

    # Create a sass_library including the scss_entry_point file, and all the Sass dependencies.
    sass_library(
        name = name + "_styles",
        srcs = [scss_entry_point],
        deps = all_sass_deps,
    )

    # Generate a "ghost" Sass stylesheet with imports for the scss_entry_point, and the stylesheets
    # with generated imports of each sk_element in sk_element_deps.
    #
    # This stylesheet recursively imports the Sass stylesheets of all the direct and indirect
    # sk_element dependencies (sk_element_deps argument).
    #
    # We will use this generated stylesheet as the entry-points for the sass_binaries below.
    generate_sass_stylesheet_with_imports(
        name = name + "_sk_element_deps_scss",
        scss_files_to_import = ([scss_entry_point] if scss_entry_point else []) + [
            make_label_target_explicit(dep) + "_sk_element_deps_scss"
            for dep in sk_element_deps
        ],
        scss_output_file = name + "__generated_entrypoint_and_sk_element_deps_imports.scss",
    )

    # Notes:
    #  - Sass compilation errors are not visible unless "bazel build" is invoked with flag
    #    "--strategy=SassCompiler=sandboxed" (now set by default in //.bazelrc). This is due to a
    #    known issue with sass_binary. For more details please see
    #    https://github.com/bazelbuild/rules_sass/issues/96.

    # Generates file development/<name>.css.
    sass_binary(
        name = "%s_css_dev" % name,
        src = name + "_sk_element_deps_scss",
        output_name = "%s/%s.css" % (DEV_OUT_DIR, name),
        deps = [name + "_styles"],
        include_paths = ["//node_modules"],
        output_style = "expanded",
        sourcemap = True,
        sourcemap_embed_sources = True,
    )

    # Generates file production/<name>.css.
    sass_binary(
        name = "%s_css_prod" % name,
        src = name + "_sk_element_deps_scss",
        output_name = "%s/%s.css" % (PROD_OUT_DIR, name),
        deps = [name + "_styles"],
        include_paths = ["//node_modules"],
        output_style = "compressed",
        sourcemap = False,
    )

    ###############
    # HTML files. #
    ###############

    # Generates file <name>.with_assets.html. Intermediate result; do not use.
    #
    # See https://www.npmjs.com/package/html-insert-assets.
    html_insert_assets(
        name = "%s_html" % name,
        outs = ["%s.with_assets.html" % name],
        args = [
            "--html=$(location %s)" % html_file,
            "--out=$@",
            "--roots=$(RULEDIR)",
            "--assets",
            # This is OK because html-insert-assets normalizes paths with successive slashes.
            "%s/%s.js" % (assets_serving_path, name),
            "%s/%s.css" % (assets_serving_path, name),
        ],
        data = [
            html_file,
            # This rule does not use the bundles directly, but by declaring them as dependencies via
            # the "data" argument, we guarantee that Bazel will rebuild <name>.with_assets.html any
            # time the bundles change. This refreshes the asset URL query parameters added by this
            # rule for cache busting purposes.
            "%s_js_dev" % name,
            "%s_js_prod" % name,
            "%s_css_dev" % name,
            "%s_css_prod" % name,
        ],
    )

    if nonce:
        # Generates file <name>.with_assets_and_nonce.html. Intermediate result; do not use.
        html_insert_nonce_attribute(
            name = "%s_html_nonce" % name,
            src = "%s.with_assets.html" % name,
            out = "%s.with_assets_and_nonce.html" % name,
            nonce = nonce,
        )

    instrumented_html = ("%s.with_assets_and_nonce.html" if nonce else "%s.with_assets.html") % name

    # Generates file development/<name>.html.
    copy_file(
        name = "%s_html_dev" % name,
        src = instrumented_html,
        dst = "%s/%s.html" % (DEV_OUT_DIR, name),
    )

    # Generates file production/<name>.html.
    copy_file(
        name = "%s_html_prod" % name,
        src = instrumented_html,
        dst = "%s/%s.html" % (PROD_OUT_DIR, name),
    )

    ###########################
    # Convenience filegroups. #
    ###########################

    # Generates all output files (that is, the development and production bundles).
    native.filegroup(
        name = name,
        srcs = [
            ":%s_dev" % name,
            ":%s_prod" % name,
        ],
    )

    # Generates the development bundle.
    native.filegroup(
        name = "%s_dev" % name,
        srcs = [
            "development/%s.html" % name,
            "development/%s.js" % name,
            "development/%s.css" % name,
            "development/%s.css.map" % name,
        ],
    )

    # Generates the production bundle.
    native.filegroup(
        name = "%s_prod" % name,
        srcs = [
            "production/%s.html" % name,
            "production/%s.js" % name,
            "production/%s.css" % name,
        ],
    )

def extract_files_from_skia_wasm_container(name, container_files, outs):
    """Extracts files from the Skia WASM container image (gcr.io/skia-public/skia-wasm-release).

    This macro takes as inputs a list of paths inside the Docker container (container_files
    argument), and a list of the same length with the destination paths for each of the files to
    extract (outs argument), relative to the directory where the macro is instantiated.

    Args:
      name: Name of the target.
      container_files: List of absolute paths inside the Docker container to extract.
      outs: Destination paths for each file to extract, relative to the target's directory.
    """

    if len(container_files) != len(outs):
        fail("Arguments container_files and outs must have the same length.")

    # Generates a .tar file with the contents of the image's filesystem (and a .json metadata file
    # which we ignore).
    #
    # See the rule implementation here:
    # https://github.com/bazelbuild/rules_docker/blob/02ad0a48fac9afb644908a634e8b2139c5e84670/container/flatten.bzl#L48
    #
    # Notes:
    #  - It is unclear whether container_flatten is part of the public API because it is not
    #    documented. But the fact that container_flatten is re-exported along with other, well
    #    documented rules in rules_docker suggests that it might indeed be part of the public API
    #    (see [1] and [2]).
    #  - If they ever make breaking changes to container_flatten, then it is probably best to fork
    #    it. The rule itself is relatively small; it is just a wrapper around a Go program that does
    #    all the heavy lifting.
    #  - This rule was chosen because most other rules in the rules_docker repository produce .tar
    #    files with layered outputs, which means we would have to do the flattening ourselves.
    #
    # [1] https://github.com/bazelbuild/rules_docker/blob/6c29619903b6bc533ad91967f41f2a3448758e6f/container/container.bzl#L28
    # [2] https://github.com/bazelbuild/rules_docker/blob/e290d0975ab19f9811933d2c93be275b6daf7427/container/BUILD#L158
    container_flatten(
        name = name + "_skia_wasm_container_filesystem",
        image = "@container_pull_skia_wasm//image",
    )

    # Name of the .tar file produced by the container_flatten target.
    skia_wasm_filesystem_tar = name + "_skia_wasm_container_filesystem.tar"

    # Shell command that returns the directory of the BUILD file where this macro was instantiated.
    #
    # This works because:
    #  - The $< variable[1] is expanded by the below genrule to the path of its only input file.
    #  - The only input file to the genrule is the .tar file produced by the above container_flatten
    #    target (see the genrule's srcs attribute).
    #  - Said .tar file is produced in the directory of the BUILD file where this macro was
    #    instantiated.
    #
    # [1] https://docs.bazel.build/versions/master/be/make-variables.html
    output_dir = "$$(dirname $<)"

    # Directory where we will untar the .tar file produced by the container_flatten target.
    skia_wasm_filesystem_dir = output_dir + "/" + skia_wasm_filesystem_tar + "_untarred"

    native.genrule(
        name = name,
        srcs = [skia_wasm_filesystem_tar],
        outs = outs,
        cmd = " && ".join([
            # Untar the .tar file produced by the container_flatten rule.
            "mkdir -p " + skia_wasm_filesystem_dir,
            "tar xf $< -C " + skia_wasm_filesystem_dir,
        ] + [
            # Copy each requested file from the container filesystem to its desired destination.
            "cp %s/%s %s/%s" % (skia_wasm_filesystem_dir, src, output_dir, dst)
            for src, dst in zip(container_files, outs)
        ]),
    )
