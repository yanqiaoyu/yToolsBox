/* src/config.h.  Generated from config.h.in by configure.  */
/* src/config.h.in.  Generated from configure.ac by autoheader.  */

/* Define if building universal (internal helper macro) */
/* #undef AC_APPLE_UNIVERSAL_BUILD */

/* What version of autogen is installed on this system */
#define AUTOGEN_VERSION ${AUTOGEN_VERSION}

/* Enable debugging code and support for the -d option */
/* #undef DEBUG */

/* Enable dmalloc function arg checking */
/* #undef DMALLOC_FUNC_CHECK */

/* Enable Electric Fence memory debugger */
/* #undef EFENCE */

/* Use 64bit packet counters */
#define ENABLE_64BITS 1

/* Enable dmalloc */
/* #undef ENABLE_DMALLOC */

/* Use shared libraries ( .so .dylib or .tbd ) */
#define ENABLE_DYNAMIC_LINK 1

/* Enable fragroute module */
/* #undef ENABLE_FRAGROUTE */

/* nls support in libopts */
#define ENABLE_NLS 1

/* Enable use of pcap_findalldevs() */
#define ENABLE_PCAP_FINDALLDEVS 1

/* Use static libraries ( .a or .A.tbd ) */
/* #undef ENABLE_STATIC_LINK */

/* Compile tcpbridge */
#define ENABLE_TCPBRIDGE 1

/* Compile tcpliveplay */
#define ENABLE_TCPLIVEPLAY 1

/* Do we have tcpdump and pcap_dump_fopen()? */
#define ENABLE_VERBOSE 1

/* Enable additional debugging code (may affect performance) */
/* #undef EXTRA_DEBUG */

/* fopen(3) accepts a 'b' in the mode flag */
#define FOPEN_BINARY_FLAG "b"

/* fopen(3) accepts a 't' in the mode flag */
#define FOPEN_TEXT_FLAG "t"

/* Are we strictly aligned? */
/* #undef FORCE_ALIGN */

/* Force using BPF for sending packet */
/* #undef FORCE_INJECT_BPF */

/* Force using libdnet for sending packets */
/* #undef FORCE_INJECT_LIBDNET */

/* Force using libpcap's pcap_inject() for sending packets */
/* #undef FORCE_INJECT_PCAP_INJECT */

/* Force using libpcap's pcap_sendpacket() for sending packets */
/* #undef FORCE_INJECT_PCAP_SENDPACKET */

/* Force using Linux's PF_PACKET for sending packets */
/* #undef FORCE_INJECT_PF */

/* Enable GNU Profiler */
/* #undef GPROF */

/* Define to 1 if you have the `alarm' function. */
#define HAVE_ALARM 1

/* Define to 1 if you have the <architecture/i386/pio.h> header file. */
/* #undef HAVE_ARCHITECTURE_I386_PIO_H */

/* Define to 1 if you have the <arpa/inet.h> header file. */
#define HAVE_ARPA_INET_H 1

/* Define to 1 if you have the `atexit' function. */
#define HAVE_ATEXIT 1

/* Do we have BPF device support? */
/* #undef HAVE_BPF */

/* Define to 1 if you have the `bzero' function. */
#define HAVE_BZERO 1

/* Define to 1 if you have the `canonicalize_file_name' function. */
#define HAVE_CANONICALIZE_FILE_NAME 1

/* Define to 1 if you have the `chmod' function. */
#define HAVE_CHMOD 1

/* Define to 1 if you have the `ctime' function. */
#define HAVE_CTIME 1

/* Building Cygwin */
/* #undef HAVE_CYGWIN */

/* Building Apple/Darwin */
/* #undef HAVE_DARWIN */

/* Define this if /dev/zero is readable device */
#define HAVE_DEV_ZERO 1

/* Define to 1 if you have the <dirent.h> header file, and it defines `DIR'.
   */
#define HAVE_DIRENT_H 1

/* Define to 1 if you have the <dlfcn.h> header file. */
#define HAVE_DLFCN_H 1

/* Does pcap.h include a header with DLT_C_HDLC? */
#define HAVE_DLT_C_HDLC 1

/* Does pcap.h include a header with DLT_LINUX_SLL? */
#define HAVE_DLT_LINUX_SLL 1

/* Does libpcap have pcap_datalink_val_to_description? */
#define HAVE_DLT_VAL_TO_DESC 1

/* Define to 1 if you have the <dnet.h> header file. */
/* #undef HAVE_DNET_H */

