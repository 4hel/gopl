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
h_dup2               | bufio.NewScanner now takes *os.File args
i_dup3               | read files at once with ioutil.ReadFile
j_exercise-1.4       | dup with print filename from map[string]string
k_lissajous          | GIF animation with scifi sinus waves
l_fetch              | curl with a few lines
m_exercise-1.7       | fetch using io.Copy to stdout
n_exercise-1.7b      | fetch auto adding http prefix to url
o_exercise-1.9       | fetch printing http status code additionaly
p_fetchall           | fetch in parallel with goroutines
q_exercise-1.10      | fetchall with writing contents to files
r_server1            | web server doing simple echo
s_server2            | web server with request count wieth Mutex
t_server3            | web server printing headers and other http infos
u_server-gif         | web server sending GIF from k\_lissajous


## 2. Program Structure

Program              | Purpose
---                  | ---
a_retpointer         | returning a pointer to local var (escape analysis)
b_echo4              | echo with flag.Parse()
c_cf                 | celsius fahrenheit type conversion using package tempconv
d_bitlevel           | shift and mask on binary numbers
popcount             | counting number of bits set to 1 in a word
tempconv             | celsius fahrenheit type conversion package


## 3. Basic Data Types

Program              | Purpose
---                  | ---
a_surface            | creating a SVG graphic with float calculations
b_exercise-3.3       | adding color gradient from red to blue
c_surfserver         | serving svg via http server


## 4. Composite Types

Program              | Purpose
---                  | ---
a_sha256             | printf sha256 sums of type [32]byte with %x
b_exercise-4.1       | counting the bits set diffently in two sha256 sums
c_exercise-4.2       | reading all STDIN and printing the sha256 sum of it
d_rev                | reverse a slice
e_exercise-4.3       | reverse array [6]int
f_exercise-4.4       | left rotate a slice
g_exercise-4.5       | remove adjacent duplicates from string slice
h_exercise-4.6       | in place squash unicode spaces into ascii space
i_charcount          | charcount NewReader(STDIN) ReadRune()
j_charclass          | count character classes
k_wordfreq           | count word frequencies with NewScanner / Split ScanWords
l_github             | github json api client with structs
m_issues             | command line issue search for github
n_exercise-4.10      | classify github issues search results by age new or old
o_exercise-4.13      | command line client for movie poster download
p_issuesreport       | issue rendering with text template
q_issueshtml         | issue rendering with html template
r_exercise-4.14      | web server rendering html issue info


## 5. Functions

Program              | Purpose
---                  | ---
a_findlinks          | parse html (stdin) and find links
b_exercise-5.1       | findlinks without for loop
c_exercise-5.2       | count types of element nodes in map
d_exercise-5.3       | parse html TextNodes and print to stdout
e_exercise-5.4       | findlinks1 + additionally find other links like images and scripts
