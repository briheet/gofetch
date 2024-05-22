#The latest version
version=7.1.0

#If the latest is not there, revert back to 5
bash_version=${BASH_VERSINFO[0]:-5}
shopt -s eval_unsafe_arith &>/dev/null

sys_locale=${LANG:-C}
XDG_CONFIG_HOME=${XDG_CONFIG_HOME:-${HOME}/.config}
PATH=$PATH:/usr/xpg4/bin:usr/sbin:/sbin:/usr/etc:/usr/home
reset='\e[0m'
shopt -s nocasematch

#Speed up
LC_ALL=C
LANG=C

read -rd '' config <<'EOF'
print_info() {
    info title
    info underline

    info "OS" distro
    info "HOST" model
    info "KERNEL" kernel
    info "UPTIME" uptime
    info "Shell" shell
    info "Resolution" resolution
    info "DE" de
    info "WM" wm
    info "WM Theme" wm_theme
    info "Theme" theme
    info "Icons" icons
    info "Terminal" term
    info "Terminal Font" term_font
    info "CPU" cpu
    info "GPU" gpu
    info "Memory" memory
}