/* Define to 1 if you don't have `vprintf' but do have `_doprnt.' */
/* #undef HAVE_DOPRNT */

/* Define to 1 if you have the <dumbnet.h> header file. */
/* #undef HAVE_DUMBNET_H */

/* Define to 1 if you have the `dup2' function. */
#define HAVE_DUP2 1

/* Define to 1 if you have the <errno.h> header file. */
#define HAVE_ERRNO_H 1

/* Define to 1 if you have the `fchmod' function. */
#define HAVE_FCHMOD 1

/* Define to 1 if you have the <fcntl.h> header file. */
#define HAVE_FCNTL_H 1

/* Define to 1 if you have the `fork' function. */
#define HAVE_FORK 1

/* Building Free BSD */
/* #undef HAVE_FREEBSD */

/* Define to 1 if fseeko (and presumably ftello) exists and is declared. */
#define HAVE_FSEEKO 1

/* Define to 1 if you have the `fstat' function. */
#define HAVE_FSTAT 1

/* Define to 1 if you have the `gethostbyname' function. */
#define HAVE_GETHOSTBYNAME 1

/* Define to 1 if you have the `getpagesize' function. */
#define HAVE_GETPAGESIZE 1

/* Define to 1 if you have the `gettimeofday' function. */
#define HAVE_GETTIMEOFDAY 1

/* Do we have inet_addr? */
#define HAVE_INET_ADDR 1

/* Do we have inet_aton? */
#define HAVE_INET_ATON 1

/* Define to 1 if you have the `inet_ntoa' function. */
#define HAVE_INET_NTOA 1

/* Do we have inet_ntop? */
#define HAVE_INET_NTOP 1

/* Do we have inet_pton? */
#define HAVE_INET_PTON 1

/* Define to 1 if the system has the type `int16_t'. */
#define HAVE_INT16_T 1

/* Define to 1 if the system has the type `int32_t'. */
#define HAVE_INT32_T 1

/* Define to 1 if the system has the type `int8_t'. */
#define HAVE_INT8_T 1

/* Define to 1 if the system has the type `intptr_t'. */
#define HAVE_INTPTR_T 1

/* Define to 1 if you have the <inttypes.h> header file. */
#define HAVE_INTTYPES_H 1

/* Define to 1 if you have the `ioperm' function. */
#define HAVE_IOPERM 1

/* Define to 1 if you have the `asan' library (-lasan). */
/* #undef HAVE_LIBASAN */

/* Do we have libdnet? */
/* #undef HAVE_LIBDNET */

/* Define to 1 if you have the `gen' library (-lgen). */
/* #undef HAVE_LIBGEN */

/* Define to 1 if you have the <libgen.h> header file. */
#define HAVE_LIBGEN_H 1

/* Define to 1 if you have the `intl' library (-lintl). */
/* #undef HAVE_LIBINTL */

/* Define to 1 if you have the <libintl.h> header file. */
#define HAVE_LIBINTL_H 1

/* Define to 1 if you have the `nsl' library (-lnsl). */
/* #undef HAVE_LIBNSL */

/* Does this version of libpcap support netmap? */
/* #undef HAVE_LIBPCAP_NETMAP */

/* Define to 1 if you have the `resolv' library (-lresolv). */
/* #undef HAVE_LIBRESOLV */

/* Define to 1 if you have the `rt' library (-lrt). */
#define HAVE_LIBRT 1

/* Define to 1 if you have the `socket' library (-lsocket). */
/* #undef HAVE_LIBSOCKET */

/* Define to 1 if you have the <limits.h> header file. */
#define HAVE_LIMITS_H 1

/* Building Linux */
#define HAVE_LINUX 1

/* Define to 1 if you have the `memmove' function. */
#define HAVE_MEMMOVE 1

/* Define to 1 if you have the <memory.h> header file. */
#define HAVE_MEMORY_H 1

/* Define to 1 if you have the `memset' function. */
#define HAVE_MEMSET 1

/* Define to 1 if you have the `mmap' function. */
#define HAVE_MMAP 1

/* Define to 1 if you have the `munmap' function. */
#define HAVE_MUNMAP 1

/* Define to 1 if you have the <ndir.h> header file, and it defines `DIR'. */
/* #undef HAVE_NDIR_H */

/* Define to 1 if you have the <netinet/in.h> header file. */
#define HAVE_NETINET_IN_H 1

/* Define to 1 if you have the <netinet/in_systm.h> header file. */
#define HAVE_NETINET_IN_SYSTM_H 1

