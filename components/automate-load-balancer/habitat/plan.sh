pkg_name=automate-load-balancer
pkg_description="internal and external load balancer and reverse proxy for Automate 2.0"
pkg_origin=chef
pkg_version="1.19.2"
pkg_maintainer="Chef Software Inc. <support@chef.io>"
pkg_license=('Chef-MLSA')
pkg_source="https://nginx.org/download/nginx-${pkg_version}.tar.gz"
pkg_upstream_url="https://www.chef.io/automate"
pkg_shasum=7c1f7bb13e79433ee930c597d272a64bc6e30c356a48524f38fd34fa88d62473
pkg_deps=(
  core/glibc
  core/libedit
  core/ncurses
  core/zlib
  core/bzip2
  core/openssl
  core/pcre
)

pkg_build_deps=(
  core/gcc
  core/make
  core/coreutils
)
pkg_lib_dirs=(lib)
pkg_bin_dirs=(sbin)
pkg_include_dirs=(include)

pkg_svc_user="root" # a-l-b needs to listen on 80 and 443
pkg_svc_group="hab" # ensures that hab user has access to all service dirs

pkg_exports=(
  [http-port]=http.port
  [https-port]=https.port
)

chef_automate_hab_binding_mode="relaxed"

pkg_binds_optional=(
  [automate-dex]="port"
  [automate-gateway]="port"
  [automate-ui]="port"
  [session-service]="port"
  [automate-cs-nginx]="port"
  [automate-workflow-nginx]="port"
  [automate-builder-api-proxy]="port"
)

do_download() {
  do_default_download
  download_file "https://github.com/openresty/headers-more-nginx-module/archive/v0.33.tar.gz" "headers-more-nginx-module-0.33.tar.gz" "3dc05d6c20e683596ddabfcc3f63c9d4e9680da75bff1c904566b5508584a6d6"

  tar zxvf ${HAB_CACHE_SRC_PATH}/headers-more-nginx-module-0.33.tar.gz
}

do_build() {
  ./configure \
    --prefix="${pkg_prefix}" \
    --conf-path="${pkg_svc_config_path}/nginx.conf" \
    --sbin-path="${pkg_prefix}/bin/nginx" \
    --pid-path="${pkg_svc_var_path}/nginx.pid" \
    --lock-path="${pkg_svc_var_path}/nginx.lock" \
    --user=hab \
    --group=hab \
    --http-log-path=/dev/stdout \
    --error-log-path=stderr \
    --http-client-body-temp-path="${pkg_svc_var_path}/client-body" \
    --http-proxy-temp-path="${pkg_svc_var_path}/proxy" \
    --http-fastcgi-temp-path="${pkg_svc_var_path}/fastcgi" \
    --http-scgi-temp-path="${pkg_svc_var_path}/scgi" \
    --http-uwsgi-temp-path="${pkg_svc_var_path}/uwsgi" \
    --with-pcre \
    --with-pcre-jit \
    --with-file-aio \
    --with-stream=dynamic \
    --with-stream_ssl_module \
    --with-mail=dynamic \
    --with-http_gunzip_module \
    --with-http_gzip_static_module \
    --with-http_realip_module \
    --with-http_v2_module \
    --with-http_ssl_module \
    --with-http_stub_status_module \
    --with-http_addition_module \
    --with-http_degradation_module \
    --with-http_flv_module \
    --with-http_mp4_module \
    --with-http_secure_link_module \
    --with-http_sub_module \
    --with-http_slice_module \
    --with-cc-opt="${CFLAGS}" \
    --with-ld-opt="${LDFLAGS}" \
    --add-module=${HAB_CACHE_SRC_PATH}/headers-more-nginx-module-0.33

  make
}

do_install() {
  make install
  mkdir -p "${pkg_prefix}/sbin"
  cp "${HAB_CACHE_SRC_PATH}/${pkg_dirname}/objs/nginx" "${pkg_prefix}/sbin"
}
