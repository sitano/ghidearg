// Copyright (C) 2015 Ivan Prisyazhnyy <john.koepi@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package ghidearg

import (
    "fmt"
    "os"
)

// hides a flag value from the process list cmd line
func HideArg(flag string, char rune) (i int, err error) {
    var s int
    i, s, err = indexValueOf(os.Args, flag)
    if i < 0 || err != nil {
        return
    }

    hideArg(i, s, char)

    return
}

// looks for an offset in the argv matched requested flag
func indexValueOf(args []string, flag string) (int, int, error) {
    for i := 0; i < len(args); i ++ {
        s := args[i]

        if len(s) == 0 || s[0] != '-' || len(s) == 1 {
            continue
        }

        num_minuses := 1
        if s[1] == '-' {
            num_minuses++
            if len(s) == 2 { // "--" terminates the flags
                i ++
                break
            }
        }

        name := s[num_minuses:]
        if len(name) == 0 || name[0] == '-' || name[0] == '=' {
            continue
        }

        // it's a flag. does it have an argument?
        has_value := false
        var j int
        for j = 1; j < len(name); j ++ { // equals cannot be first
            if name[j] == '=' {
                has_value = true
                name = name[0:j]
                break
            }
        }

        if name == flag {
            if !has_value {
                if i + 1 < len(args) {
                    // value is the current arg
                    has_value = true
                    return i + 1, 0, nil
                }

                return -1, -1, fmt.Errorf("flag needs an argument: -%s", name)
            }

            return i, num_minuses + j + 1, nil
        }
    }

    return -1, -1, fmt.Errorf("not found")
}