/* Do we have netmap support? */
/* #undef HAVE_NETMAP */

/* Does netmap have nm_open function? */
/* #undef HAVE_NETMAP_NM_OPEN */

/* Does netmap struct nmreq have nr_flags defined? */
/* #undef HAVE_NETMAP_NR_FLAGS */

/* Does netmap have NR_REG_MASK defined? */
/* #undef HAVE_NETMAP_NR_REG */

/* Does structure netmap_ring have head/tail defined? */
/* #undef HAVE_NETMAP_RING_HEAD_TAIL */

/* Define to 1 if you have the <net/bpf.h> header file. */
/* #undef HAVE_NET_BPF_H */

/* Define to 1 if you have the <net/route.h> header file. */
#define HAVE_NET_ROUTE_H 1

/* Define to 1 if you have the `ntohll' function. */
/* #undef HAVE_NTOHLL */

/* Building Open BSD */
/* #undef HAVE_OPENBSD */

/* Define this if pathfind(3) works */
/* #undef HAVE_PATHFIND */

/* Do we have libpcapnav? */
/* #undef HAVE_PCAPNAV */

/* Does libpcap have pcap_breakloop? */
/* #undef HAVE_PCAP_BREAKLOOP */

/* Does libpcap have pcap_dump_fopen? */
#define HAVE_PCAP_DUMP_FOPEN 1

/* Does libpcap have pcap_get_selectable_fd? */
#define HAVE_PCAP_GET_SELECTABLE_FD 1

/* Does libpcap have pcap_inject? */
#define HAVE_PCAP_INJECT 1

/* Does libpcap have pcap_sendpacket? */
#define HAVE_PCAP_SENDPACKET 1

/* Does libpcap have pcap_setnonblock? */
#define HAVE_PCAP_SETNONBLOCK 1

/* Does libpcap have pcap_snapshot? */
#define HAVE_PCAP_SNAPSHOT 1

/* Does libpcap have pcap_version[] */
#define HAVE_PCAP_VERSION 1

/* Do we have Linux PF_PACKET socket support? */
#define HAVE_PF_PACKET 1

/* Do we have PF_RING libpcap support? */
/* #undef HAVE_PF_RING_PCAP */

/* Define to 1 if the system has the type `pid_t'. */
#define HAVE_PID_T 1

/* Define to 1 if you have the `poll' function. */
#define HAVE_POLL 1

/* Define to 1 if you have the <poll.h> header file. */
#define HAVE_POLL_H 1

/* Define to 1 if you have the `pow' function. */
/* #undef HAVE_POW */

/* Define to 1 if the system has the type `ptrdiff_t'. */
#define HAVE_PTRDIFF_T 1

/* Define to 1 if you have the `putenv' function. */
#define HAVE_PUTENV 1

/* Define this if we have a functional realpath(3C) */
#define HAVE_REALPATH 1

/* Define to 1 if you have the `regcomp' function. */
#define HAVE_REGCOMP 1

/* Define to 1 if you have the <runetype.h> header file. */
/* #undef HAVE_RUNETYPE_H */

/* Define to 1 if you have the <sched.h> header file. */
#define HAVE_SCHED_H 1

/* Define to 1 if you have the `select' function. */
#define HAVE_SELECT 1

/* Define to 1 if you have the <setjmp.h> header file. */
#define HAVE_SETJMP_H 1

/* Define to 1 if you have the <signal.h> header file. */
#define HAVE_SIGNAL_H 1

/* Define to 1 if the system has the type `size_t'. */
#define HAVE_SIZE_T 1

/* Define to 1 if you have the `snprintf' function. */
#define HAVE_SNPRINTF 1

/* Define to 1 if you have the `socket' function. */
#define HAVE_SOCKET 1

/* Building Solaris */
/* #undef HAVE_SOLARIS */

/* Define to 1 if you have the <stdarg.h> header file. */
#define HAVE_STDARG_H 1

/* Define to 1 if you have the <stdbool.h> header file. */
#define HAVE_STDBOOL_H 1

/* Define to 1 if you have the <stddef.h> header file. */
#define HAVE_STDDEF_H 1

/* Define to 1 if you have the <stdint.h> header file. */
#define HAVE_STDINT_H 1

/* Define to 1 if you have the <stdlib.h> header file. */
#define HAVE_STDLIB_H 1

/* Define to 1 if you have the `strcasecmp' function. */
#define HAVE_STRCASECMP 1

/* Define to 1 if you have the `strchr' function. */
#define HAVE_STRCHR 1

/* Define to 1 if you have the `strcspn' function. */
#define HAVE_STRCSPN 1

/* Define to 1 if you have the `strdup' function. */
#define HAVE_STRDUP 1

/* Define to 1 if you have the `strerror' function. */
#define HAVE_STRERROR 1

/* Define this if strftime() works */
#define HAVE_STRFTIME 1

/* Define to 1 if you have the <strings.h> header file. */
#define HAVE_STRINGS_H 1

/* Define to 1 if you have the <string.h> header file. */
#define HAVE_STRING_H 1

/* Define to 1 if you have the `strlcpy' function. */
/* #undef HAVE_STRLCPY */

/* Define to 1 if you have the `strncpy' function. */
#define HAVE_STRNCPY 1

/* Define to 1 if you have the `strpbrk' function. */
#define HAVE_STRPBRK 1

/* Define to 1 if you have the `strrchr' function. */
#define HAVE_STRRCHR 1

/* Define to 1 if you have the `strsignal' function. */
#define HAVE_STRSIGNAL 1

/* Define to 1 if you have the `strspn' function. */
#define HAVE_STRSPN 1

/* Define to 1 if you have the `strstr' function. */
#define HAVE_STRSTR 1

/* Define to 1 if you have the `strtol' function. */
#define HAVE_STRTOL 1

/* Define to 1 if you have the `strtoul' function. */
#define HAVE_STRTOUL 1

/* Define to 1 if you have the `strtoull' function. */
#define HAVE_STRTOULL 1

/* Define to 1 if `tv_sec' is a member of `struct timeval'. */
#define HAVE_STRUCT_TIMEVAL_TV_SEC 1

/* Building SunOS */
/* #undef HAVE_SUNOS */

/* Define to 1 if you have the <sysexits.h> header file. */
#define HAVE_SYSEXITS_H 1

/* Define to 1 if you have the <sys/dir.h> header file, and it defines `DIR'.
   */
/* #undef HAVE_SYS_DIR_H */

/* Define to 1 if you have the <sys/file.h> header file. */
#define HAVE_SYS_FILE_H 1

/* Define to 1 if you have the <sys/ioctl.h> header file. */
#define HAVE_SYS_IOCTL_H 1

/* Define to 1 if you have the <sys/io.h> header file. */
#define HAVE_SYS_IO_H 1

/* Define to 1 if you have the <sys/limits.h> header file. */
/* #undef HAVE_SYS_LIMITS_H */

/* Define to 1 if you have the <sys/mman.h> header file. */
#define HAVE_SYS_MMAN_H 1

/* Define to 1 if you have the <sys/ndir.h> header file, and it defines `DIR'.
   */
/* #undef HAVE_SYS_NDIR_H */

/* Define to 1 if you have the <sys/param.h> header file. */
#define HAVE_SYS_PARAM_H 1

/* Define to 1 if you have the <sys/poll.h> header file. */
#define HAVE_SYS_POLL_H 1

/* Define to 1 if you have the <sys/procset.h> header file. */
/* #undef HAVE_SYS_PROCSET_H */

/* Define to 1 if you have the <sys/select.h> header file. */
#define HAVE_SYS_SELECT_H 1

/* Define to 1 if you have the <sys/socket.h> header file. */
#define HAVE_SYS_SOCKET_H 1

/* Define to 1 if you have the <sys/stat.h> header file. */
#define HAVE_SYS_STAT_H 1

/* Define to 1 if you have the <sys/stropts.h> header file. */
/* #undef HAVE_SYS_STROPTS_H */

/* Define to 1 if you have the <sys/sysctl.h> header file. */
#define HAVE_SYS_SYSCTL_H 1

/* Define to 1 if you have the <sys/systeminfo.h> header file. */
/* #undef HAVE_SYS_SYSTEMINFO_H */

/* Define to 1 if you have the <sys/time.h> header file. */
#define HAVE_SYS_TIME_H 1

/* Define to 1 if you have the <sys/types.h> header file. */
#define HAVE_SYS_TYPES_H 1

/* Define to 1 if you have the <sys/un.h> header file. */
#define HAVE_SYS_UN_H 1

/* Define to 1 if you have the <sys/wait.h> header file. */
#define HAVE_SYS_WAIT_H 1

/* Do we have tcpdump? */
#define HAVE_TCPDUMP 1

/* Do we have TUNTAP device support? */
#define HAVE_TUNTAP 1

/* Do we have Linux TX_RING socket support? */
/* #undef HAVE_TX_RING */

/* Define to 1 if the system has the type `uint16_t'. */
#define HAVE_UINT16_T 1

/* Define to 1 if the system has the type `uint32_t'. */
#define HAVE_UINT32_T 1

/* Define to 1 if the system has the type `uint8_t'. */
#define HAVE_UINT8_T 1

/* Define to 1 if the system has the type `uintptr_t'. */
#define HAVE_UINTPTR_T 1

/* Define to 1 if the system has the type `uint_t'. */
/* #undef HAVE_UINT_T */

/* Define to 1 if you have the <unistd.h> header file. */
#define HAVE_UNISTD_H 1

/* Define to 1 if you have the <utime.h> header file. */
#define HAVE_UTIME_H 1

/* Define to 1 if you have the <values.h> header file. */
/* #undef HAVE_VALUES_H */

/* Define to 1 if you have the <varargs.h> header file. */
/* #undef HAVE_VARARGS_H */

/* Define to 1 if you have the `vfork' function. */
#define HAVE_VFORK 1

/* Define to 1 if you have the <vfork.h> header file. */
/* #undef HAVE_VFORK_H */

/* Define to 1 if you have the `vprintf' function. */
#define HAVE_VPRINTF 1

/* Define to 1 if you have the `vsnprintf' function. */
#define HAVE_VSNPRINTF 1

/* Define to 1 if you have the <wchar.h> header file. */
#define HAVE_WCHAR_H 1

/* Define to 1 if the system has the type `wchar_t'. */
#define HAVE_WCHAR_T 1

/* Windows/Cygwin */
/* #undef HAVE_WIN32 */

/* Do we have WinPcap? */
/* #undef HAVE_WINPCAP */

/* Define to 1 if the system has the type `wint_t'. */
#define HAVE_WINT_T 1

/* Define to 1 if `fork' works. */
#define HAVE_WORKING_FORK 1

/* Define to 1 if `vfork' works. */
#define HAVE_WORKING_VFORK 1

/* Define to 1 if the system has the type `_Bool'. */
#define HAVE__BOOL 1

/* What is the path (if any) to the libpcap bpf header file? */
/* #undef INCLUDE_PCAP_BPF_HEADER */

/* Version of libdnet */
#define LIBDNET_VERSION ""

/* Define to 1 if `lstat' dereferences a symlink specified with a trailing
   slash. */
#define LSTAT_FOLLOWS_SLASHED_SYMLINK 1

/* Define to the sub-directory where libtool stores uninstalled libraries. */
#define LT_OBJDIR ".libs/"

/* Define to 1 if `major', `minor', and `makedev' are declared in <mkdev.h>.
   */
/* #undef MAJOR_IN_MKDEV */

/* Define to 1 if `major', `minor', and `makedev' are declared in
   <sysmacros.h>. */
#define MAJOR_IN_SYSMACROS 1

/* Define this if optional arguments are disallowed */
/* #undef NO_OPTIONAL_OPT_ARGS */

/* Name of package */
#define PACKAGE "tcpreplay"

/* Define to the address where bug reports for this package should be sent. */
#define PACKAGE_BUGREPORT "https://github.com/appneta/tcpreplay/issues"

/* Define to the full name of this package. */
#define PACKAGE_NAME "tcpreplay"

/* Define to the full name and version of this package. */
#define PACKAGE_STRING "tcpreplay 4.3.4"

/* Define to the one symbol short name of this package. */
#define PACKAGE_TARNAME "tcpreplay"

/* Define to the home page for this package. */
#define PACKAGE_URL "http://tcpreplay.sourceforge.net/"

/* Define to the version of this package. */
#define PACKAGE_VERSION "4.3.4"

/* libpcapnav's version? */
/* #undef PCAPNAV_VERSION */

/* define to a working POSIX compliant shell */
#define POSIX_SHELL "/usr/bin/bash"

/* name of regex header file */
#define REGEX_HEADER <regex.h>

/* The size of `char *', as computed by sizeof. */
#define SIZEOF_CHAR_P 8

/* The size of `int', as computed by sizeof. */
#define SIZEOF_INT 4

/* The size of `long', as computed by sizeof. */
#define SIZEOF_LONG 8

/* The size of `short', as computed by sizeof. */
#define SIZEOF_SHORT 2

/* Define to 1 if you have the ANSI C header files. */
#define STDC_HEADERS 1

/* The tcpdump binary initially used */
#define TCPDUMP_BINARY "/usr/sbin/tcpdump"

