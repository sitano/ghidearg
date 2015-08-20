// Copyright (C) 2015 Ivan Prisyazhnyy <john.koepi@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.
// +build linux

package ghidearg

/*
static int _argc;
static char** _argv;

void hide_arg(int argi, int s, char c) {
  char *arg = _argv[argi] + s;
  while (*arg) { *arg ++= c; }
}

void init(int argc, char **argv, char **envp) {
  _argc = argc;
  _argv = argv;
}

__attribute__((section(".init_array"))) typeof(init) *__init = init;
*/
import "C"

func hideArg(argi int, s int, char rune) {
    C.hide_arg(C.int(argi), C.int(s), C.char(char))
}
