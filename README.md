# golang

Exercises from the book "The Go Programming Language"


## 1. Tutorial

Program              | Purpose
---                  | ---
a_helloworld         | the classic with unicode possible
b_echo1              | echo with for i++ loop
c_echo2              | echo with for range loop
d_echo3              | echo with strings.Join
e_exercise-1.2       | echo with index printed
f_exercise-1.3       | Benchmark string concat vs join
g_dup1               | check dup lines with bufio.NewScanner
h_dup2               | bufio.Newscanner now takes *os.File args
i_dup3               | read files at once with ioutil.ReadFile
j_exercise-1.4       | dup with print filename from map[string]string
k_lissajous          | GIF animation with scifi sinus waves
l_fetch              | curl with a few lines
m_exercise-1.7       | fetch using io.Copy to stdout
n_exercise-1.7b      | fetch auto adding http prefix to url
o_exercise-1.9       | fetch printing http status code additionaly