/* Enable dumping of trace timestamps at the end of a test */
/* #undef TIMESTAMP_TRACE */

/* Define to 1 if you can safely include both <sys/time.h> and <time.h>. */
#define TIME_WITH_SYS_TIME 1

/* Version number of package */
#define VERSION "4.3.4"

/* Define if using the dmalloc debugging malloc package */
/* #undef WITH_DMALLOC */

/* Define this if a working libregex can be found */
#define WITH_LIBREGEX 1

/* Define WORDS_BIGENDIAN to 1 if your processor stores words with the most
   significant byte first (like Motorola and SPARC, unlike Intel). */
#if defined AC_APPLE_UNIVERSAL_BUILD
# if defined __BIG_ENDIAN__
#  define WORDS_BIGENDIAN 1
# endif
#else
# ifndef WORDS_BIGENDIAN
/* #  undef WORDS_BIGENDIAN */
# endif
#endif

/* Enable large inode numbers on Mac OS X 10.5.  */
#ifndef _DARWIN_USE_64_BIT_INODE
# define _DARWIN_USE_64_BIT_INODE 1
#endif

/* Number of bits in a file offset, on hosts where this is settable. */
/* #undef _FILE_OFFSET_BITS */

/* Define to 1 to make fseeko visible on some hosts (e.g. glibc 2.2). */
/* #undef _LARGEFILE_SOURCE */

/* Define for large files, on AIX-style hosts. */
/* #undef _LARGE_FILES */

/* Define for Solaris 2.5.1 so the uint32_t typedef from <sys/synch.h>,
   <pthread.h>, or <semaphore.h> is not used. If the typedef were allowed, the
   #define below would cause a syntax error. */
/* #undef _UINT32_T */

/* Define for Solaris 2.5.1 so the uint64_t typedef from <sys/synch.h>,
   <pthread.h>, or <semaphore.h> is not used. If the typedef were allowed, the
   #define below would cause a syntax error. */
/* #undef _UINT64_T */

/* Define for Solaris 2.5.1 so the uint8_t typedef from <sys/synch.h>,
   <pthread.h>, or <semaphore.h> is not used. If the typedef were allowed, the
   #define below would cause a syntax error. */
/* #undef _UINT8_T */

/* Define to empty if `const' does not conform to ANSI C. */
/* #undef const */

/* Define to `__inline__' or `__inline' if that's what the C compiler
   calls it, or to nothing if 'inline' is not supported under any name.  */
#ifndef __cplusplus
/* #undef inline */
#endif

/* Define to the type of a signed integer type of width exactly 16 bits if
   such a type exists and the standard includes do not define it. */
/* #undef int16_t */

/* Define to the type of a signed integer type of width exactly 32 bits if
   such a type exists and the standard includes do not define it. */
/* #undef int32_t */

/* Define to the type of a signed integer type of width exactly 64 bits if
   such a type exists and the standard includes do not define it. */
/* #undef int64_t */

/* Define to the type of a signed integer type of width exactly 8 bits if such
   a type exists and the standard includes do not define it. */
/* #undef int8_t */

/* Define to `long int' if <sys/types.h> does not define. */
/* #undef off_t */

/* Define to `int' if <sys/types.h> does not define. */
/* #undef pid_t */

/* Define to `unsigned int' if <sys/types.h> does not define. */
/* #undef size_t */

/* Define to `int' if <sys/types.h> does not define. */
/* #undef ssize_t */

/* Define to `uint16_t' if <sys/types.h> does not define. */
/* #undef u_int16_t */

/* Define to `uint32_t' if <sys/types.h> does not define. */
/* #undef u_int32_t */

/* Define to `uint64_t' if <sys/types.h> does not define. */
/* #undef u_int64_t */

/* Define to `uint8_t' if <sys/types.h> does not define. */
/* #undef u_int8_t */

/* Define to the type of an unsigned integer type of width exactly 16 bits if
   such a type exists and the standard includes do not define it. */
/* #undef uint16_t */

/* Define to the type of an unsigned integer type of width exactly 32 bits if
   such a type exists and the standard includes do not define it. */
/* #undef uint32_t */

/* Define to the type of an unsigned integer type of width exactly 64 bits if
   such a type exists and the standard includes do not define it. */
/* #undef uint64_t */

/* Define to the type of an unsigned integer type of width exactly 8 bits if
   such a type exists and the standard includes do not define it. */
/* #undef uint8_t */

/* Define as `fork' if `vfork' does not work. */
/* #undef vfork */
