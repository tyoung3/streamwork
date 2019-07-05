#!/bin/bash
# /mnt/net/u1/projects/development/glade/vdfdcp2/test_vdfd.ndl generated Sun Nov 22 15:44:31 EST 2009 by: vdfd-3.42.3\
exec dfd -e 1 "$0"  --  "$@"

PROCESS VDFD_EDIT lefty "-el" "2" "/home/tyoung3/dfd/vdfdcp2/vdfd.lefty";
PROCESS XSPLIT5 Split "2";
PROCESS SAVE WriteFile;
PROCESS STDERR Stderr;
ATTR STDERR lang="C";
PROCESS SPLIT0 Split "2";
PROCESS PANEL /home/tyoung3/dfd/vdfdcp2/vdfdcp;
ATTR PANEL color="purple";
(VDFD_EDIT)6  => "/dev/null";
(SPLIT0)out2  => "/tmp/dfd_edit0";
(XSPLIT5)out1 => "/tmp/dfd_edit5";
(VDFD_EDIT)5  => in(XSPLIT5);
(VDFD_EDIT)7  => in(SAVE);
(VDFD_EDIT)2  => in(STDERR);
(SPLIT0)out1  => 0(VDFD_EDIT);
(PANEL)out    => in(SPLIT0);
(XSPLIT5)out2 => cmd(PANEL);
