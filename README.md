# TOEFL word random generator

This program print a TOEFL word randomly from the word book "新东方·TOEFL词汇" by 王玉梅.

## Installation
First, [install the Go tools](https://golang.org/doc/install#install)

Second, add `$HOME/go/bin` to your `PATH`. (Otherwise, you can only run this program using its full path, for example `~/go/bin/toefl`.)

After that, run the command below in your terminal
```
$ go get github.com/ladrift/toefl
```
and type `toefl` to get a word to learn!

## Example
```
$ toefl | cowsay
 _______________________________________
/ wholesome                             \
|                                       |
| adj. 有益健康的, [同]healthful, |
\ beneficial                            /
 ---------------------------------------
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||

```

I put this command in my shell startup script, which making me learning and reviewing TOEFL words all the day. :smile:
